package jwt

import (
    "testing"
    "fmt"
)

func TestJWT(t *testing.T){
    fmt.Println("Testing function is a go")
    sE := EncodeB64("Hello")
    fmt.Println(sE)
    sD := DecodeB64(sE)
    fmt.Println(sD)

}
