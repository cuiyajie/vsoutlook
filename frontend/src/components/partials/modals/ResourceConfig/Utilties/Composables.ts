export function useUsedFormat() {
  const formatEnumMap = ref<Record<string, number>>({})
  const formatEnum = computed(() => Object.keys(formatEnumMap.value))
  function onSelected(name: string) {
    formatEnumMap.value[name] = (formatEnumMap.value[name] || 0) + 1;
  }

  function onUnselected(name: string) {
    if (formatEnumMap.value[name] === 1) {
      delete formatEnumMap.value[name];
    } else {
      formatEnumMap.value[name] -= 1;
    }
  }

  return [
    formatEnum,
    onSelected,
    onUnselected,
  ] as const
}
