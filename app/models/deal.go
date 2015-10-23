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

type BySourceName []Deal

func (a BySourceName) Len() int      { return len(a) }
func (a BySourceName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//func (a BySourceName) Less(i, j int) bool { return a[i].Source.Name < a[j].Source.Name }
func (a BySourceName) Less(i, j int) bool { return a[i].Title < a[j].Title }
