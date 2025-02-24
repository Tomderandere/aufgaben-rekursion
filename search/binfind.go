package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	if len(list) == 1 {
		return -1
	}
	result := FindSorted(list[1:], x)
	if list[0] == x {
		return result
	}
	if result == -1 {
		return -1
	}

	return result + 1

}
