package shared

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"text/template"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/adk/tool/skilltoolset"
	"google.golang.org/adk/tool/skilltoolset/skill"
	"ril.api-ia/internal/agent/config"
	"ril.api-ia/internal/agent/subagents/askcontextagent"
	sharedTools "ril.api-ia/internal/agent/subagents/shared/tools"
	"ril.api-ia/internal/agent/tools"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/infrastructure/repository/tree_agent"
)

//go:embed instructions/*.tmpl
var InstructionFiles embed.FS

//go:embed all:skills
var SharedSkillsFiles embed.FS

type AgentConfig struct {
	Name             string
	DomainPrefix     string
	SearchCategory   string
	Description      string
	Model            model.LLM
	TreeManager      *tree_agent.TreeCacheManager
	InstructionFiles fs.FS
	SkillsFiles      fs.FS
	DomainTools      []tool.Tool
}

type PromptData struct {
	Tags       []string
	Dimensions []string
}

type LookupTreeArgs struct {
	Query     string `json:"query,omitempty"`
	ID        string `json:"id,omitempty"`
	Dimension string `json:"dimension,omitempty"`
	Tag       string `json:"tag,omitempty"`
}

type LookupTreeResult struct {
	Questions []entity.QuestionTree `json:"questions"`
}

func buildSystemInstruction(data PromptData, specificFS fs.FS) (string, error) {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl := template.New("").Funcs(funcMap)

	tmpl, err := tmpl.ParseFS(InstructionFiles, "instructions/*.tmpl")
	if err != nil {
		return "", fmt.Errorf("error parseando shared templates: %w", err)
	}

	tmpl, err = tmpl.ParseFS(specificFS, "instructions/*.tmpl")
	if err != nil {
		return "", fmt.Errorf("error parseando specific templates: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "main_instruction.tmpl", data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewDomainAgent(ctx context.Context, cfg AgentConfig) (agent.Agent, error) {
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid agent config: %w", err)
	}

	dimensions, err := cfg.TreeManager.GetDimensions(ctx, cfg.DomainPrefix)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo dimensiones para %s: %w", cfg.DomainPrefix, err)
	}

	tags, err := cfg.TreeManager.GetTags(ctx, cfg.DomainPrefix)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo tags para %s: %w", cfg.DomainPrefix, err)
	}

	systemInstruction, err := buildSystemInstruction(PromptData{
		Dimensions: dimensions,
		Tags:       tags,
	}, cfg.InstructionFiles)
	if err != nil {
		return nil, err
	}

	lookupTreeTool, err := functiontool.New(functiontool.Config{
		Name: "lookup_tree_questions",
		Description: `LLAMADA OBLIGATORIA: Cada vez que el usuario introduzca un nuevo tema de gestión, DEBÉS ejecutar esta herramienta ANTES de hacer preguntas de seguimiento. Usa los criterios expertos devueltos para formular tu próxima pregunta.
            Parámetros (usá uno solo): tag (tag exacto del catálogo), dimension (dimensión exacta),
            id (número de pregunta, ej: "1", "34,35,36"), query (texto libre, último recurso).`,
	}, func(ctx tool.Context, args LookupTreeArgs) (LookupTreeResult, error) {
		preguntas, err := cfg.TreeManager.Lookup(ctx, args.ID, args.Dimension, args.Tag, args.Query, cfg.DomainPrefix)
		if err != nil {
			if errors.Is(err, tree_agent.ErrTreeNotConfigured) {
				return LookupTreeResult{}, fmt.Errorf("el árbol de criterios no está disponible todavía, intentá más tarde")
			}
			return LookupTreeResult{}, err
		}
		return LookupTreeResult{Questions: preguntas}, nil
	})
	if err != nil {
		return nil, err
	}

	toolGenerateDocument, err := functiontool.New(functiontool.Config{
		Name:        "generate_document",
		Description: "Genera un documento a partir de un prompt específico con instrucciones claras sobre formato y estructura.",
	}, tools.GenerateDocumentsToolFunc)
	if err != nil {
		return nil, err
	}

	askContext, err := askcontextagent.NewAskContextAgent(ctx)
	if err != nil {
		return nil, err
	}

	searchCfg := sharedTools.SearchToolConfig{
		ProjectID:   os.Getenv("GOOGLE_CLOUD_PROJECT"),
		Location:    os.Getenv("GOOGLE_CLOUD_LOCATION"),
		DataStoreID: "ril-inspirarme-casos_1773239632591_vista_inspirarme_casos",
	}

	getInspireCasesTool, err := sharedTools.NewGetInspireCasesTool(searchCfg, cfg.SearchCategory)
	if err != nil {
		return nil, fmt.Errorf("error creando inspire cases tool: %w", err)
	}

	allTools := []tool.Tool{
		toolGenerateDocument,
		lookupTreeTool,
		getInspireCasesTool,
		agenttool.New(askContext, &agenttool.Config{
			SkipSummarization: false,
		}),
	}
	allTools = append(allTools, cfg.DomainTools...)

	sharedSkillsFS, err := fs.Sub(SharedSkillsFiles, "skills")
	if err != nil {
		return nil, fmt.Errorf("error accediendo a shared skills: %w", err)
	}

	specificSkillsFS, err := fs.Sub(cfg.SkillsFiles, "skills")
	if err != nil {
		return nil, fmt.Errorf("error accediendo a specific skills: %w", err)
	}

	mergedSkillsFS := unionFS{
		fss: []fs.FS{specificSkillsFS, sharedSkillsFS},
	}

	mySkillToolset, err := skilltoolset.New(ctx, skilltoolset.Config{
		Source: skill.NewFileSystemSource(mergedSkillsFS),
	})
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:                cfg.Name,
		Instruction:         systemInstruction,
		GlobalInstruction:   config.GlobalInstruction,
		Description:         cfg.Description,
		Model:               cfg.Model,
		Tools:               allTools,
		Toolsets:            []tool.Toolset{mySkillToolset},
		BeforeToolCallbacks: []llmagent.BeforeToolCallback{SkipAskContextWhenDelegated(cfg.Name)},
	})
}

