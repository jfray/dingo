package lookup

import (
	"fmt"
	
	"github.com/miekg/dns"
	"github.com/jfray/dingo/conf"
)

var (
  nsAddressPort string = conf.Get("/etc/resolv.conf")
)

func SOA(domain string) (serial uint32, ttl uint32) {
	m := new(dns.Msg)
	m.RecursionDesired = true
	m.Question = make([]dns.Question, 1)
	
	m.Question[0] = dns.Question{dns.Fqdn(domain), 
			dns.TypeSOA, dns.ClassINET}
        
	c := new(dns.Client)
        in, _, _ := c.Exchange(m, nsAddressPort)

	serial = in.Answer[0].(*dns.SOA).Serial
	ttl = in.Answer[0].(*dns.SOA).Minttl
	
	return serial, ttl
}

func getKey(domain string) *dns.DNSKEY {
	//var keytag uint16
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeDNSKEY)
	m.SetEdns0(4096, true)
	r, _, err := c.Exchange(m, nsAddressPort)
	if err != nil {
		return nil
	}
	for _, k := range r.Answer {
		if k1, ok := k.(*dns.DNSKEY); ok {
//			if k1.KeyTag() == keytag {
			return k1
//			}
		}
	}
	return nil
}

func RRSIG(domain string) (values string) {
	ret := fmt.Sprintf("Hey: %v", getKey(domain))
	return ret
}
