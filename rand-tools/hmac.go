package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

// HMAC: Hashed based message authentication

// 1. value and hash generated by server with secret key is stored on user's side
// 2. when accessing server we send it in a cookie
// 3. server takes value and hash and then is generating new hash from sended value with secret key
// 4. if new hash generated from value and sended hash are the same then it's OK.

func runHmac() {
	h1 := getCode("test@exmaple.com")
	h2 := getCode("test@exampl.com")

	fmt.Print("first: " + h1 + "\n" + "second: " + h2 + "\n")
	fmt.Println("are the same? ", h1 == h2)

}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
