package app

import "testing"

func Test_RealRandomizer(t *testing.T) {

	r := NewRealRandomizer()
	got1 := r.Intn(9999999999999)
	r = NewRealRandomizer()
	got2 := r.Intn(9999999999999)
	got3 := r.Intn(9999999999999)

	if got1 == got2 {
		t.Errorf("expecting different numbers but got %v twice in a row after reinstantiating", got1)
	}

	if got2 == got3 {
		t.Errorf("expecting different numbers but got %v twice in a row with consecutive calls", got1)
	}
}
