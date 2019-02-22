package ds

type BinaryTree struct {
	Value  interface{}
	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func (b *BinaryTree) InsertLeft(value interface{}) {
	b.Left = &BinaryTree{
		Value:  value,
		Parent: b,
	}
}

func (b *BinaryTree) InsertRight(value interface{}) {
	b.Right = &BinaryTree{
		Value:  value,
		Parent: b,
	}
}
