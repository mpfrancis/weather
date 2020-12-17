package weather

type HumanReadableResponse struct {
	LocationName   string   `json:"location_name"`
	Temperature    string   `json:"temperature"`
	Wind           string   `json:"wind"`
	Cloudiness     string   `json:"cloudiness"`
	Pressure       string   `json:"pressure"`
	Humidity       string   `json:"humidity"`
	Sunrise        string   `json:"sunrise"`
	Sunset         string   `json:"sunset"`
	GeoCoordinates string   `json:"geo_coordinates"`
	RequestedTime  string   `json:"requested_time"`
	Forecast       Forecast `json:"forecast"`
}

type Forecast struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}
