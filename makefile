build:
	echo "Building for Linux..."
	@set GOOS=linux
	go build -o dist/civitdown ./main.go
	echo "Building for Windows..."
	@set GOOS=windows
	go build -o dist/civitdown.exe ./main.go