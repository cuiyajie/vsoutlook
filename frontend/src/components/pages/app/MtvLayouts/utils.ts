import { useUserSession } from "@src/stores/userSession";
import { CLockDisplayType, DateDisplayType, TallyType, AlarmLang } from "@src/utils/enums";
import parseRGB from "@src/utils/parse-rgb";
import stringRGB from '@src/utils/string-rgb';
import measureTextWidth from '@src/utils/measure-text-width';
import roundToNearestEven from "@src/utils/round-to-nearest-even";

export const defaultWinColor = 'rgb(255, 255, 255)'
export const draftWinBorder = (): LayoutWinBorder => {
  const bw = 4 / 1920
  return {
    top: bw,
    left: bw,
    right: bw,
    bottom: bw,
    color: defaultWinColor
  }
}

export const draftWin = (x: number, y: number, w: number, h: number): LayoutWin => {
  return {
    x,
    y,
    w,
    h,
    border: draftWinBorder(),
    showBorder: false
  }
}

export const draftVol = (h: number, bound: LayoutDimension): LayoutVol => {
  const { h: bh } = bound
  const sh = h / (3 * bh)
  const onew = 20 / 1920
  const g = 5 / 1920
  return {
    vols: new Array(4).fill(0).map((_, i) => ({
      start: i * 2 + 1,
      end: i * 2 + 2,
      x: i % 2 ? (1 - onew * 2 - g) : 0,
      y: i < 2 ? 0 : (1 - sh),
      one_w: onew,
      g,
      h: sh,
    })),
    alpha: 1.0,
    len: 1
  }
}

export const defaultTitleColor = {
  Text: 'rgba(200, 200, 200, 1.0)',
  Bg: 'rgba(100, 100, 100, 0.5)',
  Tally: 'rgba(100, 100, 100, 0.5)',
  TallyBorder: 'rgba(240, 240, 240, 1.0)',
}
export const draftTitle = (w: number, h: number, bound: LayoutDimension): LayoutTitle => {
  const nw = 192 / 1920
  const nh = 54 / 1080
  const tallyb = 8 / 1920
  const { w: bw, h: bh } = bound
  return {
    fontSize: 24 / 1920,
    fontFamily: getFontFamily(),
    color: defaultTitleColor.Text,
    bgColor: defaultTitleColor.Bg,
    x: (w - nw * bw) / (2 * bw),
    y: (h - nh * bh - 30) / bh,
    w: nw,
    h: nh,
    isVertial: false,
    tallyType: TallyType.Text,
    tallyBorderWidth: tallyb,
    tallyBorderColor: defaultTitleColor.TallyBorder,
    tallyBgColor: defaultTitleColor.Tally,
  }
}

export const defaultAlarmColor = {
  Text: 'rgba(255, 0, 0, 1.0)',
  Bg: 'rgba(100, 100, 100, 0.5)',
  Border: 'rgba(255, 0, 0, 1.0)',
}
export const draftAlarmBorder = (w: number, h: number, bound: LayoutDimension): LayoutAlarmBorder => {
  const bw = 4 / 1920
  const nw = 200 / 1920
  const nh = 100 / 1080
  return {
    x: (w - nw * bound.w) / bound.w / 2,
    y: (h - nh * bound.h) / bound.h / 2,
    w: nw,
    h: nh,
    width: bw,
    color: defaultAlarmColor.Border,
    interval: 1000
  }
}
export const draftAlarm = (w: number, h: number, bound: LayoutDimension): LayoutAlarm => {
  const nw = 160 / 1920
  const nh = 60 / 1080
  return {
    x: (w - nw * bound.w) / bound.w / 2,
    y: (h - nh * bound.h) / bound.h / 2,
    w: nw,
    h: nh,
    fontSize: 24 / 1920,
    bgColor: defaultAlarmColor.Bg,
    color: defaultAlarmColor.Text,
    border: draftAlarmBorder(w, h, bound),
    lang: AlarmLang.English,
  }
}

