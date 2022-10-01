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
	OriginID      string `gorm:"size:3;index"`
	DestinationID string `gorm:"size:3;index"`
	Origin        City
	Destination   City
}

type Rule struct {
	ID          uint16 `gorm:"primaryKey;index:,unique;"`
	AmountType  string
	AmountValue float64
	Routes      []Route    `gorm:"many2many:rule_route;"`
	Airlines    []Airline  `gorm:"many2many:rule_airline;"`
	Agencies    []Agency   `gorm:"many2many:rule_agency;"`
	Suppliers   []Supplier `gorm:"many2many:rule_supplier;"`
}
