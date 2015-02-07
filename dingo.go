package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jfray/dingo/lookup"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s ZONE\n", os.Args[0])
		os.Exit(1)
	}

	var oldSerial uint32

	for {
		serial, ttl := lookup.GetSOA(os.Args[1])
		if (oldSerial > 0) && (serial > oldSerial) {
			fmt.Printf("Hey dude, we updated %d to %d", oldSerial, serial)
			break
		}
        	fmt.Printf("My Serial Number is: %d & and its TTL is %d, previous serial was %d\n", serial, ttl, oldSerial)
		fmt.Printf("Sleeping %d seconds\n", ttl)
		oldSerial = serial
		time.Sleep(time.Duration(ttl) * time.Second)
	}
}
