package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {

	//step through

	ret := romanTodecimal("ix")
	fmt.Println(ret)
	err1 := errors.New("math: square root of negative number")
	log.Printf("%+v", err1)

	//flags
	serverPort := flag.String("port", ":60137", "specify the port the server listens on")
	certFile := flag.String("certfile", "cmd/self-signed-cert/cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "cmd/self-signed-cert/key.pem", "key PEM file")
	flag.Parse()

	server, err := NewServer(*serverPort)
	if err != nil {
		panic(err)
	}
	//server.Start()

	server.StartTLS(*certFile, *keyFile)

	conf := Config()
	fmt.Println(conf)
	fmt.Println("ended")

}

//placeholder
func verifyUserPass(user string, pass string) bool {
	return true
}

//
func romanTodecimal(roman string) int {

	/*
		{"i", "1"},
			{"ii", "2"},
			{"iii", "3"},
			{"iv", "4"},
			{"v", "5"},
			{"vi", "6"},
			{"x", "10"},
	*/
	m := make(map[string]int)
	m["i"] = 1
	m["v"] = 5
	m["x"] = 10
	m["l"] = 50
	m["c"] = 100
	m["d"] = 500
	m["m"] = 1000

	intRet := 0
	
	if len(roman) < 2 {
		return m[roman]
	}
	/* my good stuff */
	last := 0
	cur := 0

	//for roman numerals count right to left.  if left is small then take away other wise add
	for i := len(roman) - 1; i >= 0; i-- {
		cur = m[string(roman[i])]
		if cur >= last {
			intRet += cur
		} else {
			intRet -= cur
		}
		last = cur

	}

	//sandeeps crap -- doesnt work
	/*
	   curVal := 0
	   nextVal := 0
	   for i := 0; i <= len(roman) - 1; i++ {
	   	curVal = m[string(roman[i])]
	   	nextVal = m[string(roman[i+1])]

	   	if curVal < nextVal {
	   		intRet =+ nextVal - curVal
	   		i++
	   	} else {
	   		intRet =+ curVal
	   	}
	   }
	*/
	return intRet
}
