APP_NAME = tree
SRC      = cmd/app/main.go
ICON     = assets/icon.png

# Default target: build and run
run: build
	./$(APP_NAME)

# Build binary for current OS
build:
	go build -o $(APP_NAME) $(SRC)

# Clean build artifacts
clean:
	rm -f $(APP_NAME)
	rm -rf $(APP_NAME).app

# Format code
fmt:
	go fmt ./...

# Check for issues
lint:
	go vet ./...

# Build macOS .app bundle using fyne package
package:
	GOOS=darwin CGO_ENABLED=1 \
		go build -o $(APP_NAME).app/Contents/MacOS/$(APP_NAME) \
		$(SRC)
	fyne package -os darwin -icon $(ICON) --name $(APP_NAME) --app-id com.example.treeapp --src cmd/app

# Run with fyne if installed
fyne-run:
	fyne run $(SRC)
