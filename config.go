package weather

// Config is used for configuration and dependency injection.
type Config struct {
	BaseURL       string
	APIKey        string
	ServerAddress string
	Units         Unit
}

// Unit provides a type for setting the unit of the open weather API.
type Unit string

// List of unit types.
const (
	Metric   Unit = "metric"
	Standard Unit = "imperial"
	Imperial Unit = "standard"
)

// Symbol returns the symbol used for the given unit type.
func (u Unit) Symbol() string {
	switch u {
	case Metric:
		return "°C"
	case Standard:
		return "K"
	case Imperial:
		return "°F"
	}

	return ""
}
