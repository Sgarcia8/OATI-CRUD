import { Component, inject, OnInit, signal } from '@angular/core';
import {
  FormBuilder,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';

import { TutorialService } from '../../../core/services/tutorial.service';
import { getErrorMessage } from '../../../core/utils/http-error.util';
import {
  datetimeLocalToIso,
  toDatetimeLocalValue,
} from '../../../core/utils/date.util';
import { AlertBanner } from '../../../shared/components/alert-banner/alert-banner';
import { LoadingSpinner } from '../../../shared/components/loading-spinner/loading-spinner';

@Component({
  selector: 'app-tutorial-form',
  imports: [ReactiveFormsModule, RouterLink, AlertBanner, LoadingSpinner],
  templateUrl: './tutorial-form.html',
})
export class TutorialForm implements OnInit {
  private readonly fb = inject(FormBuilder);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly tutorialService = inject(TutorialService);

  readonly loading = signal(false);
  readonly saving = signal(false);
  readonly error = signal('');
  readonly isEditMode = signal(false);
  readonly tutorialId = signal<number | null>(null);
  readonly breadcrumbTitle = signal('Nuevo tutorial');

  readonly form = this.fb.nonNullable.group({
    title: ['', Validators.required],
    description: ['', Validators.required],
    published_at: ['', Validators.required],
  });

  ngOnInit(): void {
    const idParam = this.route.snapshot.paramMap.get('id');
    if (idParam && idParam !== 'new') {
      const id = Number(idParam);
      if (!Number.isNaN(id)) {
        this.isEditMode.set(true);
        this.tutorialId.set(id);
        this.loadTutorial(id);
      }
    } else if (!this.form.controls.published_at.value) {
      this.form.patchValue({
        published_at: toDatetimeLocalValue(new Date().toISOString()),
      });
    }
  }

  loadTutorial(id: number): void {
    this.loading.set(true);
    this.error.set('');

    this.tutorialService.getById(id).subscribe({
      next: (tutorial) => {
        this.breadcrumbTitle.set(tutorial.title);
        this.form.patchValue({
          title: tutorial.title,
          description: tutorial.description,
          published_at: toDatetimeLocalValue(tutorial.published_at),
        });
        this.loading.set(false);
      },
      error: (err) => {
        this.error.set(getErrorMessage(err));
        this.loading.set(false);
      },
    });
  }

  submit(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    const payload = {
      title: this.form.controls.title.value,
      description: this.form.controls.description.value,
      published_at: datetimeLocalToIso(this.form.controls.published_at.value),
    };

    this.saving.set(true);
    this.error.set('');

    const request$ = this.isEditMode()
      ? this.tutorialService.update(this.tutorialId()!, payload)
      : this.tutorialService.create(payload);

    request$.subscribe({
      next: (tutorial) => {
        this.saving.set(false);
        void this.router.navigate(['/tutorials', tutorial.id]);
      },
      error: (err) => {
        this.error.set(getErrorMessage(err));
        this.saving.set(false);
      },
    });
  }
}
