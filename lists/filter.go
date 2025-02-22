package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	// Gehen Sie ähnlich wie bei Remove vor:
	// Wenn die Liste leer ist, ist das Ergebnis die leere Liste.
	// Wenn das erste Element größer als key ist, ist das Ergebnis die gefilterte Restliste.
	// Wenn das erste Element kleiner oder gleich key ist, ist das Ergebnis das erste Element
	// plus die gefilterte Restliste.
	result := []int{}
	if Empty(list) {
		return []int{}
	}

	if list[0] <= key {
		result = append(result, list[0])
	}
	rest := FilterLess(list[1:], key)
	result = append(result, rest...)
	return result
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	result := []int{}
	if Empty(list) {
		return []int{}
	}

	if list[0] >= key {
		result = append(result, list[0])
	}
	rest := FilterGreater(list[1:], key)
	result = append(result, rest...)
	return result

}
