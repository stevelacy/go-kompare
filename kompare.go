package kompare

import (
  "reflect"
  "fmt"
  "github.com/mgutz/ansi"
)

type testing interface {
  Errorf(format string, args ...interface{})
}

func Fail (t testing, message string) bool {
  t.Errorf(ansi.Color(message, "red"))
  return false
}

func Success (t testing, message string) bool {
  fmt.Println(ansi.Color("✔ " + message, "green"))
  return true
}


func Type (t testing, arg1, arg2 interface{}, message string) bool {

  if reflect.TypeOf(arg1) != reflect.TypeOf(arg2) {
    return Fail(t, message)
  }

  return Success(t, message)

}

func Kind (t testing, arg1, arg2 interface{}, message string) bool {

  if reflect.ValueOf(arg1).Kind() != reflect.ValueOf(arg2).Kind() {
    return Fail(t, message)
  }

  return Success(t, message)
}
