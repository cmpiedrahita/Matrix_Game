package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("BIENVENIDO A MATRIX ESCAPE")
	fmt.Println("Neo debe escapar llegando a un teléfono")
	fmt.Println("Los agentes intentarán atraparlo")
	fmt.Println("=====================================")

	game := NewGame()
	game.Start()

	// Esperar un poco para que termine el juego
	time.Sleep(100 * time.Millisecond)
}
