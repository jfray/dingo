package lookup

import (
	"fmt"
	"github.com/miekg/dns"
	"os"
)

func GetSOA(domain string) (serial uint32, ttl uint32) {
	conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if conf == nil {
		fmt.Printf("Cannot initialize the local resolver: %s\n", err)
		os.Exit(1)
	}

	nsAddressPort := fmt.Sprintf("%s:53", conf.Servers[0])

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
