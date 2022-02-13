package entity

import "time"

type Statistics struct {
	Ram     Ram
	Storage Storage
	LoadAvg LoadAvg
	Uptime  Uptime

	UpdatedAt time.Time
}

type Ram struct {
	Total     int
	Available int
}

type Storage struct {
	Total     int
	Available int
	Free      int
}

type LoadAvg struct {
	Load1  float64
	Load5  float64
	Load15 float64
}

type Uptime struct {
	Dur int
}
