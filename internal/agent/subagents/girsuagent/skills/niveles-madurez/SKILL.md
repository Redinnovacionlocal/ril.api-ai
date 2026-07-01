---
name: niveles-madurez
description: Identifica el nivel de madurez del municipio (Bajo / Intermedio / Avanzado) usando las descripciones del árbol, luego usa la OdM del nivel detectado para recomendar y las acciones del agente para decidir cómo acompañar. El tamaño del municipio acota qué niveles vale la pena diagnosticar.
---

# Skill: Niveles de Madurez del Municipio

## 🎯 Objetivo y Activación
Esta skill define cómo trabajar con preguntas del árbol que tienen niveles de madurez. El objetivo es identificar en qué punto del camino está el municipio, recomendar el paso siguiente, y acompañar concretamente la implementación.

**Cuándo se activa:** Cuando vayas a trabajar un tema del Autodiagnóstico (AD) con un municipio. La skill cubre el ciclo completo: identificar nivel (diagnóstico), recomendar OdM, y acompañar con acciones del agente.

> **Nota técnica:** La skill espera que la pregunta del árbol tenga niveles (descripción + OdM + acciones del agente por cada uno de los 3 niveles). Si por algún motivo no los tiene, trabajá con el contenido tradicional de la pregunta.

## 🧠 Conceptos Clave: La tríada por nivel
Cada nivel de madurez (Bajo, Intermedio, Avanzado) tiene tres bloques de información, cada uno con una función distinta:

1. **Descripción (Diagnosticar):** Cómo se ve un municipio en ese nivel. El agente la lee y formula preguntas para identificar si el municipio está ahí.
2. **OdM / Oportunidad de Mejora (Recomendar):** Lo que el municipio puede hacer para avanzar desde ese nivel. Es lo que el agente le propone explícitamente al usuario.
3. **Acciones del Agente (Acompañar):** Cómo el agente apoya la implementación de la OdM (templates, cálculos, modelos de otras ciudades, indicadores, conexiones a la Red). Es lo que el agente ofrece "extra" más allá de la recomendación.

## 📄 Estructura del JSON esperado
```json
{
  "id": "P33",
  "pregunta": "...",
  "niveles": {
    "bajo": {
      "descripcion": "Cómo se ve un municipio en nivel Bajo",
      "odm": "Lo que el municipio puede hacer para avanzar",
      "acciones_agente": ["Acción 1", "Acción 2", "Acción 3"]
    },
    "intermedio": { "descripcion": "...", "odm": "...", "acciones_agente": [] },
    "avanzado":   { "descripcion": "...", "odm": "...", "acciones_agente": [] }
  },
  "nota_facilitador": "Advertencias contextuales (opcional)",
  "tags_rag": ["...", "..."]
}