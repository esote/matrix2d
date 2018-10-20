package matrix2d

// Point represents a point in the matrix, indexed by row and column.
type Point struct {
	x int
	y int
}

// Search2D searches a rectangular, sorted (column-wise and row-wise) matrix.
func Search2D(matrix [][]int, key int) (Point, bool) {
	badpt := Point{-1, -1}

	width := len(matrix)

	if width < 1 {
		return badpt, false
	}

	height := len(matrix[0])

	if height < 1 {
		return badpt, false
	}

	if height < width {
		ret, ok := Search2D(lazyTranspose(matrix), key)
		return Point{ret.y, ret.x}, ok
	}

	minc := 0
	maxr := height - 1
	diag := height / width

	for minc < width && maxr >= 0 {
		r := max(maxr-diag, 0)

		if e := matrix[minc][r]; key == e {
			return Point{minc, r}, true
		} else if key < e {
			maxr = r - 1
			continue
		}

		minrc := r + 1
		maxrc := maxr

		for minrc <= maxr {
			mid := minrc + (maxrc-minrc+1)/2

			if e := matrix[minc][mid]; key == e {
				return Point{minc, mid}, true
			} else if key < e {
				maxrc = mid - 1
				maxr = mid - 1
			} else {
				minrc = mid + 1
			}
		}

		minc++
	}

	return badpt, false
}

// Rotate a rectangular dim0 x dim1 matrix. Assumes matrix of valid size.
func lazyTranspose(matrix [][]int) [][]int {
	dim0 := len(matrix)
	dim1 := len(matrix[0])

	ret := make([][]int, dim1)

	for i := 0; i < dim1; i++ {
		ret[i] = make([]int, dim0)

		for j := 0; j < dim0; j++ {
			ret[i][j] = matrix[j][i]
		}
	}

	return ret
}

// Find maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
