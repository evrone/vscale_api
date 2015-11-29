package vscale

import (
	"testing"
)

func TestScalets(t *testing.T) {
	setup()
	defer teardown()

	scarlets, err := client.Scalets()

	if err != nil {
		t.Errorf("Can`t get scalets: %s`", err)
	}

	t.Logf("Scalets %+v", scarlets)
}
