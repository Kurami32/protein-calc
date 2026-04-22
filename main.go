package main

import (
	"fmt"
	"os"
	"strings"
)

// Tabla de masas de aminoácidos estándar (código de una letra)
// Fuente: https://www.sb-peptide.com/support/amino-acids-reference-chart
var massTable = map[byte]float64{
	'A': 89.09, 'R': 174.20, 'N': 132.12, 'D': 133.10,
	'C': 121.16, 'Q': 146.15, 'E': 147.13, 'G': 75.07,
	'H': 155.16, 'I': 131.17, 'L': 131.17, 'K': 146.19,
	'M': 149.21, 'F': 165.19, 'P': 115.13, 'S': 105.09,
	'T': 119.12, 'W': 204.23, 'Y': 181.19, 'V': 117.15,
}

// Validar si es un aminoácido
func isValidAminoAcid(c byte) bool {
	_, ok := massTable[c]
	return ok
}

// Calcular el peso molecular (Da) de una secuencia
func calcMolecularWeight(seq string) (float64, error) {
	seq = strings.ToUpper(seq) // Convertir a mayúsculas
	var total float64

	// Recorrer cada letra (Solo ASCII válido, por ejemplo si metemos "Ñ", regresará error)
	for _, r := range seq {
		c := byte(r)

		if r > 127 || !isValidAminoAcid(c) {
			return 0, fmt.Errorf("carácter inválido '%c' (no es aminoácido)", r)
		}
		total += massTable[c]
	}
	return total, nil
}

func main() {
	var seq string

	// Leer secuencia desde argumento (por ejemplo: go run main.go "AMINOACIDO") por defecto usará la hemoglobina
	// ya que es la que elegimos para el informe.
	if len(os.Args) > 1 {
		seq = os.Args[1]
	} else {
		// Hemoglobina subunidad alfa humana (141 aa, UniProt P69905)
		// https://www.uniprot.org/uniprotkb/P69905/entry
		seq = "MVLSPADKTNVKAAWGKVGAHAGEYGAEALERMFLSFPTTKTYFPHFDLSHGSAQVKGHGKKVADALTNAVAHVDDMPNALSALSDLHAHKLRVDPVNFKLLSCHLLVTLAAHLPAEFTPAVHASLDKFLASVSTVLTSKYR"
		fmt.Println("No se especificó una secuencia -- se usará la hemoglobina")
	}

	// Calcular peso molecular
	mw, err := calcMolecularWeight(seq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Mostrar resultados
	fmt.Printf("Peso molecular: %.2f Da\n", mw)
	fmt.Printf("= %.2f kDa\n", mw/1000.0)
	fmt.Println("\nPresiona Enter para salir...")
	fmt.Scanln()
}
