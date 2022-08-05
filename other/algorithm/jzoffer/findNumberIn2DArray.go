package main
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i := len(matrix)-1 //y
	j := 0 //x
	for i > -1 {
		if j < len(matrix[i]) {
			if target > matrix[i][j] {
				j++
			} else if target < matrix[i][j] {
				i--
			} else if target == matrix[i][j] {
				return true
			}
		} else {
			return false
		}
	}
	return false

}
func main() {

	matrix := [][]int{
    {1,   4,  7, 11, 15},
	{2,   5,  8, 12, 19},
	{3,   6,  9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30}}
	print(findNumberIn2DArray(matrix, 96))
}
