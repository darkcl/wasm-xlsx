.PHONY: start serve build

start: build serve

serve:
	@echo 'Start wasm-xlsx server'
	@go run server.go

build:
	@echo 'Build wasm-xlsx'
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./main.go
	@cp ./page.html ./build/index.html
	@cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./build/wasm_exec.js
