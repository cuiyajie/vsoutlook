export default function newline(str: string, sep = ' ', replace = '<br>') {
  return str.split(sep).join(replace);
}
