package main

import (
	"errors"
	"fmt"
	"math"
)

// Graph представляет собой ориентированный граф
type Graph struct {
	Vertices int
	Edges    []Edge
}

// Edge представляет собой ребро графа с начальной вершиной, конечной вершиной и весом
type Edge struct {
	From, To, Weight int
}

// NewGraph создает новый граф с заданным числом вершин
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		Edges:    []Edge{},
	}
}

// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(from, to, weight int) {
	g.Edges = append(g.Edges, Edge{From: from, To: to, Weight: weight})
}

// LevitasAlgorithm применяет алгоритм Левита для поиска кратчайших путей
func LevitasAlgorithm(graph *Graph, startVertex int) ([]int, error) {
	// Инициализация расстояний
	distance := make([]int, graph.Vertices)
	for i := range distance {
		distance[i] = int(math.MaxInt32) // Максимальное значение для расстояния (или ∞)
	}
	distance[startVertex] = 0

	// Итерации для обновления расстояний
	for i := 0; i < graph.Vertices-1; i++ {
		updated := false
		for _, edge := range graph.Edges {
			u, v, w := edge.From, edge.To, edge.Weight
			if distance[u] != int(math.MaxInt32) && distance[u]+w < distance[v] {
				distance[v] = distance[u] + w
				updated = true
			}
		}
		// Если никаких обновлений не было, прерываем итерации
		if !updated {
			break
		}
	}

	// Проверка на отрицательные циклы
	for _, edge := range graph.Edges {
		u, v, w := edge.From, edge.To, edge.Weight
		if distance[u] != int(math.MaxInt32) && distance[u]+w < distance[v] {
			return nil, errors.New("граф содержит отрицательный цикл")
		}
	}

	return distance, nil
}

func main() {
	// Пример использования алгоритма Левита
	graph := NewGraph(5)
	graph.AddEdge(0, 1, 6)
	graph.AddEdge(0, 2, 7)
	graph.AddEdge(1, 2, 8)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(1, 4, -4)
	graph.AddEdge(2, 3, -3)
	graph.AddEdge(2, 4, 9)
	graph.AddEdge(3, 2, -2)
	graph.AddEdge(4, 3, 7)

	startVertex := 0
	distance, err := LevitasAlgorithm(graph, startVertex)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Расстояния от вершины %d:\n", startVertex)
		for i, d := range distance {
			fmt.Printf("До вершины %d: %d\n", i, d)
		}
	}
}
