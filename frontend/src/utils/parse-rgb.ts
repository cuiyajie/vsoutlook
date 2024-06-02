export default function parseRGB(rgb: string) {
    // Check if the input is a valid RGB string
    const regex = /^rgb\((\d{1,3}),\s*(\d{1,3}),\s*(\d{1,3})\)$/;
    const result = regex.exec(rgb);

    if (!result) {
        throw new Error('Invalid RGB string format');
    }

    // Convert the captured string groups to integers
    const R = parseInt(result[1], 10);
    const G = parseInt(result[2], 10);
    const B = parseInt(result[3], 10);

    // Validate the range of each color component
    if (R < 0 || R > 255 || G < 0 || G > 255 || B < 0 || B > 255) {
        throw new Error('RGB values must be in the range 0-255');
    }

    return { R, G, B };
}
