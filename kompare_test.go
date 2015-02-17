package kompare_test

import (
  "testing"
  "."
)



func TestTypeString(t *testing.T) {
  kompare.Type(t, "string", "test", "string should be string")
  kompare.Type(t, " ", "--", "string should be string")
}

func TestTypeInt(t *testing.T) {
  kompare.Type(t, 1, 22222, "int should be int")
  kompare.Type(t, 1, 0, "int should be int")
}

func TestTypeBool(t *testing.T) {
  kompare.Type(t, true, false, "boolean should be boolean")
  kompare.Type(t, false, false, "boolean should be boolean")
}


func TestKind(t *testing.T) {

  kompare.Kind(t, 3, 8, "type should be type")
  kompare.Kind(t, 11, 0, "type should be type")
}
