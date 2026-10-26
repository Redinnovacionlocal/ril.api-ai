# Herramienta de seguimiento en tiempo real para la higiene urbana

## Ficha de réplica (RIL)

- **Núcleo de eficiencia:** Mapa en vivo del camión + reclamo geolocalizado derivado directo al equipo operativo.
- **Atributos de Rogers:** ventaja relativa Alta · compatibilidad Alta · complejidad Media · trialabilidad Alta · observabilidad Alta → replicabilidad **Alta** (2.77)

## Ficha completa

* **Tema:** Servicios públicos
* **Complejidad:** Media
* **Tiempo estimado:** 3 a 6 meses
* **Ciudad de origen:** Luján de Cuyo, Argentina
* **Implementación:** Implementado inicialmente en Luján de Cuyo (Argentina)

### Bajada

Un sistema que informa en tiempo real el recorrido de los camiones de recolección y poda, para que la ciudadanía conozca el horario del servicio y el gobierno local lo gestione con datos.

**Dato destacado:** Los reclamos por residuos bajaron 82% tras implementar el sistema de gestión de higiene urbana, del que el mapa de camiones forma parte.

### Qué problema resuelve

La recolección de residuos y la poda generan reclamos vecinales. Cuando el servicio no tiene trazabilidad, el gobierno local no sabe en qué punto está cada camión ni qué recorridos se cumplen, y los problemas se conocen recién cuando llega el reclamo. Sin datos, la gestión es reactiva.

### Cómo funciona

Cada camión de recolección cuenta con GPS. Desde el celular, vecinos y vecinas dan seguimiento en un mapa al recorrido de los camiones, reciben notificaciones sobre días y horarios de recolección y registran denuncias geolocalizadas. Con esa misma información, el gobierno local controla el cumplimiento de los recorridos por zona y por empresa contratista.

### Beneficios

* Datos para reducir reclamos y acortar los tiempos de respuesta.
* Control del servicio por zona y por contratista, para planificar y negociar con datos.
* Vecinos y vecinas informados sobre los horarios de recolección de su barrio.
* Una base de datos sobre la que después se pueden incorporar otras herramientas, como cámaras que detectan baches o microbasurales.

### Complejidad

**Media.** GPS y mapas en línea son tecnología estándar, pero requiere equipar a toda la flota y contar con un área operativa capaz de resolver lo que ingresa por el canal.

### Condiciones habilitantes

* Una flota de recolección, propia o contratada, a la que se le pueda instalar GPS.
* Una web o aplicación del gobierno local donde publicar el mapa y recibir reportes.
* Un equipo capaz de leer los datos y convertirlos en decisiones de gestión.
* Conectividad suficiente para operar la plataforma.
* Si el servicio está tercerizado, acceso acordado a los datos de GPS de la empresa.

### Actores involucrados

* **Área de Higiene Urbana:** Define recorridos y controla el cumplimiento del servicio.
* **Equipo de sistemas:** Instala los GPS, integra los datos y publica el mapa en línea.
* **Empresas contratistas:** Incorporan sus camiones al sistema.
* **Área de Comunicación:** Difunde la herramienta entre vecinos y vecinas.
* **Área de Atención Ciudadana:** Gestiona los reportes geolocalizados.

### Paso a paso

* Relevar la flota de recolección y sus recorridos por barrio.
* Instalar GPS en todos los camiones, incluidos los de empresas contratistas.
* Publicar un mapa en vivo en la web o aplicación del gobierno local, con el recorrido de cada camión.
* Cargar días y horarios por barrio y habilitar notificaciones para la ciudadanía.
* Incorporar en el mismo mapa la opción de reportar un problema con ubicación exacta.
* Desarrollar un tablero interno con el cumplimiento de recorridos por zona y por empresa.
* Lanzar la herramienta con una campaña de comunicación barrio por barrio.
* Responder cada reporte y medir reclamos, tiempos de respuesta y usuarios activos.

### Recomendaciones para que funcione mejor

* Comenzar por una zona o un servicio e incorporar los demás de forma gradual.
* Incorporar un chatbot que responda consultas sobre horarios y puntos de disposición de residuos.
* Con los datos ordenados, evaluar cámaras en los camiones que detecten baches y microbasurales sin esperar el reclamo.

### La idea en acción

**Luján de Cuyo, Mendoza · 172.109 habitantes · programa “Limpiando Luján”**

Luján de Cuyo recibía hasta 1.800 reclamos por año por poda y recolección, con un servicio fragmentado y sin trazabilidad. El gobierno local instaló GPS en el 100% de su flota e integró el mapa del camión en vivo a Mi Luján Digital, una plataforma que también reúne denuncias geolocalizadas, tableros de gestión y un chatbot. Actualmente avanza con la incorporación de cámaras e inteligencia artificial en los camiones para detectar baches y microbasurales.

**Impacto:**

* **−82%:** de reclamos por residuos (de 500 a 92 por año)
* **−96%:** de reclamos por poda (de 1.800 a 62 por año)
* **−65%:** en el tiempo promedio de respuesta

*Nota:* Los resultados corresponden al sistema completo de Luján de Cuyo, no solo al mapa del camión.

### Para reflexionar

Ordenar los reclamos es el primer paso. ¿Cómo va a detectar tu gobierno local los problemas que nadie reporta?

### Enlaces y recursos

* Nota en +Comunidad: https://mascomunidad.org.ar/cinco-politicas-que-reflejan-eficiencia-municipal-el-caso-de-lujan-de-cuyo/
* Historia: https://lujandecuyo.gob.ar/limpiandolujan/
* Limpiando Luján, sitio del gobierno local: https://lujandecuyo.gob.ar/limpiandolujan/

### Palabras clave de búsqueda

basura, residuos, recolección, camión de basura, poda, limpieza, higiene urbana, reclamos, gps, horarios de recolección, microbasurales, servicios públicos, contenedores
