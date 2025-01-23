package country

import "strings"

type Country struct {
	Name    string
	Alpha2  string
	Alpha3  string
	Numeric string
}

// FindByAlpha2 searches for a country by its Alpha-2 code.
//
// Example:
//
//	cz := country.FindByAlpha2("CZ")
func FindByAlpha2(alpha2 string) *Country {
	for _, country := range Countries {
		if country.Alpha2 == strings.ToUpper(alpha2) {
			return &country
		}
	}
	return nil
}

// FindByAlpha3 searches for a country by its Alpha-3 code.
//
// Example:
//
//	cz := country.FindByAlpha3("CZE")
func FindByAlpha3(alpha3 string) *Country {
	for _, country := range Countries {
		if country.Alpha3 == strings.ToUpper(alpha3) {
			return &country
		}
	}
	return nil
}

// FindByName searches for a country by its name (english).
//
// Example:
//
//	cz := country.FindByAlpha3("Czechia")
func FindByName(name string) *Country {
	for _, country := range Countries {
		if strings.ToUpper(country.Name) == strings.ToUpper(name) {
			return &country
		}
	}
	return nil
}
