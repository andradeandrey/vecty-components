GOPATH:=$(shell cd ../../../..&&pwd):$(GOPATH)
export GOPATH

run:
	gopherjs serve github.com/nobonobo/vecty-components

wasm:
	wasmgo serve github.com/nobonobo/vecty-components

deps:
	go get github.com/gopherjs/gopherjs
	go get github.com/gowasm/gopherwasm
	go get github.com/gowasm/vecty
	go get github.com/gopherjs/vecty
	go get github.com/skip2/go-qrcode
	go get github.com/vincent-petithory/dataurl
	go get github.com/dave/wasmgo
