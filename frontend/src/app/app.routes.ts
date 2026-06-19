import { Routes } from '@angular/router';

export const routes: Routes = [
  { path: '', redirectTo: 'tutorials', pathMatch: 'full' },
  {
    path: 'tutorials',
    loadComponent: () =>
      import('./features/tutorials/tutorial-list/tutorial-list').then(
        (m) => m.TutorialList,
      ),
  },
  {
    path: 'tutorials/new',
    loadComponent: () =>
      import('./features/tutorials/tutorial-form/tutorial-form').then(
        (m) => m.TutorialForm,
      ),
  },
  {
    path: 'tutorials/:id/edit',
    loadComponent: () =>
      import('./features/tutorials/tutorial-form/tutorial-form').then(
        (m) => m.TutorialForm,
      ),
  },
  {
    path: 'tutorials/:id',
    loadComponent: () =>
      import('./features/tutorials/tutorial-detail/tutorial-detail').then(
        (m) => m.TutorialDetailPage,
      ),
  },
  { path: '**', redirectTo: 'tutorials' },
];
