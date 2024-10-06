package Ping

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearLinux() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearWindows() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ConsoleClear() {
	switch runtime.GOOS {
	case "linux":
		clearLinux()
	case "windows":
		clearWindows()
	default:
		log.Fatalf("Unknow OS %s", runtime.GOOS)
	}
}

func Declare(p *Pinger) {
	for {
		time.Sleep(p.PingEvery)
		ConsoleClear()
		for _, s := range p.Pool {
			s.Lock()
			if s.Err == nil {
				fmt.Println(s.IP, s.Sent, s.Received, s.Percent)
			} else {
				fmt.Println(s.IP, s.Err.Error())
			}
			s.Unlock()
		}
	}
}
