## Телелграмм-бот для накрукти подписчиков
name: Go

on:
  push:
  pull_request:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      # - name: Build
      #   run: go build -v ./...

      # - name: Test
      #   run: go test -v ./...

  fmt:
    name: Check go fmt
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      - name: Run go fmt
        run: |
          OUTPUT=$(gofmt -l .)
          if [ -n "$OUTPUT" ]; then
            echo "The following files are not formatted:"
            echo "$OUTPUT"
            exit 1
          fi

  vet:
    name: Run go vet
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      - name: Run go vet
        run: go vet ./...

  lint:
    name: Run golangci-lint
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - uses: golangci/golangci-lint-action@v3
        with:
          version: v1.55.2
