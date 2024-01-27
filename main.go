package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	host := "smtp.ethereal.email"
	port := "587"
	from := "chyna.fay@email"
	password := "8SunfxdTMVmcs62Y"

	toList := []string{"chyna@email"}
	msg := "Subject: email testing\r\n" + "\r\n" + "!!! Hi Golang!!!"
	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("sending email successful")
}
