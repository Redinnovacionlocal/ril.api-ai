package entity

type Level struct {
	Descripcion string `json:"descripcion"`
	ODM         string `json:"odm"`
	Acciones    string `json:"acciones_agente"`
}

type QuestionTree struct {
	ID              string   `json:"id"`
	Dimension       string   `json:"dimension"`
	Pregunta        string   `json:"pregunta"`
	Opciones        string   `json:"opciones"`
	Bajo            Level    `json:"bajo"`
	Intermedio      Level    `json:"intermedio"`
	Avanzado        Level    `json:"avanzado"`
	AjusteChica     string   `json:"ajuste_chica"`
	AjusteMedianaGr string   `json:"ajuste_mediana_grande"`
	TagsRAG         []string `json:"tags_rag"`
	EsAncla         bool     `json:"es_ancla"`
	Padre           string   `json:"padre"`
	NotaFacilitador string   `json:"nota_facilitador"`
}
