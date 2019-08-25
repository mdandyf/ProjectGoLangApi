package functions

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var port string

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage! \n")
	fmt.Fprintf(w, "This is Homepage from GoLang on port 20000")
	fmt.Println("Endpoint Hit: /homePage")
}

func getJsonPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ReadJsonString())
	fmt.Println("Endpoint Hit: /getJson")
}

func HandleRequest() {
	port := strconv.Itoa(20000)
	fmt.Println("Start Server on port " + port)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/getJson", getJsonPage)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
