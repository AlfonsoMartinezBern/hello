package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift for Developers!  YUhuhuhuhuh -> KKK de la vaca"
	}
	fmt.Fprintln(w, response)
        currentTime := time.Now()	
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Servicing an impatient beginner's request.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
