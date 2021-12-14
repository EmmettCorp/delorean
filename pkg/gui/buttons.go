package gui

const (
	createButton  = "create"
	deleteButton  = "delete"
	restoreButton = "restore"
)

type points struct {
	x0 int
	x1 int
}

type buttons struct {
	createButton  points
	deleteButton  points
	restoreButton points
}

func (b *buttons) calcutalePoints() {
	indent := 1
	// order matters
	b.createButton.x0 = 0
	b.createButton.x1 = b.createButton.x0 + len(createButton) + 1

	b.restoreButton.x0 = b.createButton.x1 + indent
	b.restoreButton.x1 = b.restoreButton.x0 + len(restoreButton) + 1

	b.deleteButton.x0 = b.restoreButton.x1 + indent
	b.deleteButton.x1 = b.deleteButton.x0 + len(deleteButton) + 1
}
