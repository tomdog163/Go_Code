package day01

import (
	"fmt"
	"os"
	"os/exec"
)

func Example18() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
