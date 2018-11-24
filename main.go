package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Get hostid
	cmd, err := exec.Command("hostid").Output()
	hostid := ""
	if err != nil {
		fmt.Println("Warning: Unable to get hostid: " + err.Error())
		fmt.Println("Warning: Assuming hostid is: 00000000")
		hostid = "00000000"
	} else {
		hostid = strings.Trim(string(cmd), " \n\r\t")
	}

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Unable to get hostname: " + err.Error())
	}

	// Get and output license
	license, err := getLicense(hostid, hostname)
	if err != nil {
		log.Fatal("Unable to get license: " + err.Error())
	}

	fmt.Println("Add the following text to ~/.iourc:")
	fmt.Printf("[license]\n%s = %s;\n\n", hostname, license)
	fmt.Println("You can disable the phone home feature with something like: ")
	fmt.Println("echo '127.0.0.127 xml.cisco.com' >> /etc/hosts")
}

func getLicense(hostid string, hostname string) (license string, err error) {
	key, err := strconv.ParseInt(hostid, 16, 64)
	if err != nil {
		return "", errors.New("unable to parse hostid: " + err.Error())
	}

	for _, char := range hostname {
		key = int64(key) + int64(char)
	}

	magic1 := "4b582181567b0df321439b7eac1de68a"
	magic2 := "80000000000000000000000000000000000000000000000000000000000000000000000000000000"
	magicFinal := fmt.Sprintf("%s%s%08x%s", magic1, magic2, key, magic1)
	magicHex, err := hex.DecodeString(magicFinal)
	if err != nil {
		return "", errors.New("unable to create magic string: " + err.Error())
	}

	md5sum := md5.Sum(magicHex)
	return hex.EncodeToString(md5sum[:])[:16], nil
}
