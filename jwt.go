package jwt

import (
    b64 "encoding/base64"
    "log"
)

func EncodeB64(s string) string {
    return b64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeB64(s string) string {
    sD, err := b64.StdEncoding.DecodeString(s)
    if(err != nil){
        log.Fatal(err)
    }
    return string(sD)
}

