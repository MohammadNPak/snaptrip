package model

type City struct {
	ID        uint16
	IsActive  bool
	IATA      string `gorm:"size:3;primaryKey;index:,unique;"`
	Name      string
	CreatedAt int64 `gorm:"autoCreateTime:nano"`
}

type Airline struct {
	ICAO      string `gorm:"size:3;primaryKey;index:,unique;"`
	Name      string
	NameFarsi string
}

type Agency struct {
	ID        uint16 `gorm:"primaryKey;index:,unique;"`
	IsActive  bool
	Name      string
	CreatedAt int64 `gorm:"autoCreateTime:nano"`
}

type Supplier struct {
	ID        uint16 `gorm:"primaryKey;index:,unique;"`
	ISActive  bool
	Name      string
	NameFarsi string
	CreatedAt int64 `gorm:"autoCreateTime:nano"`
}

// type Route struct {
// 	ID              uint16 `gorm:"primaryKey;index:,unique;"`
// 	Origin          string `gorm:"size:3;default:null;"`
// 	OriginCity      City   `gorm:"foreignKey:Origin"`
// 	Destination     string `gorm:"size:3;default:null;"`
// 	DestinationCity City   `gorm:"foreignKey:Destination;"`
// 	Rules           []Rule `gorm:"many2many:rule_route;"`
// }

type Route struct {
	ID            uint16 `gorm:"primaryKey;index:,unique;"`
	OriginID      string `gorm:"size:3"`
	DestinationID string `gorm:"size:3"`
	Origin        City
	Destination   City
}

type Rule struct {
	ID          uint16 `gorm:"primaryKey;index:,unique;"`
	AmountType  string
	AmountValue float64
}

type RuleData struct {
	Rule          Rule
	RuleID        uint16
	Route         Route
	RouteId       uint16
	OriginID      string `gorm:"size:3"`
	DestinationID string `gorm:"size:3"`
	Origin        City
	Destination   City
	Airline       Airline
	AirlineID     string `gorm:"size:3"`
	Agency        Agency
	AgencyID      uint16
	Supplier      Supplier
	SupplierID    uint16
}
