package main

import (
	"fmt"
	"time"
)

// Juego principal
type Game struct {
	board        *Board
	neo          *Neo
	agent1       *Agent
	agent2       *Agent
	moveSignal   chan bool
	gameOver     chan string
	coordination chan Position // Canal para coordinación entre agentes
}

// Crear nuevo juego
func NewGame() *Game {
	board := NewBoard()
	return &Game{
		board:        board,
		neo:          NewNeo(board),
		agent1:       NewAgent(1, board),
		agent2:       NewAgent(2, board),
		moveSignal:   make(chan bool, 3), // Buffer para 3 personajes
		gameOver:     make(chan string, 1),
		coordination: make(chan Position, 2), // Canal para coordinación entre agentes
	}
}

// Iniciar el juego
func (g *Game) Start() {
	fmt.Println("Iniciando Matrix Escape...")
	g.board.Display()

	// Iniciar las goroutines para cada personaje
	go g.neo.Run(g.moveSignal, g.gameOver)
	go g.agent1.RunAgent1(g.moveSignal, g.coordination)
	go g.agent2.RunAgent2(g.moveSignal, g.coordination)

	// Bucle principal del juego
	turnCount := 0
	maxTurns := 50 // Límite de turnos para evitar juegos infinitos

	for turnCount < maxTurns {
		turnCount++
		fmt.Printf("Turno %d\n", turnCount)

		// Enviar señal de movimiento a todos los personajes
		// Todos se mueven simultáneamente (concurrencia)
		for i := 0; i < 3; i++ {
			g.moveSignal <- true
		}

		// Pequeña pausa para que se procesen los movimientos
		time.Sleep(200 * time.Millisecond)

		// Mostrar el estado actual del tablero
		g.board.Display()

		// Verificar condiciones de fin de juego
		select {
		case result := <-g.gameOver:
			fmt.Println(result)
			g.showGameStats(turnCount)
			return
		default:
			// Continuar el juego
		}

		// Pausa entre turnos para visualización
		time.Sleep(800 * time.Millisecond)
	}

	// Si llegamos aquí, el juego terminó por límite de turnos
	fmt.Println("¡Tiempo agotado! Neo logró sobrevivir pero no escapó.")
	g.showGameStats(turnCount)
}

// Mostrar estadísticas del juego
func (g *Game) showGameStats(turns int) {
	fmt.Println("\nESTADÍSTICAS DEL JUEGO:")
	fmt.Printf("Turnos jugados: %d\n", turns)

	neo, agent1, agent2 := g.board.GetPositions()
	fmt.Printf("Posición final de Neo: (%d, %d)\n", neo.X, neo.Y)
	fmt.Printf("Posición final del Agente 1: (%d, %d)\n", agent1.X, agent1.Y)
	fmt.Printf("Posición final del Agente 2: (%d, %d)\n", agent2.X, agent2.Y)

	// Calcular distancias finales
	distPhone1 := g.calculateDistance(neo, g.board.Phone1)
	distPhone2 := g.calculateDistance(neo, g.board.Phone2)
	fmt.Printf("Distancia a Teléfono 1: %.1f\n", distPhone1)
	fmt.Printf("Distancia a Teléfono 2: %.1f\n", distPhone2)

	fmt.Println("\nGracias por jugar Matrix Escape")
}

// Calcular distancia Manhattan
func (g *Game) calculateDistance(pos1, pos2 Position) float64 {
	dx := float64(pos1.X - pos2.X)
	dy := float64(pos1.Y - pos2.Y)
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}
