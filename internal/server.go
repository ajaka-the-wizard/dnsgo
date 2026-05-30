package internal

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaka-the-wizard/dnsgo/internal/handlers"
	"github.com/miekg/dns"
)

func Listen() {
	server := &dns.Server{
		Addr:    ":5353",
		Net:     "udp",
		UDPSize: 4096,
		Handler: dns.HandlerFunc(handlers.HandleReq),
	}
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		fmt.Println("\nShutting down DNS server...")
		server.Shutdown()
		os.Exit(0)
	}()
	fmt.Printf("DNS server listening on %s (UDP)\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
