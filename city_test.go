package jpareacode

import (
	"reflect"
	"testing"
)

func TestCityByName(t *testing.T) {
	e := &City{PrefCode: 13, Name: "千代田区", Code: 13101}
	if c := CityByName(13, "千代田区"); !reflect.DeepEqual(c, e) {
		t.Errorf("expected: %#v, actual: %#v", e, c)
	}
	e = nil
	if c := CityByName(12, "千代田区"); !reflect.DeepEqual(c, e) {
		t.Errorf("expected: %#v, actual: %#v", e, c)
	}
}

func TestCitiesByName(t *testing.T) {
	e := []City{{PrefCode: 13, Name: "千代田区", Code: 13101}}
	if c := CitiesByName("千代田区"); !reflect.DeepEqual(c, e) {
		t.Errorf("expected: %#v, actual: %#v", e, c)
	}
	e = nil
	if c := CitiesByName("千代田"); !reflect.DeepEqual(c, e) {
		t.Errorf("expected: %#v, actual: %#v", e, c)
	}
}

func TestCityByCode(t *testing.T) {
	e := &City{PrefCode: 13, Name: "千代田区", Code: 13101}
	if c := CityByCode(13101); !reflect.DeepEqual(c, e) {
		t.Errorf("expected: %#v, actual: %#v", e, c)
	}
	e = nil
	if c := CityByCode(1310); !reflect.DeepEqual(c, e) {
		t.Errorf("expected: %#v, actual: %#v", e, c)
	}
}

func TestFormatCityCode(t *testing.T) {
	if c := FormatCityCode(13101); c != "13101" {
		t.Errorf("expected: %s, actual: %s", "13101", c)
	}
	if c := FormatCityCode(0); c != "" {
		t.Errorf("expected: %s, actual: %s", "", c)
	}
}

func TestParseCityCode(t *testing.T) {
	if c := ParseCityCode("13101"); c != 13101 {
		t.Errorf("expected: %d, actual: %d", 13101, c)
	}
	if c := ParseCityCode(""); c != 0 {
		t.Errorf("expected: %d, actual: %d", 0, c)
	}
}

func TestValidateCityCode(t *testing.T) {
	if !ValidateCityCode(13101) {
		t.Error("expected: true, actual: false")
	}
	if ValidateCityCode(0) {
		t.Error("expected: false, actual: true")
	}
	if ValidateCityCode(50000) {
		t.Error("expected: false, actual: true")
	}
}
