build:
	@echo "----------"
	@echo "Building Go wasm Guest"
	@echo "----------"
	go run karmem.org/cmd/karmem build --golang -o "waaskm" ../../waas.km
	tinygo build -o hello.wasm -scheduler=none --no-debug -target=wasi main.go