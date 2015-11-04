package conf

import (
	"fmt"
	"github.com/miekg/dns"
)

func Get(resolvPath string) (server string) {
	conf, _ := dns.ClientConfigFromFile(resolvPath)
	nsAddressPort := fmt.Sprintf("%s:53", conf.Servers[0])
	return nsAddressPort 
}
