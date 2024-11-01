package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 {
				fmt.Print("#")
			} else {
				if j == 0 || j == width-1 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
