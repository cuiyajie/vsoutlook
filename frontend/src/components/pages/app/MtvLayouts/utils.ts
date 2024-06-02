import { useUserSession } from "@src/stores/userSession";
import parseRGB from "@src/utils/parse-rgb";
import stringRGB from '@src/utils/string-rgb';

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
    fontFamily: getFontFamily(),
    color: '#FFFFFF',
    x: (w - nw * bw) / (2 * bw),
    y: (h - nh * bh - 10) / bh,
    w: nw,
    h: nh
  }
}

export const defaultTimerColor = 'rgb(255, 0, 0)'
export const draftTimer = () => {
  return {
    fontSize: 64 / 1920,
    color: defaultTimerColor,
    x: 800 / 1920,
    y: 100 / 1920,
    showDate: true,
    dateNewLine: 0,
    dateDisplayType: 0,
    showFrame: 0,
    time24: 24,
    fontFamily: getFontFamily(),
  }
}

export const resizeVol = (vol: LayoutVol, rw: number, rh: number) => {
  return {
    ...vol,
    x: vol.x * rw,
    y: vol.y * rh,
    one_w: vol.one_w * rw,
    g: vol.g * rw,
    h: vol.h * rh,
  }
}

export const resizeTitle = (title: LayoutTitle, rw: number, rh: number) => {
  return {
    ...title,
    x: title.x * rw,
    y: title.y * rh,
    w: title.w * rw,
    h: title.h * rh,
    fontSize: title.fontSize * rw,
  }
}

export const resizeTimer = (timer: LayoutTimer, rw: number, rh: number) => {
  return {
    ...timer,
    x: timer.x * rw,
    y: timer.y * rh,
    fontSize: timer.fontSize * rw,
  }
}

export const DefaultLayouts: [number, number][][] = [
  [[1, 1]],
  [[2, 2]],
  [[3, 3]],
  [[2, 1], [4, 2]],
]

export const defaultFontSize = 24 / 1920
export const defaultFontFamily = 'Noto Serif CJK'

export function getFontFamily() {
  return useUserSession().settings.mv_template_font || defaultFontFamily
}

const round = Math.round
export function ds2db(ds: LayoutDataItem[], base: LayoutDimension) {
  const fontFamily = getFontFamily()
  return ds.map((d, i) => {
    if (d.timer) {
      return {
        index: i,
        is_clock: 1,
        clock_params: {
          show_date: d.timer.showDate ? 1 : 0,
          date_new_line: d.timer.dateNewLine,
          date_display_type: d.timer.dateDisplayType,
          showframe: d.timer.showFrame,
          "24/12": d.timer.time24,
          fontname: fontFamily,
          fontsize: round(d.timer.fontSize * base.w),
          fontcolour: parseRGB(d.timer.color),
          top: {
            x: round(d.timer.x * base.w),
            y: round(d.timer.y * base.h),
          },
        }
      }
    }
    const rects = []
    let fontSize = round(base.w * defaultFontSize)
    let pos: LayoutPos = { x: 0, y: 0 }
    if (d.win) {
      const x = round(d.win.x * base.w)
      const y = round(d.win.y * base.h)
      rects.push({
        type: 0,
        x,
        y,
        w: round(d.win.w * base.w),
        h: round(d.win.h * base.h),
      })
      pos = { x, y }
    }
    if (d.title) {
      rects.push({
        type: 1,
        x: round(d.title.x * base.w) + pos.x,
        y: round(d.title.y * base.h) + pos.y,
        w: round(d.title.w * base.w),
        h: round(d.title.h * base.h)
      })
      fontSize = round(d.title.fontSize * base.w)
    }
    if (d.vol) {
      rects.push({
        type: 2,
        x: round(d.vol.x * base.w) + pos.x,
        y: round(d.vol.y * base.h) + pos.y,
        one_w: round(d.vol.one_w * base.w),
        g: round(d.vol.g * base.w),
        h: round(d.vol.h * base.h),
        count: d.vol.cnt,
      })
    }
    return {
      index: i,
      rects,
      fontsize: fontSize,
      fontname: fontFamily
    }
  })
}

export function db2ds(db: any[], base: LayoutDimension): LayoutDataItem[] {
  return db.map(d => {
    if (d.is_clock) {
      const parmas = d.clock_params
      return {
        win: { x: 0, y: 0, w: -1, h: -1 },
        title: null,
        vol: null,
        timer: {
          x: Number(parmas.top.x / base.w),
          y: Number(parmas.top.y / base.h),
          fontSize: Number(parmas.fontsize / base.w),
          fontFamily: String(parmas.fontname),
          color: stringRGB(parmas.fontcolour),
          showDate: parmas.show_date === 1,
          dateNewLine: Number(parmas.date_new_line),
          dateDisplayType: Number(parmas.date_display_type),
          showFrame: Number(parmas.showframe),
          time24: Number(parmas["24/12"]),
        }
      }
    }
    const win = d.rects.find((r: any) => r.type === 0)
    const title = d.rects.find((r: any) => r.type === 1)
    const vol = d.rects.find((r: any) => r.type === 2)
    const fontSize = d.fontsize || round(base.w * defaultFontSize)
    const pos = win ? { x: win.x / base.w, y: win.y / base.h } : { x: 0, y: 0 }
    return {
      win: win ? {
        x: win.x / base.w,
        y: win.y / base.h,
        w: win.w / base.w,
        h: win.h / base.h,
      } : null,
      title: title ? {
        x: title.x / base.w - pos.x,
        y: title.y / base.h - pos.y,
        w: title.w / base.w,
        h: title.h / base.h,
        fontSize: fontSize / base.w,
        fontFamily: d.fontname
      } : null,
      vol: vol ? {
        x: vol.x / base.w - pos.x,
        y: vol.y / base.h - pos.y,
        one_w: vol.one_w / base.w,
        g: vol.g / base.w,
        h: vol.h / base.h,
        cnt: vol.count,
      } : null,
      timer: null
    } as LayoutDataItem
  })
}
