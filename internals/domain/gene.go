package domain

import "fmt"

type Gene struct {
	Value                     string
	Position                  *Position
	VisitedFlow               VisitedFlow
	NeighbourInRow            *Gene
	NeighbourInCol            *Gene
	NeighbourInDiagonalUpDown *Gene
	NeighbourInDiagonalDownUp *Gene
}

type Position struct {
	row int
	col int
}

type VisitedFlow struct {
	Row            bool
	Col            bool
	DiagonalUpDown bool
	DiagonalDownUp bool
}

type Search interface {
	FindInRow(dna *Dna)
	FindInCol(dna *Dna)
	FindInDiagonalUpDown(dna *Dna)
	FindInDiagonalDownUp(dna *Dna)
}

func NewGene(value string, row int, col int) *Gene {
	return &Gene{
		Value: value,
		Position: &Position{
			row,
			col,
		},
	}
}

func (g *Gene) SearchNeighbours(dna *Dna) {
	if !g.VisitedFlow.Row {
		g.FindInRow(dna)
	}

	if !g.VisitedFlow.Col {
		g.FindInCol(dna)
	}

	if !g.VisitedFlow.DiagonalUpDown {
		g.FindInDiagonalUpDown(dna)
	}

	if !g.VisitedFlow.DiagonalDownUp {
		g.FindInDiagonalDownUp(dna)
	}
}

func (g *Gene) getGreatestNeighbourhood() int {
	greatest := 0
	sizes := []int{
		g.countNeighboursInRow(),
		g.countNeighboursInCol(),
		g.countNeighboursInDiagonalDownUp(),
		g.countNeighboursInDiagonalUpDown(),
	}

	for _, s := range sizes {
		if s > greatest {
			greatest = s
		}
	}

	return greatest
}

func (g *Gene) countNeighboursInCol() int {
	count := 0
	if g.NeighbourInCol == nil {
		count++
		return count
	}
	count += g.NeighbourInCol.countNeighboursInCol() + 1
	return count
}

func (g *Gene) countNeighboursInRow() int {
	count := 0
	if g.NeighbourInRow == nil {
		count++
		return count
	}
	count += g.NeighbourInRow.countNeighboursInRow() + 1
	return count
}

func (g *Gene) countNeighboursInDiagonalDownUp() int {
	count := 0
	if g.NeighbourInDiagonalDownUp == nil {
		count++
		return count
	}
	count += g.NeighbourInDiagonalDownUp.countNeighboursInDiagonalDownUp() + 1
	return count
}

func (g *Gene) countNeighboursInDiagonalUpDown() int {
	count := 0
	if g.NeighbourInDiagonalUpDown == nil {
		count++
		return count
	}
	count += g.NeighbourInDiagonalUpDown.countNeighboursInDiagonalUpDown() + 1
	return count
}

func (g *Gene) FindInRow(dna *Dna) {
	g.VisitedFlow.Row = true
	nextPosition := fmt.Sprintf("%d;%d", g.Position.row, g.Position.col+1)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInRow = nextGene
		nextGene.FindInRow(dna)
	}
}

func (g *Gene) FindInCol(dna *Dna) {
	g.VisitedFlow.Col = true
	nextPosition := fmt.Sprintf("%d;%d", g.Position.row+1, g.Position.col)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInCol = nextGene
		nextGene.FindInCol(dna)
	}
}

func (g *Gene) FindInDiagonalUpDown(dna *Dna) {
	g.VisitedFlow.DiagonalUpDown = true
	nextPosition := fmt.Sprintf("%d;%d", g.Position.row+1, g.Position.col+1)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInDiagonalUpDown = nextGene
		nextGene.FindInDiagonalUpDown(dna)
	}
}

func (g *Gene) FindInDiagonalDownUp(dna *Dna) {
	g.VisitedFlow.DiagonalDownUp = true
	nextPosition := fmt.Sprintf("%d;%d", g.Position.row-1, g.Position.col+1)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInDiagonalDownUp = nextGene
		nextGene.FindInDiagonalDownUp(dna)
	}
}
