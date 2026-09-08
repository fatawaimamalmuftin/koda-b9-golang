package choice

import "fmt"

func BuildWindow(window uint8) error{
	if window < 5 {
		return fmt.Errorf("window must be at least 5\n")
	}

	for vertical := uint8(0); vertical < window; vertical++ {
		for horizontal := uint8(0); horizontal < window; horizontal++ {
			if vertical == 0 || vertical == window-1 || horizontal == 0 || horizontal == window-1 {
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
	return nil
}
