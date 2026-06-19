import { Component, input, output } from '@angular/core';

export type AlertType = 'error' | 'success';

@Component({
  selector: 'app-alert-banner',
  templateUrl: './alert-banner.html',
})
export class AlertBanner {
  readonly message = input.required<string>();
  readonly type = input<AlertType>('error');
  readonly dismissed = output<void>();

  dismiss(): void {
    this.dismissed.emit();
  }
}
