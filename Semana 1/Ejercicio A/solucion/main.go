package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var n int
		var err error

		for {
			fmt.Println("Introduce el número de personas (0 para salir):")
			scanner.Scan()
			n, err = strconv.Atoi(scanner.Text())

			if err == nil && n >= 0 && n < 100 {
				break
			}
			fmt.Println("Error: entrada inválida. Por favor, introduce un número válido.")
		}

		if n == 0 {
			return
		}

		var edades []int

		for {
			fmt.Printf("Introduce %d edades separadas por espacio:\n", n)
			scanner.Scan()
			edadesStrs := strings.Fields(scanner.Text())

			if len(edadesStrs) != n {
				fmt.Printf("Error: se esperaban %d edades, pero se recibieron %d. Por favor, introduce los valores nuevamente.\n", n, len(edadesStrs))
				continue
			}

			valid := true
			edades = make([]int, n)
			for i, edadStr := range edadesStrs {
				edades[i], err = strconv.Atoi(edadStr)
				if err != nil {
					fmt.Println("Error: entrada de edad inválida. Por favor, introduce los valores nuevamente.")
					valid = false
					break
				}
			}

			if valid {
				break
			}
		}

		// Ordenar edades (Bubble Sort corregido)
		for i := 0; i < len(edades)-1; i++ {
			for j := 0; j < len(edades)-1-i; j++ {
				if edades[j] > edades[j+1] {
					edades[j], edades[j+1] = edades[j+1], edades[j]
				}
			}
		}

		// Imprimir edades ordenadas
		for i, edad := range edades {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(edad)
		}
		fmt.Println()
	}
}
