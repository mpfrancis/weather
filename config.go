package weather

type Config struct {
	BaseURL string
	APIKey  string
	Units   Unit
}

type Unit string

const (
	Metric   Unit = "metric"
	Standard      = "imperial"
	Imperial      = "standard"
)

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
