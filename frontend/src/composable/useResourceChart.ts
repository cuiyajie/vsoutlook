import { useThemeColors } from '/@src/composable/useThemeColors'

const baseOptions = (themeColors: ReturnType<typeof useThemeColors>) => ({
  chart: {
    height: 180,
    type: 'radialBar',
    offsetY: 0,
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
          fontSize: '12px',
          fontWeight: 500,
          offsetY: -10,
        },
        value: {
          show: true,
          fontWeight: 900,
          color: themeColors.lightText,
          fontSize: '18px',
          offsetY: 0,
        }
      },
      hollow: {
        margin: 15,
        size: '60%',
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

  return {
    cpuOptions,
    memoryOptions,
  }
}
