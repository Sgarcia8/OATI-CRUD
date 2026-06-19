export interface Comment {
  id: number;
  content: string;
  tutorial_id: number;
  created_at: string;
  updated_at: string;
}

export interface CommentListResponse {
  data: Comment[];
  total: number;
}

export interface CreateCommentRequest {
  content: string;
}

export interface UpdateCommentRequest {
  content: string;
}
