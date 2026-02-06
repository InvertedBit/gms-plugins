PLUGIN_NAME=$(shell grep '"name":' manifest.json | cut -d'"' -f4)
VERSION=$(shell grep '"version":' manifest.json | cut -d'"' -f4)
OUTPUT=$(PLUGIN_NAME)-$(VERSION).zip

.PHONY: build test clean checksum sign all

all: test build checksum

test:
	@echo "Running tests..."
	@go test ./... -v
	@if [ $$? -ne 0 ]; then \
		echo "Tests failed! Fix errors before building."; \
		exit 1; \
	fi
	@echo "All tests passed!"

build: manifest.json test
	@echo "Building plugin $(PLUGIN_NAME) v$(VERSION)..."
	@go build -buildmode=plugin -o plugin.so .
	@zip $(OUTPUT) manifest.json plugin.so
	@if [ -d "assets" ]; then zip -r $(OUTPUT) assets/; fi
	@echo "Plugin packaged: $(OUTPUT)"
	@rm -f plugin.so

# Generate SHA-256 checksum (required)
checksum:
	@sha256sum $(OUTPUT) > $(OUTPUT).sha256
	@echo "Checksum created: $(OUTPUT).sha256"

# Generate GPG signature (optional - requires GPG key configured)
sign: checksum
	@gpg --armor --detach-sign $(OUTPUT)
	@echo "Signature created: $(OUTPUT).asc"

clean:
	rm -f $(OUTPUT) $(OUTPUT).sha256 $(OUTPUT).asc plugin.so
