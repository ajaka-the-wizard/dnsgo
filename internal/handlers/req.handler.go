package handlers

import (
	"net"

	"github.com/ajaka-the-wizard/dnsgo/internal/utils"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func HandleReq(w dns.ResponseWriter, m *dns.Msg) {
	lw := utils.GetLogger(w)
	message := new(dns.Msg)
	message.SetReply(m)
	message.Compress = true

	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			rr := &dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    300,
				},
				A: net.ParseIP("93.184.216.34"),
			}
			message.Answer = append(message.Answer, rr)
		case dns.TypeAAAA:
			rr := &dns.AAAA{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeAAAA,
					Class:  dns.ClassINET,
					Ttl:    300,
				},
				AAAA: net.ParseIP("2606:2800:220:1:248:189:25c8:1946"),
			}
			message.Answer = append(message.Answer, rr)
		default:
			message.SetRcode(m, dns.RcodeNameError)
		}
	}
	err := w.WriteMsg(message)
	if err != nil {
		lw.Info("Failed to write response: %v", zap.Any("error", err))
	}
}
