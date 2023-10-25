# Go variables
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install

# Project name and source file
BINARY_NAME = myapp
MAIN_FILE = main.go

# Output directory
OUT_DIR = bin

all: build

build:
	$(GOBUILD) -o $(OUT_DIR)/$(BINARY_NAME) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -f $(OUT_DIR)/$(BINARY_NAME)

install:
	$(GOINSTALL)

.PHONY: all build clean install
