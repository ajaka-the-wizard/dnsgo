package resolvers

import (
	"net"
	"strings"

	"github.com/ajaka-the-wizard/dnsgo/internal/config"
	"github.com/miekg/dns"
)

func (r *Resolvers) ResolveZones(q dns.Question) (dns.RR, bool) {
	records := r.getData(q.Name, q.Qtype)
	if len(records) == 0 {
		return nil, false
	}
	var answer dns.RR
	for _, record := range records {
		rr := makeRR(q.Name, record)
		if rr == nil {
			continue
		}
		answer = rr
		break
	}
	if answer == nil {
		return nil, false
	}
	return answer, true
}

func (r *Resolvers) getData(name string, t uint16) []config.Record {
	if r == nil || r.zones == nil {
		return nil
	}

	name = dns.Fqdn(strings.ToLower(name))
	records := make([]config.Record, 0)

	for zone, zoneRecords := range r.zones.AllZones {
		if !dns.IsSubDomain(dns.Fqdn(strings.ToLower(zone)), name) {
			continue
		}

		for _, record := range zoneRecords {
			recordName := dns.Fqdn(strings.ToLower(record.Name))
			recordType, ok := dns.StringToType[strings.ToUpper(record.Type)]
			if !ok || recordName != name {
				continue
			}

			if t == dns.TypeANY || t == recordType {
				records = append(records, record)
			}
		}
	}

	return records
}

func makeRR(name string, record config.Record) dns.RR {
	ttl := uint32(record.TTL)
	hdr := dns.RR_Header{
		Name:   dns.Fqdn(name),
		Rrtype: dns.StringToType[strings.ToUpper(record.Type)],
		Class:  dns.ClassINET,
		Ttl:    ttl,
	}

	switch strings.ToUpper(record.Type) {
	case "A":
		ip := net.ParseIP(record.Value).To4()
		if ip == nil {
			return nil
		}
		return &dns.A{
			Hdr: hdr,
			A:   ip,
		}
	case "AAAA":
		ip := net.ParseIP(record.Value).To16()
		if ip == nil || ip.To4() != nil {
			return nil
		}
		return &dns.AAAA{
			Hdr:  hdr,
			AAAA: ip,
		}
	case "CNAME":
		return &dns.CNAME{
			Hdr:    hdr,
			Target: dns.Fqdn(record.Value),
		}
	case "NS":
		return &dns.NS{
			Hdr: hdr,
			Ns:  dns.Fqdn(record.Value),
		}
	case "TXT":
		return &dns.TXT{
			Hdr: hdr,
			Txt: []string{record.Value},
		}
	default:
		return nil
	}
}
