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
    color: '#FFFFFF',
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

const round = Math.round
export function ds2db(ds: LayoutDataItem[], base: LayoutDimension) {
  return ds.map((d, i) => {
    const rects = []
    if (d.win) {
      rects.push({
        type: 0,
        x: round(d.win.x * base.w),
        y: round(d.win.y * base.h),
        w: round(d.win.w * base.w),
        h: round(d.win.h * base.h),
      })
    }
    if (d.title) {
      rects.push({
        type: 1,
        x: round(d.title.x * base.w),
        y: round(d.title.y * base.h),
        w: round(d.title.w * base.w),
        h: round(d.title.h * base.h),
        fontSize: round(d.title.fontSize * base.w),
        text: d.title.text,
      })
    }
    if (d.vol) {
      rects.push({
        type: 2,
        x: round(d.vol.x * base.w),
        y: round(d.vol.y * base.h),
        one_w: round(d.vol.one_w * base.w),
        g: round(d.vol.g * base.w),
        h: round(d.vol.h * base.h),
        count: d.vol.cnt,
      })
    }
    return {
      index: i,
      rects,
    }
  })
}

export function db2ds(db: any[], base: LayoutDimension): LayoutDataItem[] {
  return db.map(d => {
    const win = d.rects.find((r: any) => r.type === 0)
    const title = d.rects.find((r: any) => r.type === 1)
    const vol = d.rects.find((r: any) => r.type === 2)
    return {
      win: win ? {
        x: win.x / base.w,
        y: win.y / base.h,
        w: win.w / base.w,
        h: win.h / base.h,
      } : null,
      title: title ? {
        x: title.x / base.w,
        y: title.y / base.h,
        w: title.w / base.w,
        h: title.h / base.h,
        fontSize: title.fontSize / base.w,
        text: title.text,
      } : null,
      vol: vol ? {
        x: vol.x / base.w,
        y: vol.y / base.h,
        one_w: vol.one_w / base.w,
        g: vol.g / base.w,
        h: vol.h / base.h,
        cnt: vol.count,
      } : null,
      timer: null
    } as LayoutDataItem
  })
}
