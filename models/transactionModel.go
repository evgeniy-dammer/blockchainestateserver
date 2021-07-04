package models

//Transaction2 struct
type Transaction2 struct {
	Metadata          Metadata          `json:"metadata"`
	Address           Address           `json:"address"`
	Parcel            Parcel            `json:"parcel"`
	Structure         Structure         `json:"structure"`
	Taxes             Taxes             `json:"taxes"`
	Assessments       Assessments       `json:"assessments"`
	MarketAssessments MarketAssessments `json:"market_assessments"`
	Valuation         Valuation         `json:"valuation"`
	Owner             Owner             `json:"owner"`
	Deeds             Deeds             `json:"deeds"`
	Boundary          Boundary          `json:"boundary"`
}

//Metadata struct
type Metadata struct {
	PublishingDate string `json:"publishing_date"`
}

//Address struct
type Address struct {
	StreetNumber           string  `json:"street_number"`
	StreetPreDirection     string  `json:"street_pre_direction"`
	StreetName             string  `json:"street_name"`
	StreetSuffix           string  `json:"street_suffix"`
	StreetPostPirection    string  `json:"street_post_direction"`
	UnitType               string  `json:"unit_type"`
	UnitNumber             string  `json:"unit_number"`
	FormattedStreetAddress string  `json:"formatted_street_address"`
	City                   string  `json:"city"`
	State                  string  `json:"state"`
	ZipCode                string  `json:"zip_code"`
	ZipPlusFourCode        string  `json:"zip_plus_four_code"`
	CarrierCode            string  `json:"carrier_code"`
	Latitude               float64 `json:"latitude"`
	Longitude              float64 `json:"longitude"`
	GeocodingAccuracy      string  `json:"geocoding_accuracy"`
	CensusTract            string  `json:"census_tract"`
}

//Parcel struct
type Parcel struct {
	ApnOriginal                 string  `json:"apn_original"`
	ApnUnformatted              string  `json:"apn_unformatted"`
	ApnPrevious                 string  `json:"apn_previous"`
	FipsCode                    string  `json:"fips_code"`
	DepthFt                     float64 `json:"depth_ft"`
	FrontageFt                  float64 `json:"frontage_ft"`
	AreaSqFt                    int64   `json:"area_sq_ft"`
	AreaAcres                   float64 `json:"area_acres"`
	CountyName                  string  `json:"county_name"`
	CountyLandUseCode           string  `json:"county_land_use_code"`
	CountyLandUseDescription    string  `json:"county_land_use_description"`
	StandardizedLandUseCategory string  `json:"standardized_land_use_category"`
	StandardizedLandUseType     string  `json:"standardized_land_use_type"`
	LocationDescriptions        string  `json:"location_descriptions"`
	Zoning                      string  `json:"zoning"`
	BuildingCount               int64   `json:"building_count"`
	TaxAccountNumber            string  `json:"tax_account_number"`
	LegalDescription            string  `json:"legal_description"`
	LotCode                     string  `json:"lot_code"`
	LotNumber                   string  `json:"lot_number"`
	Subdivision                 string  `json:"subdivision"`
	Municipality                string  `json:"municipality"`
	SectionTownshipRange        string  `json:"section_township_range"`
}

//Structure struct
type Structure struct {
	YearBuilt             int64             `json:"year_built"`
	EffectiveYearBuilt    int64             `json:"effective_year_built"`
	Stories               string            `json:"stories"`
	RoomsCount            int64             `json:"rooms_count"`
	BedsCount             int64             `json:"beds_count"`
	Baths                 float64           `json:"baths"`
	PartialBathsCount     int64             `json:"partial_baths_count"`
	UnitsCount            int64             `json:"units_count"`
	ParkingType           string            `json:"parking_type"`
	ParkingSpacesCount    int64             `json:"parking_spaces_count"`
	PoolType              string            `json:"pool_type"`
	ArchitectureType      string            `json:"architecture_type"`
	ConstructionType      string            `json:"construction_type"`
	ExteriorWallType      string            `json:"exterior_wall_type"`
	FoundationType        string            `json:"foundation_type"`
	RoofMaterialType      string            `json:"roof_material_type"`
	RoofStyleType         string            `json:"roof_style_type"`
	HeatingType           string            `json:"heating_type"`
	HeatingFuelType       string            `json:"heating_fuel_type"`
	AirConditioningType   string            `json:"air_conditioning_type"`
	Fireplaces            int               `json:"fireplaces"`
	BasementType          string            `json:"basement_type"`
	Quality               string            `json:"quality"`
	Condition             string            `json:"condition"`
	FlooringTypes         string            `json:"flooring_types"`
	PlumbingFixturesCount int64             `json:"plumbing_fixtures_count"`
	InteriorWallType      string            `json:"interior_wall_type"`
	WaterType             string            `json:"water_type"`
	SewerType             string            `json:"sewer_type"`
	TotalAreaSqFt         int64             `json:"total_area_sq_ft"`
	OtherAreas            OtherAreas        `json:"other_areas"`
	OtherFeatures         OtherFeatures     `json:"other_features"`
	OtherImprovements     OtherImprovements `json:"other_improvements"`
	OtherRooms            string            `json:"other_rooms"`
	Amenities             string            `json:"amenities"`
}

