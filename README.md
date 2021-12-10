# Plantower

[![Tests](https://github.com/tomarrell/plantower/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/tomarrell/plantower/actions/workflows/test.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/tomarrell/plantower.svg)](https://pkg.go.dev/github.com/tomarrell/plantower)

This library allows you to decode data coming from the Plantower PMS5003 Digital
universal particle concentration sensor.

Currently supports models:
- PMS5003 (Active Mode)

<img src="sensor.jpg" width="50%"/>

## Usage

The library takes an `io.Reader` representing a stream of bytes from the sensor.
This allows you to decide how you read the data from the physical link.

```go
package main

import ()

func main() {
  // TODO
}
```
