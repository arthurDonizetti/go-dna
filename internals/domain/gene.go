package domain

import "fmt"

type Gene struct {
	Value                     string
	Position                  *Position
	NeighbourInRow            *Gene
	NeighbourInCol            *Gene
	NeighbourInDiagonalUpDown *Gene
	NeighbourInDiagonalDownUp *Gene

}

type Position struct {
	row int
	col int
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
	g.FindInRow(dna)
	g.FindInCol(dna)
	g.FindInDiagonalUpDown(dna)
	g.FindInDiagonalDownUp(dna)
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
	nextPosition := fmt.Sprintf("%d%d", g.Position.row, g.Position.col + 1)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInRow = nextGene
		nextGene.FindInRow(dna)
	}
}

func (g *Gene) FindInCol(dna *Dna) {
	nextPosition := fmt.Sprintf("%d%d", g.Position.row + 1, g.Position.col)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInCol = nextGene
		nextGene.FindInCol(dna)
	}
}

func (g *Gene) FindInDiagonalUpDown(dna *Dna) {
	nextPosition := fmt.Sprintf("%d%d", g.Position.row + 1, g.Position.col + 1)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInDiagonalUpDown = nextGene
		nextGene.FindInDiagonalUpDown(dna)
	}
}

func (g *Gene) FindInDiagonalDownUp(dna *Dna) {
	nextPosition := fmt.Sprintf("%d%d", g.Position.row - 1, g.Position.col + 1)
	if nextGene, ok := dna.GeneMap[nextPosition]; ok && dna.GeneMap[nextPosition].Value == g.Value {
		g.NeighbourInDiagonalDownUp = nextGene
		nextGene.FindInDiagonalDownUp(dna)
	}
}