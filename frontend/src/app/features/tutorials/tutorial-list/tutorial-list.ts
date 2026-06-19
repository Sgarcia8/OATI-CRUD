import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterLink } from '@angular/router';

import { TutorialService } from '../../../core/services/tutorial.service';
import { Tutorial } from '../../../core/models/tutorial.model';
import { getErrorMessage } from '../../../core/utils/http-error.util';
import { formatDate } from '../../../core/utils/date.util';
import { AlertBanner } from '../../../shared/components/alert-banner/alert-banner';
import { ConfirmDialog } from '../../../shared/components/confirm-dialog/confirm-dialog';
import { LoadingSpinner } from '../../../shared/components/loading-spinner/loading-spinner';

@Component({
  selector: 'app-tutorial-list',
  imports: [RouterLink, AlertBanner, ConfirmDialog, LoadingSpinner],
  templateUrl: './tutorial-list.html',
})
export class TutorialList implements OnInit {
  private readonly tutorialService = inject(TutorialService);

  readonly tutorials = signal<Tutorial[]>([]);
  readonly loading = signal(true);
  readonly error = signal('');
  readonly success = signal('');
  readonly deleteDialogOpen = signal(false);
  readonly deleteLoading = signal(false);
  readonly tutorialToDelete = signal<Tutorial | null>(null);

  readonly formatDate = formatDate;

  ngOnInit(): void {
    this.loadTutorials();
  }

  loadTutorials(): void {
    this.loading.set(true);
    this.error.set('');

    this.tutorialService.getAll().subscribe({
      next: (response) => {
        this.tutorials.set(response.data ?? []);
        this.loading.set(false);
      },
      error: (err) => {
        this.error.set(getErrorMessage(err));
        this.loading.set(false);
      },
    });
  }

  openDeleteDialog(tutorial: Tutorial): void {
    this.tutorialToDelete.set(tutorial);
    this.deleteDialogOpen.set(true);
  }

  closeDeleteDialog(): void {
    this.deleteDialogOpen.set(false);
    this.tutorialToDelete.set(null);
  }

  confirmDelete(): void {
    const tutorial = this.tutorialToDelete();
    if (!tutorial) {
      return;
    }

    this.deleteLoading.set(true);
    this.tutorialService.delete(tutorial.id).subscribe({
      next: () => {
        this.deleteLoading.set(false);
        this.closeDeleteDialog();
        this.success.set('Tutorial eliminado correctamente');
        this.loadTutorials();
        setTimeout(() => this.success.set(''), 3000);
      },
      error: (err) => {
        this.deleteLoading.set(false);
        this.closeDeleteDialog();
        this.error.set(getErrorMessage(err));
      },
    });
  }

  truncate(text: string, max = 120): string {
    if (text.length <= max) {
      return text;
    }
    return `${text.slice(0, max)}...`;
  }

  deleteMessage(): string {
    const title = this.tutorialToDelete()?.title ?? '';
    return `¿Eliminar "${title}"? Esta acción también eliminará sus comentarios.`;
  }
}
