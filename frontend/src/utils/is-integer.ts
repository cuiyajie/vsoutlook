export default function isInteger(input: unknown): boolean {
  if (typeof input === 'number') {
    return Number.isInteger(input)
  } else if (typeof input === 'string') {
    const num = Number(input)
    return !isNaN(num) && Number.isInteger(num)
  }
  return false
}
