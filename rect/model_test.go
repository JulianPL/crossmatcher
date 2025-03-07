package rect

import (
	"strings"
	"testing"
)

func TestModel_Solve(t *testing.T) {
	vRules := []string{"ab", "ab"}
	hRules := []string{"..", ".."}
	alphabet := "ab"
	candidate := []string{"..", ".."}

	model := NewModel(vRules, hRules, alphabet, candidate)
	actual := strings.Join(model.Solve(), "\n")
	expected := strings.Join([]string{"aa", "bb"}, "\n")
	if actual != expected {
		t.Errorf("Model.Solve is incorrect. Expected: %s, actual: %s", expected, actual)
	}

	vRules = []string{"ab", "ab"}
	hRules = []string{"ab", "ab"}
	alphabet = "ab"
	candidate = []string{"..", ".."}

	model = NewModel(vRules, hRules, alphabet, candidate)
	actual = strings.Join(model.Solve(), "\n")
	expected = strings.Join([]string{"##", "##"}, "\n")
	if actual != expected {
		t.Errorf("Model.Solve is incorrect. Expected: %s, actual: %s", expected, actual)
	}

	model = NewModelRandom("0", 3, 3)
	actual = strings.Join(model.Solve(), "\n")
	expected = strings.Join([]string{"000", "000", "000"}, "\n")
	if actual != expected {
		t.Errorf("Model.Solve is incorrect. Expected: %s, actual: %s", expected, actual)
	}
}
