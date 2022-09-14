package model

type City struct {
	ID        int32
	IsActive  bool
	IATA      string
	Name      string
	CreatedAt string
}

type Airline struct {
	ICAO      string
	Name      string
	NameFarsi string
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
type Route struct {
	ID          int32  `gorm:"type:int"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Rule struct {
	ID          int32
	// Routes      []Route  `json:"routes" gorm:"foreignKey:ID"`
	// Airlines    []string `json:"airlines" gorm:"foreignKey:ICAO"`
	Agencies    []string `json:"agencies" gorm:"foreignKey:Name"`
	// Suppliers   []string `json:"suppliers" gorm:"foreignKey:Name"`
	AmountType  string   `json:"amountType"`
	AmountValue float64  `json:"amountValue"`
}
