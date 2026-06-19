import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';

import { environment } from '../../../environments/environment';
import {
  CreateTutorialRequest,
  Tutorial,
  TutorialDetail,
  TutorialListResponse,
  UpdateTutorialRequest,
} from '../models/tutorial.model';

@Injectable({ providedIn: 'root' })
export class TutorialService {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = `${environment.apiUrl}/tutorials`;

  getAll(): Observable<TutorialListResponse> {
    return this.http.get<TutorialListResponse>(this.baseUrl);
  }

  getById(id: number): Observable<TutorialDetail> {
    return this.http.get<TutorialDetail>(`${this.baseUrl}/${id}`);
  }

  create(body: CreateTutorialRequest): Observable<Tutorial> {
    return this.http.post<Tutorial>(this.baseUrl, body);
  }

  update(id: number, body: UpdateTutorialRequest): Observable<Tutorial> {
    return this.http.put<Tutorial>(`${this.baseUrl}/${id}`, body);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}/${id}`, { observe: 'body' });
  }
}
