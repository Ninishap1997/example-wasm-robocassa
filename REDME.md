# Robokassa WASM Payment

WebAssembly приложение на Golang для генерации ссылок оплаты через Robokassa с веб-сервером на Go. Совместимо с устройствами Redmi для удобного использования на смартфонах.

## Установка

1. Склонируйте репозиторий:
```bash
git clone https://github.com/yourusername/robokassa-wasm.git
```
2. Скомпилируйте WASM:
```bash
GOOS=js GOARCH=wasm go build -o static/main.wasm ./cmd/wasm
```
3. Скопируйте wasm_exec.js:
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static/
```
4. Скомпилируйте сервер:
```bash
go run ./cmd/server -port=8080
```