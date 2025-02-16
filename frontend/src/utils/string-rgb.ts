export default function stringRGB(colour: { R: number, r?: number, G: number, g?: number, B: number, b?: number, A: number, a?: number }) {
  return `rgb(${colour.r || colour.R || 0}, ${colour.g || colour.G || 0}, ${colour.b || colour.B || 0}, ${colour.a || colour.A || 1.0})`;
}
