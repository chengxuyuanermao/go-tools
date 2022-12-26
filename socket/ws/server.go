package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Use() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
