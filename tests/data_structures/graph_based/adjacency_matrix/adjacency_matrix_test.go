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