package shareddomain

import (
	"testing"
)

func Test_GetDaysFromNumberOfYears_ErrorWithNegative(t *testing.T) {
	days, err := GetDaysFromNumberOfYears(-1)
	if err == nil {
		t.Errorf("It should throw an error. Expected %s, got %v", OutOfRange, err)
	}
	if days != 0 {
		t.Errorf("It should throw an error. Expected %b, got %b", 0, days)
	}
}

func Test_GetDaysFromNumberOfYears_ErrorWithZero(t *testing.T) {
	days, err := GetDaysFromNumberOfYears(0)
	if err == nil {
		t.Errorf("It should throw an error. Expected %s, got %v", OutOfRange, err)
	}
	if days != 0 {
		t.Errorf("It should throw an error. Expected %b, got %b", 0, days)
	}
}

func Test_GetDaysFromNumberOfYears_ShouldBe365(t *testing.T) {
	days, err := GetDaysFromNumberOfYears(1)
	expectedDays := 365

	if err != nil {
		t.Errorf("It should not throw an error. Expected %v, got %v", nil, err)
	}
	if days != expectedDays {
		t.Errorf("Expected %b, got %b", expectedDays, days)
	}
}
