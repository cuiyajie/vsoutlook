export const draftVol = (h: number, bound: LayoutDimension) => {
  const { h: bh } = bound
  return {
    x: 0,
    y: 0,
    one_w: 10 / 1920,
    g: 5 / 1920,
    h: h / bh,
    cnt: 2
  }
}

export const draftTitle = (w: number, h: number, bound: LayoutDimension) => {
  const nw = 192 / 1920
  const nh = 54 / 1080
  const { w: bw, h: bh } = bound
  return {
    fontSize: 24 / 1920,
    text: '示例信号',
    x: (w - nw * bw) / (2 * bw),
    y: (h - nh * bh - 10) / bh,
    w: nw,
    h: nh
  }
}

export const draftTimer = (w: number, bound: LayoutDimension) => {
  const nw = 192 / 1920
  const nh = 54 / 1080
  const { w: bw, h: bh } = bound
  return {
    fontSize: 16 / 1920,
    text: '16:00:00 2024-04-12',
    color: '#ff0000',
    x: (w - nw * bw) / (2 * bw),
    y: 10 / bh,
    w: nw,
    h: nh
  }
}

export const DefaultLayouts: [number, number][][] = [
  [[1, 1]],
  [[2, 2]],
  [[3, 3]],
  [[2, 1], [4, 2]],
]
