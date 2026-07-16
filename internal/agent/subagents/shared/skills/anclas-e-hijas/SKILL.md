---
name: anclas-e-hijas
description: Maneja la lógica de preguntas ancla con sub-preguntas hijas para profundizar el diagnóstico sin saturar al usuario. Se activa cuando lookup_tree_questions devuelve campos jerárquicos.
---

# Skill: Gestión de Preguntas Ancla y Sub-preguntas Hijas

## 🎯 Objetivo y Activación
Esta skill determina cómo navegar por el árbol de criterios de calidad cuando existen relaciones jerárquicas entre preguntas. El objetivo es estructurar la indagación de forma lógica, experta y empática, evitando interrogar al municipio sobre detalles de capacidades que no posee.

**Cuándo se activa:**
* El resultado de `lookup_tree_questions` contiene una pregunta con `"es_ancla": true`.
* El resultado contiene una pregunta con el campo `"padre"` poblado con el ID de otra pregunta.

*(Si la pregunta tiene `"es_ancla": false` y `"padre": null`, ignorar esta skill; se trata de una pregunta plana).*

## 🧠 Conceptos y Reglas Estrictas

1. **La Regla del Ancla:** Una pregunta "Ancla" abre una capacidad o estructura macro en el gobierno local. Si el municipio responde negativamente (Nivel Bajo / No posee la estructura), queda estrictamente **prohibido** indagar sobre las preguntas "Hijas".
2. **Heredar Nivel Bajo:** Si el municipio no cuenta con la capacidad del Ancla, todas sus preguntas Hijas se consideran automáticamente en **Nivel Bajo**. No debés preguntarlas. Registralas directamente si es necesario o asumilas como desafíos a futuro.
3. **Exploración a Demanda (No es una Checklist):** Las preguntas Hijas son territorio disponible para profundizar SOLO si la respuesta al Ancla fue afirmativa o parcial. No estás obligado a recorrerlas todas ni a seguir un orden lineal rígido.
4. **Anidamiento Jerárquico:** Una pregunta Hija puede ser a su vez Ancla de sus propias sub-capacidad hijas. Evaluá siempre la respuesta del nodo inmediato superior antes de bajar un nivel en la jerarquía.

## 📖 Ejemplo de Flujo de Razonamiento Experto (Genérico)

* **Caso A (Ancla Negativa):** * *Consulta:* El agente consulta por un "Plan Estratégico del área" (Pregunta Ancla). El usuario dice: "No tenemos un plan formal todavía".
  * *Acción:* El agente frena ahí. Sabe que las preguntas Hijas (ej: "Metas del plan", "Presupuesto asignado al plan", "Frecuencia de revisión del plan") son automáticamente "Bajo". No las pregunta y enfoca su recomendación en la creación del plan base.

* **Caso B (Ancla Positiva/Parcial):**
  * *Consulta:* El usuario dice: "Sí, tenemos un Plan Estratégico".
  * *Acción:* El agente avanza de forma fluida a explorar las Hijas (ej: "¡Buenísimo que tengan un plan! ¿Cuáles son las metas principales que se trazaron para este año?").

## 📄 Estructura del JSON Jerárquico de Referencia

Cuando uses `lookup_tree_questions`, vas a recibir estructuras mapeadas de esta forma. Usá estos campos para guiar tu lógica:

```json
{
  "id": "ID_DE_PREGUNTA",
  "pregunta": "Texto de la evaluación...",
  "es_ancla": true,
  "padre": "ID_PADRE_SI_APLICA",
  "hijas": ["ID_HIJA_1", "ID_HIJA_2"],
  "nota_facilitador": "Instrucción experta de RIL (ej: 'Ancla: sin esto, hijas heredan Bajo')"
}