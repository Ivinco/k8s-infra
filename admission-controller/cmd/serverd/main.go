package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"admissioncontroller/http"

	log "k8s.io/klog/v2"
)

var (
	tlscert, tlskey, port string
)

func main() {

	tlscert = getEnv("TLS_CERT_PATH", "/etc/certs/tls.crt")
	tlskey = getEnv("TLS_KEY_PATH", "/etc/certs/tls.key")
	port = getEnv("SERVER_PORT", "8443")

	flag.StringVar(&tlscert, "tlscert", tlscert, "Path to the TLS certificate")
	flag.StringVar(&tlskey, "tlskey", tlskey, "Path to the TLS key")
	flag.StringVar(&port, "port", port, "The port on which to listen")
	flag.Parse()

	server := http.NewServer(port)

	go func() {
		// listen shutdown signal
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		sig := <-signalChan
		log.Errorf("Received %s signal; shutting down...", sig)
		if err := server.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
	}()

	log.Infof("Starting server on port: %s", port)
	if err := server.ListenAndServeTLS(tlscert, tlskey); err != nil {
		log.Errorf("Failed to listen and serve: %v", err)
		os.Exit(1)
	}
}

// getEnv gets an environment variable by name and if it doesn't exist, returns a default value
func getEnv(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}
