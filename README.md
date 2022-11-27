# waas
WebAssembly as a service

wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb

sudo dpkg -i tinygo_0.26.0_amd64.deb

npm i -g assemblyscript

go run main.go --locations=us-west1 --port=8081

go run main.go --locations=us-east1 --port=8082

go run cmd/test01/main.go

go run cmd/test02/main.go --n=100

go build -o bin/test02 cmd/test02/main.go

for run in {1..10000}; do bin/test02 --n=100; done

Notes:
* WebAssembly Component Model (https://github.com/WebAssembly/component-model) probably has the same goal.
