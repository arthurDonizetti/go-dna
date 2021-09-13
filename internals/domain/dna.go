package domain

import (
	"fmt"
	"strings"
)

type Dna struct {
	Sequence []string
	SubTrees []*Tree
	GeneMap  map[string]*Gene
}

func NewDna(sequence []string) *Dna {
	dna := &Dna{
		Sequence: sequence,
		GeneMap: map[string]*Gene{},
	}

	for i, rowGene := range dna.Sequence {
		for j, colGene := range strings.Split(rowGene, "") {
			gene := NewGene(colGene, i, j)
			dna.GeneMap[fmt.Sprintf("%d%d", gene.Position.row, gene.Position.col)] = gene
		}
	}

	return dna
}

func (d *Dna) IsSimian() bool {
	d.buildSubTrees()

	count := 0

	for _, subtree := range d.SubTrees {
		if subtree.Size > 3 {
			count++
		}
	}

	return count > 1
}

func (d *Dna) buildSubTrees() {
 	for _, gene := range d.GeneMap {
		gene.SearchNeighbours(d)
		size := gene.getGreatestNeighbourhood()
		d.SubTrees = append(d.SubTrees, NewTree(gene, size))
	}
}