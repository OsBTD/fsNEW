package ascii

import (
	"fmt"
)

func PrintArt(string, []string, map[rune]([]string)) {
	if ContainsOnly(input) {
		for i := 0; i < len(input)/2; i++ {
			fmt.Println()
		}
		return
	}
	for _, line := range inputsplit {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(line); j++ {
				inputrune := rune(line[j])
				fmt.Print(Replace[inputrune][i])
			}
			fmt.Println()
		}
	}
}
