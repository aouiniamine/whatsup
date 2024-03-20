package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ConnectBody struct {
	Credential string `json:"credential"`
}

func AuthRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/auth/connect", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request has came!!")
		decoder := json.NewDecoder(r.Body)
		var body ConnectBody
		err := decoder.Decode(&body)
		fmt.Println(body.Credential)
		if err != nil {

			panic(err)
		}
		jsonRes, err := json.Marshal(body)
		if err != nil {

			panic(err)
		}

		w.Write(jsonRes)
	}).Methods("POST")
	return r
}
