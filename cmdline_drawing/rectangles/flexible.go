package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for row := 0; row < height; row++ {
		// Zeichne eine Zeile
		for column := 0; column < width; column++ {
			if row == 0 || column == 0 || row == height-1 || column == width-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()
	}
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
