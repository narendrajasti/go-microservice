package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Ice cream",
		Price: 1.00,
		SKU:   "aaa-aaaa-aaa",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
