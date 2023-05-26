package models

import "time"

type Wallet struct {
	ID            int       `json:"id"`
	DNI           string    `json:"dni_request"`
	Country       string    `json:"country_id"`
	Order_request time.Time `json:"order_request"`
}
