# FrontOati

SPA Angular 22 para consumir la API OATI-CRUD (tutoriales y comentarios).

## Integración con backend

1. Levantar el backend:

```bash
cd backend && docker compose up --build
```

2. Instalar dependencias del frontend:

```bash
cd frontend && npm install
```

3. Iniciar el servidor de desarrollo:

```bash
npm start
```

La app queda disponible en `http://localhost:4200` y consume la API en `http://localhost:8080/api/v1` (configurable en `src/environments/environment.ts`).

### Flujo de prueba manual

1. Abrir `http://localhost:4200/tutorials`
2. Crear un tutorial con título, descripción y fecha de publicación
3. Ver el detalle y agregar comentarios
4. Editar y eliminar comentarios
5. Editar y eliminar el tutorial desde listado o detalle

Si el backend no está activo, la UI muestra un mensaje de error de conexión.

## Docker

El frontend tiene su propio `docker-compose.yml`. El backend debe estar activo por separado.

1. Levantar el backend (en otra terminal):

```bash
cd backend && docker compose up --build -d
```

2. Construir y levantar el frontend:

```bash
cd frontend && docker compose up --build
```

3. Abrir `http://localhost:4200`

La imagen usa **nginx** para servir el build de producción. El navegador consume la API en `http://localhost:8080/api/v1` (puerto mapeado del backend en el host).

Para detener el frontend:

```bash
docker compose down
```

## Development server

To start a local development server, run:

```bash
ng serve
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
ng build
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Running unit tests

To execute unit tests with the [Vitest](https://vitest.dev/) test runner, use the following command:

```bash
ng test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
ng e2e
```

Angular CLI does not come with an end-to-end testing framework by default. You can choose one that suits your needs.

## Additional Resources

For more information on using the Angular CLI, including detailed command references, visit the [Angular CLI Overview and Command Reference](https://angular.dev/tools/cli) page.
