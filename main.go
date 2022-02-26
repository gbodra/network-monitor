package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gbodra/network-monitor/data"
	"github.com/schollz/progressbar/v3"
)

func loadConfig() []string {
	var subnets []string

	f, err := os.Open("subnets.cfg")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		subnets = append(subnets, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return subnets
}

func findActiveDevices(ipBase string, idScan uint) {
	bar := progressbar.Default(255)
	var devicesFound []string

	for i := 1; i < 255; i++ {
		ip := ipBase + strconv.Itoa(i)

		out, _ := exec.Command("ping", ip, "-c 1", "-t 1").Output()

		if !strings.Contains(string(out), "100.0% packet loss") {
			devicesFound = append(devicesFound, ip)
			data.InsertHost(&data.Host{IdScan: idScan, Ip: ip})
		}

		bar.Add(1)
	}

	log.Println("Found", len(devicesFound), "devices on network:", ipBase)
}

func main() {
	ips := loadConfig()

	data.MigrateDb()

	idScan := data.InsertScan(&data.Scan{Subnets: strings.Join(ips, ",")})

	for _, ip := range ips {
		findActiveDevices(ip, idScan)
	}
}
