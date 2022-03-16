package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gbodra/network-monitor/notification"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
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

func findActiveDevices(ipBase string) []string {
	bar := progressbar.Default(254)
	var devicesFound []string

	for i := 1; i < 255; i++ {
		ip := ipBase + strconv.Itoa(i)

		out, _ := exec.Command("ping", ip, "-c 1", "-t 1").Output()

		if !strings.Contains(string(out), "100.0% packet loss") {
			devicesFound = append(devicesFound, ip)
			// data.InsertHost(&data.Host{IdScan: idScan, Ip: ip})
		}

		bar.Add(1)
	}

	log.Println("Found", len(devicesFound), "devices on network:", ipBase)

	return devicesFound
}

func ScanNetwork(ips []string) {
	for _, ip := range ips {
		devices := findActiveDevices(ip)
		message := fmt.Sprint("Found ", len(devices), " devices on network ", ip, "\n", strings.Join(devices, "\n"))
		notification.SendMessageTelegram(message)
	}
}

func PrintMsg() {
	log.Println("Hello cron")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	ips := loadConfig()

	// data.MigrateDb()

	// idScan := data.InsertScan(&data.Scan{Subnets: strings.Join(ips, ",")})

	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(os.Getenv("TASK_FREQ")).Do(ScanNetwork, ips)
	scheduler.StartBlocking()

	fmt.Println("Press the any key to stop")
	fmt.Scanln()
}
