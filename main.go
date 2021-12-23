package main

import (
	"flag"
	"fmt"
)

func main() {

	//flags
	serverPort := flag.String("server port", ":60666", "specify the port the server listens on")
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
