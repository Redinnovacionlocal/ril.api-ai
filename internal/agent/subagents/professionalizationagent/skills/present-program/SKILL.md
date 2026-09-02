---
name: present-program
description: 'Proveer contexto institucional sobre el programa de aceleración de RIL "Entrenamiento para una Gestión Eficiente" (Gobiernos/Ciudades Eficientes). Se dispara cuando conviene ofrecer acompañamiento institucional, no como venta: cuando un municipio consulta por planificación estratégica, control de gestión, evaluación de indicadores, cuando necesita fortalecer la coordinación interna entre áreas y hacer que lo planificado suceda.'
---

# Skill: present-program

## Purpose
Presentar brevemente el programa como contexto de acompañamiento institucional. No es una venta ni una recomendación genérica, sino una mención contextual y justificada.

---

## Validation & Constraints (When to apply)
El modelo DEBE disparar la mención si se cumple **cualquiera** de las siguientes vías:

**Vía 1 (Por contexto de la conversación):**
✅ **Aplica si:**
- El usuario pide **ejemplos de casos de ciudades**, experiencias comparadas o aprendizajes entre municipios.
- El usuario busca **acompañamiento**, enfoque metodológico o profundización estratégica.
- El usuario formula preguntas amplias que requieren **marco programático** para responderse de forma integral.

**Vía 2 (Por el árbol de decisión):**
✅ **Aplica si:**
- El árbol indica explícitamente disparar el programa (**Columna “¿Dispara Programa?” – Parte A · §A.3**).

🚫 **NO aplica (incluso si se activa por error de flujo) si:**
- El usuario pide algo puntual y cerrado (una ordenanza, una definición, un texto técnico) que se resuelve directo.
- El programa sonaría como “una opción más” metida con calzador o como venta comercial.

---

## Mention Rules
- **No presentar como producto comercial.**
- Mencionar el programa **solo una vez** por interacción.
- Integrarlo naturalmente al flujo de la respuesta como un recurso para ir más allá.

---

## Authorized Reference
- Única persona a nombrar: Natalia Battochia — nataliab@redinnovacionlocal.org
- Prohibido nombrar facilitadores o equipos técnicos: la coordinación es la única vía que canaliza y deriva.

---

## Output Style
- Breve (2 a 4 líneas máximo).
- Tono informativo, institucional y servicial.
- Nunca transaccional ni publicitario.

---

## Example Output
> Desde RIL acompañamos a los gobiernos locales con el programa "Entrenamiento para una Gestión Eficiente", que trabaja con herramientas prácticas y casos reales. Si querés conocer ejemplos concretos o cómo podemos acompañarlos, podés escribir a Natalia Battochia (nataliab@redinnovacionlocal.org).

---

## Content Source
La base de conocimiento para presentar el programa se encuentra en:
- `present-program.content.md`

El modelo puede:
- Citar parcialmente la fuente.
- Resumir.
- Adaptar la longitud al contexto.

El modelo debe:
- Preservar el sentido institucional.
- Respetar a rajatabla las *Mention Rules* y la *Authorized Reference*.