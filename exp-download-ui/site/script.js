// Dynamically generate the HREF on the download file button.
function updateDownloadButton() {

    document.getElementById("download-button-java")
        .setAttribute("href", "/api/test-download/download.txt");

    document.getElementById("download-button-go")
        .setAttribute("href", "/go/test-download/download.txt");
}

window.addEventListener('DOMContentLoaded', () => updateDownloadButton());