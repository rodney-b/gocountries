package gocountries

import (
	"log"
	"testing"
)

func TestCountriesByName(t *testing.T) {
	countries, err := CountriesByName("italy")
	if err != nil {
		t.Errorf("Got unexpected error for requested country 'italy': %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any country with name 'italy'")
	}

	if (countries)[0].Capital != "Rome" {
		t.Errorf("Unexpected capital for country 'italy'")
	}

	if (countries)[0].RegionalBlocs[0].Acronym != "EU" {
		t.Errorf("Unexpected regional block acronym for country 'italy'")
	}

	log.Printf("Fetched: %v", (countries)[0])
}

func TestCountriesByFullName(t *testing.T) {
	countries, err := CountriesByFullName("Ital")

	if len(countries) > 0 {
		t.Errorf("Unexpected: Found country using partial name. There is no country named 'Ital'. Only full names should work.")
	}

	countries, err = CountriesByFullName("Italy")

	if err != nil {
		t.Errorf("Got unexpected error for requested country 'italy': %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any country with name 'italy'")
	}

	if (countries)[0].Capital != "Rome" {
		t.Errorf("Unexpected capital for country 'italy'")
	}

	if (countries)[0].RegionalBlocs[0].Acronym != "EU" {
		t.Errorf("Unexpected regional block acronym for country 'italy'")
	}

	log.Printf("Fetched: %v", (countries)[0])
}

func TestCountriesByCapital(t *testing.T) {
	capital := "tallinn"
	countries, err := CountriesByCapital(capital)
	if err != nil {
		t.Errorf("Got unexpected error for CountriesByCapital(\"%s\"): %v", capital, err)
		return
	}

	country := (countries)[0]

	if country.Name != "Estonia" {
		t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Estonia", country.Name)
	}

	log.Printf("Fetched: %v", country)
}

func TestAllCountries(t *testing.T) {
	countries, err := AllCountries()
	if err != nil {
		t.Errorf("Got unexpected error for GetAllCountries(): %v", err)
		return
	}

	if len(countries) < 1 {
		t.Errorf("Unexpected: couldn't find any country at all")
	}

	if countries[0].Name == "" {
		t.Errorf("Unexpected: invalid country model in GetAllCountries")
	}

	log.Printf("Fetched %d countries", len(countries))
}
