// Package weather fournit des outils pour afficher la météo.
package weather

// CurrentCondition stocke l'état actuel de la météo.
var CurrentCondition string

// CurrentLocation stocke le lieu actuel.
var CurrentLocation string

// Forecast met à jour la position et la condition, puis retourne un bulletin météo.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}