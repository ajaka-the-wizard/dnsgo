package handlers

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func HandleReq(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = true

	for _, q := range r.Question {
		fmt.Printf("Received query: %s %s\n", q.Name, dns.TypeToString[q.Qclass])
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
			m.Answer = append(m.Answer, rr)
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
			m.Answer = append(m.Answer, rr)
		default:
			m.SetRcode(r, dns.RcodeNameError)
		}
	}
	err := w.WriteMsg(m)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
