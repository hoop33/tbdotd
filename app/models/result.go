package models

import "time"

type Result struct {
	Vendor         *Vendor
	Deal           *Deal
	Name           string
	Title          string
	ImageUrl       string
	Url            string
	ExpirationDate time.Time
}

type ByDealName []Result

func (a ByDealName) Len() int           { return len(a) }
func (a ByDealName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDealName) Less(i, j int) bool { return a[i].Deal.Name < a[j].Deal.Name }
