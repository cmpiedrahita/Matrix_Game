package main

import (
	"math"
	"math/rand"
)

// Agente - los antagonistas que persiguen a Neo
type Agent struct {
	id    int
	board *Board
}

// Crear nuevo agente
func NewAgent(id int, board *Board) *Agent {
	return &Agent{id: id, board: board}
}

// Ejecutar la estrategia del Agente 1 (como goroutine)
func (a *Agent) RunAgent1(moveSignal <-chan bool, coordination chan Position) {
	for {
		select {
		case <-moveSignal:
			// Es el turno del agente para moverse
			move := a.calculateMoveAgent1(coordination)
			a.board.MoveAgent1(move)
		}
	}
}

// Ejecutar la estrategia del Agente 2 (como goroutine)
func (a *Agent) RunAgent2(moveSignal <-chan bool, coordination chan Position) {
	for {
		select {
		case <-moveSignal:
			// Es el turno del agente para moverse
			move := a.calculateMoveAgent2(coordination)
			a.board.MoveAgent2(move)
		}
	}
}

// Estrategia del Agente 1 (más agresivo, persigue directamente)
func (a *Agent) calculateMoveAgent1(coordination chan Position) Position {
	currentNeo, currentAgent1, currentAgent2 := a.board.GetPositions()
	
	// Enviar posición de Neo al otro agente para coordinación
	select {
	case coordination <- currentNeo:
	default:
		// Si el canal está lleno, continuar sin bloquear
	}
	
	// Estrategia: perseguir directamente a Neo
	possibleMoves := a.getPossibleMoves(currentAgent1)
	if len(possibleMoves) == 0 {
		return currentAgent1
	}
	
	bestMove := currentAgent1
	minDistance := math.MaxFloat64
	
	for _, move := range possibleMoves {
		// Evitar chocar con el otro agente
		if move == currentAgent2 {
			continue
		}
		
		distance := a.calculateDistance(move, currentNeo)
		if distance < minDistance {
			minDistance = distance
			bestMove = move
		}
	}
	
	return bestMove
}

// Estrategia del Agente 2 (más estratégico, intenta bloquear escape)
func (a *Agent) calculateMoveAgent2(coordination chan Position) Position {
	currentNeo, currentAgent1, currentAgent2 := a.board.GetPositions()
	
	// Intentar recibir información del otro agente
	var neoPos Position
	select {
	case neoPos = <-coordination:
		// Usar la información recibida
	default:
		// Si no hay información, usar la posición actual de Neo
		neoPos = currentNeo
	}
	
	// Estrategia: posicionarse entre Neo y el teléfono más cercano
	possibleMoves := a.getPossibleMoves(currentAgent2)
	if len(possibleMoves) == 0 {
		return currentAgent2
	}
	
	// Determinar qué teléfono está más cerca de Neo
	distPhone1 := a.calculateDistance(neoPos, a.board.Phone1)
	distPhone2 := a.calculateDistance(neoPos, a.board.Phone2)
	
	targetPhone := a.board.Phone1
	if distPhone2 < distPhone1 {
		targetPhone = a.board.Phone2
	}
	
	// Encontrar la mejor posición para interceptar
	bestMove := currentAgent2
	bestScore := float64(-1000)
	
	for _, move := range possibleMoves {
		// Evitar chocar con el otro agente
		if move == currentAgent1 {
			continue
		}
		
		score := a.evaluateInterceptMove(move, neoPos, targetPhone)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	
	return bestMove
}

// Evaluar movimiento de intercepción
func (a *Agent) evaluateInterceptMove(move, neoPos, targetPhone Position) float64 {
	score := 0.0
	
	// Bonus por estar cerca de Neo
	distToNeo := a.calculateDistance(move, neoPos)
	score += 50.0 / (distToNeo + 1)
	
	// Bonus por estar en el camino entre Neo y el teléfono
	distNeoToPhone := a.calculateDistance(neoPos, targetPhone)
	distMoveToPhone := a.calculateDistance(move, targetPhone)
	distMoveToNeo := a.calculateDistance(move, neoPos)
	
	// Si estamos aproximadamente en línea entre Neo y el teléfono
	if distMoveToNeo + distMoveToPhone <= distNeoToPhone + 1 {
		score += 30.0
	}
	
	// Añadir aleatoriedad
	score += rand.Float64() * 3
	
	return score
}

// Calcular distancia Manhattan
func (a *Agent) calculateDistance(pos1, pos2 Position) float64 {
	return math.Abs(float64(pos1.X-pos2.X)) + math.Abs(float64(pos1.Y-pos2.Y))
}

// Obtener movimientos posibles
func (a *Agent) getPossibleMoves(current Position) []Position {
	moves := []Position{}
	directions := []Position{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0}, // arriba, abajo, izquierda, derecha
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // diagonales
	}
	
	for _, dir := range directions {
		newPos := Position{current.X + dir.X, current.Y + dir.Y}
		if a.board.isValidPosition(newPos) {
			moves = append(moves, newPos)
		}
	}
	
	return moves
}