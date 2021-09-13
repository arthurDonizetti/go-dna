package domain

type Tree struct {
	Root *Gene
	Size int
}

func NewTree(gene *Gene, size int) *Tree {
	return &Tree{
		Root: gene,
		Size: size,
	}
}