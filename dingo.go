package main

import (
	"container/ring"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jfray/dingo/lookup"
)

func main() {
	debug := flag.Bool("debug", false, "Debug Mode?")
	ringSize := flag.Int("ringsize", 32, "How big should I make the ringbuffer?")
	fakeTtl := flag.Bool("fakettl", false, "Should I fake the ttl?")
	fakeTtlInterval := flag.Int("fakettl_interval", 10, "How long should the fake TTL be?")
	flag.Parse()
	
	if len(flag.Args()) == 0 {
		fmt.Printf("%s ZONE\n", flag.Arg(0))
		os.Exit(1)
	}


	r := ring.New(*ringSize)
	r.Value = 0
	r = r.Next()

	for {
		var actualTtl uint32

		serial, ttl := lookup.GetSOA(flag.Arg(0))

		if (*fakeTtl) {
			actualTtl = uint32(*fakeTtlInterval)
		} else {
			actualTtl = ttl
		}

		o := r.Prev()

		// This can't be right
		oi, _ := strconv.ParseUint(fmt.Sprintf("%d", o.Value), 0, 64)
		oldSerial := uint32(oi)
		
		r.Value = serial
		r = r.Next()

		r.Do(func(x interface{}) {
			fmt.Println(x)
		})

		if (oldSerial > 0) && (serial > oldSerial) {
			fmt.Printf("Hey dude, we updated %d to %d", oldSerial, serial)
			break
		}
		if (*debug) {
        		fmt.Printf("My Serial Number is: %d & and its TTL is %d, previous serial was %d\n", serial, actualTtl, oldSerial)
			fmt.Printf("Sleeping %d seconds\n", actualTtl)
		}
		if (*debug) {
			ticker := time.NewTicker(time.Second * 1)
		
			go func() {
        			for t := range ticker.C {
            				fmt.Println("Tick at", t)
        			}
    			}()
			
			time.Sleep(time.Duration(actualTtl) * time.Second)
			ticker.Stop()
		} else {
			time.Sleep(time.Duration(actualTtl) * time.Second)
		}
	}
}
