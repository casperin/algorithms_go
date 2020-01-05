package tree

type color int

const (
	red color = iota
	black
)

type RBTree struct {
	root    *redBlack
	compare func(int, int) int
}

type redBlack struct {
	key    int
	color  color
	left   *redBlack
	right  *redBlack
	parent *redBlack
}

func NewRedBlack() RBTree {
	return RBTree{
		root:    nil,
		compare: func(i, j int) int { return j - i },
	}
}

// Methods on the wrapper struct
func (t *RBTree) Insert(key int) {
	n := &redBlack{key: key}

	t.insertRecurse(t.root, n)

	// Repair the tree in case any of the red-black properties have been
	// violated.
	insertRepairTree(n)

	t.root = n
	for getParent(t.root) != nil {
		t.root = getParent(t.root)
	}
}

func (t *RBTree) Search(key int) bool {
	n := t.root
walk:
	if n == nil {
		return false
	}
	cmp := t.compare(key, n.key)
	if cmp == 0 {
		return true
	}
	if cmp > 0 {
		n = n.left
	} else {
		n = n.right
	}
	goto walk
}

func (t *RBTree) PreOrder() []int {
	return preOrder(t.root, []int{})
}

func (t *RBTree) Ordered() []int {
	return ordered(t.root, []int{})
}

func preOrder(n *redBlack, keys []int) []int {
	if n == nil {
		return keys
	}
	keys = append(keys, n.key)
	keys = preOrder(n.left, keys)
	keys = preOrder(n.right, keys)
	return keys
}

func ordered(n *redBlack, keys []int) []int {
	if n == nil {
		return keys
	}
	keys = ordered(n.left, keys)
	keys = append(keys, n.key)
	keys = ordered(n.right, keys)
	return keys
}

func getParent(n *redBlack) *redBlack {
	if n == nil {
		return nil
	}
	return n.parent
}

func getGrandParent(n *redBlack) *redBlack {
	return getParent(getParent(n))
}

func getSibling(n *redBlack) *redBlack {
	p := getParent(n)
	if p == nil {
		return nil
	}
	if n == p.left {
		return p.right
	}
	return p.left
}

func getUncle(n *redBlack) *redBlack {
	return getSibling(getParent(n))
}

func rotateLeft(n *redBlack) {
	nnew := n.right
	p := getParent(n)
	// Rotate
	n.right = nnew.left
	nnew.left = n
	n.parent = nnew
	// Move child/parents
	if n.right != nil {
		n.right.parent = n
	}
	// n could be root
	if p != nil {
		if n == p.left {
			p.left = nnew
		} else if n == p.right {
			p.right = nnew
		}
	}
	nnew.parent = p
}

func rotateRight(n *redBlack) {
	nnew := n.left
	p := getParent(n)
	// Rotate
	n.left = nnew.right
	nnew.right = n
	n.parent = nnew
	// Move child/parents
	if n.left != nil {
		n.left.parent = n
	}
	// n could be root
	if p != nil {
		if n == p.left {
			p.left = nnew
		} else if n == p.right {
			p.right = nnew
		}
	}
	nnew.parent = p
}

func (t *RBTree) insertRecurse(root *redBlack, n *redBlack) {
	if root != nil {
		if t.compare(n.key, root.key) > 0 {
			if root.left != nil {
				t.insertRecurse(root.left, n)
				return
			} else {
				root.left = n
			}
		} else {
			if root.right != nil {
				t.insertRecurse(root.right, n)
				return
			} else {
				root.right = n
			}
		}
	}
	n.parent = root
	n.color = red
}

func insertRepairTree(n *redBlack) {
	switch {
	case getParent(n) == nil:
		n.color = black

	case getParent(n).color == black:
		// noop, tree still valid

	case getUncle(n) != nil && getUncle(n).color == red:
		// See https://en.wikipedia.org/wiki/Red%E2%80%93black_tree#/media/File:Red-black_tree_insert_case_3.svg
		getParent(n).color = black
		getUncle(n).color = black
		getGrandParent(n).color = red
		insertRepairTree(getGrandParent(n))

	default:
		// See https://en.wikipedia.org/wiki/Red%E2%80%93black_tree#/media/File:Red-black_tree_insert_case_4.svg
		p := getParent(n)
		g := getGrandParent(n)

		if n == p.right && p == g.left {
			rotateLeft(p)
			n = n.left
		} else if n == p.left && p == g.right {
			rotateRight(p)
			n = n.right
		}
		insertCase4Step2(n)
	}
}

func insertCase4Step2(n *redBlack) {
	// See https://en.wikipedia.org/wiki/Red%E2%80%93black_tree#/media/File:Red-black_tree_insert_case_5.svg
	p := getParent(n)
	g := getGrandParent(n)

	if n == p.left {
		rotateRight(g)
	} else {
		rotateLeft(g)
	}

	p.color = black
	g.color = red
}
