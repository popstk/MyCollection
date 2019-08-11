export PATH := $(GOPATH)/bin:$(PATH)
BIN=./bin
FLAGS=-mod=vendor

ifeq ($(OS),Windows_NT)
	EXT=.exe
else
	EXT=
endif

all: backend

backend:
	go build $(FLAGS) -o $(BIN)/backend$(EXT) ./cmd/backend

