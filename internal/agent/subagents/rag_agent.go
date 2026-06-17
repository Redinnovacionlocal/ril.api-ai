package subagents

import (
	"context"
	"log"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

const SystemInstruction = "Actúa como un motor de recuperación de información (RAG) especializado en RIL. Tu función es proveer datos crudos y verificados a otro agente de IA.\n\n" +
	"REGLAS DE ORO:\n" +
	"1. FIDELIDAD TOTAL: Responde única y exclusivamente con la información recuperada de las herramientas. Si tras buscar en las herramientas no encuentras NADA relacionado, responde: \"INFORMACIÓN NO LOCALIZADA\".\n" +
	"2. CERO INFERENCIAS: No inventes datos. Usa solo el contexto recuperado.\n" +
	"3. TRAZABILIDAD: Es obligatorio citar la fuente exacta de cada dato (ej: [Fuente: buscar_en_inspirarme_casos]).\n" +
	"4. FORMATO DE SALIDA: Entrega los resultados de forma estructurada mediante listas numeradas o puntos clave.\n" +
	"5. REGLA DE BÚSQUEDA (CRÍTICA): Eres un motor semántico. Si recibes un ID numérico (ej: 6537), NO busques solo el número. Debes reformular internamente tu búsqueda para darle contexto al buscador, por ejemplo: 'caso de inspiración número 6537' o 'webinario número 6537'.\n" +
	"6. SELECCIÓN LÓGICA: Analiza la consulta para usar SOLO el Datastore más adecuado:\n" +
	"   - 'overall_knowledge_rag': marcos conceptuales y buenas prácticas.\n" +
	"   - 'buscar_en_inspirarme_casos': casos 'inspirarme' reales de municipios, soluciones implementadas y resultados. (Usa este siempre que la consulta mencione 'casos', 'inspiración' o IDs de casos. OBLIGATORIO: Al responder con un caso, incluye SIEMPRE la URL oficial que viene al final de la descripción).\n" +
	"   - 'buscar_webinarios_y_capacitaciones': Para contenido audiovisual, charlas de expertos y encuentros sincrónicos grabados. (OBLIGATORIO: Al resumir un webinario, incluye SIEMPRE la URL del portal).\n" +
	"   - 'web_reinnovacionlocal_index_rag': información institucional de RIL.\n" +
	"   - 'buscar_cursos_de_academia': Usa este para formación estructurada, rutas de aprendizaje y certificaciones. (OBLIGATORIO: Incluir siempre el Link de acceso).\n" +
	"   - 'buscar_notas_mas_comunidad': Notas y artículos de la comunidad.\n\n" +
	"OBJETIVO: Extrae y resume toda la información pertinente del documento encontrado para que el agente superior pueda responder al usuario."

type RagInput struct {
	Query string `json:"query"`
}

type RagOutput struct {
	Result string `json:"result"`
	Text   string `json:"text"`
}

func NewRagProxyTool(ctx context.Context, client *genai.Client, modelName string, saveMetadataFunc func(ctx tool.Context, meta *genai.GroundingMetadata)) (tool.Tool, error) {
	maxRagResults := int32(10)

	ragTools := []*genai.Tool{
		{
			// overall_knowledge_rag
			Retrieval: &genai.Retrieval{
				VertexAISearch: &genai.VertexAISearch{
					MaxResults: &maxRagResults,
					Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/agente-politicas-publicas-rag_1754580407685_gcs_store",
				},
			},
		},
		{
			// buscar_en_inspirarme_casos
			Retrieval: &genai.Retrieval{
				VertexAISearch: &genai.VertexAISearch{
					MaxResults: &maxRagResults,
					Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-inspirarme-casos_1773239632591_vista_inspirarme_casos",
				},
			},
		},
		{
			// buscar_webinarios_y_capacitaciones
			Retrieval: &genai.Retrieval{
				VertexAISearch: &genai.VertexAISearch{
					MaxResults: &maxRagResults,
					Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-webinarios_1773674713427_vista_webinarios",
				},
			},
		},
		{
			// web_reinnovacionlocal_index_rag
			Retrieval: &genai.Retrieval{
				VertexAISearch: &genai.VertexAISearch{
					MaxResults: &maxRagResults,
					Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/portaril-web_1754602780931",
				},
			},
		},
		{
			// buscar_cursos_de_academia
			Retrieval: &genai.Retrieval{
				VertexAISearch: &genai.VertexAISearch{
					MaxResults: &maxRagResults,
					Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-academia-cursos_1774889502369_vista_academia_cursos",
				},
			},
		},
		{
			// buscar_notas_mas_comunidad
			Retrieval: &genai.Retrieval{
				VertexAISearch: &genai.VertexAISearch{
					MaxResults: &maxRagResults,
					Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/vista-notas-mas-comunid_1781718816011_vista_notas_mas_comunidad",
				},
			},
		},
	}

	return functiontool.New(functiontool.Config{
		Name:        "consultar_bases_conocimiento_ril",
		Description: "Herramienta OBLIGATORIA para buscar información en las bases de datos internas de RIL (casos inspirarme, webinarios, cursos de academia, comunidad). Pásale la consulta del usuario y te devolverá la información verificada.",
	}, functiontool.Func[RagInput, RagOutput](func(ctx tool.Context, input RagInput) (RagOutput, error) {
		resp, err := client.Models.GenerateContent(ctx, modelName, genai.Text(input.Query), &genai.GenerateContentConfig{
			SystemInstruction: genai.NewContentFromText(SystemInstruction, "system"),
			Tools:             ragTools,
		})
		if err != nil {
			log.Printf("Error crítico en RAG Subagent Proxy: %v", err)
			return RagOutput{Text: "Error al consultar las bases de conocimiento de RIL. Por favor, intenta de nuevo."}, nil
		}
		if len(resp.Candidates) > 0 {
			candidate := resp.Candidates[0]

			if candidate.GroundingMetadata != nil && saveMetadataFunc != nil {
				saveMetadataFunc(ctx, candidate.GroundingMetadata)
			}

			if candidate.Content != nil && len(candidate.Content.Parts) > 0 {
				return RagOutput{Text: candidate.Content.Parts[0].Text}, nil
			}
		}

		return RagOutput{Text: "INFORMACIÓN NO LOCALIZADA"}, nil
	}))
}
