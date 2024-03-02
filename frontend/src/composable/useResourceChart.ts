import { useThemeColors } from '/@src/composable/useThemeColors'

const baseOptions = (themeColors: ReturnType<typeof useThemeColors>) => ({
  chart: {
    height: 140,
    type: 'radialBar',
    offsetY: -6,
  },
  plotOptions: {
    radialBar: {
      startAngle: -135,
      endAngle: 135,
      inverseOrder: true,
      dataLabels: {
        show: true,
        name: {
          show: true,
          fontSize: '9px',
          fontWeight: 900,
          offsetY: -4,
        },
        value: {
          show: true,
          fontWeight: 900,
          color: themeColors.lightText,
          fontSize: '12px',
          offsetY: 4,
        }
      },
      hollow: {
        margin: 15,
        size: '50%',
      },
      track: {
        strokeWidth: '100%',
      },
    },
  },
  stroke: {
    lineCap: 'round',
  }
})

export function useResourceCharts() {
  const themeColors = useThemeColors()
  const cpuOptions = shallowRef({
    ...baseOptions(themeColors),
    colors: [themeColors.primary],
    labels: ['CPU Usage'],
  })

  const memoryOptions = shallowRef({
    ...baseOptions(themeColors),
    colors: [themeColors.accent],
    labels: ['Memory'],
  })

  const nnrOptions = shallowRef({
    ...baseOptions(themeColors),
    colors: [themeColors.blue],
    labels: ['Rx rate'],
  })

  const ntrOptions = shallowRef({
    ...baseOptions(themeColors),
    colors: [themeColors.danger],
    labels: ['Tx rate'],
  })

  return {
    cpuOptions,
    memoryOptions,
    nnrOptions,
    ntrOptions
  }
}
