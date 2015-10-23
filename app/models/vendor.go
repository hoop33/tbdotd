package models

type Vendor struct {
	Url     string
	Sources map[string]Source
}

var Vendors = map[string]Vendor{
	"Apress": {
		Url: "http://www.apress.com/",
		Sources: map[string]Source{
			"Apress": {
				PayloadUrl: "http://www.apress.com/",
				Regex:      "\\<h2 class=\"icon\"\\>Deal of the Day\\</h2\\>(?s:.+?)\\<img .+?src=\"(.+?)\"(?s:.+?)\\<a href=\"(.+?)\"\\>(.+?)\\</a\\>",
			},
			"Springer": {
				PayloadUrl: "http://www.apress.com/",
				Regex:      "\\<h2 class=\"icon\"\\>Springer Daily Deal\\</h2\\>(?s:.+?)\\<img .+?src=\"(.+?)\"(?s:.+?)\\<a href=\"(.+?)\"\\>(.+?)\\</a\\>",
			},
		},
	},
	"InformIT": {
		Url: "http://www.informit.com/",
		Sources: map[string]Source{
			"InformIT": {
				PayloadUrl: "http://www.informit.com/deals/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
			"InformIT Video": {
				PayloadUrl: "http://www.informit.com/deals/video/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
		},
	},
	"Peachpit": {
		Url: "http://www.peachpit.com/",
		Sources: map[string]Source{
			"Peachpit": {
				PayloadUrl: "http://www.peachpit.com/deals/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
			"Peachpit Video": {
				PayloadUrl: "http://www.peachpit.com/deals/video/deal_rss.aspx",
				DateFormat: "2006-01-02T03:04:05.00-07:00",
			},
		},
	},
	"O'Reilly": {
		Url: "http://www.oreilly.com/",
		Sources: map[string]Source{
			"O'Reilly": {
				PayloadUrl: "http://feeds.feedburner.com/oreilly/ebookdealoftheday",
			},
			"O'Reilly Business": {
				PayloadUrl: "http://feeds.feedburner.com/oreilly/mspebookdeal",
			},
			"O'Reilly Video": {
				PayloadUrl: "http://feeds.feedburner.com/oreilly/videodealoftheweek",
			},
		},
	},
	"Manning": {
		Url: "https://manning.com/",
		Sources: map[string]Source{
			"Manning": {
				PayloadUrl: "https://manning.com/dotd",
				Regex:      "\\<div class=\"title\">Deal of the Day\\</div\\>(?s:.+?)\\<a href=\"(.+?)\"(?s:.+?)\\<div class=\"product-placeholder-title\"\\>\\s*(.+?)\\s*\\<(?s:.+?)\\<div style=\"background-image: url\\('(.+?)'\\)",
			},
		},
	},
	"Packt Publishing": {
		Url: "https://www.packtpub.com/",
		Sources: map[string]Source{
			"Packt Publishing": {
				PayloadUrl: "https://www.packtpub.com/books/deal-of-the-day",
				Regex:      "\\<div class=\"dotd-main-book-image(?s:.+?)\\<img src=\"(.+?)\"(?s:.+?)\\<div class=\"dotd-main-book-title\"\\>(?s:.+?)\\<h2\\>(.+?)\\</h2\\>",
			},
			"Packt Publishing Video": {
				PayloadUrl: "https://www.packtpub.com/videos/deal-of-the-week",
				Regex:      "\\<div class=\"dotd-main-book-image(?s:.+?)\\<img src=\"(.+?)\"(?s:.+?)\\<div class=\"dotw-heading\"\\>(?s:\\s+?)(.+?)(?s:\\s+?)\\</div\\>",
			},
		},
	},
}
