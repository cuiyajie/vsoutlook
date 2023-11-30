export default function downloadJsonFile(jsonData: any, fileName: string) {
  // Convert the JSON data to a Blob
  const blobData = new Blob([JSON.stringify(jsonData, null, 2)], { type: 'application/json' });

  // Create a download URL for the Blob
  const url = URL.createObjectURL(blobData);

  // Create a link element
  const link = document.createElement('a');
  link.href = url;
  link.download = fileName; // Specify the filename for the download

  // Append the link to the document body
  document.body.appendChild(link);

  // Trigger a click event to start the download
  link.click();

  // Clean up the created URL and remove the link element
  URL.revokeObjectURL(url);
  document.body.removeChild(link);
}
