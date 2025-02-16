export default function parseRGB(rgba: string, ignoreAlpha = false) {
    // Check if the input is a valid RGB string
    const regex = /^rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*([\d.]+))?\)$/;
    const result = regex.exec(rgba);

    if (!result) {
        throw new Error('Invalid RGB string format');
    }

    // Convert the captured string groups to integers
    const r = parseInt(result[1], 10);
    const g = parseInt(result[2], 10);
    const b = parseInt(result[3], 10);
    const a = result[4] ? parseFloat(result[4]) : 1.0;

    // Validate the range of each color component
    if (r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 || a < 0 || a > 1.0) {
        throw new Error('RGBA values must be in the valid range');
    }

    return ignoreAlpha ? { r, g, b } : { r, g, b, a };
}
