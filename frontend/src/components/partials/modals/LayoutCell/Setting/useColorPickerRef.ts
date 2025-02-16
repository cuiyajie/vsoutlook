export function useColorPickerRef() {
  const pickerRef = ref('body')
  onMounted(() => {
    pickerRef.value = '#layout-cell-modal'
  })
  return pickerRef
}
