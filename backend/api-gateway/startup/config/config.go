package config

import "os"

type Config struct {
	Address                     string
	UserServiceAddress          string
	AccommodationServiceAddress string
	ReservationServiceAddress   string
}

func GetConfig() Config {
	return Config{
		UserServiceAddress:          os.Getenv("USER_SERVICE_ADDRESS"),
		AccommodationServiceAddress: os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		ReservationServiceAddress:   os.Getenv("RESERVATION_SERVICE_ADDRESS"),
		Address:                     os.Getenv("GATEWAY_ADDRESS"),
	}
}