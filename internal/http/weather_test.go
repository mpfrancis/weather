package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mpfrancis/weather"
	"github.com/mpfrancis/weather/internal/mock"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	url                         string
	openWeatherResponse         string
	openWeatherForecastResponse string
	expectedResponse            string
	expectedResponseCode        int
	invoked                     bool
}

var cases = []testCase{
	// Basic successful test case
	testCase{
		url: "/weather?city=Bogota&country=co",
		openWeatherResponse: `
		{
			"coord": {
				"lon": -74.08,
				"lat": 4.61
			},
			"weather": [
				{
					"main": "Clouds",
					"description": "scattered clouds"
				}
			],
			"main": {
				"temp": 20,
				"pressure": 1025,
				"humidity": 37
			},
			"wind": {
				"speed": 2.6,
				"deg": 230
			},
			"sys": {
				"country": "CO",
				"sunrise": 1608202626,
				"sunset": 1608245303
			},
			"name": "Bogotá"
		}
		`,
		expectedResponse:     `{"location_name":"Bogotá, CO","temperature":"20 °C","wind":"Light breeze, 2.6 m/s, southwest","cloudiness":"scattered clouds","pressure":"1025 hpa","humidity":"37%","sunrise":"05:57","sunset":"17:48","geo_coordinates":"[4.61, -74.08]","requested_time":"` + time.Now().Format("2006-01-02 15:04:05") + `"}` + "\n",
		expectedResponseCode: 200,
		invoked:              true,
	},

	// Basic successful test case with forecast
	testCase{
		url: "/weather?city=Bogota&country=co&forecast=0",
		openWeatherResponse: `
		{
			"coord": {
				"lon": -74.08,
				"lat": 4.61
			},
			"weather": [
				{
					"main": "Clouds",
					"description": "scattered clouds"
				}
			],
			"main": {
				"temp": 20,
				"pressure": 1025,
				"humidity": 37
			},
			"wind": {
				"speed": 2.6,
				"deg": 230
			},
			"sys": {
				"country": "CO",
				"sunrise": 1608202626,
				"sunset": 1608245303
			},
			"name": "Bogotá"
		}
		`,
		openWeatherForecastResponse: `
		{
			"daily": [
				{
					"dt": 1608825600,
					"sunrise": 1608807628,
					"sunset": 1608850304,
					"temp": {
						"day": 19.31,
						"min": 8.89,
						"max": 19.68,
						"night": 11.64,
						"eve": 14.57,
						"morn": 9.16
					},
					"feels_like": {
						"day": 19.12,
						"night": 11.24,
						"eve": 14.99,
						"morn": 7.93
					},
					"pressure": 1014,
					"humidity": 56,
					"dew_point": 10.32,
					"wind_speed": 0.45,
					"wind_deg": 190,
					"weather": [
						{
							"id": 500,
							"main": "Rain",
							"description": "light rain",
							"icon": "10d"
						}
					],
					"clouds": 31,
					"pop": 0.97,
					"rain": 6.42,
					"uvi": 11.99
				},
				{
					"dt": 1608912000,
					"sunrise": 1608894056,
					"sunset": 1608936733,
					"temp": {
						"day": 17.67,
						"min": 10.14,
						"max": 17.74,
						"night": 10.57,
						"eve": 14.82,
						"morn": 10.28
					},
					"feels_like": {
						"day": 18,
						"night": 9.52,
						"eve": 14.78,
						"morn": 9.58
					},
					"pressure": 1013,
					"humidity": 73,
					"dew_point": 12.78,
					"wind_speed": 0.75,
					"wind_deg": 290,
					"weather": [
						{
							"id": 501,
							"main": "Rain",
							"description": "moderate rain",
							"icon": "10d"
						}
					],
					"clouds": 89,
					"pop": 1,
					"rain": 12.71,
					"uvi": 12.08
				},
				{
					"dt": 1608998400,
					"sunrise": 1608980484,
					"sunset": 1609023162,
					"temp": {
						"day": 18.45,
						"min": 9.45,
						"max": 19.2,
						"night": 9.45,
						"eve": 14.6,
						"morn": 9.92
					},
					"feels_like": {
						"day": 16.71,
						"night": 7.89,
						"eve": 13.75,
						"morn": 8.95
					},
					"pressure": 1012,
					"humidity": 54,
					"dew_point": 9.08,
					"wind_speed": 2.16,
					"wind_deg": 299,
					"weather": [
						{
							"id": 500,
							"main": "Rain",
							"description": "light rain",
							"icon": "10d"
						}
					],
					"clouds": 100,
					"pop": 0.97,
					"rain": 2.25,
					"uvi": 12.09
				},
				{
					"dt": 1609084800,
					"sunrise": 1609066912,
					"sunset": 1609109592,
					"temp": {
						"day": 16.15,
						"min": 9.17,
						"max": 18.08,
						"night": 9.7,
						"eve": 16.52,
						"morn": 9.37
					},
					"feels_like": {
						"day": 14.21,
						"night": 8.57,
						"eve": 15,
						"morn": 7.92
					},
					"pressure": 1013,
					"humidity": 60,
					"dew_point": 8.56,
					"wind_speed": 2.24,
					"wind_deg": 309,
					"weather": [
						{
							"id": 500,
							"main": "Rain",
							"description": "light rain",
							"icon": "10d"
						}
					],
					"clouds": 95,
					"pop": 0.86,
					"rain": 3.2,
					"uvi": 12.27
				},
				{
					"dt": 1609171200,
					"sunrise": 1609153339,
					"sunset": 1609196022,
					"temp": {
						"day": 17.35,
						"min": 8.17,
						"max": 19.06,
						"night": 10.8,
						"eve": 15,
						"morn": 8.17
					},
					"feels_like": {
						"day": 15.41,
						"night": 10.4,
						"eve": 14.27,
						"morn": 6.74
					},
					"pressure": 1015,
					"humidity": 49,
					"dew_point": 6.72,
					"wind_speed": 1.63,
					"wind_deg": 298,
					"weather": [
						{
							"id": 500,
							"main": "Rain",
							"description": "light rain",
							"icon": "10d"
						}
					],
					"clouds": 84,
					"pop": 0.85,
					"rain": 3.21,
					"uvi": 0.49
				},
				{
					"dt": 1609257600,
					"sunrise": 1609239766,
					"sunset": 1609282451,
					"temp": {
						"day": 16.74,
						"min": 9.58,
						"max": 16.74,
						"night": 11.48,
						"eve": 14.36,
						"morn": 9.58
					},
					"feels_like": {
						"day": 15.88,
						"night": 11.45,
						"eve": 14.2,
						"morn": 8.97
					},
					"pressure": 1014,
					"humidity": 59,
					"dew_point": 8.78,
					"wind_speed": 0.81,
					"wind_deg": 282,
					"weather": [
						{
							"id": 500,
							"main": "Rain",
							"description": "light rain",
							"icon": "10d"
						}
					],
					"clouds": 58,
					"pop": 0.88,
					"rain": 7.25,
					"uvi": 1
				},
				{
					"dt": 1609344000,
					"sunrise": 1609326193,
					"sunset": 1609368881,
					"temp": {
						"day": 15.21,
						"min": 10.37,
						"max": 15.87,
						"night": 10.37,
						"eve": 14.67,
						"morn": 10.84
					},
					"feels_like": {
						"day": 15.24,
						"night": 9.32,
						"eve": 14.79,
						"morn": 10.38
					},
					"pressure": 1013,
					"humidity": 74,
					"dew_point": 10.78,
					"wind_speed": 0.26,
					"wind_deg": 195,
					"weather": [
						{
							"id": 501,
							"main": "Rain",
							"description": "moderate rain",
							"icon": "10d"
						}
					],
					"clouds": 81,
					"pop": 0.96,
					"rain": 13.74,
					"uvi": 1
				},
				{
					"dt": 1609430400,
					"sunrise": 1609412619,
					"sunset": 1609455310,
					"temp": {
						"day": 17.18,
						"min": 8.73,
						"max": 17.18,
						"night": 10.38,
						"eve": 14.62,
						"morn": 8.73
					},
					"feels_like": {
						"day": 16.59,
						"night": 9.36,
						"eve": 14.65,
						"morn": 7.31
					},
					"pressure": 1015,
					"humidity": 61,
					"dew_point": 9.78,
					"wind_speed": 0.75,
					"wind_deg": 141,
					"weather": [
						{
							"id": 501,
							"main": "Rain",
							"description": "moderate rain",
							"icon": "10d"
						}
					],
					"clouds": 94,
					"pop": 0.99,
					"rain": 9.68,
					"uvi": 1
				}
			]
		}
	`,
		expectedResponse:     `{"location_name":"Bogotá, CO","temperature":"20 °C","wind":"Light breeze, 2.6 m/s, southwest","cloudiness":"scattered clouds","pressure":"1025 hpa","humidity":"37%","sunrise":"05:57","sunset":"17:48","geo_coordinates":"[4.61, -74.08]","requested_time":"` + time.Now().Format("2006-01-02 15:04:05") + `","forecast":{"dt":1608825600,"sunrise":1608807628,"sunset":1608850304,"temp":{"day":19.31,"min":8.89,"max":19.68,"night":11.64,"eve":14.57,"morn":9.16},"feels_like":{"day":19.12,"night":11.24,"eve":14.99,"morn":7.93},"pressure":1014,"humidity":56,"dew_point":10.32,"wind_speed":0.45,"wind_deg":190,"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"clouds":31,"pop":0.97,"rain":6.42,"uvi":11.99}}` + "\n",
		expectedResponseCode: 200,
		invoked:              true,
	},

	// Basic successful cache test case
	testCase{
		url: "/weather?city=Bogota&country=co",
		openWeatherResponse: `
		{
			"coord": {
				"lon": -74.08,
				"lat": 4.61
			},
			"weather": [
				{
					"main": "Clouds",
					"description": "scattered clouds"
				}
			],
			"main": {
				"temp": 20,
				"pressure": 1025,
				"humidity": 37
			},
			"wind": {
				"speed": 2.6,
				"deg": 230
			},
			"sys": {
				"country": "CO",
				"sunrise": 1608202626,
				"sunset": 1608245303
			},
			"name": "Bogotá"
		}
		`,
		expectedResponse:     `{"location_name":"Bogotá, CO","temperature":"20 °C","wind":"Light breeze, 2.6 m/s, southwest","cloudiness":"scattered clouds","pressure":"1025 hpa","humidity":"37%","sunrise":"05:57","sunset":"17:48","geo_coordinates":"[4.61, -74.08]","requested_time":"` + time.Now().Format("2006-01-02 15:04:05") + `"}` + "\n",
		expectedResponseCode: 200,
		invoked:              false,
	},

	// Query parameter city missing
	testCase{
		url:                  "/weather?country=co",
		expectedResponse:     "Query parameter 'city' is required\n",
		expectedResponseCode: 422,
		invoked:              false,
	},

	// Query parameter country missing
	testCase{
		url:                  "/weather?city=Bogota",
		expectedResponse:     "Query parameter 'country' is required\n",
		expectedResponseCode: 422,
		invoked:              false,
	},

	// Invalid forecast value
	testCase{
		url:                  "/weather?city=Bogota&country=co&forecast=7",
		expectedResponse:     "Query parameter 'forecast' is invalid, please provide a number between 0 and 6\n",
		expectedResponseCode: 422,
		invoked:              false,
	},

	// Invalid forecast value
	testCase{
		url:                  "/weather?city=Bogota&country=co&forecast=-1",
		expectedResponse:     "Query parameter 'forecast' is invalid, please provide a number between 0 and 6\n",
		expectedResponseCode: 422,
		invoked:              false,
	},

	// Invalid forecast value
	testCase{
		url:                  "/weather?city=Bogota&country=co&forecast=a",
		expectedResponse:     "Query parameter 'forecast' is invalid, please provide a number between 0 and 6\n",
		expectedResponseCode: 422,
		invoked:              false,
	},
}

func TestWeatherHandler(t *testing.T) {
	cfg := weather.Config{Units: weather.Metric}
	handler := NewWeatherHandler(&cfg, nil)

	for i := range cases {
		mockClient := mock.Client{}
		mockClient.GetFn = func(url string) (resp *http.Response, err error) {
			switch {
			case strings.Contains(url, "/onecall?"):
				r := ioutil.NopCloser(bytes.NewReader([]byte(cases[i].openWeatherForecastResponse)))
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			case strings.Contains(url, "/weather?"):
				r := ioutil.NopCloser(bytes.NewReader([]byte(cases[i].openWeatherResponse)))
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			}

			return nil, nil
		}

		handler.client = &mockClient

		req, err := http.NewRequest("GET", cases[i].url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, cases[i].expectedResponseCode, rr.Code)
		assert.Equal(t, cases[i].expectedResponse, rr.Body.String())
		assert.Equal(t, cases[i].invoked, mockClient.GetInvoked)
	}
}
