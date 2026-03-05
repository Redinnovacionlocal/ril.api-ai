package subagents

import (
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

const SystemInstruction = "Actúa como un motor de recuperación de información (RAG) especializado en RIL. Tu función es proveer datos crudos y verificados a otro agente de IA.\n\nREGLAS DE ORO:\n1. FIDELIDAD TOTAL: Responde única y exclusivamente con la información recuperada de las herramientas. Si la respuesta no está en las bases de datos, responde: \"INFORMACIÓN NO LOCALIZADA\".\n2. CERO INFERENCIAS: No interpretes, no supongas ni añadidas contexto externo. Prohibido alucinar o generar contenido creativo.\n3. TRAZABILIDAD: Es obligatorio citar la fuente exacta de cada dato (ej: [Fuente: inspire_case_rag]).\n4. FORMATO DE SALIDA: Entrega los resultados de forma estructurada mediante listas numeradas o puntos clave. No uses introducciones, saludos ni despedidas.\n5. SELECCIÓN LÓGICA: Analiza la consulta para usar el Datastore más adecuado (ej: casos de éxito en 'inspire_case_rag', políticas públicas en 'overall_knowledge_rag').\n\nOBJETIVO: Ser un filtro quirúrgico de información. Si los datos son contradictorios, expón ambos citando sus fuentes respectivas."

func NewRagAgent(m model.LLM) (agent.Agent, error) {
	maxRagResults := int32(10)
	return llmagent.New(llmagent.Config{
		Name:        "rilia_rag_agent",
		Description: "Agente especializado en búsqueda de información dentro de las bases de conocimiento de RIL (RAG). Su función es responder consultas específicas utilizando exclusivamente la información disponible en las bases de datos, sin generar contenido adicional ni realizar inferencias más allá de los datos encontrados.",
		Instruction: SystemInstruction,
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
			geminitool.New("inspire_case_rag", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-inspirarme-casos_1757079342527_gcs_store",
					},
				},
			}),
			geminitool.New("webinars_rag", &genai.Tool{
				Retrieval: &genai.Retrieval{
					VertexAISearch: &genai.VertexAISearch{
						MaxResults: &maxRagResults,
						Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-webinars_1759509706346_gcs_store",
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
