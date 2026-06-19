import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';

import { environment } from '../../../environments/environment';
import {
  Comment,
  CommentListResponse,
  CreateCommentRequest,
  UpdateCommentRequest,
} from '../models/comment.model';

@Injectable({ providedIn: 'root' })
export class CommentService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = environment.apiUrl;

  getByTutorialId(tutorialId: number): Observable<CommentListResponse> {
    return this.http.get<CommentListResponse>(
      `${this.apiUrl}/tutorials/${tutorialId}/comments`,
    );
  }

  create(tutorialId: number, body: CreateCommentRequest): Observable<Comment> {
    return this.http.post<Comment>(
      `${this.apiUrl}/tutorials/${tutorialId}/comments`,
      body,
    );
  }

  update(id: number, body: UpdateCommentRequest): Observable<Comment> {
    return this.http.put<Comment>(`${this.apiUrl}/comments/${id}`, body);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/comments/${id}`, {
      observe: 'body',
    });
  }
}
