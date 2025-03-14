# gandelbrot

A Golang experiment.

## Usage

```bash
go run .
```

This will:

1. Render the Mandelbrot Set at `render.png`.
1. Start the HTTP server on port 8080.

## API

- `GET /` returns `hello`.
- `GET /render` opens a websocket that doesn't do anything.
