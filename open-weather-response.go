package weather

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// OpenWeatherResponse is the object for the response from open weather's /weather endpoint.
type OpenWeatherResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

// Coord holds the location coordinate data.
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Weather provides a description of the current weather.
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main holds some general weather data.
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// Wind holds wind weather data.
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// Clouds holds weather data in regards to clouds.
type Clouds struct {
	All int `json:"all"`
}

// Sys holds some additional weather details from the open weather API.
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

// ToHumanReadable converts an open weather model to a more human readable model.
func (o *OpenWeatherResponse) ToHumanReadable(unitSymbol string) *HumanReadableResponse {
	resp := HumanReadableResponse{
		LocationName:   fmt.Sprintf("%s, %s", strings.Title(o.Name), strings.ToUpper(o.Sys.Country)),
		Temperature:    fmt.Sprintf("%g %s", o.Main.Temp, unitSymbol),
		Wind:           fmt.Sprintf("%s, %g m/s, %s", windDescription(o.Wind.Speed), o.Wind.Speed, windDirection(o.Wind.Deg)),
		Pressure:       fmt.Sprintf("%d hpa", o.Main.Pressure),
		Humidity:       fmt.Sprintf("%d%%", o.Main.Humidity),
		Sunrise:        time.Unix(o.Sys.Sunrise, 0).Format("15:04"),
		Sunset:         time.Unix(o.Sys.Sunset, 0).Format("15:04"),
		GeoCoordinates: fmt.Sprintf("[%g, %g]", o.Coord.Lat, o.Coord.Lon),
		RequestedTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	if len(o.Weather) > 0 {
		resp.Cloudiness = o.Weather[0].Description
	}

	return &resp
}

// windDescription uses the following scale to determine wind speed: https://en.wikipedia.org/wiki/Beaufort_scale
func windDescription(speed float64) string {
	switch {
	case speed <= .5:
		return "Calm"
	case speed <= 1.5:
		return "Light air"
	case speed <= 3.3:
		return "Light breeze"
	case speed <= 5.5:
		return "Gentle breeze"
	case speed <= 7.9:
		return "Moderate breeze"
	case speed <= 10.7:
		return "Fresh breeze"
	case speed <= 13.8:
		return "Strong breeze"
	case speed <= 17.1:
		return "High wind"
	case speed <= 20.7:
		return "Gale"
	case speed <= 24.4:
		return "Strong/severe gale"
	case speed <= 28.4:
		return "Storm"
	case speed <= 32.6:
		return "Violent storm"
	}

	return "Hurricane force"
}

// windDirection converts degrees to wind direction
func windDirection(direction int) string {
	direction = direction % 360

	directions := []string{
		"north",
		"north-northeast",
		"northeast",
		"east-northeast",
		"east",
		"east-southeast",
		"southeast",
		"south-southeast",
		"south",
		"south-southwest",
		"southwest",
		"west-southwest",
		"west",
		"west-northwest",
		"northwest",
		"north-northwest",
		"north",
	}

	return directions[int(math.Round(float64(direction)/22.5))]
}
