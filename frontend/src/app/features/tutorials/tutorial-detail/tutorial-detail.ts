import { Component, inject, OnInit, signal } from '@angular/core';
import {
  FormBuilder,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';

import { CommentService } from '../../../core/services/comment.service';
import { TutorialService } from '../../../core/services/tutorial.service';
import { Comment } from '../../../core/models/comment.model';
import { TutorialDetail as TutorialDetailModel } from '../../../core/models/tutorial.model';
import { getErrorMessage } from '../../../core/utils/http-error.util';
import { formatDate } from '../../../core/utils/date.util';
import { AlertBanner } from '../../../shared/components/alert-banner/alert-banner';
import { ConfirmDialog } from '../../../shared/components/confirm-dialog/confirm-dialog';
import { LoadingSpinner } from '../../../shared/components/loading-spinner/loading-spinner';

@Component({
  selector: 'app-tutorial-detail',
  imports: [
    ReactiveFormsModule,
    RouterLink,
    AlertBanner,
    ConfirmDialog,
    LoadingSpinner,
  ],
  templateUrl: './tutorial-detail.html',
})
export class TutorialDetailPage implements OnInit {
  private readonly fb = inject(FormBuilder);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly tutorialService = inject(TutorialService);
  private readonly commentService = inject(CommentService);

  readonly tutorial = signal<TutorialDetailModel | null>(null);
  readonly loading = signal(true);
  readonly error = signal('');
  readonly success = signal('');
  readonly notFound = signal(false);
  readonly savingComment = signal(false);
  readonly editingCommentId = signal<number | null>(null);
  readonly editSaving = signal(false);
  readonly deleteDialogOpen = signal(false);
  readonly deleteLoading = signal(false);
  readonly deleteTarget = signal<'tutorial' | 'comment'>('tutorial');
  readonly commentToDelete = signal<Comment | null>(null);

  readonly formatDate = formatDate;

  readonly newCommentForm = this.fb.nonNullable.group({
    content: ['', Validators.required],
  });

  readonly editCommentForm = this.fb.nonNullable.group({
    content: ['', Validators.required],
  });

  ngOnInit(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    if (Number.isNaN(id)) {
      this.notFound.set(true);
      this.loading.set(false);
      return;
    }
    this.loadTutorial(id);
  }

  loadTutorial(id: number): void {
    this.loading.set(true);
    this.error.set('');

    this.tutorialService.getById(id).subscribe({
      next: (tutorial) => {
        this.tutorial.set(tutorial);
        this.loading.set(false);
        this.notFound.set(false);
      },
      error: (err) => {
        this.loading.set(false);
        if (err.status === 404) {
          this.notFound.set(true);
        } else {
          this.error.set(getErrorMessage(err));
        }
      },
    });
  }

  createComment(): void {
    const tutorial = this.tutorial();
    if (!tutorial || this.newCommentForm.invalid) {
      this.newCommentForm.markAllAsTouched();
      return;
    }

    this.savingComment.set(true);
    this.commentService
      .create(tutorial.id, { content: this.newCommentForm.controls.content.value })
      .subscribe({
        next: () => {
          this.newCommentForm.reset();
          this.savingComment.set(false);
          this.success.set('Comentario creado');
          this.loadTutorial(tutorial.id);
          setTimeout(() => this.success.set(''), 3000);
        },
        error: (err) => {
          this.error.set(getErrorMessage(err));
          this.savingComment.set(false);
        },
      });
  }

  startEditComment(comment: Comment): void {
    this.editingCommentId.set(comment.id);
    this.editCommentForm.patchValue({ content: comment.content });
  }

  cancelEditComment(): void {
    this.editingCommentId.set(null);
    this.editCommentForm.reset();
  }

  saveComment(commentId: number): void {
    if (this.editCommentForm.invalid) {
      this.editCommentForm.markAllAsTouched();
      return;
    }

    const tutorial = this.tutorial();
    if (!tutorial) {
      return;
    }

    this.editSaving.set(true);
    this.commentService
      .update(commentId, { content: this.editCommentForm.controls.content.value })
      .subscribe({
        next: () => {
          this.editSaving.set(false);
          this.editingCommentId.set(null);
          this.success.set('Comentario actualizado');
          this.loadTutorial(tutorial.id);
          setTimeout(() => this.success.set(''), 3000);
        },
        error: (err) => {
          this.error.set(getErrorMessage(err));
          this.editSaving.set(false);
        },
      });
  }

  openDeleteTutorialDialog(): void {
    this.deleteTarget.set('tutorial');
    this.commentToDelete.set(null);
    this.deleteDialogOpen.set(true);
  }

  openDeleteCommentDialog(comment: Comment): void {
    this.deleteTarget.set('comment');
    this.commentToDelete.set(comment);
    this.deleteDialogOpen.set(true);
  }

  closeDeleteDialog(): void {
    this.deleteDialogOpen.set(false);
    this.commentToDelete.set(null);
  }

  confirmDelete(): void {
    const tutorial = this.tutorial();
    if (!tutorial) {
      return;
    }

    this.deleteLoading.set(true);

    if (this.deleteTarget() === 'tutorial') {
      this.tutorialService.delete(tutorial.id).subscribe({
        next: () => {
          this.deleteLoading.set(false);
          this.closeDeleteDialog();
          void this.router.navigate(['/tutorials']);
        },
        error: (err) => {
          this.deleteLoading.set(false);
          this.closeDeleteDialog();
          this.error.set(getErrorMessage(err));
        },
      });
      return;
    }

    const comment = this.commentToDelete();
    if (!comment) {
      return;
    }

    this.commentService.delete(comment.id).subscribe({
      next: () => {
        this.deleteLoading.set(false);
        this.closeDeleteDialog();
        this.success.set('Comentario eliminado');
        this.loadTutorial(tutorial.id);
        setTimeout(() => this.success.set(''), 3000);
      },
      error: (err) => {
        this.deleteLoading.set(false);
        this.closeDeleteDialog();
        this.error.set(getErrorMessage(err));
      },
    });
  }

  deleteDialogTitle(): string {
    return this.deleteTarget() === 'tutorial'
      ? 'Eliminar tutorial'
      : 'Eliminar comentario';
  }

  deleteDialogMessage(): string {
    if (this.deleteTarget() === 'tutorial') {
      return `¿Eliminar "${this.tutorial()?.title ?? ''}"? Esta acción también eliminará sus comentarios.`;
    }
    return '¿Eliminar este comentario? Esta acción no se puede deshacer.';
  }
}
