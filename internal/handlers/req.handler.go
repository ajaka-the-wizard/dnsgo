package handlers

import (
	"github.com/ajaka-the-wizard/dnsgo/internal/resolvers"
	"github.com/ajaka-the-wizard/dnsgo/internal/utils"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func ReqFactory(r *resolvers.Resolvers) dns.HandlerFunc {
	return dns.HandlerFunc(func(w dns.ResponseWriter, m *dns.Msg) {
		lw := utils.GetLogger(w)
		message := new(dns.Msg)
		message.SetReply(m)
		message.Compress = true

		for _, q := range m.Question {
			if rm, ok := r.ResolveZones(q); ok {
				message.Answer = append(message.Answer, rm)
			} else {
				message.SetRcode(m, dns.RcodeNameError)
			}
		}
		err := w.WriteMsg(message)
		if err != nil {
			lw.Error("Failed to write response", zap.Any("error", err))
		}
	})
}
