package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func main() {

	// runtime.GOOS = OS name  (set at compile time)

	hostFile := "/etc/hosts"

	if runtime.GOOS == "windows" {
		hostFile = os.ExpandEnv("${SystemRoot}/System32/drivers/etc/hosts")
	}

	// Check that 'myhost' exists
	contents, err := ioutil.ReadFile(hostFile)
	if err != nil {
		panic(fmt.Errorf("Unable to open host file '%s': %v", hostFile, err))
	}

	if !strings.Contains(string(contents), "myhost") {
		fmt.Printf("Please define 'myhost' in %s\n", hostFile)
	}
}
