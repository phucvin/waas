build:
	# tinygo build -ldflags="-X 'main.managedScopesFlag=us-west1'" -o hello-us-west1.wasm -scheduler=none --no-debug -target=wasi main.go
	tinygo build -ldflags="-X 'main.managedScopesFlag=us-east1'" -o hello-us-east1.wasm -scheduler=none --no-debug -target=wasi main.go
	# tinygo build -ldflags="-X 'main.managedScopesFlag=us-west1,us-east1'" -o hello-us.wasm -scheduler=none --no-debug -target=wasi main.go