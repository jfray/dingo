package main

import (
	"fmt"
	"github.com/miekg/dns"
	"os"
	"time"
)

func lookup(domain string) (serial uint32, ttl uint32) {
	m := new(dns.Msg)
	m.RecursionDesired = true
	m.Question = make([]dns.Question, 1)
	
	m.Question[0] = dns.Question{dns.Fqdn(domain), 
			dns.TypeSOA, dns.ClassINET}
        
	c := new(dns.Client)
        in, _, _ := c.Exchange(m, "8.8.8.8:53")

	serial = in.Answer[0].(*dns.SOA).Serial
	ttl = in.Answer[0].(*dns.SOA).Minttl
	
	return serial, ttl
}

func checkRecord(hostRecord string) bool {
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s ZONE\n", os.Args[0])
		os.Exit(1)
	}

	var oldSerial uint32

	for {
		serial, ttl := lookup(os.Args[1])
		if (oldSerial > 0) && (serial > oldSerial) {
			fmt.Printf("Hey dude, we updated %d to %d", oldSerial, serial)
			host := fmt.Sprintf("www.%s", os.Args[1])
			if checkRecord(host) {
				fmt.Println("The new records are in")
			}
		}
        	fmt.Printf("My Serial Number is: %d & and its TTL is %d, previous serial was %d\n", serial, ttl, oldSerial)
		fmt.Printf("Sleeping %d seconds\n", ttl)
		oldSerial = serial
		time.Sleep(time.Duration(ttl) * time.Second)
	}
}
