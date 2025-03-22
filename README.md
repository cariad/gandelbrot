# gandelbrot

![render.png](render.png)

A Golang package that renders the Mandelbrot Set.

## Usage

Install the package:

```bash
go get github.com/cariad/gandelbrot
```

In code, construct a `gandelbrot.RenderArgs` and pass it to `gandelbrot.Render()`:

```go
package main

import (
  "log"
  "os"

  "github.com/cariad/gandelbrot"
)

func main() {
  file, err := os.Create("render.png")
  if err != nil {
    log.Fatal(err)
  }

  gandelbrot.Render(&gandelbrot.RenderArgs{
    Real:         -2.5,
    Imaginary:    -2,
    ComplexWidth: 4.0,
    Writer:       file,
  })

  if err := file.Close(); err != nil {
    log.Fatal(err)
  }
}
```

See the `RenderArgs` docstrings for full usage details.

## The Author

Hello! 👋 I'm Cariad Eccleston. You can find me on [GitHub](https://github.com/cariad), [the Fediverse](https://queer.garden/@cariad) and [LinkedIn](https://www.linkedin.com/in/cariad/).
