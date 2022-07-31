package main

import (
	"fmt"

	continuum "github.com/realForbis/continuum-go-sdk"
)

func main() {
	client, err := continuum.NewClient("http://127.0.0.1:19735")
	//client, err := continuum.NewClient("ws://127.0.0.1:19736")
	if err != nil || client == nil {
		fmt.Println(err)
		return
	}

	addr, err := client.Ledger.Accounts(20, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for idx, val := range addr {
		fmt.Println(idx, "==>", val.String())
	}
}
