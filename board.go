package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Posición en el tablero
type Position struct {
	X, Y int
}

// Tablero del juego
type Board struct {
	Size      int
	Neo       Position
	Agent1    Position
	Agent2    Position
	Phone1    Position
	Phone2    Position
	mutex     sync.RWMutex // Para sincronizar acceso al tablero
}

// Crear nuevo tablero con posiciones aleatorias
func NewBoard() *Board {
	rand.Seed(time.Now().UnixNano())
	
	// Generar posiciones únicas aleatorias
	usedPositions := make(map[Position]bool)
	
	getRandomPosition := func() Position {
		for {
			pos := Position{rand.Intn(8), rand.Intn(8)}
			if !usedPositions[pos] {
				usedPositions[pos] = true
				return pos
			}
		}
	}
	
	board := &Board{
		Size:   8,
		Neo:    getRandomPosition(),
		Agent1: getRandomPosition(),
		Agent2: getRandomPosition(),
		Phone1: getRandomPosition(),
		Phone2: getRandomPosition(),
	}
	
	// Verificar que Neo no esté en un teléfono al inicio
	for board.Neo == board.Phone1 || board.Neo == board.Phone2 {
		// Regenerar posición de Neo
		delete(usedPositions, board.Neo)
		board.Neo = getRandomPosition()
	}
	
	return board
}

// Mostrar el tablero (thread-safe)
func (b *Board) Display() {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	
	fmt.Println("\n📱 MATRIX BOARD 📱")
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			pos := Position{x, y}
			symbol := "⬛" // Espacio vacío
			
			if pos == b.Neo {
				symbol = "🕴️" // Neo
			} else if pos == b.Agent1 {
				symbol = "👤" // Agente 1
			} else if pos == b.Agent2 {
				symbol = "👥" // Agente 2
			} else if pos == b.Phone1 || pos == b.Phone2 {
				symbol = "📞" // Teléfono
			}
			
			fmt.Print(symbol + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// Mover Neo (thread-safe)
func (b *Board) MoveNeo(newPos Position) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if b.isValidPosition(newPos) {
		b.Neo = newPos
		return true
	}
	return false
}

// Mover Agente 1 (thread-safe)
func (b *Board) MoveAgent1(newPos Position) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if b.isValidPosition(newPos) {
		b.Agent1 = newPos
		return true
	}
	return false
}

// Mover Agente 2 (thread-safe)
func (b *Board) MoveAgent2(newPos Position) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if b.isValidPosition(newPos) {
		b.Agent2 = newPos
		return true
	}
	return false
}

// Verificar si la posición es válida
func (b *Board) isValidPosition(pos Position) bool {
	return pos.X >= 0 && pos.X < b.Size && pos.Y >= 0 && pos.Y < b.Size
}

// Obtener posiciones actuales (thread-safe)
func (b *Board) GetPositions() (Position, Position, Position) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.Neo, b.Agent1, b.Agent2
}

// Verificar si Neo llegó a un teléfono
func (b *Board) NeoEscaped() bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.Neo == b.Phone1 || b.Neo == b.Phone2
}

// Verificar si Neo fue atrapado
func (b *Board) NeoCaught() bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.Neo == b.Agent1 || b.Neo == b.Agent2
}