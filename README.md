# zeroparse
A very simple tool that you can pipe JSON [zerolog] output into and get pretty
formatted data out. This uses zerolog's [ConsoleWriter] for the formatting, so
the output is exactly the same as using it directly, but without the in-process
overhead.

[zerolog]: https://github.com/rs/zerolog
[ConsoleWriter]: https://pkg.go.dev/github.com/rs/zerolog#readme-pretty-logging

## Usage
```
go install go.mau.fi/zeroparse@latest
tail -f zero.log | zeroparse
```
