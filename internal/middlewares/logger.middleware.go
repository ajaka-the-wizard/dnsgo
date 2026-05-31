package middlewares

import (
	"github.com/ajaka-the-wizard/dnsgo/internal/domain"
	"github.com/ajaka-the-wizard/dnsgo/internal/utils"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func LoggerMiddleware(next dns.Handler) dns.Handler {
	return dns.HandlerFunc(func(w dns.ResponseWriter, m *dns.Msg) {
		if len(m.Question) == 0 {
			next.ServeDNS(w, m)
			return
		}
		q := m.Question[0]
		reqLogger := zap.L().With(zap.Uint64("req_id", utils.GenerateID()), zap.String("qname", q.Name), zap.String("qtype", dns.TypeToString[q.Qtype]), zap.String("remote", w.RemoteAddr().String()), zap.String("proto", w.RemoteAddr().Network()))
		reqLogger.Info("Scoped logger initialized")
		writer := &domain.ContextWriters{
			ResponseWriter: w,
			Logger:         reqLogger,
		}
		next.ServeDNS(writer, m)
	})
}
