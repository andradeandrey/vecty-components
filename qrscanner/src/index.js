import QrScanner from "../lib/qr-scanner.min.js";
import worker from "raw-loader!../lib/qr-scanner-worker.min.js";
var blob = window.URL.createObjectURL(new Blob([worker], { type: "application/javascript" }), {
    type: 'application/javascript; charset=utf-8'
});
QrScanner.WORKER_PATH = blob;
window.QrScanner = QrScanner;
