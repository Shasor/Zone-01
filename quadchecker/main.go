package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	input, _ := io.ReadAll(os.Stdin) // Par exemple : quand on fait " echo A | go run . " on demande si le résultat de " echo A " est un quad, et si oui lesquels
	x := 0
	y := 0
	result := ""
	/* if len(input) == 0 { // Si paramètre vide, on quitte le programme /!\ ne fonctionne pas
		os.Exit(0)
	} */
	for i := 0; i < len(input); i++ { // Compteur x et y
		if input[i] == 10 { // Quand on rencontre un retour à la ligne (10), on augmente notre nbr de lignes (y)
			y++
		}
		if input[i] != 10 { // Sinon, on augment notre nbr de colones (x)
			x++
		}
	}
	if y != 0 { // On division le nombre de caractère total par notre nbr de ligne = nbr de colones
		x /= y
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		os.Exit(0)
	}

	for i := 'A'; i <= 'E'; i++ { // on boucle avec notre index de A à E
		quad := exec.Command("./quad"+string(i), strconv.Itoa(x), strconv.Itoa(y)) // On exécute en arrière plan (silencieux), le vrai Quad
		commandQuad, _ := quad.Output()                                            // <-- et on stock la sortie de la commande ligne 34 (juste au-dessus) dans la variable "commandQuad"
		if string(input) == string(commandQuad) {                                  // On test que le paramètre stdin (résultat de ce qu'il y a avant le pipe) est égal au vrai quad[A-E]
			if result != "" {
				result += " || "
			}
			result += "[quad" + string(i) + "] [" + strconv.Itoa(x) + "] [" + strconv.Itoa(y) + "]" // si c'est égal, on ajoute le nom, le nbr de colones et le nbr de ligne à la variable result
		}
	}

	if result == "" { // Si la variable result est vide, càd qu'aucun vrai Quad[A-E] n'est égal au paramètre stdin, alors on print un msg d'erreur
		result = "Not a quad function"
	}
	fmt.Println(result) // sinon, on print result
}
