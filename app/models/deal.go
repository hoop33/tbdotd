package models

type Deal struct {
	Vendor   *Vendor
	Title    string
	ImageUrl string
	Url      string
}

type ByVendorName []Deal

func (a ByVendorName) Len() int           { return len(a) }
func (a ByVendorName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVendorName) Less(i, j int) bool { return a[i].Vendor.Name < a[j].Vendor.Name }
