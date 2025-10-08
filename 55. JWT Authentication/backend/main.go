package main

import (
	"ecommerce/util"
	"fmt"
)

func main() {
	//cmd.Serve()

	// var s string

	// s = "hello world"
	// bytearr := []byte(s)

	// fmt.Println(s)
	// fmt.Println(bytearr)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(bytearr)
	// fmt.Println(b64Str)

	// data := []byte("hello World")

	// hash := sha256.Sum256(data)

	// fmt.Println("Hash after SHA-256 : ", hash)

	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         45,
		FirstName:   "Rakibul",
		LastName:    "Azad",
		Email:       "rakibul.azad@gmail.com",
		IsShopOwner: false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jwt)
}
