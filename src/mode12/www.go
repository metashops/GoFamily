package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"serving:%s\n",r.URL.Path)
	fmt.Printf("served:%s\n",r.Host)
}
func TimeHandler(w http.ResponseWriter,c *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
}
func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number:",PORT)
	} else {
		PORT = ":" + arguments[1]
	}
	http.HandleFunc("/time",TimeHandler)
	http.HandleFunc("/",myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
