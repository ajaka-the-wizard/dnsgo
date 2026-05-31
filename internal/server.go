package internal

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaka-the-wizard/dnsgo/internal/config"
	"github.com/ajaka-the-wizard/dnsgo/internal/handlers"
	"github.com/ajaka-the-wizard/dnsgo/internal/middlewares"
	"github.com/miekg/dns"
)

func Listen() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load env config: %v", err)
	}
	logger := config.InitializeLogger(env)
	defer logger.Sync()
	base := dns.HandlerFunc(handlers.HandleReq)
	server := &dns.Server{
		Addr:    env.ADDR,
		Net:     "udp",
		UDPSize: 4096,
	}
	server.Handler = middlewares.LoggerMiddleware(middlewares.LatencyCalculator(base))
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("\nShutting down DNS server...")
		if err = server.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
		os.Exit(0)
	}()
	log.Printf("DNS server listening on %s (UDP)\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Printf("Server stopped: %v", err)
	}
}
