package main

import (
	"github.com/miekg/dns"
	"log"
)

// replies a TXT record containing server name and version
func handleTXTVersionRequest(this *dnsRequestHandler, r, msg *dns.Msg) {
	log.Printf("TXT %s\n", msg.Question[0].Name)

	msg.Answer = append(msg.Answer, &dns.TXT{
		Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: r.Question[0].Qtype, Class: r.Question[0].Qclass, Ttl: DNSDefaultTTL},
		Txt: []string{"bind-⑨"},
	})
}
