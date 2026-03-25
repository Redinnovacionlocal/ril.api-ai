package tools

import (
	"google.golang.org/adk/tool"
)

type LookupTreeResponse struct {
	Content any
}

type LookupTreeRequest struct {
	Query string `json:"query" jsonschema:"Consulta sobre la organización de la política de seguridad ciudadana, ejemplo: '¿El gobierno local cuenta en su estructura orgánica con un área específica para el abordaje de la seguridad ciudadana?"`
}

func LookupThreeToolFunc(ctx tool.Context, args LookupTreeRequest) (LookupTreeResponse, error) {
	jsonContent := `{
  "metadata": {
    "generado": "2026-03-11T14:10:13.601657",
    "linea_accion": "Seguridad Ciudadana",
    "total_preguntas": 58,
    "completas": 22,
    "provisionales": 36,
    "por_definir": 0
  },
  "preguntas": [
    {
      "id": "P1",
      "numero": 1,
      "dimension": "Organización",
      "pregunta": "¿El gobierno local cuenta en su estructura orgánica con un área específica para el abordaje de la seguridad ciudadana?",
      "opciones": "No / Si, con rango de Dirección o Subsecretaría dependiente de otra área (Gobierno o similar) / Si, con rango de Secretaría propia",
      "doc_respaldatoria": "Organigrama y/o designación",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "área de seguridad",
        "organigrama",
        "secretaría de seguridad",
        "dirección de seguridad",
        "estructura orgánica",
        "rango del área"
      ],
      "no_odm_ids": [
        "OdM_01"
      ],
      "no_odm_textos": [
        "OdM_01: Contar con una Secretaría o dirección específica de seguridad ciudadana en el organigrama del gobierno local para institucionalizar la temática"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Huerta Grande",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Organigrama oficial publicado en web",
        "Decreto de creación del área",
        "Resolución del intendente",
        "Captura de la web institucional mostrando el área",
        "Acta del Concejo aprobando estructura orgánica"
      ],
      "si_que_hace_bueno": [
        "El área tiene funciones específicas de seguridad ciudadana (no es solo 'Gobierno')",
        "Tiene personal asignado (al menos un responsable formal)",
        "Está vigente (no es un organigrama de otra gestión)",
        "El rango permite interlocución con policía y otras áreas"
      ],
      "si_senales_alerta": [
        "Área que solo existe en el papel sin equipo real",
        "'Seguridad' es apéndice de 'Gobierno' sin funciones propias",
        "Organigrama de gestión anterior no actualizado",
        "Responsable con 5 áreas más a cargo"
      ],
      "si_que_ofrece": "• Mostrar organigramas de ciudades que certificaron\n• Sugerir modelo de decreto de creación\n• Explicar ventajas de rango Secretaría vs. Dirección\n• Conectar con ciudades similares",
      "si_problemas_odm_ids": [
        "OdM_01"
      ],
      "si_problemas_odm_textos": [
        "OdM_01: Contar con una Secretaría o dirección específica de seguridad ciudadana en el organigrama del gobierno local para institucionalizar la temática"
      ],
      "si_problemas_que_ofrece": "• Mostrar organigramas de ciudades que certificaron\n• Sugerir modelo de decreto de creación\n• Explicar ventajas de rango Secretaría vs. Dirección\n• Conectar con ciudades similares",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "facilitador_notas": "• Mostrar organigramas de ciudades que certificaron\n• Sugerir modelo de decreto de creación\n• Explicar ventajas de rango Secretaría vs. Dirección\n• Conectar con ciudades similares\nNormativa",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P2",
      "numero": 2,
      "dimension": "Organización",
      "pregunta": "El/la funcionario/a  a cargo del área de seguridad municipal: ¿es un/a profesional especializado en la materia?",
      "opciones": "No / Es policía retirado / Es civil con formación en seguridad",
      "doc_respaldatoria": "Designación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "funcionario de seguridad",
        "profesional",
        "secretario de seguridad",
        "responsable",
        "director de seguridad",
        "formación del responsable"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Documento de CV. Fotos o pdf de título u hoja de servició"
      ],
      "si_que_hace_bueno": [
        "Tiene designación formal documentada (decreto/resolución)",
        "Profesional con formación verificable en seguridad ciudadana",
        "Experiencia previa en gestión de seguridad"
      ],
      "si_senales_alerta": [
        "Sin designación formal actualizada",
        "Funcionario sin formación específica comprobable",
        "Permanencia inestable sin contrato formal"
      ],
      "si_que_ofrece": "Mostrar cómo otras ciudades definieron requisitos en ordenanzas\nOfrecer templates de designación con requisitos de formación\nSugerir caminos de profesionalización (cursos, certificaciones)\nConectar con red de referentes de ciudades del programa",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "Mostrar cómo otras ciudades definieron requisitos en ordenanzas\nOfrecer templates de designación con requisitos de formación\nSugerir caminos de profesionalización (cursos, certificaciones)\nConectar con red de referentes de ciudades del programa",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Completado por equipo",
      "facilitador_notas": "No están sistematizados (Ver docs de JM, GP, y empezar a sistamtizar)",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P3",
      "numero": 3,
      "dimension": "Organización",
      "pregunta": "¿Desde el gobierno local, se promueven instancias de formación para el equipo técnico-político a cargo de la seguridad?",
      "opciones": "No / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "Plan o programa de formación o capacitación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "capacitación equipo",
        "formación técnica",
        "cursos seguridad",
        "profesionalización",
        "COPEC",
        "diplomatura"
      ],
      "no_odm_ids": [
        "OdM_02"
      ],
      "no_odm_textos": [
        "OdM_02: Desarrollar un programa de formación permanente para los diferentes equipos (personal técnico-político, agentes de monitoreo, guardias locales, entre otros)"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Francisco",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Plan de capacitación formal",
        "Cronograma de actividades realizadas",
        "Registro de asistencia",
        "Certificados emitidos",
        "Material de capacitación (presentaciones, folletos)",
        "Publicaciones en redes/web"
      ],
      "si_que_hace_bueno": [
        "Plan de capacitación con contenidos específicos en seguridad ciudadana",
        "Se realizan al menos 2+ actividades anuales con registros",
        "Participan funcionarios y técnicos de distintas áreas involucradas",
        "Evaluación de participantes y aplicación de aprendizajes documentada"
      ],
      "si_senales_alerta": [
        "Una charla aislada contada como \"capacitación sistemática\"",
        "No participa ni el 50% del equipo",
        "Contenidos genéricos de \"seguridad\" sin aplicación local",
        "Sin registros, certificados o seguimiento de asistencia",
        "Contenidos de tránsito y seguridad vial"
      ],
      "si_que_ofrece": "Ofrecer modelos de programa anual de capacitación adaptado a municipios\nSugerir contenidos según problemáticas locales diagnosticadas\nTemplate de planificación con responsables y cronograma\nMostrar cómo medir impacto en decisiones posteriores",
      "si_problemas_odm_ids": [
        "OdM_02"
      ],
      "si_problemas_odm_textos": [
        "OdM_02: Desarrollar un programa de formación permanente para los diferentes equipos (personal técnico-político, agentes de monitoreo, guardias locales, entre otros)"
      ],
      "si_problemas_que_ofrece": "Ofrecer modelos de programa anual de capacitación adaptado a municipios\nSugerir contenidos según problemáticas locales diagnosticadas\nTemplate de planificación con responsables y cronograma\nMostrar cómo medir impacto en decisiones posteriores",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (capacitacion)",
      "facilitador_notas": "No están sistematizados (empezar a sistamtizar en un excel)",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P4",
      "numero": 4,
      "dimension": "Organización",
      "pregunta": "¿El presupuesto municipal tiene asignada una partida específica para la gestión de la seguridad ciudadana?",
      "opciones": "No / Si, pero sólo para gastos operativos (funcionamiento y personal) / Si, para gastos operativos y proyectos de prevención.",
      "doc_respaldatoria": "Informe de ejecución presupuestaria",
      "criterio_minimo_cert": "NO",
      "tags": [
        "presupuesto seguridad",
        "partida presupuestaria",
        "recursos",
        "asignación presupuestaria",
        "fondos seguridad"
      ],
      "no_odm_ids": [
        "OdM_03, OdM_05, OdM_06"
      ],
      "no_odm_textos": [
        "OdM_03: Asignar partida presupuestaria específica para contratación de personal técnico e implementación de programas de seguridad",
        "OdM_05: Definir criterios de asignación presupuestaria con enfoque territorial (por barrios o zonas de mayor conflictividad)",
        "OdM_06: Diseñar un presupuesto participativo en seguridad ciudadana"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Alicia, San Antonio de Arredondo, Noetinger, Huerta Grande, Arias y 13 más",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Presupuesto anual aprobado por Concejo Deliberante",
        "Informe de ejecución presupuestaria",
        "Partida presupuestaria identificada",
        "Planilla de asignación de recursos",
        "Captura del sistema presupuestario"
      ],
      "si_que_hace_bueno": [
        "Partida presupuestaria específica y diferenciada (código contable único)",
        "Ejecución presupuestaria mayor al 70% anual",
        "Incluye gastos de funcionamiento, personal e inversión"
      ],
      "si_senales_alerta": [
        "Presupuesto genérico sin partida diferenciada de seguridad",
        "Ejecución menor al 50% año a año",
        "Solo gastos de funcionamiento, sin inversión en prevención",
        "No se puede rastrear montos asignados a seguridad ciudadana"
      ],
      "si_que_ofrece": "Mostrar ejemplos de partidas presupuestarias de ciudades similares\nSugerir estructura de presupuesto por programa de seguridad\nAyudar a calcular inversión mínima según población y diagnóstico\nAcompañar presentación ante Concejo Deliberante",
      "si_problemas_odm_ids": [
        "OdM_03, OdM_05, OdM_06"
      ],
      "si_problemas_odm_textos": [
        "OdM_03: Asignar partida presupuestaria específica para contratación de personal técnico e implementación de programas de seguridad",
        "OdM_05: Definir criterios de asignación presupuestaria con enfoque territorial (por barrios o zonas de mayor conflictividad)",
        "OdM_06: Diseñar un presupuesto participativo en seguridad ciudadana"
      ],
      "si_problemas_que_ofrece": "Mostrar ejemplos de partidas presupuestarias de ciudades similares\nSugerir estructura de presupuesto por programa de seguridad\nAyudar a calcular inversión mínima según población y diagnóstico\nAcompañar presentación ante Concejo Deliberante",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (presupuesto)",
      "facilitador_notas": "No",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P5",
      "numero": 5,
      "dimension": "Organización",
      "pregunta": "¿El gobierno local ha recibido financiamiento (en dinero, asistencia o aporte de recursos estratégicos) provenientes de otras fuentes tales como empresas y/o organismos públicos o de cooperación internacional?",
      "opciones": "No / Hemos solicitado pero no lo recibimos / Sí, recibimos fondos de otras fuentes",
      "doc_respaldatoria": "No penaliza ni suma, solo dato",
      "criterio_minimo_cert": "NO",
      "tags": [
        "financiamiento externo",
        "organismos internacionales",
        "cooperación",
        "fondos nacionales",
        "subsidios",
        "BID",
        "CAF"
      ],
      "no_odm_ids": [
        "OdM_04"
      ],
      "no_odm_textos": [
        "OdM_04: Gestionar financiamiento ante organismos internacionales para iniciativas o programas de seguridad"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Alicia, San Antonio de Arredondo, Malvinas Argentinas, Deán Funes",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Presupuesto anual aprobado por Concejo Deliberante",
        "Informe de ejecución presupuestaria",
        "Partida presupuestaria identificada",
        "Planilla de asignación de recursos",
        "Captura del sistema presupuestario",
        "Buscar en qué formato se efectiviza una alianza p-p (convenio, etc) y un financiamiento internacional"
      ],
      "si_que_hace_bueno": [
        "Acceso a fondos de organismos de cooperación internacional",
        "Convenios vigentes con entidades públicas/privadas para inversión",
        "Documentación clara del origen y uso de recursos externos",
        "Cofinanciamiento para proyectos estratégicos identificados"
      ],
      "si_senales_alerta": [
        "Dependencia exclusiva de presupuesto municipal",
        "Solicitudes de fondos rechazadas o ignoradas",
        "Desconocimiento de líneas de financiamiento disponibles",
        "Imposibilidad de acceder a recursos por trámites complejos"
      ],
      "si_que_ofrece": "Mapear líneas de cooperación internacional disponibles para seguridad ciudadana\nAyudar a identificar y contactar potenciales cofinanciadores públicos/privados\nSugerir proyectos que faciliten acceso a fondos externos\nAcompañar gestión de convenios",
      "si_problemas_odm_ids": [
        "OdM_04"
      ],
      "si_problemas_odm_textos": [
        "OdM_04: Gestionar financiamiento ante organismos internacionales para iniciativas o programas de seguridad"
      ],
      "si_problemas_que_ofrece": "Mapear líneas de cooperación internacional disponibles para seguridad ciudadana\nAyudar a identificar y contactar potenciales cofinanciadores públicos/privados\nSugerir proyectos que faciliten acceso a fondos externos\nAcompañar gestión de convenios",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (presupuesto)",
      "facilitador_notas": "Cooperación Internacional ",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P6",
      "numero": 6,
      "dimension": "Organización",
      "pregunta": "¿El gobierno local cuenta con tasa de seguridad municipal o similar?",
      "opciones": "No / Si",
      "doc_respaldatoria": "No penaliza ni suma, solo dato",
      "criterio_minimo_cert": "NO",
      "tags": [
        "tasa de seguridad",
        "tasa municipal",
        "contribución seguridad",
        "tributo",
        "financiamiento propio"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Protocolo formal aprobado por ambas jurisdicciones (municipio-provincia)",
        "Define claramente autoridades responsables de coordinación",
        "Establece canales de comunicación operativa con policía",
        "Incluye instancias de reunión periódica documentadas"
      ],
      "si_senales_alerta": [
        "Acuerdos verbales sin documento normativo que los respalde",
        "Coordinación solo ante emergencias puntuales",
        "Desconexión operativa entre niveles de gobierno",
        "Cambios de autoridades rompen continuidad de coordinación"
      ],
      "si_que_ofrece": "Presentar modelos de protocolos de ciudades certificadas\nSugerir contenidos mínimos del protocolo según realidad provincial\nAcompañar negociación con autoridades policiales\nAyudar a documentar acuerdos en formato vinculante",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "Presentar modelos de protocolos de ciudades certificadas\nSugerir contenidos mínimos del protocolo según realidad provincial\nAcompañar negociación con autoridades policiales\nAyudar a documentar acuerdos en formato vinculante",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P7",
      "numero": 7,
      "dimension": "Organización",
      "pregunta": "¿El gobierno local está adherido a la última Ley de seguridad pública/seguridad ciudadana de su provincia?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Ordenanza",
      "criterio_minimo_cert": "NO",
      "tags": [
        "ley de seguridad",
        "adhesión",
        "normativa provincial",
        "legislación",
        "marco legal"
      ],
      "no_odm_ids": [
        "OdM_09"
      ],
      "no_odm_textos": [
        "OdM_09: Incorporar perspectiva de género en el diseño normativo de la política de seguridad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)"
      ],
      "si_que_hace_bueno": [
        "Documento aprobado formalmente con diagnóstico basado en datos",
        "Define ejes, objetivos específicos y acciones concretas",
        "Incluye indicadores medibles con metas por año",
        "Presupuesto asignado e identificación de responsables"
      ],
      "si_senales_alerta": [
        "Listado de deseos sin diagnóstico fundamentado",
        "Objetivos vagos (\"mejorar la seguridad\") sin acciones específicas",
        "Sin indicadores cuantitativos ni forma de medir",
        "Diseñado por una sola persona sin participación interárea"
      ],
      "si_que_ofrece": "Acompañar proceso de diagnóstico participativo con datos\nSugerir estructura: ejes estratégicos, objetivos, acciones, indicadores\nTemplate de plan adaptable a municipios de distintos tamaños\nMostrar planes de otras ciudades como referencia",
      "si_problemas_odm_ids": [
        "OdM_09"
      ],
      "si_problemas_odm_textos": [
        "OdM_09: Incorporar perspectiva de género en el diseño normativo de la política de seguridad"
      ],
      "si_problemas_que_ofrece": "Acompañar proceso de diagnóstico participativo con datos\nSugerir estructura: ejes estratégicos, objetivos, acciones, indicadores\nTemplate de plan adaptable a municipios de distintos tamaños\nMostrar planes de otras ciudades como referencia",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA y actualizado por eqiuipo",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P8",
      "numero": 8,
      "dimension": "Organización",
      "pregunta": "¿El gobierno local cuenta con una normativa propia  (ordenanza/decreto) que delimite su alcance y responsabilidades en materia de seguridad ciudadana?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Ordenanza",
      "criterio_minimo_cert": "NO",
      "tags": [
        "ordenanza de seguridad",
        "normativa propia",
        "regulación municipal",
        "marco normativo local"
      ],
      "no_odm_ids": [
        "OdM_07, OdM_09"
      ],
      "no_odm_textos": [
        "OdM_07: Poseer una normativa que regule funciones, alcances y responsabilidades del Sistema de Seguridad local y sus interacciones con otros actores estatales y de la comunidad",
        "OdM_09: Incorporar perspectiva de género en el diseño normativo de la política de seguridad"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Colonia Caroya",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Ordenanza del Concejo Deliberante",
        "Decreto del intendente",
        "Resolución municipal",
        "Convenio con la provincia",
        "Capítulo de la Carta Orgánica municipal"
      ],
      "si_que_hace_bueno": [
        "Define explícitamente qué puede y qué no puede hacer el municipio",
        "Distingue funciones municipales de policiales",
        "Vigente y actualizada (últimos 5 años)",
        "Consistente con ley provincial"
      ],
      "si_senales_alerta": [
        "Normativa genérica sin especificar alcance",
        "Ordenanza de más de 10 años",
        "Contradicción con ley provincial",
        "Solo existe como borrador nunca aprobado"
      ],
      "si_que_ofrece": "• Mostrar modelos de ordenanza de ciudades certificadas\n• Explicar contenido según ley provincial vigente\n• Ofrecer template adaptable\n• Señalar artículos imprescindibles vs. opcionales",
      "si_problemas_odm_ids": [
        "OdM_07, OdM_09"
      ],
      "si_problemas_odm_textos": [
        "OdM_07: Poseer una normativa que regule funciones, alcances y responsabilidades del Sistema de Seguridad local y sus interacciones con otros actores estatales y de la comunidad",
        "OdM_09: Incorporar perspectiva de género en el diseño normativo de la política de seguridad"
      ],
      "si_problemas_que_ofrece": "• Mostrar modelos de ordenanza de ciudades certificadas\n• Explicar contenido según ley provincial vigente\n• Ofrecer template adaptable\n• Señalar artículos imprescindibles vs. opcionales",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "facilitador_notas": "• Mostrar modelos de ordenanza de ciudades certificadas\n• Explicar contenido según ley provincial vigente\n• Ofrecer template adaptable\n• Señalar artículos imprescindibles vs. opcionales\nNormativa ",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P9",
      "numero": 9,
      "dimension": "Organización",
      "pregunta": "El código de faltas municipal: ¿es una norma tradicional orientada a la inspección, control y sanción o es una norma moderna que incorpora también la prevención y resolución alternativa de controversias vecinales?",
      "opciones": "Es un código tradicional / Es un código moderno",
      "doc_respaldatoria": "Ordenanza",
      "criterio_minimo_cert": "NO",
      "tags": [
        "código de faltas",
        "código de convivencia",
        "contravenciones",
        "faltas municipales",
        "inspección"
      ],
      "no_odm_ids": [
        "OdM_08"
      ],
      "no_odm_textos": [
        "OdM_08: Contar con un código de convivencia con procedimientos para la inspección, control y sanción así como prevención y resolución alternativa de controversias vecinales"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Cruz del Eje, San Antonio de Arredondo, Malvinas Argentinas, Deán Funes",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Ordenanza del Concejo Deliberante"
      ],
      "si_que_hace_bueno": [
        "Unifica normativa vinculada con la seguridad",
        "Incorpora mecanismos de resolución de conflictos",
        "Contempla instancias de participación ciudadana",
        "Incluye tareas comunitarias en su sistema de sanciones",
        "Clarifica derechos y deberes ciudadanos con lenguaje claro y enfoque pedagógico",
        "Está alineado a la normativa provincial",
        "Vigente y actualizado (últimos 5 años)"
      ],
      "si_senales_alerta": [
        "No incorpora a la seguridad ciudadana como una co responsabilidad del municipio",
        "Se enfoca en el control y sanción de faltas",
        "Se limita a temáticas de seguridad vial e inspecciones"
      ],
      "si_que_ofrece": "• Mostrar modelos de Códigos de Convivencia consolidades y actualizados\n• Explicar contenido según ley provincial vigente\n• Ofrecer template adaptable\n• Señalar artículos imprescindibles vs. opcionales\n• Adapta los objetivos a las capacidades de cada ciudad",
      "si_problemas_odm_ids": [
        "OdM_08"
      ],
      "si_problemas_odm_textos": [
        "OdM_08: Contar con un código de convivencia con procedimientos para la inspección, control y sanción así como prevención y resolución alternativa de controversias vecinales"
      ],
      "si_problemas_que_ofrece": "• Mostrar modelos de Códigos de Convivencia consolidades y actualizados\n• Explicar contenido según ley provincial vigente\n• Ofrecer template adaptable\n• Señalar artículos imprescindibles vs. opcionales\n• Adapta los objetivos a las capacidades de cada ciudad",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "facilitador_notas": "Códigos de Convivencia: primeros pasos",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P10",
      "numero": 10,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿La seguridad ciudadana es un eje estratégico dentro del plan general de gestión del gobierno local?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Plan de Gestión",
      "criterio_minimo_cert": "NO",
      "tags": [
        "plan de gobierno",
        "eje estratégico",
        "seguridad en plan",
        "prioridad política",
        "agenda de gobierno"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Discursos o conferencias del intendente/a",
        "Resolución municipal"
      ],
      "si_que_hace_bueno": [
        "Que sea una temática recurrente en discursos oficiales"
      ],
      "si_senales_alerta": [],
      "si_que_ofrece": "",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P11",
      "numero": 11,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿El gobierno local cuenta con un plan específico de seguridad local?  (escrito/publico/formalizado)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Plan de Gestión",
      "criterio_minimo_cert": "NO",
      "tags": [
        "plan de seguridad",
        "plan local",
        "plan específico",
        "estrategia de seguridad",
        "planificación"
      ],
      "no_odm_ids": [
        "OdM_10, OdM_12"
      ],
      "no_odm_textos": [
        "OdM_10: Diseñar un plan de seguridad local basado en datos y evidencia en función de un diagnóstico de factores de riesgo, capacidades locales y la situación socio/ambiental/delictual de la ciudad",
        "OdM_12: Formalizar normativamente el plan de seguridad local y sus programas a través de ordenanza/s o instrumento similar para garantizar consenso y sostenibilidad en el tiempo"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Cruz del Eje, San Antonio de Arredondo, Malvinas Argentinas, Deán Funes",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal aprobado por decreto",
        "PDF publicado en web",
        "Documento Word en borrador avanzado",
        "Capítulo dentro del plan de gestión general",
        "PPT usada como plan operativo"
      ],
      "si_que_hace_bueno": [
        "Tiene diagnóstico basado en datos",
        "Define ejes, objetivos y acciones concretas",
        "Incluye indicadores y metas medibles",
        "Define responsables y plazos",
        "Elaborado con más de un área"
      ],
      "si_senales_alerta": [
        "Solo un listado de deseos sin diagnóstico",
        "Sin indicadores ni forma de medir",
        "Lo hizo una sola persona",
        "Desactualizado (gestión anterior)",
        "Copia de otro municipio sin adaptar"
      ],
      "si_que_ofrece": "• Guiar elaboración paso a paso\n• Mostrar estructura estándar con ejemplos\n• Ofrecer banco de 358 líneas de acción\n• Comparar con planes de ciudades similares",
      "si_problemas_odm_ids": [
        "OdM_10, OdM_12"
      ],
      "si_problemas_odm_textos": [
        "OdM_10: Diseñar un plan de seguridad local basado en datos y evidencia en función de un diagnóstico de factores de riesgo, capacidades locales y la situación socio/ambiental/delictual de la ciudad",
        "OdM_12: Formalizar normativamente el plan de seguridad local y sus programas a través de ordenanza/s o instrumento similar para garantizar consenso y sostenibilidad en el tiempo"
      ],
      "si_problemas_que_ofrece": "• Guiar elaboración paso a paso\n• Mostrar estructura estándar con ejemplos\n• Ofrecer banco de 358 líneas de acción\n• Comparar con planes de ciudades similares",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P12",
      "numero": 12,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿Dicho plan es elaborado en conjunto entre el área de seguridad y las demás secretarías municipales vinculadas a la gestión de la conflictividad? (secretaría de desarrollo social, educación, salud, espacio público, entre otras).\nAclaración: si el plan no ha sido elaborado marque no",
      "opciones": "No, lo elabora principalmente el área de seguridad / Si, se elabora en conjunto con todas las áreas involucradas",
      "doc_respaldatoria": "Plan de Gestión",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "plan participativo",
        "plan interáreas",
        "elaboración conjunta",
        "gabinete",
        "coordinación interna"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal aprobado por decreto",
        "PDF publicado en web",
        "Documento Word en borrador avanzado",
        "Capítulo dentro del plan de gestión general",
        "PPT usada como plan operativo"
      ],
      "si_que_hace_bueno": [
        "Plan diseñado con participación de salud, educación, desarrollo social, gobierno",
        "Incluye responsables interáreas con compromisos formales",
        "Establece instancias de coordinación periódica documentadas",
        "Reconoce interdependencia entre áreas en objetivos comunes"
      ],
      "si_senales_alerta": [
        "Plan diseñado unilateralmente por área de seguridad",
        "Otras áreas se enteran del plan pero no participaron",
        "Responsables puntuales sin coordinación operativa",
        "Abordaje fragmentado sin visión de complejidad"
      ],
      "si_que_ofrece": "Acompañar convocatoria a áreas clave para participación\nFacilitar mesas de trabajo para definición conjunta de objetivos\nTemplate de plan con espacios de coordinación interárea\nAyudar a sistematizar acuerdos en documentos vinculantes",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "Acompañar convocatoria a áreas clave para participación\nFacilitar mesas de trabajo para definición conjunta de objetivos\nTemplate de plan con espacios de coordinación interárea\nAyudar a sistematizar acuerdos en documentos vinculantes",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (plan_especifico)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P13",
      "numero": 13,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿Dicho plan fue diseñado en función a un diagnóstico previo que contemple entre otros elementos los factores socio-delictuales de los habitantes, el perfil urbanístico del territorio, la\ndinámica delictual y las «zonas calientes» así como las principales demandas y percepciones de la ciudadanía sobre el delito y la violencia en la ciudad?\nAclaración: si el plan no ha sido elaborado marque no",
      "opciones": "No / Si",
      "doc_respaldatoria": "Informe diagnóstico",
      "criterio_minimo_cert": "NO",
      "tags": [
        "diagnóstico previo",
        "plan basado en datos",
        "evidencia",
        "línea base",
        "análisis previo"
      ],
      "no_odm_ids": [
        "OdM_10"
      ],
      "no_odm_textos": [
        "OdM_10: Diseñar un plan de seguridad local basado en datos y evidencia en función de un diagnóstico de factores de riesgo, capacidades locales y la situación socio/ambiental/delictual de la ciudad"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Cruz del Eje, San Antonio de Arredondo, Malvinas Argentinas, Deán Funes",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Informe de encuesta/relevamiento propio",
        "Estudio de consultora externa",
        "Encuesta online (Google Forms)",
        "Datos de aplicación o plataforma digital",
        "Informe dentro de medición más amplia"
      ],
      "si_que_hace_bueno": [
        "Metodología definida y documentada (encuesta, grupos focales, consultas)",
        "Muestra representativa intencional de distintos sectores",
        "Realizado más de una vez para identificar cambios",
        "Resultados sistematizados y utilizados en decisiones posteriores",
        "Integral: que identifique problemas securitarios, factores de riesgo, factores de protección."
      ],
      "si_senales_alerta": [
        "Consulta informal sin metodología establecida",
        "Se realizó una única vez hace más de 3 años",
        "Participación sesgada (siempre mismos actores)",
        "Datos no tabulados ni sistematizados"
      ],
      "si_que_ofrece": "Template de metodología participativa para municipios chicos\nSugerir herramientas gratuitas (Google Forms, Ushahidi, Konceptos)\nAcompañar sistematización de resultados\nMostrar cómo lo hicieron otras ciudades del programa\nGuias para hacer diagnósticos",
      "si_problemas_odm_ids": [
        "OdM_10"
      ],
      "si_problemas_odm_textos": [
        "OdM_10: Diseñar un plan de seguridad local basado en datos y evidencia en función de un diagnóstico de factores de riesgo, capacidades locales y la situación socio/ambiental/delictual de la ciudad"
      ],
      "si_problemas_que_ofrece": "Template de metodología participativa para municipios chicos\nSugerir herramientas gratuitas (Google Forms, Ushahidi, Konceptos)\nAcompañar sistematización de resultados\nMostrar cómo lo hicieron otras ciudades del programa\nGuias para hacer diagnósticos",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (encuestas_relevamientos)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P14",
      "numero": 14,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿Dicho plan cuenta con objetivos, indicadores y metas?\nAclaración: si el plan no ha sido elaborado marque no",
      "opciones": "No / Si",
      "doc_respaldatoria": "Documento de objetivos y metas",
      "criterio_minimo_cert": "NO",
      "tags": [
        "indicadores",
        "metas",
        "objetivos del plan",
        "tablero de gestión",
        "seguimiento del plan"
      ],
      "no_odm_ids": [
        "OdM_11, OdM_14, OdM_20"
      ],
      "no_odm_textos": [
        "OdM_11: Incorporar metas e indicadores de corto, mediano y largo plazo en el plan de seguridad",
        "OdM_14: Contar con un tablero de gestión para monitorear los programas y proyectos de seguridad",
        "OdM_20: Vincular los tableros de gestión con decisiones presupuestarias y operativas"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Antonio de Arredondo, Malvinas Argentinas, Deán Funes",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Creada por acto administrativo vigente (decreto, ordenanza)",
        "Funciona operativamente con reuniones periódicas documentadas",
        "Conocida y reconocida por actores municipales y policiales",
        "Se mantiene activa y se adapta a nuevas necesidades"
      ],
      "si_senales_alerta": [
        "Existe solo en papel sin operatividad real",
        "Desactualizada de gestiones anteriores",
        "Solo la conoce quien la creó",
        "Se ha dejado de usar sin actualización"
      ],
      "si_que_ofrece": "Presentar modelos de sistemas de ciudades certificadas\nSugerir cronograma de funcionamiento (reuniones mensuales/bimestrales)\nTemplate de minuta de reuniones\nAyudar a visibilizar avances para mantener continuidad política",
      "si_problemas_odm_ids": [
        "OdM_11, OdM_14, OdM_20"
      ],
      "si_problemas_odm_textos": [
        "OdM_11: Incorporar metas e indicadores de corto, mediano y largo plazo en el plan de seguridad",
        "OdM_14: Contar con un tablero de gestión para monitorear los programas y proyectos de seguridad",
        "OdM_20: Vincular los tableros de gestión con decisiones presupuestarias y operativas"
      ],
      "si_problemas_que_ofrece": "Presentar modelos de sistemas de ciudades certificadas\nSugerir cronograma de funcionamiento (reuniones mensuales/bimestrales)\nTemplate de minuta de reuniones\nAyudar a visibilizar avances para mantener continuidad política",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "Si",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P15",
      "numero": 15,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿Realizan actividades de monitoreo a los programas y/o proyectos de seguridad local?",
      "opciones": "No / Si, pero sin un tablero de gestión / Si, a través de un tablero de gestión",
      "doc_respaldatoria": "Reporte de seguimiento",
      "criterio_minimo_cert": "NO",
      "tags": [
        "monitoreo",
        "seguimiento",
        "control de gestión",
        "avance de proyectos",
        "reportes de avance"
      ],
      "no_odm_ids": [
        "OdM_13, OdM_16, OdM_18"
      ],
      "no_odm_textos": [
        "OdM_13: Realizar actividades frecuentes de monitoreo de programas y proyectos del plan de seguridad local",
        "OdM_16: Dar seguimiento a los programas en mesas multiagenciales",
        "OdM_18: Dar seguimiento y evaluar iniciativas de prevención contemplando la perspectiva de los ciudadanos destinatarios"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Antonio de Arredondo, Balnearia, Agua de Oro, Pilar, Almafuerte y 5 más",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Tablero en Excel/Google Sheets",
        "Dashboard digital (PowerBI, Data Studio)",
        "Reportes periódicos en Word/PDF",
        "Minutas de reuniones con datos",
        "Planilla simple de avance de metas"
      ],
      "si_que_hace_bueno": [
        "Se actualiza periódicamente (mensual/trimestral)",
        "Tiene indicadores cuantitativos",
        "Se usa para tomar decisiones",
        "Lo ven más personas que el responsable"
      ],
      "si_senales_alerta": [
        "Tablero armado una vez y nunca actualizado",
        "Solo datos cualitativos ('avanzamos bien')",
        "Nadie lo mira ni lo usa",
        "Indicadores imposibles de medir.• No tiene legitimidad o no tiene periocididad"
      ],
      "si_que_ofrece": "• Ofrecer template de tablero adaptado a su plan\n• Sugerir indicadores medibles\n• Mostrar dashboards de otras ciudades\n• Ayudar a definir frecuencia y responsable",
      "si_problemas_odm_ids": [
        "OdM_13, OdM_16, OdM_18"
      ],
      "si_problemas_odm_textos": [
        "OdM_13: Realizar actividades frecuentes de monitoreo de programas y proyectos del plan de seguridad local",
        "OdM_16: Dar seguimiento a los programas en mesas multiagenciales",
        "OdM_18: Dar seguimiento y evaluar iniciativas de prevención contemplando la perspectiva de los ciudadanos destinatarios"
      ],
      "si_problemas_que_ofrece": "• Ofrecer template de tablero adaptado a su plan\n• Sugerir indicadores medibles\n• Mostrar dashboards de otras ciudades\n• Ayudar a definir frecuencia y responsable",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "Es complicado seguir la temporalidad y el avance de un proyecto o plan con minuta de reuniones, lo repensaría",
      "facilitador_notas": "SI",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P16",
      "numero": 16,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿Realizan evaluaciones de los proyectos de seguridad?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Informe de evaluación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "evaluación",
        "evaluación de impacto",
        "resultados",
        "efectividad",
        "medición de resultados"
      ],
      "no_odm_ids": [
        "OdM_15, OdM_19"
      ],
      "no_odm_textos": [
        "OdM_15: Evaluar programas y/o proyectos de prevención a través de metodologías que permitan medir el alcance de los objetivos",
        "OdM_19: Incorporar evaluación de impacto en programas relevantes"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Tablero en Excel/Google Sheets",
        "Dashboard digital (PowerBI, Data Studio)",
        "Reportes periódicos en Word/PDF",
        "Minutas de reuniones con datos",
        "Planilla simple de avance de metas"
      ],
      "si_que_hace_bueno": [
        "Reporte periódico documentado (anual como mínimo)",
        "Analiza avance en indicadores versus metas propuestas",
        "Identifica cambios respecto a diagnóstico inicial",
        "Utilizado para ajustar acciones siguiente período"
      ],
      "si_senales_alerta": [
        "Reportes informales o inexistentes",
        "Solo datos que lucen bien, sin autocrítica",
        "Desconexión entre reporte y ejecución real",
        "No hay mecanismos para ajustar en función de resultados"
      ],
      "si_que_ofrece": "Template de reporte anual de seguridad ciudadana\nMostrar cómo lo hacen otras ciudades del programa\nAcompañar sistematización de datos para reportes\nAyudar a comunicar resultados a ciudadanía",
      "si_problemas_odm_ids": [
        "OdM_15, OdM_19"
      ],
      "si_problemas_odm_textos": [
        "OdM_15: Evaluar programas y/o proyectos de prevención a través de metodologías que permitan medir el alcance de los objetivos",
        "OdM_19: Incorporar evaluación de impacto en programas relevantes"
      ],
      "si_problemas_que_ofrece": "Template de reporte anual de seguridad ciudadana\nMostrar cómo lo hacen otras ciudades del programa\nAcompañar sistematización de datos para reportes\nAyudar a comunicar resultados a ciudadanía",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (monitoreo)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P17",
      "numero": 17,
      "dimension": "Planificación y seguimiento",
      "pregunta": "¿Realizan rendición de cuentas de los resultados del plan local de seguridad ante la ciudadanía?",
      "opciones": "Nunca / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "Informes publicados, pág web, nota periódico, nota diario local",
      "criterio_minimo_cert": "NO",
      "tags": [
        "rendición de cuentas",
        "transparencia",
        "audiencia pública",
        "informe a la ciudadanía"
      ],
      "no_odm_ids": [
        "OdM_17"
      ],
      "no_odm_textos": [
        "OdM_17: Rendir cuentas de la gestión en seguridad ciudadana periódicamente con la ciudadanía"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Informe anual publicado",
        "Página web con resultados",
        "Nota en diario/medio local",
        "Presentación ante Concejo Deliberante",
        "Audiencia pública de rendición"
      ],
      "si_que_hace_bueno": [
        "Realizada con periodicidad definida (anual como mínimo)",
        "Presenta datos cuantitativos y cualitativos concretos",
        "Accesible para ciudadanía común (lenguaje, formato)",
        "Incluye avance en metas y cambios versus período anterior"
      ],
      "si_senales_alerta": [
        "Se hizo una vez y no se repitió",
        "Solo presenta datos positivos sin información de dificultades",
        "Formato inaccesible (jerga técnica, tablas complejas)",
        "No incluye indicadores ni datos concretos"
      ],
      "si_que_ofrece": "Template de informe de rendición de cuentas con ejemplos\nSugerir medios de divulgación accesibles (web, redes, asambleas)\nMostrar cómo presentan otras ciudades certificadas\nAcompañar sesión pública de presentación",
      "si_problemas_odm_ids": [
        "OdM_17"
      ],
      "si_problemas_odm_textos": [
        "OdM_17: Rendir cuentas de la gestión en seguridad ciudadana periódicamente con la ciudadanía"
      ],
      "si_problemas_que_ofrece": "Template de informe de rendición de cuentas con ejemplos\nSugerir medios de divulgación accesibles (web, redes, asambleas)\nMostrar cómo presentan otras ciudades certificadas\nAcompañar sesión pública de presentación",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (rendicion_cuentas)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P18",
      "numero": 18,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Participa el municipio de alguna instancia de coordinación de la política de seguridad local con la Provincia? (ej.  Mesa conjunta con autoridades ministeriales, judiciales y policiales provinciales)",
      "opciones": "No / Si, pero se reúne ante eventualidades / Si, existe un espacio institucionalizado con encuentros frecuentes",
      "doc_respaldatoria": "Informes publicados, Notas periodísticas / acta del área si existe una mesa",
      "criterio_minimo_cert": "NO",
      "tags": [
        "mesa de seguridad",
        "coordinación policial",
        "consejo de seguridad provincial",
        "reunión con policía",
        "articulación"
      ],
      "no_odm_ids": [
        "OdM_24, OdM_25, OdM_27"
      ],
      "no_odm_textos": [
        "OdM_24: Establecer instancias institucionalizadas y regulares de reunión con los actores del gabinete municipal para coordinar acciones en materia de seguridad",
        "OdM_25: Promover instancias institucionalizadas y periódicas de coordinación con las agencias judiciales y policiales que operan en el territorio",
        "OdM_27: Establecer protocolos conjuntos ante emergencias de seguridad o conflictos multiactorales"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Francisco",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "* Minutas de reuniones",
        "* Ordenanza o decreto del intendente con institucionalización de una Mesa de trabajo",
        "* Políticas delineadas en conjunto municipio-provincia",
        "* Convenios interinstitucionales",
        "*  Notas periodisticas de medios oficiales, publciadas",
        "* Informes internos municipales de gestión donde se reporte la actividad de la mesa para todo el gabinete / Comunicaciones oficiales",
        "* Calendario o plan de trabajo"
      ],
      "si_que_hace_bueno": [
        "Minutas: Figura objetivo, participantes con cargo, compromisos, plazos y responsables",
        "Ordenanza Publicada en Boletín Oficial y aprobada por el Concejo Deliberante",
        "en Políticas conjuntas ambos niveles de gobierno figuran expresamente como co-autores o firmantes",
        "Los convenios firmado por ambas partes, con objeto claro, roles y mecanismo de seguimiento",
        "Las notas periodísticas con Fecha, medio identificable, cita a funcionarios de ambos niveles y descripción del contenido",
        "Los informes internos de gestión con mención explícita de la mesa, fechas, temas y acuerdos en canal oficial",
        "Los calendarios de trabajo que tengan fechas definidas, actores convocados, temario tentativo y aprobación formal"
      ],
      "si_senales_alerta": [
        "Minutas que son tomas de notas simples sin estructura, sin firmas ni fecha",
        "Proyectos de ordenanza sin aprobación del Concejo",
        "Políticas unilaterales que mencionan a la provincia de paso, sin firma o adhesión provincial",
        "Convenios institucionales vencidos, sin firma de una de las partes o sin cláusula de implementación concreta",
        "Notas periodísticas que son gacetillas del propio municipio, o que anuncian reuniones futuras sin confirmar que ocurrieron",
        "Decretos internos sin publicación oficial, o que crean la mesa pero no definen su funcionamiento ni periodicidad",
        "Calendarios informales sin respaldo institucional, o planes de trabajo que nunca se actualizan",
        "Comunicaciones por WhatsApp o correos personales sin membrete ni sistema de archivo formal"
      ],
      "si_que_ofrece": "• Generar modelos de minutas estandarizadas con estructura completa para completar\n* Adaptar las capacidades del municipio al objetivo de formalización de vínculos planteado\n* Identificar lo que ya existe en manteria de mesas de trabajo, políticas conjuntas o comunicación institucional.\n• Acercar modelos de ordenanzas de institucionalización adaptables al municipio\n•Verificar si la política o acuerdo acredita autoría conjunta real de ambos niveles\n* revisar organigrama para identificar responsable/s posibles que hagan seguimiento\n•Generar modelo de convenio marco interinstitucional adaptable\n•Acercar modelos de decretos ejecutivos municipales para creación de mesas\n•Generar modelo de plan de trabajo anual de la mesa en consonancia con los objetivos de la gestión\n",
      "si_problemas_odm_ids": [
        "OdM_24, OdM_25, OdM_27"
      ],
      "si_problemas_odm_textos": [
        "OdM_24: Establecer instancias institucionalizadas y regulares de reunión con los actores del gabinete municipal para coordinar acciones en materia de seguridad",
        "OdM_25: Promover instancias institucionalizadas y periódicas de coordinación con las agencias judiciales y policiales que operan en el territorio",
        "OdM_27: Establecer protocolos conjuntos ante emergencias de seguridad o conflictos multiactorales"
      ],
      "si_problemas_que_ofrece": "• Generar modelos de minutas estandarizadas con estructura completa para completar\n* Adaptar las capacidades del municipio al objetivo de formalización de vínculos planteado\n* Identificar lo que ya existe en manteria de mesas de trabajo, políticas conjuntas o comunicación institucional.\n• Acercar modelos de ordenanzas de institucionalización adaptables al municipio\n•Verificar si la política o acuerdo acredita autoría conjunta real de ambos niveles\n* revisar organigrama para identificar responsable/s posibles que hagan seguimiento\n•Generar modelo de convenio marco interinstitucional adaptable\n•Acercar modelos de decretos ejecutivos municipales para creación de mesas\n•Generar modelo de plan de trabajo anual de la mesa en consonancia con los objetivos de la gestión\n",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "facilitador_notas": "Capsula EN ARMADO sobre Gobernanza y participación ciudadana NO TERMINADO: \nhttps://docs.google.com/document/d/1qGcAX3CV0GrebACEldyrmWRGUi7oEBb5X9Ec5AwAyvQ/edit?usp=sharing ",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P19",
      "numero": 19,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Participa el municipio de alguna instancia de coordinación de la política de seguridad local con la Nación? (ej. Mesa conjunta con autoridades ministeriales, judiciales y policiales nacionales)",
      "opciones": "No / Si, pero se reúne ante eventualidades / Si, existe un espacio institucionalizado con encuentros frecuentes",
      "doc_respaldatoria": "Informes publicados, Minuta o temario reunión, captura mail, pág web",
      "criterio_minimo_cert": "NO",
      "tags": [
        "coordinación federal",
        "fuerzas federales",
        "gendarmería",
        "prefectura",
        "articulación multinivel"
      ],
      "no_odm_ids": [
        "OdM_24, OdM_27"
      ],
      "no_odm_textos": [
        "OdM_24: Establecer instancias institucionalizadas y regulares de reunión con los actores del gabinete municipal para coordinar acciones en materia de seguridad",
        "OdM_27: Establecer protocolos conjuntos ante emergencias de seguridad o conflictos multiactorales"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Minutas de reuniones",
        "Ordenanza o decreto de institucionalización de mesa de trabajo",
        "Políticas delineadas en conjunto",
        "Convenios interinstitucionales",
        "Notas periodísticas de actividades conjuntas"
      ],
      "si_que_hace_bueno": [
        "* El municipio accede a información policial y judicial que de otro modo no tiene: estadísticas de delito, causas en trámite, operativos planificados",
        "* Permite alinear recursos: evita duplicar esfuerzos o que municipio y provincia trabajen en sentidos opuestos",
        "* Abre canales para gestionar problemas concretos: solicitar más presencia policial, coordinar respuesta ante eventos, articular con fiscalía",
        "* Mejora la capacidad de respuesta ante crisis porque los vínculos ya están construidos antes de que ocurran",
        "* El municipio puede influir en la agenda de seguridad provincial en lugar de solo recibirla",
        "* Genera confianza entre instituciones que comparten territorio pero tienen dependencias distintas",
        "* Permite aprender de lo que la provincia ya sabe sobre el territorio: patrones delictivos, zonas críticas, perfil de conflictividad",
        "\""
      ],
      "si_senales_alerta": [
        "Minutas sin estructura, sin firmas ni fechas",
        "Ordenanza solo como proyecto sin aprobación",
        "Políticas unilaterales del municipio sin respaldo provincial",
        "Solo coordinación ante emergencias específicas"
      ],
      "si_que_ofrece": "Mostrar ordenanzas de ciudades con coordinación formalizada\nTemplate de protocolo de coordinación municipio-policía\nAcompañar presentación ante autoridades policiales\nConectar con ciudades que lograron coordinación operativa",
      "si_problemas_odm_ids": [
        "OdM_24, OdM_27"
      ],
      "si_problemas_odm_textos": [
        "OdM_24: Establecer instancias institucionalizadas y regulares de reunión con los actores del gabinete municipal para coordinar acciones en materia de seguridad",
        "OdM_27: Establecer protocolos conjuntos ante emergencias de seguridad o conflictos multiactorales"
      ],
      "si_problemas_que_ofrece": "Mostrar ordenanzas de ciudades con coordinación formalizada\nTemplate de protocolo de coordinación municipio-policía\nAcompañar presentación ante autoridades policiales\nConectar con ciudades que lograron coordinación operativa",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Completado por equipo",
      "facilitador_notas": "Capsula EN ARMADO sobre Gobernanza y participación ciudadana NO TERMINADO: \nhttps://docs.google.com/document/d/1qGcAX3CV0GrebACEldyrmWRGUi7oEBb5X9Ec5AwAyvQ/edit?usp=sharing ",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P20",
      "numero": 20,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿El gobierno local cuenta con despliegue de fuerzas federales en su territorio?",
      "opciones": "No / Si",
      "doc_respaldatoria": "No penaliza ni suma, solo dato",
      "criterio_minimo_cert": "NO",
      "tags": [
        "despliegue fuerzas federales",
        "presencia federal",
        "operativo federal",
        "refuerzo policial"
      ],
      "no_odm_ids": [
        "OdM_27"
      ],
      "no_odm_textos": [
        "OdM_27: Establecer protocolos conjuntos ante emergencias de seguridad o conflictos multiactorales"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Operativos conjuntos (interfuerza)",
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "* Que los operativos interfuerza esten delineados previamente y con participación municipal",
        "* que reporten al ministerio de seguridad y al municipio las novedades"
      ],
      "si_senales_alerta": [],
      "si_que_ofrece": "Template de registro de actores con variables clave\nSugerir herramientas para mantenerlo actualizado\nAcompañar mapeo participativo de actores locales\n* ayudar a delinear operativos interfuerza",
      "si_problemas_odm_ids": [
        "OdM_27"
      ],
      "si_problemas_odm_textos": [
        "OdM_27: Establecer protocolos conjuntos ante emergencias de seguridad o conflictos multiactorales"
      ],
      "si_problemas_que_ofrece": "Template de registro de actores con variables clave\nSugerir herramientas para mantenerlo actualizado\nAcompañar mapeo participativo de actores locales\n* ayudar a delinear operativos interfuerza",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Completado por equipo",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P21",
      "numero": 21,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿El gobierno local promueve instancias de participación ciudadana en materia de seguridad? (Ej. foros vecinales de seguridad, juntas de participación ciudadanas, etc.)",
      "opciones": "Nunca / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "Ordenanza. Pieza comunicacional",
      "criterio_minimo_cert": "NO",
      "tags": [
        "participación ciudadana",
        "consejo barrial",
        "junta vecinal",
        "audiencia pública seguridad",
        "foro ciudadano"
      ],
      "no_odm_ids": [
        "OdM_28"
      ],
      "no_odm_textos": [
        "OdM_28: Conformar instancias de participación ciudadana en materia de seguridad (consejos, comités o juntas locales y/o barriales)"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "* Minutas de reuniones",
        "* Ordenanza o decreto del intendente con institucionalización de una Mesa de trabajo",
        "* Políticas delineadas en conjunto municipio-provincia",
        "* Convenios interinstitucionales",
        "*  Notas periodisticas de medios oficiales, publciadas",
        "* Informes internos municipales de gestión donde se reporte la actividad de la mesa para todo el gabinete / Comunicaciones oficiales",
        "* Calendario o plan de trabajo"
      ],
      "si_que_hace_bueno": [
        "* El municipio accede a información policial y judicial que de otro modo no tiene: estadísticas de delito, causas en trámite, operativos planificados",
        "* Permite alinear recursos: evita duplicar esfuerzos o que municipio y provincia trabajen en sentidos opuestos",
        "* Abre canales para gestionar problemas concretos: solicitar más presencia policial, coordinar respuesta ante eventos, articular con fiscalía",
        "* Mejora la capacidad de respuesta ante crisis porque los vínculos ya están construidos antes de que ocurran",
        "* El municipio puede influir en la agenda de seguridad provincial en lugar de solo recibirla",
        "* Genera confianza entre instituciones que comparten territorio pero tienen dependencias distintas",
        "* Permite aprender de lo que la provincia ya sabe sobre el territorio: patrones delictivos, zonas críticas, perfil de conflictividad"
      ],
      "si_senales_alerta": [
        "Minutas que son tomas de notas simples sin estructura, sin firmas ni fecha",
        "Proyectos de ordenanza sin aprobación del Concejo",
        "Políticas unilaterales que mencionan a la provincia de paso, sin firma o adhesión provincial",
        "Convenios institucionales vencidos, sin firma de una de las partes o sin cláusula de implementación concreta",
        "Notas periodísticas que son gacetillas del propio municipio, o que anuncian reuniones futuras sin confirmar que ocurrieron",
        "Decretos internos sin publicación oficial, o que crean la mesa pero no definen su funcionamiento ni periodicidad",
        "Calendarios informales sin respaldo institucional, o planes de trabajo que nunca se actualizan",
        "Comunicaciones por WhatsApp o correos personales sin membrete ni sistema de archivo formal"
      ],
      "si_que_ofrece": "Generar modelos de minutas estandarizadas con estructura completa para completar\n* Adaptar las capacidades del municipio al objetivo de formalización de vínculos planteado\n* Identificar lo que ya existe en manteria de mesas de trabajo, políticas conjuntas o comunicación institucional.\nAcercar modelos de ordenanzas de institucionalización adaptables al municipio\nVerificar si la política o acuerdo acredita autoría conjunta real de ambos niveles\n* revisar organigrama para identificar responsable/s posibles que hagan seguimiento\nGenerar modelo de convenio marco interinstitucional adaptable\nAcercar modelos de decretos ejecutivos municipales para creación de mesas\nGenerar modelo de plan de trabajo anual de la mesa en consonancia con los objetivos de la gestión",
      "si_problemas_odm_ids": [
        "OdM_28"
      ],
      "si_problemas_odm_textos": [
        "OdM_28: Conformar instancias de participación ciudadana en materia de seguridad (consejos, comités o juntas locales y/o barriales)"
      ],
      "si_problemas_que_ofrece": "Generar modelos de minutas estandarizadas con estructura completa para completar\n* Adaptar las capacidades del municipio al objetivo de formalización de vínculos planteado\n* Identificar lo que ya existe en manteria de mesas de trabajo, políticas conjuntas o comunicación institucional.\nAcercar modelos de ordenanzas de institucionalización adaptables al municipio\nVerificar si la política o acuerdo acredita autoría conjunta real de ambos niveles\n* revisar organigrama para identificar responsable/s posibles que hagan seguimiento\nGenerar modelo de convenio marco interinstitucional adaptable\nAcercar modelos de decretos ejecutivos municipales para creación de mesas\nGenerar modelo de plan de trabajo anual de la mesa en consonancia con los objetivos de la gestión",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Completado por equipo",
      "facilitador_notas": "Capsula de participación ciudadana EN ARMADO  https://docs.google.com/document/d/1qGcAX3CV0GrebACEldyrmWRGUi7oEBb5X9Ec5AwAyvQ/edit?usp=sharing \nBUSCAR MODELOS DE MINUTAS",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P22",
      "numero": 22,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "Dichas instancias de participación ciudadana: ¿son representativas de los diversos grupos poblacionales? (mujeres, colectivo LGBTQ+, adultos mayores, población migrante, personas con discapacidad, etc.)",
      "opciones": "No / Algunas veces / Siempre",
      "doc_respaldatoria": "Minuta de reuniones, Listado de participantes (de contar con ellos)",
      "criterio_minimo_cert": "NO",
      "tags": [
        "representatividad",
        "diversidad",
        "género",
        "jóvenes",
        "barrios vulnerables",
        "inclusión"
      ],
      "no_odm_ids": [
        "OdM_29"
      ],
      "no_odm_textos": [
        "OdM_29: Propiciar que los consejos locales de seguridad incluyan representación sectorial"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Ordenanza de creación del espacio participativo",
        "Actas de reuniones",
        "Listado de participantes",
        "Pieza comunicacional de convocatoria",
        "Minuta con resultados"
      ],
      "si_que_hace_bueno": [
        "Realización periódica de consultas sobre percepción de inseguridad",
        "Participan ciudadanos de distintos barrios y sectores",
        "Resultados se utilizan para ajustar estrategias",
        "Se documentan cambios en percepción período a período"
      ],
      "si_senales_alerta": [
        "Una única consulta sin repetición",
        "Solo participan ciudadanos de zonas céntricas",
        "Resultados no se utilizan en decisiones posteriores",
        "Sin sistematización de datos"
      ],
      "si_que_ofrece": "Template de encuesta de percepción de inseguridad\nSugerir metodología participativa para validar resultados\nMostrar herramientas tecnológicas simples (Google Forms)\nAcompañar análisis e incorporación en diagnóstico",
      "si_problemas_odm_ids": [
        "OdM_29"
      ],
      "si_problemas_odm_textos": [
        "OdM_29: Propiciar que los consejos locales de seguridad incluyan representación sectorial"
      ],
      "si_problemas_que_ofrece": "Template de encuesta de percepción de inseguridad\nSugerir metodología participativa para validar resultados\nMostrar herramientas tecnológicas simples (Google Forms)\nAcompañar análisis e incorporación en diagnóstico",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (participacion_ciudadana)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P23",
      "numero": 23,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "Dichas instancias de participación ciudadana: ¿tienen impacto real en la toma de decisiones?",
      "opciones": "No, se toma principalmente como insumo para  el diagnóstico / Si, se toma de insumo para diagnóstico y toma de decisiones",
      "doc_respaldatoria": "Minutas reuniones",
      "criterio_minimo_cert": "NO",
      "tags": [
        "impacto participación",
        "incidencia ciudadana",
        "efecto en políticas",
        "resultados de participación"
      ],
      "no_odm_ids": [
        "OdM_30"
      ],
      "no_odm_textos": [
        "OdM_30: Conformar instancias de participación ciudadana con actores específicos para que participen del diagnóstico e implementación de programas de prevención"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Ordenanza de creación del espacio participativo",
        "Actas de reuniones",
        "Listado de participantes",
        "Pieza comunicacional de convocatoria",
        "Minuta con resultados"
      ],
      "si_que_hace_bueno": [
        "Instancias formales donde ciudadanía contribuye a diagnóstico y decisiones",
        "Participan ciudadanos de distintos sectores y territorios",
        "Los aportes se documentan y se comunica cómo fueron considerados",
        "Hay mecanismos de retroalimentación"
      ],
      "si_senales_alerta": [
        "Participación solo para recopilación de datos",
        "Nunca se consulta sobre decisiones de política",
        "Los aportes se reciben pero no se comunica resultado",
        "Participación meramente informativa"
      ],
      "si_que_ofrece": "Acompañar diseño de espacios con poder de incidencia real\nTemplate de protocolo de participación ciudadana\nMostrar cómo lo hacen ciudades con participación activa\nAyudar a comunicar cómo se incorporan aportes",
      "si_problemas_odm_ids": [
        "OdM_30"
      ],
      "si_problemas_odm_textos": [
        "OdM_30: Conformar instancias de participación ciudadana con actores específicos para que participen del diagnóstico e implementación de programas de prevención"
      ],
      "si_problemas_que_ofrece": "Acompañar diseño de espacios con poder de incidencia real\nTemplate de protocolo de participación ciudadana\nMostrar cómo lo hacen ciudades con participación activa\nAyudar a comunicar cómo se incorporan aportes",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (participacion_ciudadana)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P24",
      "numero": 24,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Cuentan con un mapeo de los actores  estratégicos que conforman el ecosistema de seguridad en la ciudad?",
      "opciones": "No / Los  conocemos pero no hay un registro exhaustivo / Contamos con un registro exhaustivo de todos los actores",
      "doc_respaldatoria": "Hoja tablero o registro",
      "criterio_minimo_cert": "NO",
      "tags": [
        "mapeo de actores",
        "ecosistema de seguridad",
        "actores estratégicos",
        "red de actores",
        "stakeholders"
      ],
      "no_odm_ids": [
        "OdM_21"
      ],
      "no_odm_textos": [
        "OdM_21: Realizar mapeo de los actores que forman parte del ecosistema de seguridad en la ciudad (gubernamentales, judiciales, policiales, académicos, privados, sociedad civil) en vistas al establecimiento de alianzas"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Mapa/listado de actores con datos de contacto",
        "Documento de análisis de stakeholders",
        "Base de datos de organizaciones",
        "Directorio actualizado"
      ],
      "si_que_hace_bueno": [
        "Registro exhaustivo actualizado de actores clave",
        "Incluye datos de contacto, roles y capacidades de cada uno",
        "Actores de diversos sectores (público, privado, sociedad civil)",
        "Se actualiza periódicamente y se utiliza para convocatoria"
      ],
      "si_senales_alerta": [
        "Información dispersa en correos o archivos personales",
        "Datos desactualizados de gestión anterior",
        "Solo se conocen actores de trabajos previos",
        "No hay sistemática de actualización"
      ],
      "si_que_ofrece": "Template de mapeo de actores de seguridad ciudadana\nSugerir frecuencia de actualización (semestral)\nMostrar cómo lo hacen redes de seguridad en otras ciudades\nConectar con plataforma de actores del programa",
      "si_problemas_odm_ids": [
        "OdM_21"
      ],
      "si_problemas_odm_textos": [
        "OdM_21: Realizar mapeo de los actores que forman parte del ecosistema de seguridad en la ciudad (gubernamentales, judiciales, policiales, académicos, privados, sociedad civil) en vistas al establecimiento de alianzas"
      ],
      "si_problemas_que_ofrece": "Template de mapeo de actores de seguridad ciudadana\nSugerir frecuencia de actualización (semestral)\nMostrar cómo lo hacen redes de seguridad en otras ciudades\nConectar con plataforma de actores del programa",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (mapeo_actores)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P25",
      "numero": 25,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Articulan la política de seguridad local con las iniciativas de actores territoriales cuya labor pueda contribuir a disminuir la conflictividad social en los barrios?  (clubes, ONG, iglesias, etc)",
      "opciones": "No / Algunas veces/ informalmente / Frecuentemente y de manera formal",
      "doc_respaldatoria": "Listado de actores",
      "criterio_minimo_cert": "NO",
      "tags": [
        "articulación educación",
        "articulación salud",
        "escuelas",
        "centros de salud",
        "prevención integral"
      ],
      "no_odm_ids": [
        "OdM_23"
      ],
      "no_odm_textos": [
        "OdM_23: Incluir actores del sistema educativo y de salud en acciones preventivas"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del programa con objetivos",
        "Convenio con instituciones",
        "Decreto de creación",
        "Informe de actividades",
        "Registro de beneficiarios/participantes"
      ],
      "si_que_hace_bueno": [
        "Focalizado en población o problemática identificada",
        "Tiene continuidad (no evento aislado)",
        "Involucra más de un área municipal",
        "Tiene indicadores de resultado",
        "Basado en diagnóstico previo"
      ],
      "si_senales_alerta": [
        "Un taller aislado que se hizo una vez",
        "No focalizado: para \"toda la ciudadanía\" sin criterio",
        "Sin indicadores de resultado",
        "Lo lleva una sola persona",
        "Confunde actividad con programa"
      ],
      "si_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "si_problemas_odm_ids": [
        "OdM_23"
      ],
      "si_problemas_odm_textos": [
        "OdM_23: Incluir actores del sistema educativo y de salud en acciones preventivas"
      ],
      "si_problemas_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (programas_proyectos)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P26",
      "numero": 26,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Existe alguna instancia de coordinación o intercambio con los gobiernos locales vecinos en materia de seguridad?",
      "opciones": "No / Si,  coordinamos ante eventualidades / Si, existe un espacio institucionalizado con encuentros frecuentes",
      "doc_respaldatoria": "Minuta o temario reunión, captura mail, pág web",
      "criterio_minimo_cert": "NO",
      "tags": [
        "coordinación municipios vecinos",
        "intermunicipal",
        "región",
        "mancomunidad",
        "cooperación local"
      ],
      "no_odm_ids": [
        "OdM_26"
      ],
      "no_odm_textos": [
        "OdM_26: Promover instancias institucionalizadas y periódicas de coordinación con los gobiernos locales próximos a la ciudad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Minutas de reuniones",
        "Ordenanza o decreto de institucionalización de mesa de trabajo",
        "Políticas delineadas en conjunto",
        "Convenios interinstitucionales",
        "Notas periodísticas de actividades conjuntas"
      ],
      "si_que_hace_bueno": [
        "Minutas con objetivo, participantes con cargo, compromisos, plazos y responsables",
        "Ordenanza publicada y aprobada",
        "Ambos niveles de gobierno firman documentos conjuntos",
        "Encuentros periódicos (no solo ante emergencias)"
      ],
      "si_senales_alerta": [
        "Minutas sin estructura, sin firmas ni fecha",
        "Proyectos de ordenanza sin aprobación",
        "Políticas unilaterales que mencionan a la otra jurisdicción de paso",
        "Solo se coordinan ante emergencias"
      ],
      "si_que_ofrece": "1. Generar modelos de minutas estandarizadas\n2. Adaptar capacidades del municipio al objetivo de formalización\n3. Identificar lo que ya existe y puede formalizarse\n4. Conectar con experiencias de otras ciudades",
      "si_problemas_odm_ids": [
        "OdM_26"
      ],
      "si_problemas_odm_textos": [
        "OdM_26: Promover instancias institucionalizadas y periódicas de coordinación con los gobiernos locales próximos a la ciudad"
      ],
      "si_problemas_que_ofrece": "1. Generar modelos de minutas estandarizadas\n2. Adaptar capacidades del municipio al objetivo de formalización\n3. Identificar lo que ya existe y puede formalizarse\n4. Conectar con experiencias de otras ciudades",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (coordinacion_interjurisdiccional)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P27",
      "numero": 27,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Existe alguna instancia de coordinación de la política de seguridad local con los prestadores de seguridad privada?",
      "opciones": "No / Si,  coordinamos ante eventualidades / Si, existe un espacio institucionalizado con encuentros frecuentes",
      "doc_respaldatoria": "Minuta o temario reunión, captura mail, pág web",
      "criterio_minimo_cert": "NO",
      "tags": [
        "coordinación sector privado",
        "cámaras comerciales",
        "empresas",
        "seguridad privada",
        "público-privada"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Minutas de reuniones",
        "Ordenanza o decreto de institucionalización de mesa de trabajo",
        "Políticas delineadas en conjunto",
        "Convenios interinstitucionales",
        "Notas periodísticas de actividades conjuntas"
      ],
      "si_que_hace_bueno": [
        "Minutas con objetivo, participantes con cargo, compromisos, plazos y responsables",
        "Ordenanza publicada y aprobada",
        "Ambos niveles de gobierno firman documentos conjuntos",
        "Encuentros periódicos (no solo ante emergencias)"
      ],
      "si_senales_alerta": [
        "Minutas sin estructura, sin firmas ni fecha",
        "Proyectos de ordenanza sin aprobación",
        "Políticas unilaterales que mencionan a la otra jurisdicción de paso",
        "Solo se coordinan ante emergencias"
      ],
      "si_que_ofrece": "1. Generar modelos de minutas estandarizadas\n2. Adaptar capacidades del municipio al objetivo de formalización\n3. Identificar lo que ya existe y puede formalizarse\n4. Conectar con experiencias de otras ciudades",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Generar modelos de minutas estandarizadas\n2. Adaptar capacidades del municipio al objetivo de formalización\n3. Identificar lo que ya existe y puede formalizarse\n4. Conectar con experiencias de otras ciudades",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (coordinacion_interjurisdiccional)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P28",
      "numero": 28,
      "dimension": "Gobernanza y participación ciudadana",
      "pregunta": "¿Existe alguna instancia de coordinación de la política de seguridad local con el sector productivo? (cámaras empresariales, comerciales, áreas industriales, sector agropecuario, etc)",
      "opciones": "No / Si,  coordinamos ante eventualidades / Si, existe un espacio institucionalizado con encuentros frecuentes",
      "doc_respaldatoria": "Listado de actores",
      "criterio_minimo_cert": "NO",
      "tags": [
        "coordinación organizaciones sociales",
        "ONG",
        "sociedad civil",
        "iglesias",
        "clubes",
        "organizaciones barriales"
      ],
      "no_odm_ids": [
        "OdM_22"
      ],
      "no_odm_textos": [
        "OdM_22: Promover iniciativas de seguridad de gestión público/privada, sumando actores relevantes del sector productivo local y otras áreas del municipio o del gobierno provincial"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Minutas de reuniones",
        "Ordenanza o decreto de institucionalización de mesa de trabajo",
        "Políticas delineadas en conjunto",
        "Convenios interinstitucionales",
        "Notas periodísticas de actividades conjuntas"
      ],
      "si_que_hace_bueno": [
        "Minutas con objetivo, participantes con cargo, compromisos, plazos y responsables",
        "Ordenanza publicada y aprobada",
        "Ambos niveles de gobierno firman documentos conjuntos",
        "Encuentros periódicos (no solo ante emergencias)"
      ],
      "si_senales_alerta": [
        "Minutas sin estructura, sin firmas ni fecha",
        "Proyectos de ordenanza sin aprobación",
        "Políticas unilaterales que mencionan a la otra jurisdicción de paso",
        "Solo se coordinan ante emergencias"
      ],
      "si_que_ofrece": "1. Generar modelos de minutas estandarizadas\n2. Adaptar capacidades del municipio al objetivo de formalización\n3. Identificar lo que ya existe y puede formalizarse\n4. Conectar con experiencias de otras ciudades",
      "si_problemas_odm_ids": [
        "OdM_22"
      ],
      "si_problemas_odm_textos": [
        "OdM_22: Promover iniciativas de seguridad de gestión público/privada, sumando actores relevantes del sector productivo local y otras áreas del municipio o del gobierno provincial"
      ],
      "si_problemas_que_ofrece": "1. Generar modelos de minutas estandarizadas\n2. Adaptar capacidades del municipio al objetivo de formalización\n3. Identificar lo que ya existe y puede formalizarse\n4. Conectar con experiencias de otras ciudades",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (coordinacion_interjurisdiccional)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P29",
      "numero": 29,
      "dimension": "Gestión de la información",
      "pregunta": "¿El gobierno local cuenta con  un organismo (observatorio o similar) para la producción de conocimiento, sistematización de información en materia delictual y el apoyo en el diseño de políticas públicas basadas en evidencia?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Ordenanza y/o reporte de datos, informes publicados",
      "criterio_minimo_cert": "NO",
      "tags": [
        "observatorio",
        "sistematización datos",
        "producción de información",
        "análisis datos seguridad",
        "centro de datos"
      ],
      "no_odm_ids": [
        "OdM_31, OdM_33, OdM_36, OdM_37"
      ],
      "no_odm_textos": [
        "OdM_31: Impulsar un observatorio o similar que promueva el acceso, la producción, el análisis de información relevante en materia de seguridad ciudadana",
        "OdM_33: Contar con personal técnico especializado para la producción, sistematización y análisis de información en materia de seguridad",
        "OdM_36: Producir datos sistemáticos para el control de gestión de las áreas y equipos de seguridad ciudadana",
        "OdM_37: Realizar informes periódicos de gestión para socializar con el gabinete y demás actores del ecosistema de seguridad"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Francisco",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Captura del sistema en funcionamiento",
        "Captura de página web institucional",
        "PDF con Organigrama de dicho organismo/designación de personal",
        "PDF con ordenanza de creación",
        "Informe técnico del sistema",
        "Contrato o licitación de adquisición",
        "Manual de usuario",
        "Convenio de uso si es de terceros"
      ],
      "si_que_hace_bueno": [
        "Tiene marco normativo de respaldo",
        "Estructura proporcional al territorio/necesidad",
        "Operativo y en uso regular (no solo instalado)",
        "Provee insumos que son utilizados en la gestión del área"
      ],
      "si_senales_alerta": [
        "Estructura creada pero sin funcionamiento activo",
        "Sin publicación periódica de información.",
        "Sin integración de diferentes fuentes de información"
      ],
      "si_que_ofrece": "1. Dimensionar según territorio y necesidades\n2. Modelo de normativa de respaldo\n3. Ejemplo de soluciones de otras ciudades",
      "si_problemas_odm_ids": [
        "OdM_31, OdM_33, OdM_36, OdM_37"
      ],
      "si_problemas_odm_textos": [
        "OdM_31: Impulsar un observatorio o similar que promueva el acceso, la producción, el análisis de información relevante en materia de seguridad ciudadana",
        "OdM_33: Contar con personal técnico especializado para la producción, sistematización y análisis de información en materia de seguridad",
        "OdM_36: Producir datos sistemáticos para el control de gestión de las áreas y equipos de seguridad ciudadana",
        "OdM_37: Realizar informes periódicos de gestión para socializar con el gabinete y demás actores del ecosistema de seguridad"
      ],
      "si_problemas_que_ofrece": "1. Dimensionar según territorio y necesidades\n2. Modelo de normativa de respaldo\n3. Ejemplo de soluciones de otras ciudades",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Similitud Seguridad (sistema_tecnologico)+ Complementado por equipo",
      "facilitador_notas": "No",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P30",
      "numero": 30,
      "dimension": "Gestión de la información",
      "pregunta": "¿El gobierno local cuenta con alguna plataforma tecnológica para georreferenciar eventos en el territorio? (mapas del delito, conflictos, faltas, etc.)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Sistema o plataforma",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "plataforma tecnológica",
        "sistema de información",
        "software",
        "base de datos",
        "SIG",
        "georreferenciación"
      ],
      "no_odm_ids": [
        "OdM_32, OdM_35, OdM_41"
      ],
      "no_odm_textos": [
        "OdM_32: Promover interoperabilidad de sistemas de información entre áreas municipales",
        "OdM_35: Incorporar IA para el análisis de información",
        "OdM_41: Procesar y geolocalizar la información social/ambiental/contravencional/delictual de manera conjunta con otras áreas"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Dashboard digital (PowerBI, Data Studio)",
        "Softwares de información geográfica (Gqis, MyMaps, Shiny)",
        "Desarrollos tecnologicos y plataformas de georreferencia (Multiagencia, bots, agentes de mapeo)"
      ],
      "si_que_hace_bueno": [
        "Que se actualice periodicamente y/o sistemáticamente",
        "Que tenga un único criterio de división geoespacial del territorio",
        "En caso de tener información de varios canales, que no se repita."
      ],
      "si_senales_alerta": [
        "Que no se actualice regularmente",
        "Información mal cruzada o compartimentada por área",
        "No exista un único criterio (alguna información dividada por \"barrios\" y otra por \"zonas\")"
      ],
      "si_que_ofrece": "Facilita informai",
      "si_problemas_odm_ids": [
        "OdM_32, OdM_35, OdM_41"
      ],
      "si_problemas_odm_textos": [
        "OdM_32: Promover interoperabilidad de sistemas de información entre áreas municipales",
        "OdM_35: Incorporar IA para el análisis de información",
        "OdM_41: Procesar y geolocalizar la información social/ambiental/contravencional/delictual de manera conjunta con otras áreas"
      ],
      "si_problemas_que_ofrece": "Facilita informai",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P31",
      "numero": 31,
      "dimension": "Gestión de la información",
      "pregunta": "¿Acceden a información estadística delictual sobre su jurisdicción elaborada por algún ente público provincial? (Ministerio Público, Policía u Observatorio)",
      "opciones": "No accedemos a esas estadísticas / Accedemos solo a requerimiento puntual / Accedemos de manera frecuente",
      "doc_respaldatoria": "",
      "criterio_minimo_cert": "NO",
      "tags": [
        "estadísticas delictuales",
        "datos policiales",
        "información criminal",
        "denuncias",
        "registros delictivos"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos compartida",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P32",
      "numero": 32,
      "dimension": "Gestión de la información",
      "pregunta": "¿Cuentan con un convenio o similar para asegurar el acceso a los datos delictuales producidos por otras jurisdicciones?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Convenio",
      "criterio_minimo_cert": "NO",
      "tags": [
        "convenio datos",
        "acceso información",
        "intercambio datos",
        "acuerdo con policía",
        "compartir información"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal convenio firmado",
        "Informe o reporte escrito de datos compartidos",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos compartidos"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P33",
      "numero": 33,
      "dimension": "Gestión de la información",
      "pregunta": "¿Realizan mapeo y análisis de factores de riesgo y protección sociodelictuales presentes en el territorio?  (condiciones individuales, familiares o sociales que estimulan o evitan el desarrollo del comportamiento delictivo)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Listado de indicadores o reporte",
      "criterio_minimo_cert": "NO",
      "tags": [
        "factores de riesgo",
        "factores de protección",
        "diagnóstico territorial",
        "mapeo de riesgos",
        "vulnerabilidades"
      ],
      "no_odm_ids": [
        "OdM_36, OdM_38, OdM_39, OdM_41"
      ],
      "no_odm_textos": [
        "OdM_36: Producir datos sistemáticos para el control de gestión de las áreas y equipos de seguridad ciudadana",
        "OdM_38: Realizar diagnósticos sobre conflictos, faltas, delitos, factores sociales y ambientales vinculados a seguridad",
        "OdM_39: Realizar diagnósticos puntuales sobre problemáticas de seguridad priorizadas",
        "OdM_41: Procesar y geolocalizar la información social/ambiental/contravencional/delictual de manera conjunta con otras áreas"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Las Varillas",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Mapa/listado de personas intervenidas por el municipio",
        "Documento de análisis de factores de riesgo/territoriales",
        "Base de datos de organizaciones",
        "Informe de redes y articulaciones",
        "evidencia de reuniones con áreas sociales"
      ],
      "si_que_hace_bueno": [
        "Incluye actores de distintos sectores (público intermunicipal, privado, sociedad civil)",
        "Tiene datos de contacto actualizados",
        "Identifica roles y capacidades de cada actor",
        "Se actualiza periódicamente",
        "Se usa para convocar y articular"
      ],
      "si_senales_alerta": [
        "Lista informal incompleta",
        "Datos de contacto desactualizados",
        "Solo incluye actores con los que ya se trabaja",
        "No distingue roles ni capacidades",
        "Conocimiento en la cabeza de una persona"
      ],
      "si_que_ofrece": "1. Template de mapeo de actores\n2. Metodología para identificación sistemática\n3. Mostrar mapeos de otras ciudades\n4. Sugerir herramienta digital para mantenerlo actualizado",
      "si_problemas_odm_ids": [
        "OdM_36, OdM_38, OdM_39, OdM_41"
      ],
      "si_problemas_odm_textos": [
        "OdM_36: Producir datos sistemáticos para el control de gestión de las áreas y equipos de seguridad ciudadana",
        "OdM_38: Realizar diagnósticos sobre conflictos, faltas, delitos, factores sociales y ambientales vinculados a seguridad",
        "OdM_39: Realizar diagnósticos puntuales sobre problemáticas de seguridad priorizadas",
        "OdM_41: Procesar y geolocalizar la información social/ambiental/contravencional/delictual de manera conjunta con otras áreas"
      ],
      "si_problemas_que_ofrece": "1. Template de mapeo de actores\n2. Metodología para identificación sistemática\n3. Mostrar mapeos de otras ciudades\n4. Sugerir herramienta digital para mantenerlo actualizado",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (mapeo_actores)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P34",
      "numero": 34,
      "dimension": "Gestión de la información",
      "pregunta": "¿Realizan mapeo y análisis  sobre factores situacional-ambientales vinculados a la seguridad? (iluminación, requerimiento de poda, terrenos baldíos, basurales, etc.)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Listado de indicadores o reporte",
      "criterio_minimo_cert": "NO",
      "tags": [
        "factores situacionales",
        "CPTED",
        "espacio público",
        "iluminación",
        "terrenos baldíos",
        "prevención ambiental"
      ],
      "no_odm_ids": [
        "OdM_38, OdM_42"
      ],
      "no_odm_textos": [
        "OdM_38: Realizar diagnósticos sobre conflictos, faltas, delitos, factores sociales y ambientales vinculados a seguridad",
        "OdM_42: Promover instancias de consulta y/o participación ciudadana sobre las problemáticas de seguridad"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Las Varillas",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Mapa/listado de actores con datos de contacto",
        "Documento de análisis de stakeholders",
        "Base de datos de organizaciones",
        "Informe de redes y articulaciones",
        "Directorio actualizado"
      ],
      "si_que_hace_bueno": [
        "Incluye actores de distintos sectores (público, privado, sociedad civil)",
        "Tiene datos de contacto actualizados",
        "Identifica roles y capacidades de cada actor",
        "Se actualiza periódicamente",
        "Se usa para convocar y articular"
      ],
      "si_senales_alerta": [
        "Lista informal incompleta",
        "Datos de contacto desactualizados",
        "Solo incluye actores con los que ya se trabaja",
        "No distingue roles ni capacidades",
        "Conocimiento en la cabeza de una persona"
      ],
      "si_que_ofrece": "1. Template de mapeo de actores\n2. Metodología para identificación sistemática\n3. Mostrar mapeos de otras ciudades\n4. Sugerir herramienta digital para mantenerlo actualizado",
      "si_problemas_odm_ids": [
        "OdM_38, OdM_42"
      ],
      "si_problemas_odm_textos": [
        "OdM_38: Realizar diagnósticos sobre conflictos, faltas, delitos, factores sociales y ambientales vinculados a seguridad",
        "OdM_42: Promover instancias de consulta y/o participación ciudadana sobre las problemáticas de seguridad"
      ],
      "si_problemas_que_ofrece": "1. Template de mapeo de actores\n2. Metodología para identificación sistemática\n3. Mostrar mapeos de otras ciudades\n4. Sugerir herramienta digital para mantenerlo actualizado",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (mapeo_actores)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P35",
      "numero": 35,
      "dimension": "Gestión de la información",
      "pregunta": "¿Realizan encuestas de victimización y percepción del delito?",
      "opciones": "Nunca / Alguna vez / Frecuentemente",
      "doc_respaldatoria": "Reporte de resultados",
      "criterio_minimo_cert": "NO",
      "tags": [
        "encuesta victimización",
        "percepción del delito",
        "encuesta seguridad",
        "sensación de inseguridad"
      ],
      "no_odm_ids": [
        "OdM_40"
      ],
      "no_odm_textos": [
        "OdM_40: Realizar encuestas de victimización"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Informe de encuesta propia",
        "Estudio de consultora externa",
        "Encuesta online (Google Forms)",
        "Datos de app con módulo percepción",
        "Encuesta dentro de medición más amplia"
      ],
      "si_que_hace_bueno": [
        "Tiene metodología definida",
        "Muestra representativa o al menos intencional",
        "Se hace periodicamente lo que permite su comparación",
        "Los resultados se usaron para algo"
      ],
      "si_senales_alerta": [
        "Consulta informal en reunión barrial",
        "Se hizo una vez hace varios años y no se repitió",
        "Muestra sesgada (algún grupo/barrio/etc sobre o subrepresentado)",
        "No se tabularon resultados"
      ],
      "si_que_ofrece": "• Ofrecer modelo adaptado a municipios chicos\n• Sugerir herramientas gratuitas (Google Forms + guía)\n• Mostrar cómo lo hicieron otras ciudades\n• Ayudar a tabular y analizar\n• Asistencia para construir la muestra",
      "si_problemas_odm_ids": [
        "OdM_40"
      ],
      "si_problemas_odm_textos": [
        "OdM_40: Realizar encuestas de victimización"
      ],
      "si_problemas_que_ofrece": "• Ofrecer modelo adaptado a municipios chicos\n• Sugerir herramientas gratuitas (Google Forms + guía)\n• Mostrar cómo lo hicieron otras ciudades\n• Ayudar a tabular y analizar\n• Asistencia para construir la muestra",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P36",
      "numero": 36,
      "dimension": "Gestión de la información",
      "pregunta": "¿Se brinda acceso a la ciudadanía a la información vinculada a las problemáticas de seguridad del distrito? (a través de página web o similar)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Pág WEB",
      "criterio_minimo_cert": "NO",
      "tags": [
        "datos abiertos",
        "transparencia datos",
        "acceso público información",
        "publicación datos",
        "portal datos"
      ],
      "no_odm_ids": [
        "OdM_34, OdM_42"
      ],
      "no_odm_textos": [
        "OdM_34: Publicar los datos que genera el observatorio o área de sistematización de información sobre seguridad",
        "OdM_42: Promover instancias de consulta y/o participación ciudadana sobre las problemáticas de seguridad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Link a página web/ portal de datos abiertos",
        "Captura de página",
        "Publicación en redes sociales oficiales",
        "Presentación de datos en medios (radio; tv; diario)"
      ],
      "si_que_hace_bueno": [
        "* Tienen publicación periódica (semestral o con mayor frecuencia)",
        "Son abiertos y accesibles (se pueden descargar y utilizar)",
        "Los resultados se presentan en instancias que permiten la participación ciudadana",
        "Se publican siempre los mismos tipos de datos o categorías (permite comparación)",
        "Se integran fuentes de información (Denuncias MPF/Policía; Registro de Guardia Local, etc.)"
      ],
      "si_senales_alerta": [
        "Los datos no son accesibles por la ciudadanía",
        "Se publican con frecuencias variables",
        "No siempre se publica la misma información",
        "Se presentan datos sin especificar la fuente"
      ],
      "si_que_ofrece": "* Propone recategorizaciones que permitan la integración de la información \n• Integra diferentes fuentes en una única base",
      "si_problemas_odm_ids": [
        "OdM_34, OdM_42"
      ],
      "si_problemas_odm_textos": [
        "OdM_34: Publicar los datos que genera el observatorio o área de sistematización de información sobre seguridad",
        "OdM_42: Promover instancias de consulta y/o participación ciudadana sobre las problemáticas de seguridad"
      ],
      "si_problemas_que_ofrece": "* Propone recategorizaciones que permitan la integración de la información \n• Integra diferentes fuentes en una única base",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P37",
      "numero": 37,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El gobierno local cuenta con un Centro o Comando de Operaciones que articule a diversos actores vinculados a la seguridad en la localidad?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Ordenanza o plan de coordinación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "centro de operaciones",
        "comando operaciones",
        "COEM",
        "sala de situación",
        "monitoreo operativo"
      ],
      "no_odm_ids": [
        "OdM_43, OdM_44"
      ],
      "no_odm_textos": [
        "OdM_43: Promover la conformación de un Comando Unificado de Operaciones con las fuerzas de seguridad provinciales, nacionales y locales",
        "OdM_44: Protocolizar la actuación del Comando unificado"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Existe formalmente (aprobado/creado por acto administrativo)",
        "Está operativo y en uso regular",
        "Es conocido por los actores relevantes",
        "Se actualiza o mantiene periódicamente"
      ],
      "si_senales_alerta": [
        "Existe en el papel pero no opera en la práctica",
        "Está desactualizado o en desuso",
        "Solo lo conoce quien lo creó",
        "No se ha adaptado a las necesidades actuales"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [
        "OdM_43, OdM_44"
      ],
      "si_problemas_odm_textos": [
        "OdM_43: Promover la conformación de un Comando Unificado de Operaciones con las fuerzas de seguridad provinciales, nacionales y locales",
        "OdM_44: Protocolizar la actuación del Comando unificado"
      ],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P38",
      "numero": 38,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El área de seguridad municipal participa en la definición de intervenciones en el espacio público? (tareas de desmalezamiento, iluminación, planeamiento urbano, etc?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Captura tablero de incidentes y derivaciones/ misiones y funciones del área",
      "criterio_minimo_cert": "NO",
      "tags": [
        "intervención urbana",
        "espacio público",
        "CPTED",
        "urbanismo seguro",
        "diseño urbano",
        "plazas",
        "iluminación"
      ],
      "no_odm_ids": [
        "OdM_45, OdM_46, OdM_47"
      ],
      "no_odm_textos": [
        "OdM_45: Promover desde el área de seguridad municipal intervenciones urbanas en la ciudad",
        "OdM_46: Diseñar intervenciones urbanas con participación comunitaria",
        "OdM_47: Evaluar el impacto de intervenciones urbanas sobre percepción de seguridad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [
        "OdM_45, OdM_46, OdM_47"
      ],
      "si_problemas_odm_textos": [
        "OdM_45: Promover desde el área de seguridad municipal intervenciones urbanas en la ciudad",
        "OdM_46: Diseñar intervenciones urbanas con participación comunitaria",
        "OdM_47: Evaluar el impacto de intervenciones urbanas sobre percepción de seguridad"
      ],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P39",
      "numero": 39,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El municipio contrata servicios adicionales de vigilancia y seguridad a la policía provincial? (para patrullaje, seguridad en escuelas y espacios públicos u otros).",
      "opciones": "No / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "",
      "criterio_minimo_cert": "NO",
      "tags": [
        "vigilancia privada",
        "seguridad privada",
        "vigilador",
        "guardia privada",
        "servicio de vigilancia"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P40",
      "numero": 40,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿Cuenta el gobierno local con un cuerpo civil de prevención propio? (guardia urbana, local o similar)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Ordenenza",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "guardia urbana",
        "cuerpo de prevención",
        "agentes municipales",
        "patrullaje",
        "prevención urbana",
        "caños de escape",
        "tránsito municipal"
      ],
      "no_odm_ids": [
        "OdM_49, OdM_51"
      ],
      "no_odm_textos": [
        "OdM_49: Poner en marcha un cuerpo civil de prevención, capacitado y con móviles para el patrullaje",
        "OdM_51: Desarrollar instancias ciudadanas de presentación de las guardias locales a los vecinos para dar a conocer su rol y desempeño"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron Rosario, Noetinger, San Francisco",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Ordenanza del Concejo Deliberante",
        "Decreto de creación",
        "Resolución que designa personal",
        "Nómina actualizada de agentes",
        "Registro fotográfico",
        "Convenio si es con terceros",
        "Capacitaciones realizadas por los agentes"
      ],
      "si_que_hace_bueno": [
        "Ordenanza o decreto vigente que especifique función del cuerpo de prevención",
        "Dotación proporcional (ref: 1 cada 3.000-5.000 hab)",
        "Personal capacitado",
        "Protocolos de actuación escritos",
        "Cobertura horaria definida (no solo días hábiles 8-14)"
      ],
      "si_senales_alerta": [
        "12 agentes para 50.000 habitantes",
        "Sin protocolos formales",
        "Sin capacitación en seguridad",
        "Solo función de tránsito",
        "No operan fines de semana/noche",
        "Ordenanza sin reglamentar"
      ],
      "si_que_ofrece": "• Calcular dotación sugerida según población\n• Modelo de ordenanza de ciudades certificadas\n• Template de protocolos de patrullaje\n• Calcular costo estimado\n• Conectar con ciudades de referencia",
      "si_problemas_odm_ids": [
        "OdM_49, OdM_51"
      ],
      "si_problemas_odm_textos": [
        "OdM_49: Poner en marcha un cuerpo civil de prevención, capacitado y con móviles para el patrullaje",
        "OdM_51: Desarrollar instancias ciudadanas de presentación de las guardias locales a los vecinos para dar a conocer su rol y desempeño"
      ],
      "si_problemas_que_ofrece": "• Calcular dotación sugerida según población\n• Modelo de ordenanza de ciudades certificadas\n• Template de protocolos de patrullaje\n• Calcular costo estimado\n• Conectar con ciudades de referencia",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "Dotación personal: 1agente/1000 habitantes segun evidencia internacional ¿de donde salio 1 cada 3.000-5.000?\nRegistro fotográfico no creo que sea un formato válido para determinar si tienen o no cuerpo de prevención, salvo que sea el formato del documento q le pedimos que carguen la odenanza o decreto (por ej.: JPG.PNG.)",
      "facilitador_notas": "- Para calcular la dotación sugeridad tiene que tener acceso a la población de la ciudad en la base de datos del censo 2022. \n- Modelo de Ordenanzas de ciudades certificadas tenemos, pero no están agrupadas en una misma carpeta y la mayoría están en proceso de validación. Este es el Estatuto de la Guardia de San Francisco: Estatuto Guardia Local 7324.pdf \nEn esta carpeta hay modelos de normativas de otras ciudades argetinas: Guardias Locales \n- No hay template de protocolos de patrullaje. Avanzamos en recomendaciones en este documento GUARDIA URBANA \n- Fórmula de cálculo de costo de patrullaje lo comenzamos a idear pero no está verificado en este documento: GUARDIA URBANA \n- Información de contacto con ciudades de refrencia se puede obtener en portal RIL; en comunidad de Seguridad Ciudadana y en saleforce los datos de contacto de funcionarios de seguridad ciudadana",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P41",
      "numero": 41,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "Dicho cuerpo de prevención: ¿posee móviles destinados a patrullaje?",
      "opciones": "No / Si",
      "doc_respaldatoria": "copia licitación/acta de entrega",
      "criterio_minimo_cert": "NO",
      "tags": [
        "móviles",
        "patrullaje",
        "vehículos",
        "motos",
        "bicicletas",
        "flota",
        "recorrido"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        ".Subejecución"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n4. Conectar con experiencias y recursos de la Red\ncruzar gps del móvil con mapa del delito para analiar el buen uso del móvil, cantidad suficiente, etc.",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n4. Conectar con experiencias y recursos de la Red\ncruzar gps del móvil con mapa del delito para analiar el buen uso del móvil, cantidad suficiente, etc.",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Generado por IA revisado por equipo",
      "facilitador_notas": "No",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P42",
      "numero": 42,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El gobierno local promueve instancias de formación para el personal de guardia urbana? (propias o en colaboración con otro nivel de gobierno o instituciones)",
      "opciones": "No / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "Plan o programa de formación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "formación guardia",
        "capacitación agentes",
        "entrenamiento",
        "protocolo actuación",
        "academia"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Plan de capacitación formal",
        "Cronograma de actividades realizadas",
        "Registro de asistencia",
        "Certificados emitidos",
        "Material de capacitación (presentaciones, folletos)",
        "Publicaciones en redes/web"
      ],
      "si_que_hace_bueno": [
        "Tiene contenido específico de seguridad ciudadana (no genérico)",
        "Se hizo más de una vez en el año",
        "Llega a distintos públicos (no siempre los mismos)",
        "Fortalece la capacidad de resolución de conflictos del cuerpo de prevención",
        "Responde a una demanda del propio cuerpo como oportunidad de mejora",
        "Permite profesionalizar a los agentes"
      ],
      "si_senales_alerta": [
        "Una charla aislada cuenta como \"capacitación\"",
        "Siempre van las mismas personas",
        "Contenido genérico sin adaptar al territorio o a problemas de seguridad",
        "No hay registro de participantes",
        "Se confunde difusión con capacitación"
      ],
      "si_que_ofrece": "1. Ofrecer modelos de programa de capacitación\n2. Sugerir contenidos según problemáticas locales\n3. Template de planificación de actividades\n4. Mostrar cómo medir impacto de capacitaciones",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Ofrecer modelos de programa de capacitación\n2. Sugerir contenidos según problemáticas locales\n3. Template de planificación de actividades\n4. Mostrar cómo medir impacto de capacitaciones",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (capacitacion)",
      "facilitador_notas": "Recomendación de capacitaciones para cuerpos de prevención realizamos x la comunidad de seguridad ciudadana y enviamos como VR (rolo)",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P43",
      "numero": 43,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El gobierno local cuenta con un sistema público de videovigilancia urbana?",
      "opciones": "No / Si, monitoreado por la policía / Si, monitoreado por el municipio",
      "doc_respaldatoria": "Licitación",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "videovigilancia",
        "cámaras",
        "CCTV",
        "monitoreo",
        "sistema de cámaras",
        "centro de monitoreo"
      ],
      "no_odm_ids": [
        "OdM_52, OdM_56"
      ],
      "no_odm_textos": [
        "OdM_52: Contar con un sistema público de videovigilancia que articule la respuesta ante emergencias",
        "OdM_56: Coordinar sistemas de videovigilancia con Municipios cercanos para el diseño de corredores seguros"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Francisco",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Ordenanza de creación del sistema",
        "Convenio con policía",
        "Licitación o contrato",
        "Informe técnico con ubicación de cámaras",
        "Captura del centro de monitoreo"
      ],
      "si_que_hace_bueno": [
        "Tiene marco normativo",
        "Cámaras proporcionales al territorio",
        "Centro de monitoreo operativo (no solo graba)",
        "Protocolo de privacidad",
        "Mantenimiento activo",
        "Imágenes se usan para investigaciones"
      ],
      "si_senales_alerta": [
        "Cámaras sin centro de monitoreo activo",
        "50%+ fuera de servicio",
        "Sin protocolo de privacidad",
        "Solo graban, nadie mira",
        "Sin convenio con fiscalía para uso"
      ],
      "si_que_ofrece": "• Dimensionar según territorio y puntos críticos\n• Modelo de ordenanza de videovigilancia\n• Template de protocolo de privacidad\n• Costo estimado\n• Conectar con ciudades con sistemas activos",
      "si_problemas_odm_ids": [
        "OdM_52, OdM_56"
      ],
      "si_problemas_odm_textos": [
        "OdM_52: Contar con un sistema público de videovigilancia que articule la respuesta ante emergencias",
        "OdM_56: Coordinar sistemas de videovigilancia con Municipios cercanos para el diseño de corredores seguros"
      ],
      "si_problemas_que_ofrece": "• Dimensionar según territorio y puntos críticos\n• Modelo de ordenanza de videovigilancia\n• Template de protocolo de privacidad\n• Costo estimado\n• Conectar con ciudades con sistemas activos",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "facilitador_notas": "Ordenanzas (victoria rapida con compilación de ordenanzas \"modelo\" para coronel moldes) Normativas para Central de Monitoreo ",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P44",
      "numero": 44,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿Reciben las imágenes un tratamiento de confidencialidad, resguardo de autenticidad y reserva de la información para su potencial utilización en causas judiciales?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Ordenanza",
      "criterio_minimo_cert": "NO",
      "tags": [
        "confidencialidad imágenes",
        "cadena de custodia",
        "protección datos",
        "privacidad",
        "resguardo material"
      ],
      "no_odm_ids": [
        "OdM_55"
      ],
      "no_odm_textos": [
        "OdM_55: Resguardar la autenticidad, reserva y confidencialidad del material de videovigilancia para poder aportar pruebas en causas judiciales"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Plan de capacitación formal",
        "Cronograma de actividades realizadas",
        "Registro de asistencia",
        "Certificados emitidos",
        "Material de capacitación (presentaciones, folletos)",
        "Publicaciones en redes/web"
      ],
      "si_que_hace_bueno": [
        "Tiene contenido específico de seguridad ciudadana (no genérico)",
        "Se hizo más de una vez en el año",
        "Llega a distintos públicos (no siempre los mismos)",
        "Los participantes evalúan la actividad",
        "Contenido basado en diagnóstico local"
      ],
      "si_senales_alerta": [
        "Una charla aislada cuenta como \"capacitación\"",
        "Siempre van las mismas personas",
        "Contenido genérico sin adaptar al territorio",
        "No hay registro de participantes",
        "Se confunde difusión con capacitación"
      ],
      "si_que_ofrece": "1. Ofrecer modelos de programa de capacitación\n2. Sugerir contenidos según problemáticas locales\n3. Template de planificación de actividades\n4. Mostrar cómo medir impacto de capacitaciones",
      "si_problemas_odm_ids": [
        "OdM_55"
      ],
      "si_problemas_odm_textos": [
        "OdM_55: Resguardar la autenticidad, reserva y confidencialidad del material de videovigilancia para poder aportar pruebas en causas judiciales"
      ],
      "si_problemas_que_ofrece": "1. Ofrecer modelos de programa de capacitación\n2. Sugerir contenidos según problemáticas locales\n3. Template de planificación de actividades\n4. Mostrar cómo medir impacto de capacitaciones",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (capacitacion)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P45",
      "numero": 45,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El gobierno local participa de procesos investigativos aportando pruebas o evidencias en colaboración con las autoridades judiciales?",
      "opciones": "No / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "norma regulatoria cámaras/ misión y funciones del área",
      "criterio_minimo_cert": "NO",
      "tags": [
        "aporte investigativo",
        "judicialización",
        "peritaje",
        "fiscalía",
        "prueba judicial",
        "imágenes judiciales"
      ],
      "no_odm_ids": [
        "OdM_55"
      ],
      "no_odm_textos": [
        "OdM_55: Resguardar la autenticidad, reserva y confidencialidad del material de videovigilancia para poder aportar pruebas en causas judiciales"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [
        "OdM_55"
      ],
      "si_problemas_odm_textos": [
        "OdM_55: Resguardar la autenticidad, reserva y confidencialidad del material de videovigilancia para poder aportar pruebas en causas judiciales"
      ],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P46",
      "numero": 46,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿Cuenta con protocolos o procedimientos de trabajo documentados sobre la actuación de cuerpo civil de prevención, monitoreo de cámaras, entre otros?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Documentos/manual de procedimientos",
      "criterio_minimo_cert": "NO",
      "tags": [
        "protocolos videovigilancia",
        "procedimientos",
        "manual operativo",
        "SOP",
        "protocolo de monitoreo"
      ],
      "no_odm_ids": [
        "OdM_48"
      ],
      "no_odm_textos": [
        "OdM_48: Promover protocolos de actuación documentados para cada una de las áreas de seguridad ciudadana"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Tablero en Excel/Google Sheets",
        "Dashboard digital (PowerBI, Data Studio)",
        "Reportes periódicos en Word/PDF",
        "Minutas de reuniones con datos",
        "Planilla simple de avance de metas"
      ],
      "si_que_hace_bueno": [
        "Se actualiza periódicamente (mensual/trimestral)",
        "Tiene indicadores cuantitativos",
        "Se usa para tomar decisiones",
        "Lo ven más personas que el responsable"
      ],
      "si_senales_alerta": [
        "Tablero armado una vez y nunca actualizado",
        "Solo datos cualitativos (\"avanzamos bien\")",
        "Nadie lo mira ni lo usa para decidir",
        "Indicadores imposibles de medir"
      ],
      "si_que_ofrece": "1. Ofrecer template de tablero adaptado al plan de seguridad ciudadana\n2. Sugerir indicadores medibles y relevantes\n3. Mostrar dashboards de otras ciudades\n4. Ayudar a definir frecuencia y responsable de actualización",
      "si_problemas_odm_ids": [
        "OdM_48"
      ],
      "si_problemas_odm_textos": [
        "OdM_48: Promover protocolos de actuación documentados para cada una de las áreas de seguridad ciudadana"
      ],
      "si_problemas_que_ofrece": "1. Ofrecer template de tablero adaptado al plan de seguridad ciudadana\n2. Sugerir indicadores medibles y relevantes\n3. Mostrar dashboards de otras ciudades\n4. Ayudar a definir frecuencia y responsable de actualización",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (monitoreo)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P47",
      "numero": 47,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿Hay coordinación operativa entre el sistema de seguridad local y los servicios de las empresas de seguridad privada?",
      "opciones": "No / Si",
      "doc_respaldatoria": "minuta o temario reunión, notas en medios, etc",
      "criterio_minimo_cert": "NO",
      "tags": [
        "coordinación vigilancia",
        "patrullaje integrado",
        "guardia + policía",
        "operativo conjunto"
      ],
      "no_odm_ids": [
        "OdM_50"
      ],
      "no_odm_textos": [
        "OdM_50: Coordinar las tareas de vigilancia y patrullaje de la guardia local con la policía y con los servicios privados de seguridad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [
        "OdM_50"
      ],
      "si_problemas_odm_textos": [
        "OdM_50: Coordinar las tareas de vigilancia y patrullaje de la guardia local con la policía y con los servicios privados de seguridad"
      ],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P48",
      "numero": 48,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El sistema de videovigilancia pública integra las cámaras del sector privado?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Captura de registro/normativa;",
      "criterio_minimo_cert": "NO",
      "tags": [
        "cámaras privadas",
        "registro cámaras",
        "integración cámaras",
        "alerta comunitaria",
        "botón de pánico"
      ],
      "no_odm_ids": [
        "OdM_53"
      ],
      "no_odm_textos": [
        "OdM_53: Armar registro voluntario de cámaras de seguridad privadas que se integre a la estrategia de la central de monitoreo"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [
        "OdM_53"
      ],
      "si_problemas_odm_textos": [
        "OdM_53: Armar registro voluntario de cámaras de seguridad privadas que se integre a la estrategia de la central de monitoreo"
      ],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P49",
      "numero": 49,
      "dimension": "Prevención situacional y gestión del espacio público",
      "pregunta": "¿El gobierno local cuenta con  tecnologías aplicadas a la seguridad? (reconocimiento facial, lectura de matrículas,  Inteligencia Artificial, otras.)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Sistemas tecnológicos ()",
      "criterio_minimo_cert": "NO",
      "tags": [
        "inteligencia artificial",
        "IA",
        "analítica video",
        "reconocimiento",
        "detección automática",
        "tecnología aplicada"
      ],
      "no_odm_ids": [
        "OdM_54"
      ],
      "no_odm_textos": [
        "OdM_54: Incorporar la Inteligencia Artificial al sistema de cámaras"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Captura del sistema en funcionamiento",
        "Informe técnico del sistema",
        "Contrato o licitación de adquisición",
        "Manual de usuario",
        "Convenio de uso si es de terceros"
      ],
      "si_que_hace_bueno": [
        "Tiene marco normativo de respaldo",
        "Sistema proporcional al territorio/necesidad",
        "Operativo y en uso regular (no solo instalado)",
        "Protocolo de uso documentado",
        "Mantenimiento activo"
      ],
      "si_senales_alerta": [
        "Sistema instalado pero no utilizado",
        "Porcentaje significativo fuera de servicio",
        "Sin protocolo de uso o privacidad",
        "Sin mantenimiento ni actualización",
        "No se integra con otros sistemas"
      ],
      "si_que_ofrece": "1. Dimensionar según territorio y necesidades\n2. Modelo de normativa de respaldo\n3. Template de protocolo de uso\n4. Comparativa de soluciones de otras ciudades",
      "si_problemas_odm_ids": [
        "OdM_54"
      ],
      "si_problemas_odm_textos": [
        "OdM_54: Incorporar la Inteligencia Artificial al sistema de cámaras"
      ],
      "si_problemas_que_ofrece": "1. Dimensionar según territorio y necesidades\n2. Modelo de normativa de respaldo\n3. Template de protocolo de uso\n4. Comparativa de soluciones de otras ciudades",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (sistema_tecnologico)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P50",
      "numero": 50,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿El gobierno local implementa proyectos de prevención social del delito? (acciones focalizadas sobre grupos considerados en riesgo para su integración social)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Listado de proyectos",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "prevención social",
        "factores sociales",
        "inclusión",
        "programas sociales",
        "abordaje territorial",
        "vulnerabilidad"
      ],
      "no_odm_ids": [
        "OdM_57, OdM_58"
      ],
      "no_odm_textos": [
        "OdM_57: Poseer iniciativas y coordinar con áreas que atienden factores sociales de la inseguridad",
        "OdM_58: Trabajar en los factores sociales en función de diagnósticos, basados en evidencia"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del programa/proyecto con objetivos directamente vinculados a la prevención de delitos y/o violencias",
        "Convenio con instituciones educativas/salud/culturales/deportivas/tercer sector",
        "Decreto de creación y/o reglamentación del programa/ proyecto de prevención social del delito y/o violencias",
        "Informe de actividades",
        "Registro de beneficiarios",
        "Publicación en página web oficial del municipio del programa/proyecto",
        "Portal de gestión virtual del programa/proyecto"
      ],
      "si_que_hace_bueno": [
        "Focalizado en población de riesgo identificada",
        "Tiene continuidad (no evento aislado)",
        "Involucra más de un área municipal y/o provincial",
        "Tiene indicador de resultado",
        "Basado en diagnóstico de factores de riesgo",
        "Tiene indicadores de impacto vinculados a la reducción del delito/violencias",
        "Cuenta con recursos para su sostenibilidad en el tiempo: recursos humanos capacitados e idóneos y partida presupuestaria asignada y actualizada"
      ],
      "si_senales_alerta": [
        "Un taller aislado que se hizo una vez",
        "No focalizado: para 'toda la ciudadanía'",
        "Sin indicadores",
        "Lo lleva una sola persona",
        "Confunde prevención con comunicación",
        "No cuenta con los recursos humanos suficientes",
        "Los recursos humanos no están capacitados en la temática",
        "Los recursos presupuestarios no sonsificientes ni se actualizan anualmente"
      ],
      "si_que_ofrece": "• Modelos de programas de otras ciudades\n• Ayudar a diseñar programa focalizado\n• Sugerir indicadores de proceso y resultado\n• Guía de articulación interáreas\n",
      "si_problemas_odm_ids": [
        "OdM_57, OdM_58"
      ],
      "si_problemas_odm_textos": [
        "OdM_57: Poseer iniciativas y coordinar con áreas que atienden factores sociales de la inseguridad",
        "OdM_58: Trabajar en los factores sociales en función de diagnósticos, basados en evidencia"
      ],
      "si_problemas_que_ofrece": "• Modelos de programas de otras ciudades\n• Ayudar a diseñar programa focalizado\n• Sugerir indicadores de proceso y resultado\n• Guía de articulación interáreas",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "Idealmente es preferible el documento del proyecto/programa antes que una publicación web, pero hay municipios que publican estos en sus páginas web oficiales; este formato es válido también. Lo excluyente es que sean programas con ibjetivos de prevención del delito y violencias, no políticas sociales universales. ",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P51",
      "numero": 51,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿El gobierno local cuenta con una instancia municipal que preste servicios de acceso a la justicia y mediación comunitaria?\n(asesoramiento en temas vinculados con el acceso a derechos, resolución de conflictos vecinales, etc)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Documento del programa o creción del área",
      "criterio_minimo_cert": "NO",
      "tags": [
        "mediación",
        "resolución de conflictos",
        "justicia restaurativa",
        "mediación vecinal",
        "acceso a justicia"
      ],
      "no_odm_ids": [
        "OdM_60, OdM_61"
      ],
      "no_odm_textos": [
        "OdM_60: Poseer o articular iniciativas para la resolución pacífica de conflictos vecinales",
        "OdM_61: Poseer o articular iniciativas de promoción de acceso a la justicia"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del programa con objetivos",
        "Convenio con instituciones",
        "Decreto/ordenanza de creación",
        "Informe de actividades"
      ],
      "si_que_hace_bueno": [
        "Focalizado en población o problemática identificada",
        "Tiene continuidad (no evento aislado)",
        "Involucra más de un área municipal",
        "Tiene indicadores de resultado",
        "Basado en diagnóstico previo",
        "Cuenta con un protocolo de articulación entre áreas"
      ],
      "si_senales_alerta": [
        "Sin indicadores de resultado",
        "Lo lleva una sola persona",
        "Confunde actividad con programa",
        "Depende de otro nivel de gobierno y el municipio no tiene ingerencia"
      ],
      "si_que_ofrece": "1. Modelos de programas/ áreas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Guía de articulación interáreas",
      "si_problemas_odm_ids": [
        "OdM_60, OdM_61"
      ],
      "si_problemas_odm_textos": [
        "OdM_60: Poseer o articular iniciativas para la resolución pacífica de conflictos vecinales",
        "OdM_61: Poseer o articular iniciativas de promoción de acceso a la justicia"
      ],
      "si_problemas_que_ofrece": "1. Modelos de programas/ áreas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Guía de articulación interáreas",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente": "Completado por equipo",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P52",
      "numero": 52,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿El gobierno local  cuenta con dispositivos de asistencia a la víctima?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Documento del programa o creción del área",
      "criterio_minimo_cert": "NO",
      "tags": [
        "asistencia a víctimas",
        "contención",
        "acompañamiento",
        "violencia de género",
        "centro de asistencia"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del programa con objetivos",
        "Convenio con instituciones",
        "Decreto de creación",
        "Informe de actividades",
        "Registro de beneficiarios/participantes"
      ],
      "si_que_hace_bueno": [
        "Focalizado en población o problemática identificada",
        "Tiene continuidad (no evento aislado)",
        "Involucra más de un área municipal",
        "Tiene indicadores de resultado",
        "Basado en diagnóstico previo"
      ],
      "si_senales_alerta": [
        "Un taller aislado que se hizo una vez",
        "No focalizado: para \"toda la ciudadanía\" sin criterio",
        "Sin indicadores de resultado",
        "Lo lleva una sola persona",
        "Confunde actividad con programa"
      ],
      "si_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (programas_proyectos)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P53",
      "numero": 53,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿El gobierno local tiene un programa y equipo en el territorio dedicado al trabajo de prevención de consumos problemáticos?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Documento del programa o creción del área",
      "criterio_minimo_cert": "NO",
      "tags": [
        "prevención comunitaria",
        "trabajo territorial",
        "operador barrial",
        "promotor comunitario",
        "consumos problemáticos"
      ],
      "no_odm_ids": [
        "OdM_58, OdM_59, OdM_62"
      ],
      "no_odm_textos": [
        "OdM_58: Trabajar en los factores sociales en función de diagnósticos, basados en evidencia",
        "OdM_59: Planificar programas de prevención social del delito con otros actores del ecosistema de seguridad y promoviendo la participación comunitaria",
        "OdM_62: Poseer o articular iniciativas de prevención de consumos problemáticos"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Francisco",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del proyecto",
        ".Página web que muestra el proyecto/acción/servicio"
      ],
      "si_que_hace_bueno": [
        "El área tiene funciones específicas de seguridad ciudadana (no es genérica)",
        "Tiene personal asignado (al menos un responsable formal)",
        "Está vigente (no es un organigrama de otra gestión)",
        "El rango permite interlocución con otras áreas relevantes"
      ],
      "si_senales_alerta": [
        "Área que solo existe en el papel sin equipo real",
        "Es apéndice de otra área sin funciones propias",
        "Organigrama de gestión anterior no actualizado",
        "Responsable con múltiples áreas a cargo sin dedicación"
      ],
      "si_que_ofrece": "1. Mostrar organigramas de ciudades que certificaron en seguridad ciudadana\n2. Sugerir modelo de decreto de creación del área\n3. Explicar ventajas de diferentes rangos jerárquicos\n4. Conectar con ciudades similares que avanzaron",
      "si_problemas_odm_ids": [
        "OdM_58, OdM_59, OdM_62"
      ],
      "si_problemas_odm_textos": [
        "OdM_58: Trabajar en los factores sociales en función de diagnósticos, basados en evidencia",
        "OdM_59: Planificar programas de prevención social del delito con otros actores del ecosistema de seguridad y promoviendo la participación comunitaria",
        "OdM_62: Poseer o articular iniciativas de prevención de consumos problemáticos"
      ],
      "si_problemas_que_ofrece": "1. Mostrar organigramas de ciudades que certificaron en seguridad ciudadana\n2. Sugerir modelo de decreto de creación del área\n3. Explicar ventajas de diferentes rangos jerárquicos\n4. Conectar con ciudades similares que avanzaron",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (area_dedicada)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P54",
      "numero": 54,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿Cuenta el gobierno local con alguna política de prevención de la reincidencia y la reintegración social post penitenciaria?",
      "opciones": "No / Si",
      "doc_respaldatoria": "Documento del programa o creción del área",
      "criterio_minimo_cert": "NO",
      "tags": [
        "reintegración social",
        "post-penitenciario",
        "reinserción",
        "personas liberadas",
        "inclusión laboral"
      ],
      "no_odm_ids": [
        "OdM_63"
      ],
      "no_odm_textos": [
        "OdM_63: Poseer o articular iniciativas de reintegración social post-penitenciaria"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento formal (decreto, resolución, ordenanza)",
        "Informe o reporte escrito",
        "Captura de pantalla o registro digital",
        "Planilla o base de datos",
        "Registro fotográfico o audiovisual"
      ],
      "si_que_hace_bueno": [
        "Tiene respaldo documental verificable",
        "Es consistente con el plan de seguridad ciudadana",
        "Involucra a los actores relevantes",
        "Se puede medir o verificar su implementación"
      ],
      "si_senales_alerta": [
        "Sin respaldo documental",
        "Desconectado de la estrategia general",
        "Implementación parcial o inconsistente",
        "No se puede verificar objetivamente"
      ],
      "si_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "si_problemas_odm_ids": [
        "OdM_63"
      ],
      "si_problemas_odm_textos": [
        "OdM_63: Poseer o articular iniciativas de reintegración social post-penitenciaria"
      ],
      "si_problemas_que_ofrece": "1. Mostrar cómo lo resolvieron otras ciudades del programa\n2. Ofrecer templates o modelos adaptables\n3. Sugerir pasos concretos para avanzar\n4. Conectar con experiencias y recursos de la Red",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Generado por IA",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P55",
      "numero": 55,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿El gobierno local  tiene un plan de comunicación a la ciudadanía en materia de seguridad que incluya la difusión de iniciativas del gobierno?",
      "opciones": "No / Forma parte de la estrategia de comunicación general / Tenemos una estrategia específica de comunicación sobre el tema",
      "doc_respaldatoria": "Plan de comunicación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "plan de comunicación",
        "estrategia comunicacional",
        "prensa seguridad",
        "difusión",
        "medios"
      ],
      "no_odm_ids": [
        "OdM_65, OdM_67, OdM_68"
      ],
      "no_odm_textos": [
        "OdM_65: Establecer la comunicación en seguridad como una estrategia general de comunicación del gobierno",
        "OdM_67: Desarrollar una estrategia de comunicación específica del plan local de seguridad y los programas de prevención",
        "OdM_68: Llevar adelante un plan de comunicación en materia de seguridad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del programa con objetivos",
        "Convenio con instituciones",
        "Decreto de creación",
        "Informe de actividades",
        "Registro de beneficiarios/participantes"
      ],
      "si_que_hace_bueno": [
        "Focalizado en población o problemática identificada",
        "Tiene continuidad (no evento aislado)",
        "Involucra más de un área municipal",
        "Tiene indicadores de resultado",
        "Basado en diagnóstico previo"
      ],
      "si_senales_alerta": [
        "Un taller aislado que se hizo una vez",
        "No focalizado: para \"toda la ciudadanía\" sin criterio",
        "Sin indicadores de resultado",
        "Lo lleva una sola persona",
        "Confunde actividad con programa"
      ],
      "si_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "si_problemas_odm_ids": [
        "OdM_65, OdM_67, OdM_68"
      ],
      "si_problemas_odm_textos": [
        "OdM_65: Establecer la comunicación en seguridad como una estrategia general de comunicación del gobierno",
        "OdM_67: Desarrollar una estrategia de comunicación específica del plan local de seguridad y los programas de prevención",
        "OdM_68: Llevar adelante un plan de comunicación en materia de seguridad"
      ],
      "si_problemas_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (programas_proyectos)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P56",
      "numero": 56,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿La ciudadanía tiene algún medio de comunicación distinto al 911 para informar al municipio una emergencia, evento delictivo o denuncia? (aplicación telefónica, WhatsApp, buzones ciudadanos, herramienta de gobierno electrónico u otra)",
      "opciones": "No / Si",
      "doc_respaldatoria": "Sistema o canal utilizado",
      "criterio_minimo_cert": "SÍ",
      "tags": [
        "canal de denuncia",
        "línea municipal",
        "app seguridad",
        "WhatsApp seguridad",
        "comunicación ciudadana",
        "alerta vecinal"
      ],
      "no_odm_ids": [
        "OdM_66"
      ],
      "no_odm_textos": [
        "OdM_66: Establecer un medio de comunicación distinto al 911 para informar al municipio una emergencia, evento delictivo o denuncia"
      ],
      "no_que_ofrece": "• Benchmark: mostrar cómo lo resolvieron San Francisco (Ojos en Alerta)",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "App municipal con botón de alerta",
        "Línea telefónica municipal",
        "Grupo de WhatsApp con protocolo",
        "Sistema tipo 'Ojos en alerta'",
        "Plataforma web de reporte",
        "Chatbot de atención"
      ],
      "si_que_hace_bueno": [
        "Funciona 24/7",
        "Tiene protocolo de respuesta con tiempos",
        "La ciudadanía lo conoce y lo usa",
        "Se registran reportes para análisis",
        "Hay alguien que responde"
      ],
      "si_senales_alerta": [
        "App que nadie descargó",
        "Solo atiende lunes a viernes 8-14",
        "Sin protocolo de respuesta",
        "No se sistematizan reportes",
        "Se confunde con canal de reclamos general"
      ],
      "si_que_ofrece": "• Evaluar opciones según presupuesto\n• Comparativa de sistemas de otras ciudades\n• Ayudar a diseñar protocolo de respuesta\n• Sugerir cómo difundir el canal\n• Métricas de efectividad",
      "si_problemas_odm_ids": [
        "OdM_66"
      ],
      "si_problemas_odm_textos": [
        "OdM_66: Establecer un medio de comunicación distinto al 911 para informar al municipio una emergencia, evento delictivo o denuncia"
      ],
      "si_problemas_que_ofrece": "• Evaluar opciones según presupuesto\n• Comparativa de sistemas de otras ciudades\n• Ayudar a diseñar protocolo de respuesta\n• Sugerir cómo difundir el canal\n• Métricas de efectividad",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "Si la ciudad es chica y bajo delito, no tiene que  brindar el servicio 24/7 como buena acción. (No le va a dar la esctructura y será en vano) Analizar alternativas de atención para proponer como buena práctica.",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P57",
      "numero": 57,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿Se impulsan iniciativas de sensibilización y difusión para fomentar una cultura de convivencia ciudadana?",
      "opciones": "No / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "Plan de sensibilización o difusión",
      "criterio_minimo_cert": "NO",
      "tags": [
        "sensibilización",
        "campañas",
        "difusión",
        "cultura de prevención",
        "concientización",
        "educación ciudadana"
      ],
      "no_odm_ids": [
        "OdM_64, OdM_67, OdM_68"
      ],
      "no_odm_textos": [
        "OdM_64: Desarrollar campañas de sensibilización y difusión para fomentar una cultura de convivencia ciudadana",
        "OdM_67: Desarrollar una estrategia de comunicación específica del plan local de seguridad y los programas de prevención",
        "OdM_68: Llevar adelante un plan de comunicación en materia de seguridad"
      ],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Documento del programa con objetivos",
        "Convenio con instituciones",
        "Decreto de creación",
        "Informe de actividades",
        "Registro de beneficiarios/participantes"
      ],
      "si_que_hace_bueno": [
        "Focalizado en población o problemática identificada",
        "Tiene continuidad (no evento aislado)",
        "Involucra más de un área municipal",
        "Tiene indicadores de resultado",
        "Basado en diagnóstico previo"
      ],
      "si_senales_alerta": [
        "Un taller aislado que se hizo una vez",
        "No focalizado: para \"toda la ciudadanía\" sin criterio",
        "Sin indicadores de resultado",
        "Lo lleva una sola persona",
        "Confunde actividad con programa"
      ],
      "si_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "si_problemas_odm_ids": [
        "OdM_64, OdM_67, OdM_68"
      ],
      "si_problemas_odm_textos": [
        "OdM_64: Desarrollar campañas de sensibilización y difusión para fomentar una cultura de convivencia ciudadana",
        "OdM_67: Desarrollar una estrategia de comunicación específica del plan local de seguridad y los programas de prevención",
        "OdM_68: Llevar adelante un plan de comunicación en materia de seguridad"
      ],
      "si_problemas_que_ofrece": "1. Modelos de programas de otras ciudades\n2. Ayudar a diseñar programa focalizado\n3. Sugerir indicadores de proceso y resultado\n4. Guía de articulación interáreas",
      "fuentes_rag": "",
      "estado": "Provisional",
      "notas": "",
      "fuente": "Similitud Seguridad (programas_proyectos)",
      "facilitador_notas": "",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    },
    {
      "id": "P58",
      "numero": 58,
      "dimension": "Gestión socio-comunitaria y comunicación",
      "pregunta": "¿El gobierno local capacita a la ciudadanía en materia de seguridad ciudadana? (incluyendo recomendaciones sobre medidas de autoprotección como horarios de retiro de basura, desmalezado y otras medidas de cuidado)",
      "opciones": "No / Ocasionalmente / Frecuentemente",
      "doc_respaldatoria": "Plan de capacitación",
      "criterio_minimo_cert": "NO",
      "tags": [
        "capacitación ciudadana",
        "talleres seguridad",
        "formación ciudadana",
        "educación comunitaria",
        "prevención desde la comunidad"
      ],
      "no_odm_ids": [],
      "no_odm_textos": [],
      "no_que_ofrece": "",
      "si_datos_pide": "",
      "si_formatos_validos": [
        "Plan de capacitación formal",
        "Cronograma de actividades realizadas",
        "Registro de asistencia",
        "Certificados emitidos",
        "Material de capacitación (presentaciones, folletos)",
        "Publicaciones en redes/web con fotos"
      ],
      "si_que_hace_bueno": [
        "Tiene contenido específico de seguridad (no genérico)",
        "Se hizo más de una vez en el año",
        "Llega a distintos públicos (no siempre los mismos)",
        "Los participantes evalúan la actividad",
        "El contenido se actualiza"
      ],
      "si_senales_alerta": [
        "Una charla aislada cuenta como 'capacitación'",
        "Siempre van las mismas 20 personas",
        "Contenido genérico sin adaptar al territorio",
        "No hay registro de quién participó",
        "Se confunde difusión/campaña con capacitación"
      ],
      "si_que_ofrece": "• Ofrecer modelos de programa de capacitación\n• Sugerir contenidos según problemáticas locales\n• Template de planificación de actividades\n• Mostrar cómo medir impacto de capacitaciones\n• Distinguir capacitación vs. campaña de comunicación",
      "si_problemas_odm_ids": [],
      "si_problemas_odm_textos": [],
      "si_problemas_que_ofrece": "• Ofrecer modelos de programa de capacitación\n• Sugerir contenidos según problemáticas locales\n• Template de planificación de actividades\n• Mostrar cómo medir impacto de capacitaciones\n• Distinguir capacitación vs. campaña de comunicación",
      "fuentes_rag": "",
      "estado": "Completo",
      "notas": "",
      "fuente_original": "Completado por equipo",
      "fecha_actualizacion": "2026-03-11T14:10:13.600220"
    }
  ]
}`
	lookupThreeResponse := LookupTreeResponse{
		Content: jsonContent,
	}
	return lookupThreeResponse, nil
}
