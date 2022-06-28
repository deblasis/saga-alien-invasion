package app

import "testing"

func Test_RealRandomizer(t *testing.T) {
	r := NewRealRandomizer()
	got1 := r.Intn(9999999)
	r = NewRealRandomizer()
	got2 := r.Intn(9999999)
	got3 := r.Intn(9999999)

	if got1 == got2 {
		t.Errorf("expecting different numbers but got %v twice in a row after reinstantiating", got1)
	}

	if got2 == got3 {
		t.Errorf("expecting different numbers but got %v twice in a row with consecutive calls", got1)
	}
}

func Test_FakeRandomizer(t *testing.T) {
	r := NewFakeRandomizer()
	got1 := r.Intn(9999999)
	r = NewFakeRandomizer()
	got2 := r.Intn(9999999)
	got3 := r.Intn(9999999)

	if got1 != got2 {
		t.Errorf("expecting the same numbers but got %v and then %v after reinstantiating", got1, got2)
	}

	if got2 == got3 {
		t.Errorf("expecting different numbers but got %v twice in a row with consecutive calls", got1)
	}
}
