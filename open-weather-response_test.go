package weather

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type DirectionCase struct {
	degree            int
	expectedDirection string
}

func TestWindDirection(t *testing.T) {
	cases := []DirectionCase{
		{10, "north"},
		{20, "north-northeast"},
		{30, "north-northeast"},
		{40, "northeast"},
		{50, "northeast"},
		{60, "east-northeast"},
		{70, "east-northeast"},
		{80, "east"},
		{90, "east"},
		{100, "east"},
		{110, "east-southeast"},
		{120, "east-southeast"},
		{130, "southeast"},
		{140, "southeast"},
		{150, "south-southeast"},
		{160, "south-southeast"},
		{170, "south"},
		{180, "south"},
		{190, "south"},
		{200, "south-southwest"},
		{210, "south-southwest"},
		{220, "southwest"},
		{230, "southwest"},
		{240, "west-southwest"},
		{250, "west-southwest"},
		{260, "west"},
		{270, "west"},
		{280, "west"},
		{290, "west-northwest"},
		{300, "west-northwest"},
		{310, "northwest"},
		{320, "northwest"},
		{330, "north-northwest"},
		{340, "north-northwest"},
		{350, "north"},
		{360, "north"},
	}

	for i := range cases {
		direction := windDirection(cases[i].degree)
		if direction != cases[i].expectedDirection {
			t.Fatalf("Expected %s but got %s", cases[i].expectedDirection, direction)
		}
	}
}

type DescriptionCase struct {
	speed               float64
	expectedDescription string
}

func TestWindDescription(t *testing.T) {
	cases := []DescriptionCase{
		{0, "Calm"},
		{.5, "Calm"},
		{.6, "Light air"},
		{1.5, "Light air"},
		{1.6, "Light breeze"},
		{3.3, "Light breeze"},
		{3.4, "Gentle breeze"},
		{5.5, "Gentle breeze"},
		{5.6, "Moderate breeze"},
		{7.9, "Moderate breeze"},
		{8, "Fresh breeze"},
		{10.7, "Fresh breeze"},
		{10.8, "Strong breeze"},
		{13.8, "Strong breeze"},
		{13.9, "High wind"},
		{17.1, "High wind"},
		{17.2, "Gale"},
		{20.7, "Gale"},
		{20.8, "Strong/severe gale"},
		{24.4, "Strong/severe gale"},
		{24.5, "Storm"},
		{28.4, "Storm"},
		{28.5, "Violent storm"},
		{32.6, "Violent storm"},
		{32.7, "Hurricane force"},
		{50, "Hurricane force"},
	}

	for i := range cases {
		description := windDescription(cases[i].speed)
		if description != cases[i].expectedDescription {
			t.Fatalf("Expected %s but got %s", cases[i].expectedDescription, description)
		}
	}
}

type HumanReadableCase struct {
	input  OpenWeatherResponse
	output *HumanReadableResponse
}

func TestToHumanReadable(t *testing.T) {
	cases := []HumanReadableCase{
		{
			OpenWeatherResponse{
				Name: "Bogota",
				Sys: Sys{
					Country: "CO",
					Sunrise: 1608202626,
					Sunset:  1608245303,
				},
				Main: Main{
					Temp:     20,
					Pressure: 1000,
					Humidity: 50,
				},
				Wind: Wind{
					Speed: 3,
					Deg:   10,
				},
				Coord: Coord{
					Lat: 4.61,
					Lon: -74.08,
				},
				Weather: []Weather{{Description: "Scattered clouds"}},
			},
			&HumanReadableResponse{
				LocationName:   "Bogota, CO",
				Temperature:    fmt.Sprintf("%g %s", 20.0, Metric.Symbol()),
				Wind:           "Light breeze, 3 m/s, north",
				Pressure:       "1000 hpa",
				Humidity:       "50%",
				Sunrise:        "05:57",
				Sunset:         "17:48",
				GeoCoordinates: "[4.61, -74.08]",
				RequestedTime:  time.Now().Format("2006-01-02 15:04:05"),
				Cloudiness:     "Scattered clouds",
			},
		},
	}

	for i := range cases {
		hr := cases[i].input.ToHumanReadable(Metric.Symbol())
		if !reflect.DeepEqual(cases[i].output, hr) {
			t.Fatalf("Expected %+v but got %+v", cases[i].output, hr)
		}
	}
}
