# FrontOati

SPA Angular 22 para gestionar tutoriales y comentarios. Consume la API REST documentada en [`../backend/README.md`](../backend/README.md).

## Despliegue

Aplicación en producción: **https://oati-crud-front.onrender.com**

## Stack tecnológico

- Angular 22 (standalone components, lazy loading)
- TypeScript
- Tailwind CSS 4
- nginx (contenedor Docker de producción)

## Estructura del proyecto

```
src/app/
├── core/
│   ├── models/          # Tipos Tutorial, Comment y errores de API
│   ├── services/        # TutorialService, CommentService
│   └── utils/           # Formateo de fechas y errores HTTP
├── features/tutorials/
│   ├── tutorial-list/   # Listado con acciones CRUD
│   ├── tutorial-form/   # Crear y editar tutorial
│   └── tutorial-detail/ # Detalle con comentarios anidados
└── shared/components/   # Header, spinner, alertas, diálogo de confirmación
```

## Rutas

| Ruta | Pantalla |
|---|---|
| `/tutorials` | Listado de tutoriales |
| `/tutorials/new` | Crear tutorial |
| `/tutorials/:id` | Detalle con comentarios |
| `/tutorials/:id/edit` | Editar tutorial |

## Configuración de la API

La URL base se define por entorno:

| Archivo | Uso | Valor por defecto |
|---|---|---|
| `src/environments/environment.ts` | Desarrollo | `http://localhost:8080/api/v1` |
| `src/environments/environment.prod.ts` | Producción | `https://oati-crud.onrender.com/api/v1` |

Los servicios en `core/services/` leen `environment.apiUrl` para todas las peticiones HTTP.

## Desarrollo local

Requisitos: Node.js 22+ y npm.

```bash
npm install
npm start
```

La app queda en **http://localhost:4200**. Para que el CRUD funcione, la API debe estar disponible en la URL configurada en `environment.ts` (consulta [`../backend/README.md`](../backend/README.md) para levantarla).

### Flujo manual

1. Ir a `/tutorials`
2. Crear un tutorial (título, descripción, fecha de publicación)
3. Abrir el detalle y gestionar comentarios
4. Editar o eliminar desde listado o detalle

Si la API no responde, la interfaz muestra un banner de error de conexión.

## Docker

Build multi-etapa (Node + nginx) con enrutamiento SPA vía `try_files`:

```bash
docker compose up --build
```

Abrir **http://localhost:4200**. El contenedor sirve el build de producción; la API debe estar accesible desde el navegador en el host (por defecto `http://localhost:8080/api/v1`).

Detener:

```bash
docker compose down
```

## Compilación

```bash
npm run build
```

Salida en `dist/`. Para despliegues estáticos (Render, S3, etc.), usar el contenido de `dist/front-oati/browser/`.

En Render, configurar una regla de reescritura `/*` → `/index.html` para que las rutas del SPA funcionen al recargar la página.

## Pruebas

```bash
npm test
```

Incluye una prueba básica de arranque en `app.spec.ts`.
