package jwt

import (
    "testing"
    "fmt"
)

func TestJWT(t *testing.T){
    fmt.Println("Testing function is a go")
    sE := EncodeB64URL("Hello")
    fmt.Println(sE)
    sD := DecodeB64URL(sE)
    fmt.Println(sD)

}
