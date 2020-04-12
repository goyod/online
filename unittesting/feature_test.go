package feature

import "testing"

func TestOneAddTwo(t *testing.T) {
	a, b := 1, 2
	want := 3
	get := Add(a, b)
	if want != get {
		t.Errorf("%d+%d is %d but get %d\n", a, b, want, get)
	}
}

func Test_One_Add_Three(t *testing.T) {
	a, b := 1, 3
	want := 4
	get := Add(a, b)
	if want != get {
		t.Errorf("%d+%d is %d but get %d\n", a, b, want, get)
	}
}

func TestAdd(t *testing.T) {
	t.Run("case 2+2 = 4", func(t *testing.T) {
		a, b := 2, 2
		want := 4
		get := Add(a, b)
		if want != get {
			t.Errorf("%d+%d is %d but get %d\n", a, b, want, get)
		}
	})

	t.Run("case 4+4 = 8", func(t *testing.T) {
		a, b := 4, 4
		want := 8
		get := Add(a, b)
		if want != get {
			t.Errorf("%d+%d is %d but get %d\n", a, b, want, get)
		}
	})
}
