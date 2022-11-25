# waas
WebAssembly as a service

wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb

sudo dpkg -i tinygo_0.26.0_amd64.deb

go run main.go --locations=us-west1 --port-8081

go run main.go --locations=us-east1 --port-8082

go run cmd/test01/main.go