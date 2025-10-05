package main

import (
	"fmt"
	"ntpTime/ntp"
	"os"
)

func main() {
	t, err := ntpTime.GetTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(t)
}