type unionFS struct {
	fss []fs.FS
}

func (u unionFS) Open(name string) (fs.File, error) {
	for _, f := range u.fss {
		file, err := f.Open(name)
		if err == nil {
			return file, nil
		}
		if !errors.Is(err, fs.ErrNotExist) {
			return nil, err
		}
	}
	return nil, fs.ErrNotExist
}

func (u unionFS) ReadDir(name string) ([]fs.DirEntry, error) {
	var entries []fs.DirEntry
	seen := make(map[string]bool)
	found := false

	for _, f := range u.fss {
		dirs, err := fs.ReadDir(f, name)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return nil, err
		}
		found = true
		for _, d := range dirs {
			if !seen[d.Name()] {
				seen[d.Name()] = true
				entries = append(entries, d)
			}
		}
	}

	if !found {
		return nil, fs.ErrNotExist
	}
	return entries, nil
}

func (cfg AgentConfig) validate() error {
	if cfg.Name == "" {
		return errors.New("name is required")
	}
	if cfg.DomainPrefix == "" {
		return errors.New("domain prefix is required")
	}
	if cfg.Model == nil {
		return errors.New("model is required")
	}
	if cfg.TreeManager == nil {
		return errors.New("tree manager is required")
	}
	if cfg.InstructionFiles == nil {
		return errors.New("instruction files are required")
	}
	if cfg.SkillsFiles == nil {
		return errors.New("skills files are required")
	}
	return nil
}

func SkipAskContextWhenDelegated(agentName string) llmagent.BeforeToolCallback {
	return func(ctx tool.Context, t tool.Tool, args map[string]any) (map[string]any, error) {
		if t.Name() != "ask_context_agent" {
			return nil, nil
		}

		root, _ := ctx.State().Get("root_agent")
		log.Printf("SkipAskContextWhenDelegated: root=%v, agent=%s", root, agentName)
		if root == agentName {
			return nil, nil
		}

		return map[string]any{
			"ask_context_skipped": true,
		}, nil
	}
}
