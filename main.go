package main

import "net/http"
import "time"
import "fmt"

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://higher.team"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	fmt.Println("start...")
	go func() {
		if err := http.ListenAndServe("0.0.0.0:80", http.HandlerFunc(redirectTLS)); err != nil {
			fmt.Println(err)
		}
	}()
	for {
		time.Sleep(time.Second)
	}
}
