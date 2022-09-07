package model

type City struct {
	ID         uint16
	IsActive   bool
	IATA       string
	Name       string
	created_at string
}

type Airline struct {
	ICAO       string
	name       string
	name_farsi string
}

type Agencie struct {
	Id        int32
	IsActive  bool
	Name      string
	CreatedAt string
}

type Supplier struct {
	ID        int32
	ISActive  bool
	Name      string
	NameFarsi string
	CreatedAt string
}
type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Rule struct {
	ID          int32
	Routes      []Route  `json:"routes"`
	Airlines    []string `json:"airlines"`
	Agencies    []string `json:"agencies"`
	Suppliers   []string `json:"suppliers"`
	AmountType  string   `json:"amountType"`
	AmountValue float64  `json:"amountValue"`
}

type Ticket struct {
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Airline      string `json:"airline"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	BasePrice    int64  `json:"basePrice"`
	Markup       int64  `json:"markup"`
	PayablePrice int64  `json:"payablePrice"`
}
