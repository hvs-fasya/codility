package matrix_longest_increasing_path

type MatrixItem struct {
	I int
	J int
}

func (m Matrix) Item(i, j int) int {
	return m[i][j]
}

type Matrix [][]int

func Solution(matrix [][]int) int {
	store := make(map[MatrixItem][]MatrixItem)
	maxLen := 0
	for i := range matrix {
		for j := range matrix[i] {
			paths := Matrix(matrix).Search(MatrixItem{i, j}, store)
			if len(paths) > maxLen {
				maxLen = len(paths)
			}
		}
	}
	return maxLen
}

func (m Matrix) Search(item MatrixItem, store map[MatrixItem][]MatrixItem) []MatrixItem {
	if _, ok := store[item]; ok {
		return store[item]
	}
	neighbours := getNeighbours(item.I, item.J, m)
	neighboursPath := []MatrixItem{}
	for _, neighbour := range neighbours {
		path := m.Search(neighbour, store)
		if len(path) > len(neighboursPath) {
			neighboursPath = append([]MatrixItem{}, path...)
		}
	}
	store[item] = append([]MatrixItem{item}, neighboursPath...)
	return store[item]
}

func getNeighbours(i, j int, m Matrix) []MatrixItem {
	res := make([]MatrixItem, 0)
	if j < len(m[i])-1 && m[i][j+1] > m[i][j] {
		res = append(res, MatrixItem{i, j + 1})
	}
	if i < len(m)-1 && m[i+1][j] > m[i][j] {
		res = append(res, MatrixItem{i + 1, j})
	}
	if i > 0 && m[i-1][j] > m[i][j] {
		res = append(res, MatrixItem{i - 1, j})
	}
	if j > 0 && m[i][j-1] > m[i][j] {
		res = append(res, MatrixItem{i, j - 1})
	}
	return res
}
