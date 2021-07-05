package model_test

import "testing"

func diffError(t *testing.T, name, diff string) {
	t.Errorf("Name:%s mismatch (-want +got):\n%s", name, diff)
}
