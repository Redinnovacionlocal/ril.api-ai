---
name: ajuste-por-tamano
description: El tamaño del municipio (chico, mediano, grande) define qué niveles de madurez tiene sentido diagnosticar y modula el tono de las recomendaciones. Una misma capacidad ausente puede ser esperable en un municipio chico y una brecha en uno grande. Usar siempre que se elabore una recomendación.
---

# Skill: Ajuste por tamaño del municipio

Esta skill describe cómo el tamaño del municipio condiciona dos cosas fundamentales en la interacción:
1. Qué niveles de madurez tiene sentido diagnosticar para un tema.
2. Cómo modular el tono y la urgencia de las recomendaciones.

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

## 📏 Categorías de tamaño
* **Chica:** menos de 10.000 habitantes (algunas líneas usan <25.000).
* **Mediana:** entre 10.000 y 100.000 habitantes.
* **Grande:** más de 100.000 habitantes.

*Nota: Algunas preguntas del árbol usan solo dos categorías (chica vs. mediana-grande), otras tres. La skill lee lo que hay en el JSON sin imponer un esquema.*

---

## ⚙️ Procedimiento

### Paso 1 — Obtener el tamaño
Buscar en memoria: `get_user_memory(record_type="contexto_municipio", key="tamanio_ciudad")` o `key="poblacion"`.
* Si no está, obtenerlo de la conversación. Si el usuario dio población cuantitativa ("45.000 habitantes"), usarla. 
* Si no, preguntar: *"¿De qué tamaño es la ciudad, más o menos?"*.
* Guardar inmediatamente con `save_user_memory` para no volver a preguntar.

### Paso 2 — Definir el rango de niveles diagnosticables
Para una pregunta del árbol, el rango de niveles relevantes depende del tamaño:

| Tamaño | Niveles diagnosticables principales | Nivel "horizonte" |
| :--- | :--- | :--- |
| **Chico** | Bajo, Intermedio | Avanzado se menciona pero no se diagnostica activamente |
| **Mediano** | Bajo, Intermedio, Avanzado | — |
| **Grande** | Intermedio, Avanzado | Bajo es una brecha clara |

*Nota: Esto no es rígido. El campo `ajuste_por_tamano` de la pregunta puede sugerir un rango distinto para un tema específico. El agente debe leer ambas cosas y decidir.*

### Paso 3 — Modular el tono de la recomendación

**Si el municipio es chico:**
* **Tono:** Comprensivo. *"Es esperable que no tengan X dado el tamaño."*
* **Acciones:** Priorizan primeros pasos accesibles, no estructuras formales costosas.
* **Diagnóstico:** Si está en Bajo, presentar como punto de partida, no como brecha.
* **Estrategia:** Sugerir colaboración regional cuando aplique.

**Si el municipio es mediano o grande:**
* **Tono:** Más directivo. *"Hay escala para tenerlo y es una oportunidad/brecha."*
* **Acciones:** Más ambiciosas (estructuras formales, ordenanzas propias, procesos sistematizados).
* **Diagnóstico:** Si está en Bajo en una capacidad transversal, nombrar la brecha sin minimizar.

---

## 📝 Ejemplo: Mismo tema, ajuste distinto

**Contexto:** Sin cuerpo civil de prevención (P33 nivel Bajo).

> **Municipio chico (8.000 hab):**
> *"En municipios del tamaño de ustedes, no tener cuerpo civil propio es bastante común — no siempre hay escala para sostenerlo. Lo que sí podemos pensar es cómo está clarificado quién hace qué entre seguridad, tránsito y la policía provincial."*

> **Municipio mediano (60.000 hab):**
> *"Para una ciudad de su tamaño, no tener cuerpo civil propio es una brecha importante. Hay escala suficiente y los recursos suelen estar disponibles. Te propongo empezar por una ordenanza modelo y dimensionar la dotación inicial."*

El contenido objetivo es similar, pero el tono, la urgencia y la expectativa son completamente distintos.

---

## ⚠️ Casos particulares

* **Tamaño en frontera entre categorías:** Si el municipio tiene ~9.000 hab (frontera chica/mediana), preguntar al usuario cómo se vive desde el municipio. A veces uno de 9.000 funciona como mediano por su rol regional, o uno de 12.000 funciona como chico por su ruralidad.
* **Usuario rechaza el ajuste:** Si el usuario dice *"sí, ya sé que somos chicos, pero queremos hacer esto igual"*, respetar la decisión. No insistir con que la capacidad es ambiciosa. Apoyar el camino que el municipio elige.
* **No hay info de tamaño y la conversación es breve:** Si el usuario hizo una pregunta puntual y no es natural pedir el tamaño, dar la respuesta en tono neutral. Mencionar al final algo como *"esto puede variar según el tamaño de la ciudad"* para abrir la posibilidad.

---

## ❌ Lo que esta skill NO hace
* No reemplaza la lógica de niveles (skill separada).
* No reemplaza la lógica de anclas (skill separada).
* No es un filtro: el contenido del árbol no se oculta, solo se prioriza según el tamaño.
* No define las categorías exactas de tamaño en el árbol (eso lo define cada pregunta en el JSON).