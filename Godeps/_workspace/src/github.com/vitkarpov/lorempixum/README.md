# Lorempixum

Lorempixum is a golang package for generating stub images with the given sizes.

## Install

```
$ go get github.com/gorilla/mux
```

## Overview

Let's take a look at simple usage example.

```go
package main

import (
    "github.com/vitkarpov/lorempixum"
    "os"
)

func main() {
    img := lorempixum.GetImage(100, 100)
    lorempixum.StreamImage(os.Stdout, img, "jpeg")
}
```

This tiny program writes a new jpeg image to stdout:

```
$ go build ./example/main.go && ./main > foo.jpeg
```

Lorempixum supports jpeg, png and gif for animated stubs.

## API

#### func  GetImage

```go
func GetImage(width, height int) *image.Paletted
```

#### func  StreamImage

```go
func StreamImage(out io.Writer, img *image.Paletted, format string) error
```
