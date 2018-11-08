# vecty-components

Component set for [Vecty](https://github.com/gopherjs/vecty)

## Supported environment

- GOOS=js GOARCH=wasm Go1.11~
- GopherJS

## Dependency

```sh
make deps
```

- go get github.com/gopherjs/gopherjs
- go get github.com/gowasm/gopherwasm
- go get github.com/gowasm/vecty
- go get github.com/gopherjs/vecty
- go get github.com/skip2/go-qrcode
- go get github.com/vincent-petithory/dataurl
- go get github.com/dave/wasmgo

## Sample

GopherJS version
```sh
make run
```

open http://localhost:8080


WebAssembly version
```sh
make wasm
```
(automatic open http://localhost:8080)

## Components

### Camera

- WebRTC getUserMedia

### QR Scanner

- nimiq/qr-scanner wrapper

### QR code

- QR code renderer

### Toggle

- iOS like toggle switch

### Spinner

- Simple SVG base spinner
