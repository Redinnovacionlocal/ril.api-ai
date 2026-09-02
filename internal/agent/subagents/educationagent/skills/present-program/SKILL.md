---
name: present-program
description: 'Proveer contexto institucional y mencionar el programa "Ciudades de la Educación" de RIL. Se dispara cuando un municipio o funcionario consulta por: planificación o diseño de políticas educativas locales, creación del área de educación desde cero, metodologías para la revinculación o retención escolar, fortalecimiento del ecosistema educativo local, diagnóstico edilicio o de infraestructura, o cuando solicita ejemplos de casos de éxito y acompañamiento estratégico para ordenar su gestión educativa municipal.'
---

# Skill: present-program

## Purpose
Presentar brevemente el programa "Ciudades de la Educación" como un contexto de acompañamiento institucional. No es un discurso de venta ni una recomendación comercial genérica, sino una mención contextualizada y justificada para profundizar la gestión educativa local.

---

## Validation & Constraints (When to apply)
El modelo DEBE disparar la mención si se cumple **cualquiera** de las siguientes vías:

**Vía 1 (Por contexto de la conversación):**
✅ **Aplica si:**
- El usuario pide **ejemplos de casos de ciudades**, experiencias comparadas, buenas prácticas o aprendizajes entre municipios en materia educativa.
- El usuario busca **acompañamiento**, enfoque metodológico o una hoja de ruta para **diseñar, ordenar o transformar la agenda educativa** local.
- El usuario está **creando el área de educación desde cero** o busca articular a la municipalidad con escuelas y actores del ecosistema educativo.
- El usuario formula preguntas complejas o amplias sobre temas clave (ej. deserción escolar, infraestructura, formación docente) que requieren un marco programático para responderse integralmente.

🚫 **NO aplica si:**
- El usuario solicita un dato o insumo normativo puntual y cerrado (una ordenanza específica, una definición conceptual, un modelo de convenio técnico).
- La pregunta se responde de forma directa con la base de conocimiento sin requerir contexto institucional.
- La mención del programa suena forzada, publicitaria o como "un contenido metido con calzador".

**Vía 2 (Por el árbol de decisión):**
✅ **Aplica si:**
- El árbol indica explícitamente disparar el programa (**Columna “¿Dispara Programa?” – Parte A · §A.3**).
---

## Mention Rules
- **Prohibido el tono comercial/venta:** Presentar el programa como un espacio de red, metodologías de trabajo y aprendizaje entre pares.
- **Sin redundancias:** Mencionar el nombre del programa ("Ciudades de la Educación") una sola vez en el mensaje.
- **Integración fluida:** Colocar la mención de forma natural al cierre o durante la respuesta, ofreciendo el programa como el paso lógico para ir más allá.

---

## Authorized Reference
- **Persona autorizada:** Eduardo Viñales — eduardo@redinnovacionlocal.org
- **Cargo:** Coordinación del programa Ciudades de la Educación.
- **Restricción estricta:** Queda prohibido nombrar a facilitadores, consultores externos o equipos técnicos. Toda derivación o canalización pasa de forma exclusiva por la coordinación.

---

## Output Style
- **Extensión:** Breve (2 a 4 líneas máximo).
- **Tono:** Profesional, servicial, institucional y enfocado en la gestión pública.
- **Enfoque:** Orientado a la acción colaborativa (aprender de otras ciudades, metodologías probadas).

---

## Example Output
> Para profundizar en el diseño de esta política o conocer cómo otros municipios resolvieron desafíos similares, desde RIL acompañamos a los gobiernos locales a través del programa **Ciudades de la Educación**, aportando herramientas de gestión y conexión con buenas prácticas del país.
> 
> Si querés explorar un acompañamiento institucional o ver casos aplicados a tu municipio, podés escribirle directamente a Eduardo Viñales (eduardo@redinnovacionlocal.org), coordinador/a del programa.

---

## Content Source
La base de conocimiento para detallar el valor del programa se encuentra en:
- `present-program.content.md`

**Instrucciones de uso del contenido:**
- El modelo puede resumir o adaptar la cita según el tono de la consulta del usuario.
- El modelo debe respetar a rajatabla las *Mention Rules* y los datos de la *Authorized Reference*.