//OtherAreas struct
type OtherAreas struct {
	Type string `json:"type"`
	SqFt string `json:"sq_ft"`
}

//OtherFeatures struct
type OtherFeatures struct {
	Type string `json:"type"`
	SqFt string `json:"sq_ft"`
}

//OtherImprovements struct
type OtherImprovements struct {
	Type string `json:"type"`
	SqFt string `json:"sq_ft"`
}

//Taxes struct
type Taxes struct {
	Year         int64  `json:"year"`
	Amount       int64  `json:"amount"`
	Exemptions   string `json:"exemptions"`
	RateCodeArea string `json:"rate_code_area"`
}

//Assessments struct
type Assessments struct {
	Year             int64 `json:"year"`
	LandValue        int64 `json:"land_value"`
	ImprovementValue int64 `json:"improvement_value"`
	TotalValue       int64 `json:"total_value"`
}

//MarketAssessments struct
type MarketAssessments struct {
	Year             int64 `json:"year"`
	LandValue        int64 `json:"land_value"`
	ImprovementValue int64 `json:"improvement_value"`
	TotalValue       int64 `json:"total_value"`
}

//Valuation struct
type Valuation struct {
	Value                     int64  `json:"value"`
	High                      int64  `json:"high"`
	Low                       int64  `json:"low"`
	ForecastStandardDeviation int64  `json:"forecast_standard_deviation"`
	Date                      string `json:"date"`
}

//Owner struct
type Owner struct {
	Name                   string `json:"name"`
	SecondName             string `json:"second_name"`
	FormattedStreetAddress string `json:"formatted_street_address"`
	UnitType               string `json:"unit_type"`
	UnitNumber             string `json:"unit_number"`
	City                   string `json:"city"`
	State                  string `json:"state"`
	ZipCode                string `json:"zip_code"`
	ZipPlusFourCode        string `json:"zip_plus_four_code"`
	OwnerOccupied          string `json:"owner_occupied"`
}

//Deeds struct
type Deeds struct {
	DocumentType          string  `json:"document_type"`
	RecordingDate         string  `json:"recording_date"`
	OriginalContractDate  string  `json:"original_contract_date"`
	DeedBook              string  `json:"deed_book"`
	DeedBage              string  `json:"deed_page"`
	DocumentId            string  `json:"document_id"`
	SalePrice             int64   `json:"sale_price"`
	SalePriceDescription  string  `json:"sale_price_description"`
	TransferTax           float64 `json:"transfer_tax"`
	DistressedSale        bool    `json:"distressed_sale"`
	RealEstateOwned       string  `json:"real_estate_owned"`
	SellerFirstName       string  `json:"seller_first_name"`
	SellerLastName        string  `json:"seller_last_name"`
	Seller2FirstName      string  `json:"seller2_first_name"`
	Seller2LastName       string  `json:"seller2_last_name"`
	SellerAddress         string  `json:"seller_address"`
	SellerUnitNumber      string  `json:"seller_unit_number"`
	SellerCity            string  `json:"seller_city"`
	SellerState           string  `json:"seller_state"`
	SellerZipCode         string  `json:"seller_zip_code"`
	SellerZipPlusFourCode string  `json:"seller_zip_plus_four_code"`
	BuyerFirstName        string  `json:"buyer_first_name"`
	BuyerLastName         string  `json:"buyer_last_name"`
	Buyer2FirstName       string  `json:"buyer2_first_name"`
	Buyer2LastName        string  `json:"buyer2_last_name"`
	BuyerAddress          string  `json:"buyer_address"`
	BuyerUnitType         string  `json:"buyer_unit_type"`
	BuyerUnitNumber       string  `json:"buyer_unit_number"`
	BuyerCity             string  `json:"buyer_city"`
	BuyerState            string  `json:"buyer_state"`
	BuyerZipCode          string  `json:"buyer_zip_code"`
	BuyerZipPlusFourCode  string  `json:"buyer_zip_plus_four_code"`
	LenderName            string  `json:"lender_name"`
	LenderType            string  `json:"lender_type"`
	LoanAmount            int64   `json:"loan_amount"`
	LoanType              string  `json:"loan_type"`
	LoanDueDate           string  `json:"loan_due_date"`
	LoanFinanceType       string  `json:"loan_finance_type"`
	LoanInterestRate      float64 `json:"loan_interest_rate"`
}

//Boundary struct
type Boundary struct {
	Wkt     string  `json:"wkt"`
	Geojson Geojson `json:"geojson"`
}

//Geojson struct
type Geojson struct {
	Type        string        `json:"type"`
	Coordinates []Coordinates `json:"coordinates"`
}

//Coordinates struct
type Coordinates struct {
	Longitude float64 `json:"longetude"`
	Latitude  float64 `json:"latitude"`
}
