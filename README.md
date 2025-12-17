# ril.api-ia

Mini descripción

Proyecto backend en Go que expone una API HTTP para gestión de sesiones y ejecución de un "runner" mediante SSE. El proyecto usa los paquetes `google.golang.org/adk` (session, artifact, runner) y un agente personalizado ubicado en `internal/agent`. Repositorios en memoria están disponibles en `internal/infrastructure/repository/memory` y se cargan datos de prueba cuando `APP_ENV=local`.

Requisitos

- macOS (u otro sistema con Go instalado)
- Go 1.20+
- git (opcional)

Archivos/paths clave

- `cmd/api/main.go` - entrada de la aplicación
- `internal/agent` - agente personalizado usado por el runner
- `internal/application/usecase/session_use_case.go` - reglas de negocio de sesiones
- `internal/infrastructure/repository/memory` - repositorios en memoria (útiles para desarrollo)
- `.env` / `.env.example` - variables de entorno

Variables de entorno (ejemplo)

Las variables pueden colocarse en un archivo `.env` en la raíz del proyecto. Ejemplo mínimo:

APP_NAME=ril-api
APP_ENV=local
PORT=8080

Setup rápido

1. Clonar el repositorio:

```bash
git clone <repo-url>
cd <repo-folder>
```

2. Instalar dependencias y preparar el entorno:

```bash
go mod tidy
cp .env.example .env # editar .env según necesidades
```

3. (Opcional) Ajustar variables en `.env`.

Ejecución en desarrollo

- Ejecutar:

```bash
APP_ENV=local go run ./cmd/api
```

- O desde el binario compilado:

```bash
go build -o bin/ril-api ./cmd/api
./bin/ril-api
```

- Por defecto el servidor corre en `:8080`. Puedes configurar el `PORT` en `.env`.

Seeds / datos de desarrollo

Si `APP_ENV=local`, en `main.go` se ejecuta `FillMockUsers` que crea usuarios de prueba y escribe sus tokens en los logs.

Endpoints principales (revisar handlers para detalles)

- POST /sessions         - crear sesión
- GET  /sessions         - listar sesiones
- GET  /sessions/:id     - obtener sesión
- DELETE /sessions/:id   - eliminar sesión
- POST /run-sse          - iniciar ejecución vía SSE

Tests

```bash
go test ./...
```

Build / Release

```bash
go build -o bin/ril-api ./cmd/api
```

Notas

- Se usa `github.com/joho/godotenv` para cargar `.env`.
- Los repositorios por defecto son en memoria; para producción sustituir por implementación SQL en `internal/infrastructure/repository/sql`.
- `.env.example` está presente pero vacío; copia y rellena `.env` con los valores necesarios.

Contacto / Más info

Revisa los handlers en `internal/infrastructure/http/handler` y la configuración del runner en `cmd/api/main.go` para entender el flujo del runner y la integración con el agente.

