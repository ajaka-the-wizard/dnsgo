package middlewares

import (
	"time"

	"github.com/ajaka-the-wizard/dnsgo/internal/utils"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func LatencyCalculator(next dns.Handler) dns.Handler {
	return dns.HandlerFunc(func(w dns.ResponseWriter, m *dns.Msg) {
		lw := utils.GetLogger(w)
		lw.Info("query started")

		start := time.Now()
		next.ServeDNS(w, m)

		latencyUs := time.Since(start).Microseconds()

		lw.With(zap.Int64("latency_us", latencyUs)).Info("Query completed")
	})
}
