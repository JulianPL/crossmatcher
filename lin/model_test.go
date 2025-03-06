package lin

import "testing"

func TestModel_Solve(t *testing.T) {
	rule := "ab(a|b)+ba"
	alphabet := "ab"
	candidate := "......"

	model := NewModel(rule, alphabet, candidate)
	actual := model.Solve()
	expected := "ab..ba"
	if actual != expected {
		t.Errorf("Model.Solve is incorrect. Expected: %s, actual: %s", expected, actual)
	}
	candidate = "b....."
	model = NewModel(rule, alphabet, candidate)
	actual = model.Solve()
	expected = "######"
	if actual != expected {
		t.Errorf("Model.Solve is incorrect. Expected: %s, actual: %s", expected, actual)
	}
}
