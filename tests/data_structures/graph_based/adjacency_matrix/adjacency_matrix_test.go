package data_structures_tests

import (
	"testing"

	ds "dsa/src/data_structures/graph_based/adjacency_matrix"
)

func TestNewAdjacencyMatrix(t *testing.T) {
	am := ds.NewAdjacencyMatrix()
	if len(am.GetVertices()) != 0 {
		t.Errorf("Expected 0 vertices, got %d", len(am.GetVertices()))
	}
}

func TestAddVertex(t *testing.T) {
	am := ds.NewAdjacencyMatrix()
	am.AddVertex("A")
	vertices := am.GetVertices()
	if len(vertices) != 1 || vertices[0] != "A" {
		t.Errorf("Expected vertex A, got %v", vertices)
	}
}

func TestAddEdge(t *testing.T) {
	am := ds.NewAdjacencyMatrix()
	am.AddVertex("A")
	am.AddVertex("B")
	am.AddEdge("A", "B")

	edges := am.GetEdges("A")
	if len(edges) != 1 || edges[0] != "B" {
		t.Errorf("Expected edge from A to B, got %v", edges)
	}
}

func TestGetEdgesNonExistentVertex(t *testing.T) {
	am := ds.NewAdjacencyMatrix()
	edges := am.GetEdges("NonExistent")
	if edges != nil {
		t.Errorf("Expected nil, got %v", edges)
	}
}

func TestRemoveVertex(t *testing.T) {
	am := ds.NewAdjacencyMatrix()
	am.AddVertex("A")
	am.AddVertex("B")
	am.AddEdge("A", "B")
	am.RemoveVertex("A")

	if len(am.GetVertices()) != 1 {
		t.Errorf("Expected 1 vertex, got %d", len(am.GetVertices()))
	}

	edges := am.GetEdges("A")
	if edges != nil {
		t.Errorf("Expected nil, got %v", edges)
	}
}

func TestRemoveEdge(t *testing.T) {
	am := ds.NewAdjacencyMatrix()
	am.AddVertex("A")
	am.AddVertex("B")
	am.AddEdge("A", "B")
	am.RemoveEdge("A", "B")

	edges := am.GetEdges("A")
	if len(edges) != 0 {
		t.Errorf("Expected 0 edges, got %d", len(edges))
	}
}
