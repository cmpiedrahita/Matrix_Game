# 🕴️ Matrix Escape Game

<div align="center">

![Matrix](https://img.shields.io/badge/Matrix-Escape-green?style=for-the-badge&logo=matrix&logoColor=white)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Concurrency](https://img.shields.io/badge/Concurrency-Goroutines-blue?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)

*Un juego de escape inspirado en The Matrix, implementado en Go con programación concurrente*

</div>

## 🎯 Descripción

**Matrix Escape** es un juego de estrategia en tiempo real donde Neo debe escapar de los Agentes llegando a uno de los teléfonos disponibles en el tablero. El juego utiliza goroutines para simular el movimiento simultáneo de todos los personajes, creando una experiencia de juego dinámica y desafiante.

### 🎮 Características Principales

- **🔄 Concurrencia Real**: Todos los personajes se mueven simultáneamente usando goroutines
- **🧠 IA Inteligente**: Los agentes tienen diferentes estrategias de persecución
- **📱 Tablero Dinámico**: Tablero 8x8 con posiciones aleatorias al inicio
- **🎲 Elementos Aleatorios**: Cada partida es única
- **📊 Estadísticas**: Seguimiento detallado del progreso del juego

## 🎭 Personajes

| Personaje | Emoji | Descripción | Estrategia |
|-----------|-------|-------------|------------|
| **Neo** | 🕴️ | El protagonista que debe escapar | Busca el teléfono más cercano evitando agentes |
| **Agente 1** | 👤 | Perseguidor agresivo | Persigue directamente a Neo |
| **Agente 2** | 👥 | Estratega defensivo | Intenta bloquear el escape de Neo |
| **Teléfonos** | 📞 | Puntos de escape | Objetivos de Neo para ganar |

## 🚀 Instalación y Ejecución

### Prerrequisitos

- **Go 1.21+** instalado en tu sistema
- Terminal o línea de comandos

### Pasos de Instalación

1. **Clona el repositorio**
   ```bash
   git clone <tu-repositorio>
   cd Matrix_Game
   ```

2. **Inicializa el módulo Go** (si es necesario)
   ```bash
   go mod init matrix-game
   ```

3. **Ejecuta el juego**
   ```bash
   go run .
   ```

## 🎲 Cómo Jugar

### Objetivo
- **Neo**: Llegar a cualquiera de los dos teléfonos (📞) para escapar
- **Agentes**: Atrapar a Neo antes de que escape

### Mecánicas del Juego

1. **Inicio**: Todos los personajes aparecen en posiciones aleatorias
2. **Turnos**: Cada turno, todos los personajes se mueven simultáneamente
3. **Movimiento**: Los personajes pueden moverse en 8 direcciones (incluidas diagonales)
4. **Victoria**: 
   - Neo gana si llega a un teléfono
   - Los Agentes ganan si atrapan a Neo
   - Empate si se agotan los turnos (50 turnos máximo)

### Ejemplo de Tablero

```
📱 MATRIX BOARD 📱
⬛ ⬛ 📞 ⬛ ⬛ ⬛ ⬛ ⬛ 
⬛ 👤 ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ 
⬛ ⬛ ⬛ 🕴️ ⬛ ⬛ ⬛ ⬛ 
⬛ ⬛ ⬛ ⬛ ⬛ ⬛ 👥 ⬛ 
⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ 
⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ 
⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ 📞 
⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ ⬛ 
```

## 🏗️ Arquitectura del Código

### Estructura de Archivos

```
Matrix_Game/
├── main.go          # Punto de entrada del programa
├── game.go          # Lógica principal del juego
├── board.go         # Gestión del tablero y posiciones
├── neo.go           # Comportamiento e IA de Neo
├── agent.go         # Comportamiento e IA de los Agentes
├── go.mod           # Configuración del módulo Go
└── README.md        # Este archivo
```

### Componentes Principales

#### 🎮 Game (`game.go`)
- Orquesta todo el juego
- Maneja la comunicación entre goroutines
- Controla el flujo de turnos
- Genera estadísticas finales

#### 🏁 Board (`board.go`)
- Gestiona el tablero 8x8
- Controla las posiciones de todos los personajes
- Implementa thread-safety con mutex
- Valida movimientos

#### 🕴️ Neo (`neo.go`)
- IA que busca el teléfono más cercano
- Evita a los agentes usando evaluación de riesgo
- Toma decisiones basadas en distancias Manhattan

#### 👤 Agent (`agent.go`)
- **Agente 1**: Estrategia agresiva de persecución directa
- **Agente 2**: Estrategia defensiva de intercepción
- Coordinación entre agentes usando canales

## 🔧 Características Técnicas

### Concurrencia
- **Goroutines**: Cada personaje ejecuta en su propia goroutine
- **Canales**: Comunicación segura entre goroutines
- **Mutex**: Protección de datos compartidos (tablero)
- **Sincronización**: Movimientos simultáneos coordinados

### Algoritmos de IA
- **Distancia Manhattan**: Cálculo de distancias en el tablero
- **Evaluación de Movimientos**: Sistema de puntuación para decisiones
- **Coordinación**: Los agentes comparten información estratégica

### Seguridad de Hilos
- Todas las operaciones del tablero son thread-safe
- Uso de `sync.RWMutex` para lecturas/escrituras concurrentes
- Canales con buffer para evitar bloqueos

## 📊 Estadísticas del Juego

Al final de cada partida, el juego muestra:
- Número de turnos jugados
- Posiciones finales de todos los personajes
- Distancias finales a los teléfonos
- Resultado de la partida

## 🛠️ Personalización

### Modificar el Tamaño del Tablero
En `board.go`, cambia el valor de `Size`:
```go
board := &Board{
    Size: 10, // Cambiar de 8 a 10 para un tablero 10x10
    // ...
}
```

### Ajustar Límite de Turnos
En `game.go`, modifica `maxTurns`:
```go
maxTurns := 100 // Cambiar de 50 a 100 turnos
```

### Personalizar Estrategias de IA
Modifica las funciones de evaluación en `neo.go` y `agent.go` para cambiar el comportamiento de los personajes.

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si tienes ideas para mejorar el juego:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Ideas para Futuras Mejoras

- [ ] Interfaz gráfica con GUI
- [ ] Múltiples niveles de dificultad
- [ ] Más tipos de agentes con diferentes estrategias
- [ ] Obstáculos en el tablero
- [ ] Modo multijugador
- [ ] Sistema de puntuación
- [ ] Replay de partidas
- [ ] Configuración personalizable

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 🎬 Inspiración

Este juego está inspirado en la icónica película "The Matrix" (1999), donde Neo debe escapar de los Agentes en un mundo digital. El juego captura la esencia de persecución y escape de la película en un formato de juego de estrategia.

---

<div align="center">

**¿Te gustó el proyecto? ¡Dale una ⭐ al repositorio!**

*Hecho con ❤️ y Go*

</div>