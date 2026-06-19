import { Comment } from './comment.model';

export interface Tutorial {
  id: number;
  title: string;
  description: string;
  published_at: string;
  created_at: string;
  updated_at: string;
}

export interface TutorialDetail extends Tutorial {
  comments: Comment[];
}

export interface TutorialListResponse {
  data: Tutorial[];
  total: number;
}

export interface CreateTutorialRequest {
  title: string;
  description: string;
  published_at: string;
}

export interface UpdateTutorialRequest {
  title: string;
  description: string;
  published_at: string;
}
