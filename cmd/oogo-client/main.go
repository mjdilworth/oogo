package main

import (
	"fmt"

	"github.com/mjdilworth/oogo"
	"github.com/mjdilworth/oogo/clientlib"
)

func main() {
	fmt.Println(oogo.Config())
	fmt.Println(clientlib.Hello())
}
