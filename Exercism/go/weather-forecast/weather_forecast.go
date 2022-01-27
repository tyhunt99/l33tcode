/*
Package weather implements a sinple function to output the weather condition.
*/
package weather

// CurrentCondition is set to the value passed into Forecast().
var CurrentCondition string

// CurrentLocation set to the value passed into Forecast().
var CurrentLocation string

// Forecast: given a location and condition return a string about the weather condition in the given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
