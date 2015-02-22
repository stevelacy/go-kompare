# go-kompare
> Minimal go-lang type compare suite

[Documentation](http://godoc.org/github.com/stevelacy/go-kompare)

## Installation

```shell

$ go get github.com/stevelacy/go-kompare

```

## Example

```go

func TestTypeString(t *testing.T) {
  kompare.Type(t, "string", "test", "string should be string")
  kompare.Type(t, "true", true, "compare will fail")
}

```


#### Screenshot of this package's test:

![Testing](https://raw.githubusercontent.com/stevelacy/go-kompare/master/screenshot.png)
