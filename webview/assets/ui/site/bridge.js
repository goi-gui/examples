// The bridge channel differs by engine: WebKit exposes
// window.webkit.messageHandlers.goi, WebView2 exposes window.chrome.webview.
// Every page here posts through this one helper.
const goiPost = window.chrome?.webview?.postMessage
  ? (message) => window.chrome.webview.postMessage(message)
  : (message) => window.webkit.messageHandlers.goi.postMessage(message);
