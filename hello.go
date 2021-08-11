package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
	"io/ioutil"
	"encoding/json"
)

func main() {
	b, err := ioutil.ReadFile("email.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    // fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    // fmt.Println(str) // print the content as a 'string'
/*
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`
*/
	r := strings.NewReader(str)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)


	json, err := json.MarshalIndent(header, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)

	/*
	received := header.Get("Received")
	out := strings.TrimLeft(strings.TrimRight(receivedstr[1],"\sby\s"),"from\s")
    fmt.Println(out)
*/

	// var addresses []string
	// addresses = header.Get("Received")
    fmt.Println(header["Received"][1])
	out := header["Received"][1]
	out1 := strings.TrimLeft(out,"from\t")
	out2 := strings.TrimRight(out1,"\t by \t")
    fmt.Println(out2)
	// fmt.Println(addresses[0].String())
}