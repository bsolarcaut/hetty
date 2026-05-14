package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	// defaultAddr is the default address the HTTP proxy and admin interface listen on.
	// Changed from :8080 to 127.0.0.1:8080 to avoid binding on all interfaces by default.
	defaultAddr = "127.0.0.1:8080"
	// defaultAdminPath is the default path prefix for the admin interface.
	defaultAdminPath = "/hetty/"
	// defaultDbPath is the default path for the database file.
	// Using a local file by default instead of in-memory so sessions persist between restarts.
	defaultDbPath = "hetty.db"
)

// version is set at build time via ldflags.
var version = "dev"

func main() {
	// Parse command-line flags.
	addr := flag.String("addr", defaultAddr, "Address to listen on (e.g. \"127.0.0.1:8080\")")
	adminPath := flag.String("adminPath", defaultAdminPath, "Path prefix for the admin interface")
	dbPath := flag.String("db", defaultDbPath, "Path to the database file (default: hetty.db)")
	certFile := flag.String("cert", "", "Path to the CA certificate file (PEM format)")
	keyFile := flag.String("key", "", "Path to the CA private key file (PEM format)")
	printVersion := flag.Bool("version", false, "Print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: hetty [options]\n\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nHetty is an HTTP toolkit for security research.")
	}

	flag.Parse()

	if *printVersion {
		fmt.Printf("hetty %s\n", version)
		os.Exit(0)
	}

	log.Printf("[INFO] Starting hetty %s", version)
	log.Printf("[INFO] Listening on %s", *addr)

	// Log configuration details.
	if *dbPath != "" {
		log.Printf("[INFO] Using database at: %s", *dbPath)
	} else {
		log.Printf("[INFO] Using in-memory database")
	}

	if *certFile != "" && *keyFile != "" {
		log.Printf("[INFO] Using CA certificate: %s", *certFile)
	} else {
		log.Printf("[INFO] No CA certificate provided; HTTPS interception will use auto-generated certs")
	}

	log.Printf("[INFO] Admin interface available at http://%s%s", *addr, *adminPath)

	// TODO: Initialize proxy, admin interface, and database.
	// This will be wired up as additional packages are implemented.
	_ = adminPath
	_ = dbPath
	_ = certFile
	_ = keyFile

	log.Fatal("[ERROR] Server not yet implemented")
}
