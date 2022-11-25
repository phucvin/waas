build:
	go run karmem.org/cmd/karmem build --golang -o "waaskm" ../../waas.km
	tinygo build -o capitalize.wasm -scheduler=none --no-debug -target=wasi main.go