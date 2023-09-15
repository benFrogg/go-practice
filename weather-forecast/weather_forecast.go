// Package weather calls another file where it's function will be used here.
package weather

// CurrentCondition declares a string variable.
var CurrentCondition string

// CurrentLocation declares a string variable.
var CurrentLocation string

// Forecast takse in 2 string parameters and impose onto the newly declared variables and concat them.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
