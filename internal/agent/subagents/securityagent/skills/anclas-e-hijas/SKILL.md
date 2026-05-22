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
* **Regla de herencia:** Si la ancla está en nivel Bajo (el municipio no tiene la capacidad), las hijas heredan Bajo automáticamente sin necesidad de preguntarlas. La `nota_facilitador` en el árbol indica esto explícitamente: *"Ancla: sin X, P##-P## heredan nivel Bajo"*.
* **Anidamiento:** Una hija puede ser a su vez ancla de otras hijas. Ejemplo: P38 (videovigilancia) es ancla con hijas P39-P42; P40 (protocolo de monitoreo) es hija de P38 y a su vez tiene su propio rol estructural.
* **Exploración a demanda:** Las hijas no se recorren todas ni en orden. Son territorio disponible para profundizar, no una checklist obligatoria.

### ⚓ Las 6 anclas confirmadas en seguridad

| Ancla | Tema | Hijas |
| :--- | :--- | :--- |
| **P1** | Área de seguridad | P2, P3, P4 |
| **P10** | Plan de seguridad | P11, P12, P13 |
| **P23** | Canal de alertas ciudadanas | P24 |
| **P33** | Cuerpo civil de prevención | P34, P35, P36, P37 |
| **P38** | Sistema de videovigilancia | P39, P40, P41, P42 |
| **P40** | Protocolo de central de monitoreo | *(también es hija de P38)* |

## 📄 Estructura del JSON Relevante

**Si la pregunta es ancla:**
```json
{
  "id": "P33",
  "pregunta": "...",
  "es_ancla": true,
  "padre": null,
  "hijas": ["P34", "P35", "P36", "P37"],
  "niveles": { "...": "..." },
  "nota_facilitador": "Ancla: sin cuerpo propio, P34-P37 heredan nivel Bajo."
}