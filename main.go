package main

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"syscall"
)

func main() {
	input_key := "keySecret" //minimum three(3) characters
	input_data := "here you will find the text"

	inbyte_ik, _ := syscall.ByteSliceFromString(input_key)
	inbyte_id, _ := syscall.ByteSliceFromString(input_data)

	a := jwt.EncodeSegment(inbyte_ik)
	b := jwt.EncodeSegment(inbyte_id)

	fmt.Println("Cifrado a: ", a)
	fmt.Println("Cifrado b: ", b)

	var c string
	for i := 0; i <= len(inbyte_ik); i += 3 {
		c += b[i:i+3] + a[i:i+3]
		fmt.Println(a[i : i+3])
	}
	c += b[len(inbyte_ik)+3:]

	fmt.Println(c)
	fmt.Print(input_data)

	//Testing output
	niceSecurity, _ := jwt.DecodeSegment(c)
	normalDecoder, _ := jwt.DecodeSegment(b)

	fmt.Println("niceSecurity: ", bytes.NewBuffer(niceSecurity).String())
	fmt.Println("normalDecoder: ", bytes.NewBuffer(normalDecoder).String())

}
