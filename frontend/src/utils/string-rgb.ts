export default function stringRGB(colour: { R: number, G: number, B: number }) {
  return `rgb(${colour.R || 0}, ${colour.G || 0}, ${colour.B || 0})`;
}
