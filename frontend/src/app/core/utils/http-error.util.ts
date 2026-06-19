import { HttpErrorResponse } from '@angular/common/http';

import { ApiError } from '../models/api-error.model';

export function getErrorMessage(err: unknown): string {
  if (err instanceof HttpErrorResponse) {
    const body = err.error as ApiError | string | null;
    if (body && typeof body === 'object' && 'message' in body && body.message) {
      return body.message;
    }
    if (err.status === 0) {
      return 'No se pudo conectar con el servidor. Verifica que el backend esté activo.';
    }
    return `Error ${err.status}: ${err.statusText || 'solicitud fallida'}`;
  }

  if (err instanceof Error) {
    return err.message;
  }

  return 'Ocurrió un error inesperado';
}
