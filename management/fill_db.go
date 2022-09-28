package management

import (
	"encoding/csv"
	"os"
	"snapptrip/model"
	"strconv"

	"gorm.io/gorm"
)

func ReadCsvFile(filename string) ([][]string, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func loadCity(db *gorm.DB) *[]model.City {
	csvData, err := ReadCsvFile("bootcamp/data/city.csv")
	if err != nil {
		panic(err)
	}
	var cities []model.City
	for _, line := range csvData {
		id, _ := strconv.ParseUint(line[0], 10, 16)
		is_active, _ := strconv.ParseBool(line[1])
		cities = append(cities, model.City{
			ID:       uint16(id),
			IsActive: is_active,
			IATA:     line[2],
			Name:     line[4]})
	}
	db.CreateInBatches(cities, len(cities))

	return &cities
}

func loadAgency(db *gorm.DB) *[]model.Agency {
	csvData, err := ReadCsvFile("bootcamp/data/agency.csv")
	if err != nil {
		panic(err)
	}
	var agencies []model.Agency
	for _, line := range csvData {
		id, _ := strconv.ParseUint(line[0], 10, 16)
		is_active, _ := strconv.ParseBool(line[1])
		agencies = append(agencies, model.Agency{
			ID:       uint16(id),
			IsActive: is_active,
			Name:     line[2]})
	}
	db.CreateInBatches(agencies, len(agencies))

	return &agencies
}

func loadAirline(db *gorm.DB) *[]model.Airline {
	csvData, err := ReadCsvFile("bootcamp/data/airline.csv")
	if err != nil {
		panic(err)
	}
	var airlines []model.Airline
	for _, line := range csvData {
		airlines = append(airlines, model.Airline{
			ICAO:      line[0],
			Name:      line[1],
			NameFarsi: line[2],
		})
	}
	db.CreateInBatches(airlines, len(airlines))
	return &airlines
}

func loadSupplier(db *gorm.DB) *[]model.Supplier {
	csvData, err := ReadCsvFile("bootcamp/data/supplier.csv")
	if err != nil {
		panic(err)
	}
	var suppliers []model.Supplier
	for _, line := range csvData {
		id, _ := strconv.ParseUint(line[0], 10, 16)
		is_active, _ := strconv.ParseBool(line[1])
		suppliers = append(suppliers, model.Supplier{
			ID:        uint16(id),
			ISActive:  is_active,
			Name:      line[2],
			NameFarsi: line[3],
		})
	}
	db.CreateInBatches(suppliers, len(suppliers))
	return &suppliers
}

func LoadData(db *gorm.DB) {
	loadCity(db)
	loadAgency(db)
	loadAirline(db)
	loadSupplier(db)

}
