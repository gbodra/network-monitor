package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Scan struct {
	gorm.Model
	Subnets string
}

type Host struct {
	gorm.Model
	IdScan uint
	Ip     string
}

type Report struct {
	gorm.Model
	Scans                   uint
	NewDevicesSinceLastScan uint
}

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("network-monitor.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func MigrateDb() {
	db := ConnectDatabase()
	db.AutoMigrate(&Scan{})
	db.AutoMigrate(&Host{})
	db.AutoMigrate(&Report{})
}

func InsertScan(scan *Scan) uint {
	db := ConnectDatabase()

	db.Create(scan)

	return scan.ID
}

func InsertHost(host *Host) uint {
	db := ConnectDatabase()

	db.Create(host)

	return host.ID
}

func InsertReport(report *Report) uint {
	db := ConnectDatabase()

	db.Create(report)

	return report.ID
}
