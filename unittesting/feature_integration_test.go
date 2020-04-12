// +build integration

package feature_test

import (
	"testing"

	feature "github.com/goyod/online/unittesting"
)

func TestNextFibonacci(t *testing.T) {
	t.Run("0,1 => 1,1", func(t *testing.T) {
		a, b := 0, 1
		want := 4
		c, d := feature.NextFibonacci(a, b)
		if b != c && want != d {
			t.Errorf("%d,%d => %d,%d but get %d,%d\n", a, b, b, want, c, d)
		}
	})
	t.Run("1,1 => 1,2", func(t *testing.T) {
		a, b := 1, 1
		want := 2
		c, d := feature.NextFibonacci(a, b)
		if b != c && want != d {
			t.Errorf("%d,%d => %d,%d but get %d,%d\n", a, b, b, want, c, d)
		}
	})
	t.Run("1,2 => 2,3", func(t *testing.T) {
		a, b := 1, 2
		want := 3
		c, d := feature.NextFibonacci(a, b)
		if b != c && want != d {
			t.Errorf("%d,%d => %d,%d but get %d,%d\n", a, b, b, want, c, d)
		}
	})
	t.Run("2,3 => 3,5", func(t *testing.T) {
		a, b := 2, 3
		want := 5
		c, d := feature.NextFibonacci(a, b)
		if b != c && want != d {
			t.Errorf("%d,%d => %d,%d but get %d,%d\n", a, b, b, want, c, d)
		}
	})
	t.Run("3,5 => 5,8", func(t *testing.T) {
		a, b := 3, 5
		want := 8
		c, d := feature.NextFibonacci(a, b)
		if b != c && want != d {
			t.Errorf("%d,%d => %d,%d but get %d,%d\n", a, b, b, want, c, d)
		}
	})
	t.Run("5,8 => 8,13", func(t *testing.T) {
		a, b := 5, 8
		want := 13
		c, d := feature.NextFibonacci(a, b)
		if b != c && want != d {
			t.Errorf("%d,%d => %d,%d but get %d,%d\n", a, b, b, want, c, d)
		}
	})
}
