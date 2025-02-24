export function isValidNumberRanges(input: string): boolean {
  // 分割字符串为多个子范围
  const ranges = input.split(',');

  for (const range of ranges) {
    const trimmed = range.trim();

    // 忽略空字符串
    if (trimmed === '') continue;

    // 分割范围为起始和结束数字
    const parts = trimmed.split('-');

    // 多个 '-' 或无效分割时返回 false
    if (parts.length > 2) return false;

    const start = parseInt(parts[0], 10);
    const end = parts.length === 2 ? parseInt(parts[1], 10) : start;

    // 验证数字有效性及顺序
    if (isNaN(start) || (parts.length === 2 && isNaN(end))) return false;
    if (parts.length === 2 && end < start) return false;
  }

  return true;
}

export function numberRangeToArray(input: string): number[] {
  // 原始转换逻辑（与之前一致）
  const ranges = input.split(',');
  const result: number[] = [];
  for (const range of ranges) {
    const trimmed = range.trim();
    if (trimmed.includes('-')) {
      const [start, end] = trimmed.split('-').map(Number);
      for (let num = start; num <= end; num++) {
        result.push(num);
      }
    } else {
      result.push(Number(trimmed));
    }
  }

  // 使用 Set 去重（自动去除重复项）
  const uniqueResult = [...new Set(result)];

  // 排序（Set 不保证顺序，需显式排序）
  return uniqueResult.sort((a, b) => a - b);
}
