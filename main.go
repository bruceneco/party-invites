package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	port := v.Get("server.port").(string)

	sm := mux.NewRouter()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gopher!"))
	})
	log.Printf("[INFO] Server started at port :%s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), sm)
	if err != nil {
		log.Fatalln(err)
	}

}
