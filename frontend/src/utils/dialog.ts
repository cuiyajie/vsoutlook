export interface ConfirmDialogOptions {
  title: string,
  cancelText?: string,
  confirmText?: string,
  content?: string,
  html?: string,
  onConfirm?: (hide: () => void) => void
}
export interface AlertDialogOptions {
  title: string,
  cancelText?: string,
  content?: string
}

export function confirm(options: ConfirmDialogOptions) {
  bus.trigger(Signal.OpenConfirmDialog, options)
}

export function valert(options: AlertDialogOptions) {
  bus.trigger(Signal.OpenAlertDialog, options)
}
