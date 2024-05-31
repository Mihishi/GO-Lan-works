package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Simple Web Server</title>
		</head>
		<body>
			<h1>Welcome to my simple web server!</h1>
			<p>This is a static HTML page served by a Go web server.</p>
		</body>
		</html>
		`
		fmt.Fprint(w, html)
	})

	
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
