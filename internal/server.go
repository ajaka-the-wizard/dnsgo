package internal

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaka-the-wizard/dnsgo/internal/config"
	"github.com/ajaka-the-wizard/dnsgo/internal/customs"
	"github.com/ajaka-the-wizard/dnsgo/internal/handlers"
	"github.com/ajaka-the-wizard/dnsgo/internal/middlewares"
	"github.com/ajaka-the-wizard/dnsgo/internal/resolvers"
	"github.com/miekg/dns"
)

func Listen() {
	env := config.LoadEnv()
	z := config.LoadZones()
	r := config.LoadRoots()

	resolvers := resolvers.CreateResolver(z, r)
	gigantic := customs.InitializeGigantic()
	logger := config.InitializeLogger(env)
	defer func() { _ = logger.Sync() }()
	base := dns.HandlerFunc(handlers.ReqFactory(resolvers))
	server := &dns.Server{
		Addr:    env.ADDR,
		Net:     "udp",
		UDPSize: 4096,
	}
	server.Handler = middlewares.LoggerMiddleware(middlewares.LatencyCalculator(base), gigantic)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("\nShutting down DNS server...")
		if err := server.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
		os.Exit(0)
	}()
	log.Printf("DNS server listening on %s (UDP)\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server stopped: %v", err)
	}
}
