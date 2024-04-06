package nscan

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

func netScan(addr string) string {
	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			return "timeout"
		} else {
			return "refuse"
		}
	}
	_ = conn.Close()
	return "success"
}

func comsumer(ich <-chan res, och chan<- res) {
	for v := range ich {
		v.result = netScan(v.addr)
		och <- v
	}
}

type res struct {
	addr   string
	result string
}

func Run(cmd *cobra.Command, args []string) {
	ich := make(chan res, 1024)
	och := make(chan res, 1024)
	sch := make(chan struct{})

	for i := 0; i < 200; i++ {
		go func() {
			comsumer(ich, och)
		}()
	}

	fmt.Println("start")
	go func() {
		for end := 1; end < 256; end++ {
			ipStr := fmt.Sprintf("192.168.1.%d", end)
			for port := 6000; port < 10000; port++ {
				addr := fmt.Sprintf("%s:%d", ipStr, port)
				ich <- res{
					addr: addr,
				}
			}
		}
	}()

	var s []string
	go func() {
		for r := range och {
			fmt.Printf("addr:%s,res:%s\n", r.addr, r.result)
			if r.result == "success" {
				s = append(s, r.addr)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("loop,time:%s,%d,%d\n", time.Now().Format(time.DateTime), len(ich), len(och))
			if len(ich) == 0 && len(och) == 0 {
				sch <- struct{}{}
				break
			}
		}
	}()

	<-sch
	for {
		time.Sleep(time.Second)
		if len(ich) == 0 && len(och) == 0 {
			close(ich)
			close(och)
			break
		}
	}

	fmt.Println("end")
	for _, v := range s {
		fmt.Println("res:" + v)
	}
}
