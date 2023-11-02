package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "run":
		if err := run(); err != nil {
			slog.Error(err.Error())
		}
	case "child":
		child()
	default:
		slog.Error("unexpected argument")
	}

}

func run() error {
	fmt.Println(os.Args[2], os.Args[3:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	return cmd.Run()
}

func child() error {
	syscall.Sethostname([]byte("devfest"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
