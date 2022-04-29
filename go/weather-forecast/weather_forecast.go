// Package weather provides tools about weather conditions.
package weather


// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the city.
var CurrentLocation string

// Forecast() is a function that returns a string that tells the current weather condition at the currentLocation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