export const draftVector = (w: number, h: number, bound: LayoutDimension): LayoutArea => {
  const nw = 100 / 1920
  const nh = 100 / 1080
  return {
    x: (w - nw * bound.w - 60) / bound.w,
    y: (h - nh * bound.h - 60) / bound.h,
    w: nw,
    h: nh,
    alpha: 0.9
  }
}

export const draftOscillogram = (w: number, h: number, bound: LayoutDimension): LayoutArea => {
  const nw = 100 / 1920
  const nh = 200 / 1080
  return {
    x: 60 / bound.w,
    y: (h - nh * bound.h - 60) / bound.h,
    w: nw,
    h: nh,
    alpha: 0.9
  }
}

export const defaultMetaColor = {
  Text: 'rgba(200, 200, 200, 1.0)',
  Bg: 'rgba(100, 100, 100, 0.5)',
}
export const draftMeta = (w: number, h: number, bound: LayoutDimension): LayoutMeta => {
  const nw = 200 / 1920
  const nh = 100 / 1080
  return {
    x: (w - nw * bound.w - 60) / bound.w,
    y: 30 / bound.h,
    w: nw,
    h: nh,
    fontSize: 24 / 1920,
    bgColor: defaultMetaColor.Bg,
    color: defaultMetaColor.Text,
  }
}

export const defaultTimerColor = 'rgba(255, 0, 0, 1.0)'
export const draftTimer: () => LayoutTimer = () => {
  const fontFamily = getFontFamily()
  const fontSize = 64 / 1920
  const dateWidth = measureTextWidth('2021年01月01', fontFamily, `${fontSize}px`)
  return {
    fontSize,
    color: defaultTimerColor,
    x: 0,
    y: 100 / 1920,
    showDate: true,
    dateDisplayType: DateDisplayType.Chinese,
    clockDisplayType: CLockDisplayType.Digital,
    time24: 24,
    fontFamily,
    fontHeight: fontSize,
    fontGap: 10 / 1920,
    timeColor: defaultTimerColor,
    timePosition: {
      x: dateWidth + 12 / 1920,
      y: 100 / 1920
    }
  }
}

export const defaultTextColor = 'rgba(255, 255, 255, 1.0)'
export const defaultTextBgColor = 'rgba(255, 0, 0, 1.0)'
export const draftText: () => LayoutText = () => {
  const fontFamily = getFontFamily()
  const fontSize = 60 / 1920
  return {
    x: 640 / 1920,
    y: 120 / 1080,
    w: 640 / 1920,
    h: 120 / 1080,
    gap: 20 / 1920,
    rect1: {
      fontSize,
      fontFamily,
      color: defaultTextColor,
      bgColor: defaultTextBgColor,
    },
    rect2: {
      fontSize,
      fontFamily,
      color: defaultTextColor,
      bgColor: defaultTextBgColor,
    }
  }
}

export const resizeWinBorder = (win: LayoutWin, rw: number): LayoutWin => {
  return {
    ...win,
    border: {
      ...win.border,
      top: win.border.top * rw,
      left: win.border.left * rw,
      right: win.border.right * rw,
      bottom: win.border.bottom * rw,
    }
  }
}

export const resizeVol = (vol: LayoutVol, rw: number, rh: number): LayoutVol => {
  return {
    ...vol,
    vols: vol.vols.map(v => ({
      ...v,
      x: v.x * rw,
      y: v.y * rh,
      one_w: v.one_w * rw,
      g: v.g * rw,
      h: v.h * rh,
    }))
  }
}

export const resizeTitle = (title: LayoutTitle, rw: number, rh: number): LayoutTitle => {
  return {
    ...title,
    x: title.x * rw,
    y: title.y * rh,
    w: title.w * rw,
    h: title.h * rh,
    fontSize: title.fontSize * rw,
    tallyBorderWidth: title.tallyBorderWidth * rw,
  }
}

export const DefaultLayouts: [number, number][][] = [
  [[1, 1]],
  [[2, 2]],
  [[3, 3]],
  [[2, 1], [4, 2]],
  [[3, 4]],
  [[4, 4]],
  [[5, 5]],
]

