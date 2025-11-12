// Package weather provides weather forecast for cities in Goblinocus.
package weather

// CurrentCondition holds the value of current weather.
var CurrentCondition string
// CurrentLocation holds the city for which the weather is displayed.
var CurrentLocation string

// Forecast returns string that shows current location and weather right now.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
