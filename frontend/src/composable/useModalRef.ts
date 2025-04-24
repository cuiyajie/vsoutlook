export function useModalRef(selector: string) {
  const modalRef = ref<HTMLElement>(document.body)
  onMounted(() => {
    modalRef.value = document.getElementById(selector) || document.body
  })
  return modalRef
}
