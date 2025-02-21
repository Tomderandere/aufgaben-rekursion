package lists

// Liefert das Produkt aller Elemente in list.
// Wenn list leer ist, wird 1 zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func Product(list []int) int {
	if Empty(list) {
		return 1
	}

	return Product(list[1:]) * list[0]
}
