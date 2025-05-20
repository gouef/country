package tests

import (
	"github.com/gouef/country"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByAlpha2(t *testing.T) {
	result := country.FindByAlpha2("cz")
	assert.NotNil(t, result)
	assert.Equal(t, "Czechia", result.Name)
	assert.Equal(t, "CZ", result.Alpha2)

	result = country.FindByAlpha2("USF")
	assert.Nil(t, result)
}

func TestFindByAlpha3(t *testing.T) {
	result := country.FindByAlpha3("CZE")
	assert.NotNil(t, result)
	assert.Equal(t, "Czechia", result.Name)
	assert.Equal(t, "CZE", result.Alpha3)

	result = country.FindByAlpha3("USAF")
	assert.Nil(t, result)
}

func TestFindByName(t *testing.T) {
	result := country.FindByName("Czechia")
	assert.NotNil(t, result)
	assert.Equal(t, "Czechia", result.Name)

	result = country.FindByName("USA")
	assert.Nil(t, result)
}

func TestFlagEmoji(t *testing.T) {
	tests := []struct {
		name     string
		alpha2   string
		expected string
	}{
		{"Germany", "DE", "🇩🇪"},
		{"United States", "US", "🇺🇸"},
		{"Czechia lowercase", "cz", "🇨🇿"},
		{"Invalid one letter", "X", ""},
		{"Invalid three letters", "USA", ""},
		{"Empty", "", ""},
	}

	for _, tt := range tests {
		c := country.Country{
			Name:   tt.name,
			Alpha2: tt.alpha2,
		}
		assert.Equal(t, tt.expected, c.FlagEmoji(), "case: %s", tt.name)
	}
}
