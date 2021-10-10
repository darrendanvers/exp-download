// Dynamically generate the HREF on the download file button.
function updateDownloadButton() {
    document.getElementById("download-button")
        .setAttribute("href", "/api/test-download/download.txt");
}

window.addEventListener('DOMContentLoaded', () => updateDownloadButton());