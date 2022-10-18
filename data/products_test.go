package data

import "testing"

func TestChecksValidation(t *testing.T){
	p := Product{
		Name: "ice tea",
		Price: 1.00,
		SKU: "abs--shb",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}