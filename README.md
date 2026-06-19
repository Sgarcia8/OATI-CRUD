# OATI-CRUD

Monorepo con API REST y SPA para gestionar **tutoriales** y **comentarios** (relación 1:N). Incluye CRUD completo, eliminación lógica, Swagger, Docker y despliegue en Render.

## Estructura

```
OATI-CRUD/
├── backend/    # API Go + Beego + PostgreSQL
└── frontend/   # SPA Angular 22
```

| Documentación | Descripción |
|---|---|
| [backend/README.md](backend/README.md) | API, arquitectura, endpoints, Swagger, CORS, Docker |
| [frontend/README.md](frontend/README.md) | Rutas, entornos, desarrollo y Docker del SPA |

## Despliegue (producción)

| Recurso | URL |
|---|---|
| Frontend | https://oati-crud-front.onrender.com |
| API | https://oati-crud.onrender.com |
| Swagger | https://oati-crud.onrender.com/swagger/ |

## Stack

| Capa | Tecnologías |
|---|---|
| Backend | Go 1.26, Beego v2, PostgreSQL 16, Beego ORM |
| Frontend | Angular 22, TypeScript, Tailwind CSS 4 |

## Inicio rápido (local)

### 1. Backend

```bash
cd backend
docker compose up --build
```

API disponible en **http://localhost:8080** · Swagger en **http://localhost:8080/swagger/**

### 2. Frontend

En otra terminal:

```bash
cd frontend
npm install
npm start
```

App disponible en **http://localhost:4200** (consume `http://localhost:8080/api/v1`).

### Con Docker (solo frontend)

El backend debe estar activo por separado:

```bash
cd frontend && docker compose up --build
```

## Modelo de datos (resumen)

- **tutorials** — título, descripción, fecha de publicación
- **comments** — contenido, FK `tutorial_id`
- Ambas tablas usan `is_deleted` para eliminación lógica

Diagrama y detalle en [backend/README.md](backend/README.md#modelo-de-datos).

## Requisitos

- **Backend:** Go 1.26+ o Docker
- **Frontend:** Node.js 22+ y npm
