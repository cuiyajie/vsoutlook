export interface ConfirmDialogOptions {
  title: string,
  cancelText?: string,
  confirmText?: string,
  content?: string,
  html?: string,
  onConfirm?: (hide: () => void) => void
}

export function confirm(options: ConfirmDialogOptions) {
  bus.trigger(Signal.OpenConfirmDialog, options)
}
