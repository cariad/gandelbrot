# gandelbrot

![render.png](render.png)

A Golang module that renders the Mandelbrot Set.

## Usage

Install the module:

```bash
go install github.com/cariad/gandelbrot
```

In code, construct a `RenderArgs` and pass it to `Render()`:

```go
package main

import (
  "log"
  "os"
)

func main() {
  file, err := os.Create("render.png")
  if err != nil {
    log.Fatal(err)
  }

  Render(&RenderArgs{
    real:         -2.5,
    imaginary:    -2,
    complexWidth: 4.0,
    writer:       file,
  })

  if err := file.Close(); err != nil {
    log.Fatal(err)
  }
}
```

See the `RenderArgs` docstrings for full usage details.

## The Author

Hello! 👋 I'm Cariad Eccleston. You can find me on [GitHub](https://github.com/cariad), [the Fediverse](https://queer.garden/@cariad) and [LinkedIn](https://www.linkedin.com/in/cariad/).