export const defLayoutItem = (x: number, y: number, w: number, h: number) => ({
  win: draftWin(x, y, w, h),
  title: null,
  vol: null,
  timer: null,
  text: null,
  vector: null,
  oscillogram: null,
  alarm: null,
  meta: null,
} as LayoutDataItem)

export const defaultFontSize = 24 / 1920
export const defaultFontFamily = 'Noto Serif CJK'

export function getFontFamily() {
  return useUserSession().settings.mv_template_font || defaultFontFamily
}

const round = Math.round
export function ds2db(ds: LayoutDataItem[], base: LayoutDimension) {
  const fontFamily = getFontFamily()
  let delta = 0
  return ds.map((d, i) => {
    if (d.timer) {
      const sharedProperties: any = {
        line_time_top: {
          x: roundToNearestEven(d.timer.timePosition.x * base.w),
          y: roundToNearestEven(d.timer.timePosition.y * base.h),
        },
        line_date_top: {
          x: roundToNearestEven(d.timer.x * base.w),
          y: roundToNearestEven(d.timer.y * base.h),
        }
      }
      return [{
        index: i + delta,
        windows_type: 1,
        clock_params: {
          show_date: d.timer.showDate ? 1 : 0,
          display_style: d.timer.clockDisplayType === CLockDisplayType.LED ? 2 : d.timer.dateDisplayType,
          "24/12": d.timer.time24,
          style_01_params: {
            fontname: fontFamily,
            fontsize: round(d.timer.fontSize * base.w),
            fontcolour: parseRGB(d.timer.color),
            ...sharedProperties
          },
          style_2_params: {
            font_height: round(d.timer.fontHeight * base.w),
            font_gap: round(d.timer.fontGap * base.w),
            fontcolour: parseRGB(d.timer.timeColor),
            ...sharedProperties
          }
        }
      }]
    } else if (d.text) {
      delta++
      const w = roundToNearestEven((d.text.w - d.text.gap) * base.w / 2)
      const h = roundToNearestEven(d.text.h * base.h)
      return [{
        index: i,
        windows_type: 2,
        fontname: d.text.rect1.fontFamily,
        fontsize: round(d.text.rect1.fontSize * base.w),
        rects: [{
          type: 1,
          description: "text area",
          x: roundToNearestEven(d.text.x * base.w),
          y: roundToNearestEven(d.text.y * base.h),
          w,
          h,
          bg_color: parseRGB(d.text.rect1.bgColor),
          text_color: parseRGB(d.text.rect1.color),
        }]
      }, {
        index: i + delta,
        windows_type: 2,
        fontname: d.text.rect2.fontFamily,
        fontsize: round(d.text.rect2.fontSize * base.w),
        rects: [{
          type: 1,
          description: "text area",
          x: roundToNearestEven((d.text.x + d.text.gap) * base.w + w),
          y: roundToNearestEven(d.text.y * base.h),
          w,
          h,
          bg_color: parseRGB(d.text.rect2.bgColor),
          text_color: parseRGB(d.text.rect2.color),
        }]
      }]
    }
    const rects = []
    let fontSize = round(base.w * defaultFontSize)
    let fontIsVertial = false
    let tallytype = TallyType.Light
    let displayOscilloscope = false
    let scaleRatio = 1
    let lang = AlarmLang.English
    let pos: LayoutPos = { x: 0, y: 0 }
    if (d.win) {
      const x = round(d.win.x * base.w)
      const y = round(d.win.y * base.h)
      const w = round(d.win.w * base.w)
      const h = round(d.win.h * base.h)
      const wb = {} as any
      if (d.win.showBorder && d.win.border) {
        wb.window_border = {
          x: 0,
          y: 0,
          top_border_width: round(d.win.border.top * base.w),
          left_border_width: round(d.win.border.left * base.w),
          right_border_width: round(d.win.border.right * base.w),
          bottom_border_width: round(d.win.border.bottom * base.w),
          border_colour_1: parseRGB(d.win.border.color, true),
        }
      }
      if (d.title) {
        wb.tally_border = {
          border_width: round(d.title.tallyBorderWidth * base.w),
        }
      }
      if (d.alarm) {
        wb.alarm_border = {
          x: round(d.alarm.border.x * base.w),
          y: round(d.alarm.border.y * base.h),
          w: round(d.alarm.border.w * base.w),
          h: round(d.alarm.border.h * base.h),
          border_width: round(d.alarm.border.width * base.w),
          border_colour: parseRGB(d.alarm.border.color),
          interval_time_ms: d.alarm.border.interval,
        }
      }
      rects.push(Object.assign({
        type: 0,
				description: "video area",
        x,
        y,
        w,
        h,
      }, wb))
      pos = { x, y }
      scaleRatio = Math.min(round(w / base.w), 1)
    }
    if (d.title) {
      rects.push({
        type: 1,
				description: "text area",
        x: round(d.title.x * base.w) + pos.x,
        y: round(d.title.y * base.h) + pos.y,
        w: round(d.title.w * base.w),
        h: round(d.title.h * base.h),
        bg_color: parseRGB(d.title.bgColor),
        text_color: parseRGB(d.title.color),
        w_tally: round(d.title.h * base.h),
        tally_bg_color: parseRGB(d.title.tallyBgColor),
      })
      fontSize = round(d.title.fontSize * base.w)
      fontIsVertial = d.title.isVertial
      tallytype = d.title.tallyType
    }
    if (d.vol) {
      rects.push({
        type: 2,
        description: "audio vu area",
        vus: d.vol.vols.map((v: LayoutSingleVol, vi) => ({
          index: vi,
          channelstart: v.start,
          channelend: v.end,
          x: round(v.x * base.w) + pos.x,
          y: round(v.y * base.h) + pos.y,
          one_w: round(v.one_w * base.w),
          g: round(v.g * base.w),
          h: round(v.h * base.h),
          a: round(d.vol!.alpha * 10) / 10,
        })).slice(0, d.vol.len),
      })
    }
    if (d.alarm) {
      rects.push({
        type: 3,
        description: "alarm text area",
        x: round(d.alarm.x * base.w) + pos.x,
        y: round(d.alarm.y * base.h) + pos.y,
        w: round(d.alarm.w * base.w),
        h: round(d.alarm.h * base.h),
        bg_color: parseRGB(d.alarm.bgColor),
        text_color: parseRGB(d.alarm.color),
        font_size: round(d.alarm.fontSize * base.w),
      })
      lang = d.alarm.lang
    }
    if (d.oscillogram) {
      rects.push({
        type: 4,
        description: "oscillogram area",
        x: round(d.oscillogram.x * base.w) + pos.x,
        y: round(d.oscillogram.y * base.h) + pos.y,
        w: round(d.oscillogram.w * base.w),
        h: round(d.oscillogram.h * base.h),
        a: round(d.oscillogram.alpha * 100) / 100,
      })
      displayOscilloscope = true
    }
    if (d.vector) {
      rects.push({
        type: 5,
        description: "vector area",
        x: round(d.vector.x * base.w) + pos.x,
        y: round(d.vector.y * base.h) + pos.y,
        w: round(d.vector.w * base.w),
        h: round(d.vector.h * base.h),
        a: round(d.vector.alpha * 100) / 100,
      })
    }
    if (d.meta) {
      rects.push({
        type: 6,
        description: "meta info area",
        x: round(d.meta.x * base.w) + pos.x,
        y: round(d.meta.y * base.h) + pos.y,
        w: round(d.meta.w * base.w),
        h: round(d.meta.h * base.h),
        font_size: round(d.meta.fontSize * base.w),
        text_color: parseRGB(d.meta.color),
        bg_color: parseRGB(d.meta.bgColor),
      })
    }
    return [{
      index: i + delta,
      windows_type: 0,
      rects,
      fontsize: fontSize,
      language_ch: lang,
      fontname: fontFamily,
      font_is_vertical: fontIsVertial,
      display_oscilloscope: displayOscilloscope,
      scale_ratio: scaleRatio,
      tallytype,
    }]
  }).flat()
}

