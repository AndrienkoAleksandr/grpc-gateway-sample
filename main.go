package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"path"
)

const (
	tlsPath = "/etc/tls"
)

func main() {

	// Create a new instance of http.ServeMux
	mux := http.NewServeMux()

	// Register your routes
	mux.HandleFunc("/", helloHandler)

	// Create a new TLS configuration
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12, // Minimum TLS version to support
		// You can customize other TLS settings here, such as ciphers, curve preferences, etc.
	}

	// Create a new HTTP server with TLS configuration
	server := &http.Server{
		Addr:      ":443",     // Listen on port 443 (HTTPS default)
		Handler:   mux,        // Set the ServeMux as the server's handler
		TLSConfig: tlsConfig,  // Set the TLS configuration
	}

	// Start the server with TLS enabled
	log.Printf("Server listening on https://localhost%s\n", server.Addr)
	err := server.ListenAndServeTLS(path.Join(tlsPath, "tls.crt"), path.Join(tlsPath, "tls.key"))
	if err != nil {
		log.Fatal(err)
	}
}

// Handler function for the root route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, TLS!")
}
