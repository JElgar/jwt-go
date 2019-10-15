package jwt

import (
    b64 "encoding/base64"
    "log"
    "crypto/rsa"
    "fmt"
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

func GenerateSignature(headerAndPayload string) string{
   signer, err := loadPrivateKey("private.pem")
	if err != nil {
		fmt.Errorf("signer is damaged: %v", err)
	}

	signed, err := signer.Sign([]byte(headerAndPayload))
    fmt.Println(signed)
}

// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPublicKey(path string) (Unsigner, error) {
        data, err := ioutil.ReadFile(path)

        if err != nil {
                return nil, err
        }
        return parsePublicKey(data)
}

// parsePublicKey parses a PEM encoded private key.
func parsePrivateKey(pemBytes []byte) (Signer, error) {
        block, _ := pem.Decode(pemBytes)
        if block == nil {
                return nil, errors.New("ssh: no key found")
        }

        var rawkey interface{}
        switch block.Type {
        case "RSA PRIVATE KEY":
                rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
                if err != nil {
                        return nil, err
                }
                rawkey = rsa
        default:
                return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
        }
        return newSignerFromKey(rawkey)
}
