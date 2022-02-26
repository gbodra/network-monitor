package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

func getLocalIp() []string {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)

	}

	defer conn.Close()
	ip := conn.LocalAddr().(*net.UDPAddr).String()
	s := strings.Split(ip, ":")
	octets := strings.Split(s[0], ".")
	return octets
}

func findActiveDevices(octets []string) {
	ipBase := strings.Join(octets[:3], ".")

	for i := 1; i <= 255; i++ {
		ip := ipBase + "." + strconv.Itoa(i)

		out, _ := exec.Command("ping", ip, "-c 1").Output()

		if strings.Contains(string(out), "100.0% packet loss") {
			log.Println(ip, "| TANGO DOWN")
		} else {
			hostname, err := net.LookupAddr(ip)

			if err != nil {
				log.Println(ip, "|", err, "| IT'S ALIVEEE")
			} else {
				log.Println(ip, "|", hostname, "| IT'S ALIVEEE")
			}
		}
	}
}

func main() {
	localIp := getLocalIp()
	log.Println("Hostname:", localIp)

	findActiveDevices(localIp)
}
