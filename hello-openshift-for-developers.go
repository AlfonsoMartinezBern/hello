package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
        "net/url"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	
        currentTime := time.Now()	
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Chispas 123!"
	}
	fmt.Fprintln(w, "--> ", "Prueba servicio con Go 1.19")
	fmt.Fprintln(w, "--> ", response)
	fmt.Fprintln(w, "--> ", currentTime.For1mat("2006-01-02 15:04:05"))
    	testLoop:for val := 1; val < 7; val++ {
        	fmt.Printf("%d", val)
        	switch {
        	case val == 1:
            		fmt.Fprintln(w, ".... %d -> Start"  , val)
        	case val == 5:
            		fmt.Fprintln(w, ".... %d -> Break"  , val)
            		break testLoop
        	case val > 2:
            		fmt.Fprintln(w, ".... %d -> Running", val)
            		break 
        	default:
            		fmt.Fprintln(w, ".... %d -> Progress", val)
        }
    }	
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
