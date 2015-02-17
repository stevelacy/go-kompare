# go-kompare
> Minimal go-lang type compare suite

[Documentation](http://godoc.org/github.com/stevelacy/go-kompare)

## Installation

```shell

$ go get github.com/stevelacy/go-kompare

```

## Example

```go

kompare.Type("string", "other string", "this should not fail")

kompare.Type("string", true, "this will fail")


```


#### Screenshot of this package's test:

![Testing](https://raw.githubusercontent.com/stevelacy/go-kompare/master/screenshot.png)
