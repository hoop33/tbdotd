package models

import "time"

type Deal struct {
	Vendor         *Vendor
	Source         *Source
	Name           string
	Title          string
	ImageUrl       string
	Url            string
	ExpirationDate time.Time
}
