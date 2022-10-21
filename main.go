package main

import (
	"fmt"

	"go.mau.fi/whatsmeow"
)

func main() {
	id := whatsmeow.GenerateMessageID()
	fmt.Println(id)
}
