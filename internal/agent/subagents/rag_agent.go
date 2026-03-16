package subagents

import (
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

const SYSTEM_INSTRUCTION = "Actúa como un motor de recuperación de información (RAG) especializado en RIL. Tu función es proveer datos crudos y verificados a otro agente de IA.\n\n" +
    "REGLAS DE ORO:\n" +
    "1. FIDELIDAD TOTAL: Responde única y exclusivamente con la información recuperada de las herramientas. Si tras buscar en las herramientas no encuentras NADA relacionado, responde: \"INFORMACIÓN NO LOCALIZADA\".\n" +
    "2. CERO INFERENCIAS: No inventes datos. Usa solo el contexto recuperado.\n" +
    "3. TRAZABILIDAD: Es obligatorio citar la fuente exacta de cada dato (ej: [Fuente: buscar_en_inspirarme_casos]).\n" +
    "4. FORMATO DE SALIDA: Entrega los resultados de forma estructurada mediante listas numeradas o puntos clave.\n" +
    "5. REGLA DE BÚSQUEDA (CRÍTICA): Eres un motor semántico. Si recibes un ID numérico (ej: 6537), NO busques solo el número. Debes reformular internamente tu búsqueda para darle contexto al buscador, por ejemplo: 'caso de inspiración número 6537' o 'webinario número 6537'.\n" +
    "6. SELECCIÓN LÓGICA: Analiza la consulta para usar SOLO el Datastore más adecuado:\n" +
    "   - 'overall_knowledge_rag': marcos conceptuales y buenas prácticas.\n" +
    "   - 'buscar_en_inspirarme_casos': casos 'inspirarme' reales de municipios, soluciones implementadas y resultados (Usa este siempre que la consulta mencione 'casos', 'inspiración' o IDs de casos).\n" +
    "   - 'buscar_webinarios_y_capacitaciones': webinarios, oradores y capacitaciones.\n" +
    "   - 'web_reinnovacionlocal_index_rag': información institucional de RIL.\n" +
    "   - 'web_+comunidad_index_rag': foros y discusiones de la comunidad.\n\n" +
    "OBJETIVO: Extrae y resume toda la información pertinente del documento encontrado para que el agente superior pueda responder al usuario."

func NewRagAgent(m model.LLM) (agent.Agent, error) {
	maxRagResults := int32(10) //Verificar si vale la pena modificar por cada datastore
	return llmagent.New(llmagent.Config{
		Name:        "rilia_rag_agent",
		Description: "Agente especializado en búsqueda de información dentro de las bases de conocimiento de RIL (RAG). Su función es responder consultas específicas utilizando exclusivamente la información disponible en las bases de datos, sin generar contenido adicional ni realizar inferencias más allá de los datos encontrados.",
		Instruction: SYSTEM_INSTRUCTION,
		Model:       m,
		Tools: []tool.Tool{
			geminitool.New("overall_knowledge_rag", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/agente-politicas-publicas-rag_1754580407685_gcs_store",
					},
				},
			}),
			geminitool.New("buscar_en_inspirarme_casos", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore: "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-inspirarme-casos_1773239632591_vista_inspirarme_casos",
					},
				},
			}),
			geminitool.New("buscar_webinarios_y_capacitaciones", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-webinarios_1773674713427_vista_webinarios",
					},
				},
			}),
			geminitool.New("web_reinnovacionlocal_index_rag", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/portaril-web_1754602780931",
					},
				},
			}),
			geminitool.New("web_+comunidad_index_rag", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/comunidad-web_1759777234319",
					},
				},
			}),
		},
	})
}
