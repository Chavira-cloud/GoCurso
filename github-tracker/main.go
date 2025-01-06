package main 

import (
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux" //mux es una dependencia de go
) 

func postHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Recived POST request!")
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request")
		return
	}

	fmt.Printf(string(body))

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello",postHandler).Methods("POST") //para saber que path vamos a declarar

	fmt.Println("Server listering on port 8080") //inicializar el puerto 8080 que va ser nuestro servidor
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}

 