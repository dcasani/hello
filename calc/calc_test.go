package calc

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	} 
}

func TestDivide(t *testing.T) {
    got := Divide(10, 2)
    want := 5    
    if got != want {
        t.Errorf("Divide(10, 2) = %d; want %d", got, want)
    }    
}