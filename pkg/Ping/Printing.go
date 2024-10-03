package Ping

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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
		ConsoleClear()
		for _, s := range p.Pool {
			if s.Err == nil {
				fmt.Println(s.IP.String(), s.Sent, s.Received, s.Percent)
			} else {
				fmt.Println(s.IP, s.Err.Error())
				os.Exit(1)
			}
		}
	}
}
