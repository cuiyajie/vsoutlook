import { useThemeColors } from '/@src/composable/useThemeColors'

export function useResourceCharts() {
  const themeColors = useThemeColors()
  const usageOptions = shallowRef({
    series: [57],
    chart: {
      height: 120,
      type: 'radialBar',
      offsetY: -10,
    },
    colors: [themeColors.primary],
    plotOptions: {
      radialBar: {
        startAngle: -135,
        endAngle: 135,
        inverseOrder: true,
        dataLabels: {
          show: true,
          name: {
            show: true,
            fontSize: '14px',
            fontWeight: 500,
            offsetY: -10,
          },
          value: {
            show: true,
            fontWeight: 600,
            color: themeColors.lightText,
            fontSize: '16px',
            offsetY: -5,
          },
          total: {
            show: true,
            fontSize: '14px',
            fontWeight: 500,
            color: themeColors.lightText,
          },
        },
        hollow: {
          margin: 15,
          size: '75%',
        },
        track: {
          strokeWidth: '300%',
        },
      },
    },

    stroke: {
      lineCap: 'round',
    },
    labels: ['Productivity'],
  })

  return {
    usageOptions,
  }
}
