# Weather API
This API provides the weather based on a provided location utilizing https://openweathermap.org/api. The package structure is based on https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp.

## Running the Code
### Starting API
```
WEATHER_BASEURL=http://api.openweathermap.org/data/2.5 WEATHER_APIKEY=1508a9a4840a5574c822d70ca2132032 go run cmd/main.go
```

### Example curl command
```
curl 'http://localhost:10000/weather?city=Bogota&country=co&forecast=0'
```

### Configuration Options and Examples
```
WEATHER_BASEURL=http://api.openweathermap.org/data/2.5
WEATHER_APIKEY=abc123
WEATHER_UNITS=metric
SERVER_ADDRESS=:10000
CACHE_EXPIRATION=2m
```

## Get Weather

Get current weather information and optional forecast information. The city and country code are required.

**URL** : `/weather`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

**Required Query Parameters** : city, country

**Optional Query Parameters** : forecast

### Success Response

**Code** : `200 OK`

**Content examples**

For Bogota, CO without a forecast requested.

```json
{
  "location_name": "Bogotá, CO",
  "temperature": "18 °C",
  "wind": "Light breeze, 3.1 m/s, east-northeast",
  "cloudiness": "broken clouds",
  "pressure": "1024 hpa",
  "humidity": "48%",
  "sunrise": "05:57",
  "sunset": "17:48",
  "geo_coordinates": "[4.61, -74.08]",
  "requested_time": "2020-12-17 17:00:50"
}
```

For Bogota, CO with forecast data requested.

```json
{
  "location_name": "Bogotá, CO",
  "temperature": "18 °C",
  "wind": "Light breeze, 3.1 m/s, east-northeast",
  "cloudiness": "broken clouds",
  "pressure": "1024 hpa",
  "humidity": "48%",
  "sunrise": "05:57",
  "sunset": "17:48",
  "geo_coordinates": "[4.61, -74.08]",
  "requested_time": "2020-12-17 17:01:24",
  "forecast": {
    "dt": 1608220800,
    "sunrise": 1608202626,
    "sunset": 1608245303,
    "temp": {
      "day": 18.55,
      "min": 8.97,
      "max": 18.76,
      "night": 10.04,
      "eve": 18,
      "morn": 8.97
    },
    "feels_like": {
      "day": 17.37,
      "night": 8.99,
      "eve": 16.83,
      "morn": 7.15
    },
    "pressure": 1014,
    "humidity": 53,
    "dew_point": 9,
    "wind_speed": 1.3,
    "wind_deg": 148,
    "weather": [
      {
        "id": 500,
        "main": "Rain",
        "description": "light rain",
        "icon": "10d"
      }
    ],
    "clouds": 100,
    "pop": 0.93,
    "rain": 2.51,
    "uvi": 10.76
  }
}
```

### Notes

* The forecast query parameter accepts 0 through 6, with 0 being today. If not provided, no forecast data will be provided.