export function db2ds(db: any[], base: LayoutDimension, bounding: LayoutDimension): LayoutDataItem[] {
  let reservedText: any = null
  return db.map(d => {
    if (d.windows_type === 1) {
      const parmas = d.clock_params
      const timerParams = {
        showDate: parmas.show_date === 1,
        time24: Number(parmas['24/12']),
      } as LayoutTimer
      if (parmas.display_style === 2) {
        timerParams.clockDisplayType = CLockDisplayType.LED
        timerParams.dateDisplayType = DateDisplayType.Hyphen
      } else {
        timerParams.clockDisplayType = CLockDisplayType.Digital
        timerParams.dateDisplayType = parmas.display_style
      }
      const style2Params = parmas.style_2_params
      timerParams.fontHeight = Number(style2Params.font_height / base.w)
      timerParams.fontGap = Number(style2Params.font_gap / base.w)
      timerParams.timeColor = stringRGB(style2Params.fontcolour)
      const style1Params = parmas.style_01_params
      timerParams.fontSize = Number(style1Params.fontsize / base.w)
      timerParams.color = stringRGB(style1Params.fontcolour)
      timerParams.fontFamily = style1Params.fontname
      timerParams.x = Number(style1Params.line_date_top.x / base.w)
      timerParams.y = Number(style1Params.line_date_top.y / base.h)
      timerParams.timePosition = {
        x: Number(style1Params.line_time_top.x / base.w),
        y: Number(style1Params.line_time_top.y / base.h)
      }
      return {
        ...defLayoutItem(0, 0, -1, -1),
        timer: timerParams
      }
    } else if (d.windows_type === 2) {
      if (!reservedText) {
        reservedText = d
        return null
      } else {
        let text1
        let text2
        if (reservedText.rects[0].x < d.rects[0].x) {
          text1 = reservedText
          text2 = d
        } else {
          text1 = d
          text2 = reservedText
        }
        const rect1 = text1.rects[0]
        const rect2 = text2.rects[0]
        const textItem =  {
          ...defLayoutItem(0, 0, -1, -1),
          text: {
            x: rect1.x / base.w,
            y: rect1.y / base.h,
            w: (rect2.x - rect1.x + rect2.w) / base.w,
            h: rect1.h / base.h,
            gap: (rect2.x - rect1.x - rect1.w) / base.w,
            rect1: {
              fontSize: text1.fontsize / base.w,
              fontFamily: text1.fontname,
              color: stringRGB(rect1.text_color),
              bgColor: stringRGB(rect1.bg_color),
            },
            rect2: {
              fontSize: text2.fontsize / base.w,
              fontFamily: text2.fontname,
              color: stringRGB(rect2.text_color),
              bgColor: stringRGB(rect2.bg_color),
            }
          }
        } as LayoutDataItem
        reservedText = null
        return textItem
      }
    }
    d.rects = d.rects || []
    const win = d.rects.find((r: any) => r.type === 0)
    const title = d.rects.find((r: any) => r.type === 1)
    const vol = d.rects.find((r: any) => r.type === 2)
    const alarm = d.rects.find((r: any) => r.type === 3)
    const oscillogram = d.rects.find((r: any) => r.type === 4)
    const vector = d.rects.find((r: any) => r.type === 5)
    const meta = d.rects.find((r: any) => r.type === 6)
    const fontSize = (d.fontsize / base.w) || defaultFontSize
    const pos = win ? { x: win.x / base.w, y: win.y / base.h } : { x: 0, y: 0 }
    const result = defLayoutItem(0.25, 0.25, 0.5, 0.5)
    if (win) {
      result.win = {
        x: win.x / base.w,
        y: win.y / base.h,
        w: win.w / base.w,
        h: win.h / base.h,
        border: draftWinBorder(),
        showBorder: false
      }
      const wb = win.window_border
      if (wb) {
        result.win.border = {
          top: wb.top_border_width / base.w,
          left: wb.left_border_width / base.w,
          right: wb.right_border_width / base.w,
          bottom: wb.bottom_border_width / base.w,
          color: stringRGB(wb.border_colour_1)
        }
        result.win.showBorder = true
      }
    }
    if (vol) {
      const draft = draftVol(bounding.h * (result.win.h || 0.5), bounding)
      const vus = vol.vus || []
      result.vol = {
        vols: vus.map((v: any) => ({
          start: v.channelstart,
          end: v.channelend,
          x: v.x / base.w - pos.x,
          y: v.y / base.h - pos.y,
          one_w: v.one_w / base.w,
          g: v.g / base.w,
          h: v.h / base.h,
        })).concat(draft.vols.slice(vus.length - 4)),
        len: vus.length || 1,
        alpha: vus[0]?.a || 1.0,
      }
    }
    if (title) {
      const draft = draftTitle(bounding.w * (result.win.w || 0.5), bounding.h * (result.win.h || 0.5), bounding)
      result.title = {
        x: title.x / base.w - pos.x,
        y: title.y / base.h - pos.y,
        w: title.w / base.w,
        h: title.h / base.h,
        fontSize,
        fontFamily: d.fontname || getFontFamily(),
        color: title.text_color ? stringRGB(title.text_color) : defaultTitleColor.Text,
        bgColor: title.bg_color ? stringRGB(title.bg_color) : defaultTitleColor.Bg,
        isVertial: d.font_is_vertical || false,
        tallyType: d.tallytype || TallyType.Light,
        tallyBorderWidth: win?.tally_border?.border_width ? (win.tally_border.border_width / base.w) : draft.tallyBorderWidth,
        tallyBorderColor: defaultTitleColor.TallyBorder,
        tallyBgColor: title.tally_bg_color ? stringRGB(title.tally_bg_color) : defaultTitleColor.Tally,
      }
    }
    if (alarm) {
      result.alarm = {
        x: alarm.x / base.w - pos.x,
        y: alarm.y / base.h - pos.y,
        w: alarm.w / base.w,
        h: alarm.h / base.h,
        fontSize,
        color: alarm.text_color ? stringRGB(alarm.text_color) : defaultAlarmColor.Text,
        bgColor: alarm.bg_color ? stringRGB(alarm.bg_color) : defaultAlarmColor.Bg,
        lang: d.language_ch || AlarmLang.English,
        border: draftAlarmBorder(bounding.w * (result.win.w || 0.5), bounding.h * (result.win.h || 0.5), bounding),
      }
      const ab = win.alarm_border
      if (ab) {
        result.alarm.border = {
          x: ab.x / base.w,
          y: ab.y / base.h,
          w: ab.w / base.w,
          h: ab.h / base.h,
          width: ab.border_width / base.w,
          color: ab.border_colour ? stringRGB(ab.border_colour) : defaultAlarmColor.Border,
          interval: ab.interval_time_ms,
        }
      }
    }
    if (oscillogram) {
      result.oscillogram = {
        x: oscillogram.x / base.w - pos.x,
        y: oscillogram.y / base.h - pos.y,
        w: oscillogram.w / base.w,
        h: oscillogram.h / base.h,
        alpha: oscillogram.a || 0.9
      }
    }
    if (vector) {
      result.vector = {
        x: vector.x / base.w - pos.x,
        y: vector.y / base.h - pos.y,
        w: vector.w / base.w,
        h: vector.h / base.h,
        alpha: vector.a || 0.9
      }
    }
    if (meta) {
      result.meta = {
        x: meta.x / base.w - pos.x,
        y: meta.y / base.h - pos.y,
        w: meta.w / base.w,
        h: meta.h / base.h,
        fontSize: meta.font_size / base.w,
        bgColor: meta.bg_color ? stringRGB(meta.bg_color) : defaultMetaColor.Bg,
        color: meta.text_color ? stringRGB(meta.text_color) : defaultMetaColor.Text,
      }
    }
    return result
  }).filter<LayoutDataItem>(v => v)
}
