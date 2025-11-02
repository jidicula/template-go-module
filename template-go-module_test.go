package templategomod

import "testing"

func TestFoo(t *testing.T) {
	want := "Hello, world!"
	got := Foo()

	if want != got {
		t.Errorf("wrong message: got %v want %v", got, want)
	}
}
