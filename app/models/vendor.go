package models

type Vendor struct {
	Name  string
	Url   string
	Deals []Deal
}

func VendorWithName(name string) Vendor {
	for _, vendor := range Vendors {
		if vendor.Name == name {
			return vendor
		}
	}
	return Vendor{}
}

var Vendors = []Vendor{
	{
		Name: "Apress",
		Url:  "http://www.apress.com/",
		Deals: []Deal{
			{
				Name:       "Apress",
				PayloadUrl: "http://www.apress.com/",
				Regex:      "\\<h2 class=\"icon\"\\>Deal of the Day\\</h2\\>(?s:.+?)\\<img .+?src=\"(.+?)\"(?s:.+?)\\<a href=\"(.+?)\"\\>(.+?)\\</a\\>",
			},
			{
				Name:       "Springer",
				PayloadUrl: "http://www.apress.com/",
				Regex:      "\\<h2 class=\"icon\"\\>Springer Daily Deal\\</h2\\>(?s:.+?)\\<img .+?src=\"(.+?)\"(?s:.+?)\\<a href=\"(.+?)\"\\>(.+?)\\</a\\>",
			},
		},
	},
	{
		Name: "InformIT",
		Url:  "http://www.informit.com/",
		Deals: []Deal{
			{
				Name:       "InformIT",
				PayloadUrl: "http://www.informit.com/deals/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
			{
				Name:       "InformIT Video",
				PayloadUrl: "http://www.informit.com/deals/video/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
		},
	},
	{
		Name: "Peachpit",
		Url:  "http://www.peachpit.com/",
		Deals: []Deal{
			{
				Name:       "Peachpit",
				PayloadUrl: "http://www.peachpit.com/deals/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
			{
				Name:       "Peachpit Video",
				PayloadUrl: "http://www.peachpit.com/deals/video/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
		},
	},
	{
		Name: "O'Reilly",
		Url:  "http://www.oreilly.com/",
		Deals: []Deal{
			{
				Name:       "O'Reilly",
				PayloadUrl: "http://feeds.feedburner.com/oreilly/ebookdealoftheday",
			},
			{
				Name:       "O'Reilly Business",
				PayloadUrl: "http://feeds.feedburner.com/oreilly/mspebookdeal",
			},
			{
				Name:       "O'Reilly Video",
				PayloadUrl: "http://feeds.feedburner.com/oreilly/videodealoftheweek",
			},
		},
	},
	{
		Name: "Manning",
		Url:  "https://manning.com/",
		Deals: []Deal{
			{
				Name:       "Manning",
				PayloadUrl: "https://manning.com/dotd",
				Regex:      "\\<div class=\"title\">Deal of the Day\\</div\\>(?s:.+?)\\<a href=\"(.+?)\"(?s:.+?)\\<div class=\"product-placeholder-title\"\\>\\s*(.+?)\\s*\\<(?s:.+?)\\<div style=\"background-image: url\\('(.+?)'\\)",
			},
		},
	},
	{
		Name: "Packt Publishing",
		Url:  "https://www.packtpub.com/",
		Deals: []Deal{
			{
				Name:       "Packt Publishing",
				PayloadUrl: "https://www.packtpub.com/books/deal-of-the-day",
				Regex:      "\\<div class=\"dotd-main-book-image(?s:.+?)\\<img src=\"(.+?)\"(?s:.+?)\\<div class=\"dotd-main-book-title\"\\>(?s:.+?)\\<h2\\>(.+?)\\</h2\\>",
			},
			{
				Name:       "Packt Publishing Video",
				PayloadUrl: "https://www.packtpub.com/videos/deal-of-the-week",
				Regex:      "\\<div class=\"dotd-main-book-image(?s:.+?)\\<img src=\"(.+?)\"(?s:.+?)\\<div class=\"dotw-heading\"\\>(?s:\\s+?)(.+?)(?s:\\s+?)\\</div\\>",
			},
		},
	},
}
