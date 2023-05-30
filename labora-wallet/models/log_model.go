package models

import "time"

type Log struct {
	ID             int       `json:"id"`
	DNI            string    `json:"dni_request"`
	Country        string    `json:"country_id"`
	Status_request string    `json:"status_request"`
	Date_request   time.Time `json:"date_request"`
	Type_request   string    `json:"type_request"`
}
