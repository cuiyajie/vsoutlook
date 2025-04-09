export default function randomStr(
  length = 6,
  chars = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
) {
  let str = ''
  const clen = chars.length
  for (let i = length; i >= 0; i--) {
    str += chars[Math.floor(Math.random() * clen)]
  }
  return str
}
