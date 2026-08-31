package tools

import (
	"fmt"
	"log"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/api/iterator"
)

type SearchToolConfig struct {
	ProjectID   string
	Location    string
	DataStoreID string
}

type GetInspireCasesArgs struct {
	Query string `json:"query" description:"Consulta o palabras clave para buscar casos de inspiración y éxito"`
}

func NewGetInspireCasesTool(cfg SearchToolConfig, categoriaAgente string) (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name: "buscar_casos_inspiracion",
		Description: "Busca casos de éxito e inspiración en la base de datos municipal " +
			"relacionados con la temática del agente actual.",
	}, func(ctx tool.Context, args GetInspireCasesArgs) ([]string, error) {

		client, err := discoveryengine.NewSearchClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("error al crear el cliente de discoveryengine: %w", err)
		}
		defer client.Close()

		servingConfig := fmt.Sprintf(
			"projects/%s/locations/%s/collections/default_collection/dataStores/%s/servingConfigs/default_search",
			cfg.ProjectID, cfg.Location, cfg.DataStoreID,
		)

		filterExpression := fmt.Sprintf(`caso_categorias_filter: ANY("%s")`, categoriaAgente)
		log.Printf("[SEARCH INIT] Query: %s | Filter: %s | ServingConfig: %s", args.Query, filterExpression, servingConfig)

		req := &discoveryenginepb.SearchRequest{
			ServingConfig: servingConfig,
			Query:         args.Query,
			Filter:        filterExpression,
			PageSize:      5,
		}

		it := client.Search(ctx, req)
		resultados := make([]string, 0)

		for {
			resp, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Printf("[SEARCH ERROR] Error iterando resultados: %v", err)
				return nil, fmt.Errorf("error iterando resultados de búsqueda: %w", err)
			}

			var documentData map[string]interface{}
			if resp.GetDocument().GetStructData() != nil {
				documentData = resp.GetDocument().GetStructData().AsMap()
			} else if resp.GetDocument().GetDerivedStructData() != nil {
				documentData = resp.GetDocument().GetDerivedStructData().AsMap()
			}

			if desc, ok := documentData["description"].(string); ok {
				resultados = append(resultados, desc)
			} else {
				resultados = append(resultados, fmt.Sprintf("%v", documentData))
			}
		}

		log.Printf("[SEARCH SUCCESS] Total resultados procesados: %d", len(resultados))
		return resultados, nil
	})
}
