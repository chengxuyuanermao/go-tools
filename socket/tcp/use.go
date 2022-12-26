package tcp

import (
	"flag"
	"fmt"
)

func Use() {
	param := flag.String("type", "", "client or server")
	flag.Parse()
	fmt.Println(*param)
	if *param == "client" {
		client()
	} else {
		server()
	}

}
