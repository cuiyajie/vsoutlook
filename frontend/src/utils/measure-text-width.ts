// measure text width in dom, regard of font size, family, and text content, letter spacing, using dom to test
export default function measureTextWidth(
    text: string,
    fontFamily: string,
    fontSize: string,
    letterSpacing: string = 'normal'
): number {
    // 创建一个临时元素来测量文本宽度
    const tempElement = document.createElement('div');
    tempElement.style.position = 'absolute'; // 确保不影响布局
    tempElement.style.visibility = 'hidden'; // 隐藏元素
    tempElement.style.whiteSpace = 'nowrap'; // 防止文本换行
    tempElement.style.fontFamily = fontFamily; // 设置字体族
    tempElement.style.fontSize = fontSize; // 设置字号
    tempElement.style.letterSpacing = letterSpacing; // 设置字母间距
    tempElement.textContent = text; // 设置文本内容

    // 将元素添加到 DOM 中
    document.body.appendChild(tempElement);

    // 测量元素的宽度
    const width = tempElement.offsetWidth;

    // 从 DOM 中移除元素
    document.body.removeChild(tempElement);

    return width;
}
