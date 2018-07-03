package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

const title = `
   ▄▄▄▄    ██▀███   ▄▄▄       ██▓ ███▄    █ 
  ▓█████▄ ▓██ ▒ ██▒▒████▄    ▓██▒ ██ ▀█   █ 
  ▒██▒ ▄██▓██ ░▄█ ▒▒██  ▀█▄  ▒██▒▓██  ▀█ ██▒
  ▒██░█▀  ▒██▀▀█▄  ░██▄▄▄▄██ ░██░▓██▒  ▐▌██▒
  ░▓█  ▀█▓░██▓ ▒██▒ ▓█   ▓██▒░██░▒██░   ▓██░
  ░▒▓███▀▒░ ▒▓ ░▒▓░ ▒▒   ▓▒█░░▓  ░ ▒░   ▒ ▒ 
  ▒░▒   ░   ░▒ ░ ▒░  ▒   ▒▒ ░ ▒ ░░ ░░   ░ ▒░
   ░    ░   ░░   ░   ░   ▒    ▒ ░   ░   ░ ░ 
   ░         ░           ░  ░ ░           ░ 
        ░                                   
    ██████  ▄████▄   ▄▄▄       ███▄    █    
  ▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █    
  ░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒   
    ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒   
  ▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░   
  ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒    
  ░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░░ ░░   ░ ▒░   
  ░  ░  ░  ░          ░   ▒      ░   ░ ░    
        ░  ░ ░            ░  ░         ░    
           ░                                
                                made by wdhg

`

var wg = &sync.WaitGroup{}
var host = flag.String("h", "", "target")
var ms = flag.Int64("t", 500, "timeout in milliseconds")
var portLower = flag.Int("pl", 0, "lower port bound")
var portUpper = flag.Int("pu", 1000, "upper port bound")

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] host\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	fmt.Print(title)
	flag.Parse()
	if flag.NFlag() < 1 || *host == "" {
		usage()
		return
	}
}

// basic connect to port
func scanPort(address string, port int, wait time.Duration) {
	conn, err := net.DialTimeout("tcp", address+":"+strconv.Itoa(port), wait)
	if err == nil {
		fmt.Printf("%v/tcp		open\n", port)
		conn.Close()
	}
	wg.Done()
}

func main() {
	portRange := *portUpper - *portLower
	wg.Add(portRange)
	fmt.Printf("PORT		STATUS\n")
	for i := 0; i < portRange; i++ {
		go scanPort(
			*host,
			*portLower+i,
			time.Duration(*ms)*time.Millisecond,
		)
	}

	wg.Wait()
}
