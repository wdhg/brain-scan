package main

import (
	"flag"
	"fmt"
	"net"
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
var target = flag.String("t", "", "target")
var timeout = flag.Int64("d", 500, "timeout delay in milliseconds")
var portLower = flag.Int("pl", 0, "lower port bound")
var portUpper = flag.Int("pu", 1000, "upper port bound")

func init() {
	fmt.Print(title)
	flag.Parse()
	if flag.NFlag() < 1 || *target == "" {
		flag.Usage()
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
	for i := 0; i <= portRange; i++ {
		go scanPort(
			*target,
			*portLower+i,
			time.Duration(*timeout)*time.Millisecond,
		)
	}

	wg.Wait()
}
