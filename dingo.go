package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jfray/dingo/lookup"
)

func main() {

	var oldSerial uint32
	debug := flag.Bool("debug", false, "Debug Mode?")
	flag.Parse()
	
	if len(flag.Args()) == 0 {
		fmt.Printf("%s ZONE\n", flag.Arg(0))
		os.Exit(1)
	}

	for {
		serial, ttl := lookup.GetSOA(flag.Arg(0))
		if (oldSerial > 0) && (serial > oldSerial) {
			fmt.Printf("Hey dude, we updated %d to %d", oldSerial, serial)
			break
		}
		if (*debug) {
        		fmt.Printf("My Serial Number is: %d & and its TTL is %d, previous serial was %d\n", serial, ttl, oldSerial)
			fmt.Printf("Sleeping %d seconds\n", ttl)
		}
		oldSerial = serial
		if (*debug) {
			ticker := time.NewTicker(time.Second * 1)
		
			go func() {
        			for t := range ticker.C {
            				fmt.Println("Tick at", t)
        			}
    			}()
			time.Sleep(time.Duration(ttl) * time.Second)
			ticker.Stop()
		} else {
			time.Sleep(time.Duration(ttl) * time.Second)
		}
	}
}
