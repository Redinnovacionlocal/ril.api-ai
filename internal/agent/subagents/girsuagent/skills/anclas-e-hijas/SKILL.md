---
name: anclas-e-hijas
description: Maneja la lógica de preguntas ancla con sub-preguntas hijas. Una ancla es una pregunta que abre una capacidad amplia; sus hijas profundizan en sub-aspectos. Si la ancla está en nivel Bajo (el municipio no tiene la capacidad), las hijas heredan Bajo automáticamente. Usar cuando lookup_tree_questions devuelve una pregunta con es_ancla=true o con el campo padre poblado.
---

# Skill: Preguntas ancla y sus hijas

## 🎯 Objetivo y Activación
Esta skill describe cómo manejar preguntas del árbol que tienen relaciones jerárquicas (anclas con hijas). El objetivo es que el agente profundice donde corresponde sin convertir la conversación en un interrogatorio.

**Cuándo se activa:**
* Cuando `lookup_tree_questions` devuelve una pregunta con `es_ancla: true`.
* Cuando devuelve una pregunta con `padre` poblado (es decir, es hija de una ancla).
* *Si la pregunta tiene `es_ancla: false` y `padre` vacío, ignorar esta skill (es una pregunta plana sin relaciones).*

## 🧠 Conceptos Clave
* **Ancla:** Pregunta que abre una capacidad amplia. Si la respuesta a la ancla es "no tienen X", no tiene sentido profundizar en cómo es ese X.
* **Hija:** Pregunta que profundiza un aspecto de la ancla. Solo es viable si la ancla recibió respuesta afirmativa.
* **Regla de herencia:** Si la ancla está en nivel Bajo (el municipio no tiene la capacidad), las hijas heredan Bajo automáticamente sin necesidad de preguntarlas. La `nota_facilitador` en el árbol indica esto explícitamente
* **Anidamiento:** Una hija puede ser a su vez ancla de otras hijas. Ejemplo: 38 (videovigilancia) es ancla con hijas 39-42; 40 (protocolo de monitoreo) es hija de 38 y a su vez tiene su propio rol estructural.
* **Exploración a demanda:** Las hijas no se recorren todas ni en orden. Son territorio disponible para profundizar, no una checklist obligatoria.

### ⚓ Las 6 anclas confirmadas en seguridad

| Ancla | Tema | Hijas |
| :--- | :--- | :--- |
| **P1** | Área de seguridad | 2, 3, 4 |
| **P10** | Plan de seguridad | 11, 12, 13 |
| **P23** | Canal de alertas ciudadanas | 24 |
| **P33** | Cuerpo civil de prevención | 34, 35, 36, 37 |
| **P38** | Sistema de videovigilancia | 39, 40, 41, 42 |
| **P40** | Protocolo de central de monitoreo | *(también es hija de 38)* |

## 📄 Estructura del JSON Relevante

**Si la pregunta es ancla:**
```json
{
  "id": "33",
  "pregunta": "...",
  "es_ancla": true,
  "padre": null,
  "hijas": ["34", "35", "36", "37"],
  "niveles": { "...": "..." },
  "nota_facilitador": "Ancla: sin cuerpo propio, 34-37 heredan nivel Bajo."
}