package jwt

import (
    b64 "encoding/base64"
    "log"
)

func EncodeB64URL(s string) string {
    return b64.URLEncoding.EncodeToString([]byte(s))
}

func DecodeB64URL(s string) string {
    sD, err := b64.URLEncoding.DecodeString(s)
    if(err != nil){
        log.Fatal(err)
    }
    return string(sD)
}

