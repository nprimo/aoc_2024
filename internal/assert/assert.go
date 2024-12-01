package assert

import "testing"

func Equal(t *testing.T, want any, got any) {
	if want != got {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
