// Package weather - A package to get the weather forecast for a given city.
package weather

// CurrentCondition - A store of the current condition of the city.
var CurrentCondition string

// CurrentLocation - A store of the current location of the city.
var CurrentLocation string

// Forecast - Display the weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
