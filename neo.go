package main

import (
	"math"
	"math/rand"
	"time"
)

// Neo - el protagonista que debe escapar
type Neo struct {
	board *Board
}

// Crear nuevo Neo
func NewNeo(board *Board) *Neo {
	return &Neo{board: board}
}

// Ejecutar la estrategia de Neo (como goroutine)
func (n *Neo) Run(moveSignal <-chan bool, gameOver chan<- string) {
	for {
		select {
		case <-moveSignal:
			// Es el turno de Neo para moverse
			n.makeMove()

			// Verificar si Neo escapó
			if n.board.NeoEscaped() {
				gameOver <- "¡NEO ESCAPÓ! ¡Ha llegado al teléfono!"
				return
			}

			// Verificar si Neo fue atrapado
			if n.board.NeoCaught() {
				gameOver <- "¡NEO FUE ATRAPADO! Los agentes ganaron."
				return
			}
		}
	}
}

// Estrategia de movimiento de Neo
func (n *Neo) makeMove() {
	currentNeo, agent1, agent2 := n.board.GetPositions()

	// Calcular distancias a los teléfonos
	distPhone1 := n.calculateDistance(currentNeo, n.board.Phone1)
	distPhone2 := n.calculateDistance(currentNeo, n.board.Phone2)

	// Elegir el teléfono más cercano como objetivo
	target := n.board.Phone1
	if distPhone2 < distPhone1 {
		target = n.board.Phone2
	}

	// Generar posibles movimientos
	possibleMoves := n.getPossibleMoves(currentNeo)

	if len(possibleMoves) == 0 {
		return // No hay movimientos posibles
	}

	// Evaluar cada movimiento
	bestMove := currentNeo
	bestScore := float64(-1000)

	for _, move := range possibleMoves {
		score := n.evaluateMove(move, target, agent1, agent2)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	// Realizar el mejor movimiento
	n.board.MoveNeo(bestMove)
}

// Calcular distancia Manhattan entre dos posiciones
func (n *Neo) calculateDistance(pos1, pos2 Position) float64 {
	return math.Abs(float64(pos1.X-pos2.X)) + math.Abs(float64(pos1.Y-pos2.Y))
}

// Obtener movimientos posibles
func (n *Neo) getPossibleMoves(current Position) []Position {
	moves := []Position{}
	directions := []Position{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0}, // arriba, abajo, izquierda, derecha
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // diagonales
	}

	for _, dir := range directions {
		newPos := Position{current.X + dir.X, current.Y + dir.Y}
		if n.board.isValidPosition(newPos) {
			moves = append(moves, newPos)
		}
	}

	return moves
}

// Evaluar qué tan bueno es un movimiento
func (n *Neo) evaluateMove(move, target, agent1, agent2 Position) float64 {
	score := 0.0

	// Bonus por acercarse al teléfono objetivo
	distToTarget := n.calculateDistance(move, target)
	score += 100.0 / (distToTarget + 1)

	// Penalización por acercarse a los agentes
	distToAgent1 := n.calculateDistance(move, agent1)
	distToAgent2 := n.calculateDistance(move, agent2)

	if distToAgent1 < 2 {
		score -= 50.0 / (distToAgent1 + 0.1)
	}
	if distToAgent2 < 2 {
		score -= 50.0 / (distToAgent2 + 0.1)
	}

	// Añadir algo de aleatoriedad para hacer el juego más interesante
	score += rand.Float64() * 5

	return score
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
