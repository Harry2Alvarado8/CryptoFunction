package main

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"syscall"
)

func main() {
	input_key := "keySecret" //minimum three(3) characters
	input_data := "here you will find the text"

	//testing
	e := encode_p(input_key, input_data)

	fmt.Println("Encode:::: ", e)

	d := decoded_p("keySecret", e)

	fmt.Println("Decode:::: ", d)

}

func encode_p(keySecret string, str string) string {
	inbyte_ik, _ := syscall.ByteSliceFromString(keySecret)
	inbyte_id, _ := syscall.ByteSliceFromString(str)

	a := jwt.EncodeSegment(inbyte_ik)
	b := jwt.EncodeSegment(inbyte_id)

	var ec string
	for i := 0; i <= len(inbyte_ik); i += 3 {
		ec += b[i:i+3] + a[i:i+3]
	}
	ec += b[len(inbyte_ik)+2:]

	return ec
}

func decoded_p(keySecret string, encode_p string) string {
	inbyte_ik, _ := syscall.ByteSliceFromString(keySecret)

	a := jwt.EncodeSegment(inbyte_ik)

	for i := 0; i < (len(a)/3)*3; i += 3 {
		encode_p = strings.Replace(encode_p, a[i:i+3], "", 1)
	}
	de, _ := jwt.DecodeSegment(encode_p)

	return bytes.NewBuffer(de).String()
}
