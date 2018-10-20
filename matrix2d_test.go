package matrix2d

import (
	"testing"
)

// Valid matrix.
func TestSearch2D(t *testing.T) {
	// 3x5 matrix
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 5, 7},
		{1, 4, 9, 16, 25},
	}

	item := 3
	want := Point{0, 2}

	got, ok := Search2D(matrix, item)

	if !ok {
		t.Errorf("want: %t, got %t", true, ok)
	}

	if got != want {
		t.Errorf("want: (%d, %d), got: (%d, %d)", want.x, want.y, got.x, got.y)
	}

	if e := matrix[got.x][got.y]; e != item {
		t.Errorf("want: %d, got: %d", item, e)
	}
}

// Valid matrix, but needs to be transposed.
func TestSearch2DRotated(t *testing.T) {
	// 5x3 matrix
	matrix := [][]int{
		{1, 1, 1},
		{2, 2, 4},
		{3, 3, 9},
		{16, 5, 4},
		{25, 7, 5},
	}

	item := 16
	want := Point{3, 0}

	got, ok := Search2D(matrix, item)

	if !ok {
		t.Errorf("want: %t, got %t", true, ok)
	}

	if got != want {
		t.Errorf("want: (%d, %d), got: (%d, %d)", want.x, want.y, got.x, got.y)
	}

	if e := matrix[got.x][got.y]; e != item {
		t.Errorf("want: %d, got: %d", item, e)
	}
}

// Invalid width (zero columns).
func TestSearch2DWidth(t *testing.T) {
	var matrix [][]int

	got, ok := Search2D(matrix, 0)

	want := Point{-1, -1}

	if ok {
		t.Errorf("want: %t, got %t", false, ok)
	}

	if got != want {
		t.Errorf("want: (%d, %d), got: (%d, %d)", want.x, want.y, got.x, got.y)
	}
}

// Invalid height (zero rows).
func TestSearch2DHeight(t *testing.T) {
	matrix := make([][]int, 3)
	matrix[0] = make([]int, 0)

	got, ok := Search2D(matrix, 0)

	want := Point{-1, -1}

	if ok {
		t.Errorf("want: %t, got %t", false, ok)
	}

	if got != want {
		t.Errorf("want: (%d, %d), got: (%d, %d)", want.x, want.y, got.x, got.y)
	}
}

// Not found.
func TestSearch2DNotFound(t *testing.T) {
	// 3x5 matrix
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 5, 7},
		{1, 4, 9, 16, 25},
	}

	item := 15
	want := Point{-1, -1}

	got, ok := Search2D(matrix, item)

	if ok {
		t.Errorf("want: %t, got %t", false, ok)
	}

	if got != want {
		t.Errorf("want: (%d, %d), got: (%d, %d)", want.x, want.y, got.x, got.y)
	}
}

// After middle, checking branch used with find-middle on even lengths.
func TestSearch2DAfterMid(t *testing.T) {
	// 3x6 matrix
	matrix := [][]int{
		{1, 2, 3, 4, 5, 9},
	}

	item := 5
	want := Point{0, 4}
	got, ok := Search2D(matrix, item)

	if !ok {
		t.Errorf("want: %t, got %t", false, ok)
	}

	if got != want {
		t.Errorf("want: (%d, %d), got: (%d, %d)", want.x, want.y, got.x, got.y)
	}

	if e := matrix[got.x][got.y]; e != item {
		t.Errorf("want: %d, got: %d", item, e)
	}
}
