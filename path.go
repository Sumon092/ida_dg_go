package main

import (
	"fmt"
	"net/http"
)

func Chain() {
	paths := []string{"/path0", "/path1", "/path2"}

	for _, path := range paths {
		http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, path)
		})
	}

	// Create a local variable inside the loop

	// http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, localPath) // Use the local variable
	// })

	http.ListenAndServe(":8080", nil)
}
