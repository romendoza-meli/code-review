package domain

type VehicleAttributes struct {
	Brand        string
	Model        string
	Registration string
	Year         int
	Color        string
	MaxSpeed     int
	FuelType     string
	Transmission string
	Passengers   int
	Height       float64
	Width        float64
	Weight       float64
}

type Vehicle struct {
	Id         int
	Attributes VehicleAttributes
}
