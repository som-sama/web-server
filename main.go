package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//adding a new multiplexer to route incoming http requests

	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	//adding a port
	const addr = ":8000"

	//adding a server
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server started on port ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = `<html>
		<body>
		<p>Khati is a papi	</p>
		</body>
		</html>
		`

	w.WriteHeader(200)
	w.Write([]byte(page))

}
