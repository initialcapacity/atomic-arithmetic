# Atomic Arithmetic

Inspired by [Tailwind CSS](https://tailwindcss.com/).

A utility-first arithmetic framework packed with functions like `Add23`, `Subtract8921`, `MultiplyBy5522`
and `DivideBy7` that can be composed to compute any number, directly in your Go code.
Atomic Arithmetic is unapologetically modern, and takes advantage of all the latest and greatest Go features
to make the developer experience as enjoyable as possible.

Build whatever you want, without touching pesky [arithmetic operators](https://go.dev/ref/spec#Arithmetic_operators).
Currently supports numbers up to 10,000, which should be enough for most applications.

## Usage

Import the operations package to use the atomic arithmetic functions.
Fast, built-in autocomplete in most IDEs!

```go
package main

import "github.com/initialcapacity/atomic-arithmetic/pkg/operations"

func main() {
	operations.Add365(17)
	operations.Subtract714(893)
	operations.MultiplyBy199(87)
	operations.DivideBy342(744)
	operations.Mod17(654)
}
```

Run the [demo](./cmd/demo/main.go) to see these examples in action.

```shell
go run ./cmd/demo
```

## Development

1.  Re-generate the operations package.
    ```shell
    go ./cmd/generate
    ```

1.  Run the extensive test suite.
    ```shell
    go test ./...
    ```
