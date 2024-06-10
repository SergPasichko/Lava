package domain

import "time"

type Plant struct {
	Id          uint64
	UserId      uint64
	Name        string
	Address     string
	Lat         float64
	Lon         float64
	Type        PlantType
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type PlantType string

const (
	SolarType PlantType = "SOLAR"
	WindType  PlantType = "WIND"
)
