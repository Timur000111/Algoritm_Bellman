package main

import (
	"errors"
	"fmt"
	"math"
)

// Edge представляет ребро графа с исходной вершиной, конечной вершиной и весом.
type Edge struct {
	From, To, Weight int
}

// Graph представляет граф, состоящий из множества вершин и рёбер.
type Graph struct {
	Vertices int
	Edges    []Edge
}

// NewGraph создает новый граф с заданным числом вершин.
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		Edges:    []Edge{},
	}
}

// BellmanFord реализует алгоритм Беллмана-Форда для поиска кратчайших путей в графе.
func BellmanFord(graph *Graph, startVertex int) ([]int, error) {
	distance := make([]int, graph.Vertices)
	for i := range distance {
		distance[i] = int(^uint(0) >> 1) // Максимальное значение для расстояния (или ∞)
	}
	distance[startVertex] = 0

	// Выполняем |V| - 1 итераций
	for i := 0; i < graph.Vertices-1; i++ {
		for _, edge := range graph.Edges {
			u, v, w := edge.From, edge.To, edge.Weight
			if distance[u] != int(^uint(0)>>1) && distance[u]+w < distance[v] {
				distance[v] = distance[u] + w
			}
		}
	}

	// Проверка на наличие отрицательных циклов
	for _, edge := range graph.Edges {
		u, v, w := edge.From, edge.To, edge.Weight
		if distance[u] != int(^uint(0)>>1) && distance[u]+w < distance[v] {
			return nil, errors.New("граф содержит отрицательный цикл")
		}
	}

	return distance, nil
}

func main() {
	// Создание нового графа
	graph := NewGraph(5)

	// Добавление рёбер
	graph.Edges = append(graph.Edges, Edge{From: 0, To: 1, Weight: 6})
	graph.Edges = append(graph.Edges, Edge{From: 0, To: 2, Weight: 7})
	graph.Edges = append(graph.Edges, Edge{From: 1, To: 2, Weight: 8})
	graph.Edges = append(graph.Edges, Edge{From: 1, To: 3, Weight: 5})
	graph.Edges = append(graph.Edges, Edge{From: 1, To: 4, Weight: -4})
	graph.Edges = append(graph.Edges, Edge{From: 2, To: 3, Weight: -3})
	graph.Edges = append(graph.Edges, Edge{From: 2, To: 4, Weight: 9})
	graph.Edges = append(graph.Edges, Edge{From: 3, To: 1, Weight: -2})
	graph.Edges = append(graph.Edges, Edge{From: 4, To: 3, Weight: 7})

	// Выполнение алгоритма Беллмана-Форда
	distance, err := BellmanFord(graph, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод кратчайших путей
	fmt.Println("Кратчайшие пути от вершины 0:")
	for i, d := range distance {
		fmt.Printf("до вершины %d: %d\n", i, d)
	}
}
