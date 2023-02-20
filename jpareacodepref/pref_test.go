package jpareacodepref

import (
	"testing"
)

func TestPrefectureCodeInt(t *testing.T) {
	if c := PrefectureCodeInt("北海道"); c != 1 {
		t.Errorf("expected: %d, actual: %d", 1, c)
	}
	if c := PrefectureCodeInt("北海"); c != 0 {
		t.Errorf("expected: %d, actual: %d", 0, c)
	}
}

func TestPrefectureCodeInts(t *testing.T) {
	c := PrefectureCodeInts("東京都", "北海道")
	if len(c) != 2 || c[0] != 13 || c[1] != 1 {
		t.Errorf("expected: %s, actual: %d,%d", "13,1", c[0], c[1])
	}
	c = PrefectureCodeInts("東京都", "北海")
	if len(c) != 2 || c[0] != 13 || c[1] != 0 {
		t.Errorf("expected: %s, actual: %d,%d", "13,0", c[0], c[1])
	}
}

func TestPrefectureCodeString(t *testing.T) {
	if c := PrefectureCodeString("北海道"); c != "01" {
		t.Errorf("expected: %s, actual: %s", "01", c)
	}
	if c := PrefectureCodeString("北海"); c != "" {
		t.Errorf("expected: %s, actual: %s", "", c)
	}
}

func TestPrefectureCodeStrings(t *testing.T) {
	c := PrefectureCodeStrings("東京都", "北海道")
	if len(c) != 2 || c[0] != "13" || c[1] != "01" {
		t.Errorf("expected: %s, actual: %s,%s", "13,1", c[0], c[1])
	}
	c = PrefectureCodeStrings("東京都", "北海")
	if len(c) != 2 || c[0] != "13" || c[1] != "" {
		t.Errorf("expected: %s, actual: %s,%s", "13,", c[0], c[1])
	}
}

func TestPrefectureName(t *testing.T) {
	if c := PrefectureName(1); c != "北海道" {
		t.Errorf("expected: %s, actual: %s", "北海道", c)
	}
	if c := PrefectureName(0); c != "" {
		t.Errorf("expected: %s, actual: %s", "", c)
	}
}

func TestPrefectureNames(t *testing.T) {
	c := PrefectureNames(13, 1)
	if len(c) != 2 || c[0] != "東京都" || c[1] != "北海道" {
		t.Errorf("expected: %s, actual: %s,%s", "東京都,北海道", c[0], c[1])
	}
	c = PrefectureNames(13, 0)
	if len(c) != 2 || c[0] != "東京都" || c[1] != "" {
		t.Errorf("expected: %s, actual: %s,%s", "東京都,", c[0], c[1])
	}
}

func TestFormatPrefectureCode(t *testing.T) {
	if c := FormatPrefectureCode(1); c != "01" {
		t.Errorf("expected: %s, actual: %s", "01", c)
	}
	if c := FormatPrefectureCode(0); c != "" {
		t.Errorf("expected: %s, actual: %s", "", c)
	}
}

func TestParsePrefectureCode(t *testing.T) {
	if c := ParsePrefectureCode("01"); c != 1 {
		t.Errorf("expected: %d, actual: %d", 1, c)
	}
	if c := ParsePrefectureCode("00"); c != 0 {
		t.Errorf("expected: %d, actual: %d", 0, c)
	}
}

func TestValidatePrefectureCode(t *testing.T) {
	if !ValidatePrefectureCode(1) {
		t.Error("expected: true, actual: false")
	}
	if ValidatePrefectureCode(0) {
		t.Error("expected: false, actual: true")
	}
}
