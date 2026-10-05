---
name: ajuste-por-tamano
description: El tamaño del municipio (chico, mediano/grande) define qué niveles de madurez tiene sentido diagnosticar y modula el tono de las recomendaciones en Primera Infancia. Una misma capacidad ausente puede ser esperable en un municipio chico y una brecha en uno mediano o grande. Usar siempre que se elabore una recomendación.
---

# Skill: Ajuste por tamaño del municipio (Primera Infancia)

Esta skill describe cómo el tamaño del municipio condiciona dos cosas fundamentales en la interacción:
1. Qué niveles de madurez tiene sentido diagnosticar para un tema.
2. Cómo modular el tono y la expectativa de las recomendaciones.

## 🎯 Cuándo se activa
**Siempre que el agente esté por elaborar una recomendación específica al municipio.** Es un modulador transversal, no una respuesta separada.
*Nota: Si no hay información de tamaño en memoria ni en la conversación, obtenerlo primero (ver Paso 1).*

---

## 🧠 Conceptos clave

* **Tamaño acota niveles diagnosticables:** Para un municipio chico, diagnosticar contra el nivel Avanzado de muchas preguntas puede ser desproporcionado. El tamaño define qué niveles vale la pena explorar.
* **Capacidad ausente "esperable":** En una ciudad chica, no tener cierta estructura puede ser normal (no hay escala para sostenerla).
* **Capacidad ausente "brecha":** En una ciudad mediana o grande, no tener cierta estructura suele indicar una falta concreta.
* **Modulación:** El ajuste modula el tono y la expectativa, pero **no cambia** el contenido objetivo de las acciones disponibles.

---

## 📏 Corte de tamaño (definido con el equipo, fijo para Primera Infancia)

* **Chica:** hasta 50.000 habitantes.
* **Mediana:** de 50.001 a 500.000 habitantes.
* **Grande:** más de 500.000 habitantes.

Este corte está cargado y revisado en las 72 preguntas del árbol, en las columnas
**"Ajuste ciudad chica"** y **"Ajuste ciudad mediana/grande"** (campos `ajuste_chica` /
`ajuste_mediana_grande` de cada pregunta). A diferencia de otros agentes, en Primera
Infancia este criterio de tamaño **no varía por pregunta**: es el mismo corte para
las 72 preguntas del árbol.

*Nota: Mediana y grande comparten el mismo campo de ajuste (`ajuste_mediana_grande`).
No hay un ajuste diferenciado entre mediana y grande: para efectos de esta skill,
tratalas como un mismo bloque frente a chica.*

---

## ⚙️ Procedimiento

### Paso 1 — Obtener el tamaño
Buscar en memoria: `get_user_memory(record_type="contexto_municipio", key="tamanio_ciudad")` o `key="poblacion"`.
* Si no está, obtenerlo de la conversación. Si el usuario dio población cuantitativa ("45.000 habitantes"), usarla.
* Si no, preguntar: *"¿De qué tamaño es la ciudad, más o menos?"*.
* Guardar inmediatamente con `save_user_memory` para no volver a preguntar.

### Paso 2 — Leer el ajuste de la pregunta
Para una pregunta del árbol, leer el campo que corresponda según el tamaño:
* Población ≤ 50.000 → usar `ajuste_chica`.
* Población > 50.000 (mediana o grande) → usar `ajuste_mediana_grande`.

Ese texto define qué niveles son relevantes para ese tema y cómo formular la
recomendación. No hay lógica de fallback a tabla genérica: el ajuste viene siempre
cargado en la pregunta.

### Paso 3 — Modular el tono de la recomendación

**Si el municipio es chico (≤50.000 hab.):**
* **Tono:** Comprensivo. *"Es esperable que no tengan X dado el tamaño."*
* **Acciones:** Priorizan primeros pasos accesibles, no estructuras formales costosas.
* **Diagnóstico:** Si está en Bajo, presentar como punto de partida, no como brecha.
* **Estrategia:** Sugerir colaboración regional o con otras áreas (salud, desarrollo social) cuando aplique.

**Si el municipio es mediano o grande (>50.000 hab.):**
* **Tono:** Más directivo. *"Hay escala para tenerlo y es una oportunidad/brecha."*
* **Acciones:** Más ambiciosas y acordes a la escala del municipio.
* **Diagnóstico:** Si está en Bajo en una capacidad transversal, nombrar la brecha sin minimizar.

---

## 📝 Ejemplo: Mismo tema, ajuste distinto

**Contexto:** Organización del área municipal de primera infancia.

> **Municipio chico (30.000 hab.):**
> *"En municipios del tamaño de ustedes, no necesariamente hace falta una estructura grande. Lo importante es tener un referente claro, con responsabilidades definidas y capacidad de coordinar las acciones de primera infancia."*

> **Municipio mediano/grande (120.000 hab.):**
> *"Para una ciudad de este tamaño, ya hay escala para contar con una estructura más dedicada. Si la gestión de primera infancia depende solo de una persona sin responsabilidades formalizadas, puede ser una brecha que conviene ordenar."*

---

## ⚠️ Casos particulares

* **Tamaño en frontera (ej. 49.000 o 51.000 hab.):** Usar el corte tal cual está definido (50.000 es el límite exacto de "chica"). No negociar el umbral; si hace falta, preguntar al usuario cómo se vive desde el municipio.
* **Usuario rechaza el ajuste:** Si el usuario dice *"sí, ya sé que somos chicos, pero queremos hacer esto igual"*, respetar la decisión. No insistir con que la capacidad es ambiciosa. Apoyar el camino que el municipio elige.
* **No hay info de tamaño y la conversación es breve:** Si el usuario hizo una pregunta puntual y no es natural pedir el tamaño, dar la respuesta en tono neutral. Mencionar al final algo como *"esto puede variar según el tamaño de la ciudad"* para abrir la posibilidad.

---

## ❌ Lo que esta skill NO hace
* No reemplaza la lógica de niveles (skill `niveles-madurez`, separada).
* No reemplaza la lógica de anclas (skill `anclas-e-hijas`, separada). Una ancla en Bajo bloquea sus hijas independientemente del tamaño del municipio.
* No es un filtro: el contenido del árbol no se oculta, solo se prioriza según el tamaño.
* No define cortes distintos por pregunta: en Primera Infancia el corte (50.000 / 500.000) es único para las 72 preguntas.
