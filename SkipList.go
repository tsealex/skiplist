package skiplist

func flip() bool {
	return true
}

type Element struct {
	next, prev, down *Element
	value            int
}

func (e *Element) insertNext(i int, down *Element) *Element {
	res := &Element{value: i, next: e.next, prev: e, down: down}
	e.next.prev = res
	e.next = res
	return res
}

func (e *Element) insertPrev(i int, down *Element) *Element {
	res := &Element{value: i, next: e, prev: e.prev, down: down}
	e.prev.next = res
	e.prev = res
	return res
}

type SkipList struct {
	root *Element
	len  int
}

func (sl *SkipList) Contains(i int) bool {
	visited := sl.search(i)
	return visited != nil && visited[len(visited)-1].value == i
}

func (sl *SkipList) search(i int) []*Element {
	var visited []*Element
	if sl.len == 0 {
		return nil
	} else {
		node := sl.root
		for node != nil {
			for node.value < i && node.next != nil {
				node = node.next
			}
			for node.value > i && node.prev != nil {
				node = node.prev
			}
			visited = append(visited, node)
			node = node.down
		}
	}
	return visited
}

func (sl *SkipList) Insert(i int) {
	sl.insert(i)
}

func (sl *SkipList) insert(i int) {
	visited := sl.search(i)
	if visited == nil {
		sl.root = &Element{value: i}
	} else {
		var down *Element
		for j := len(visited) - 1; j >= 0; j-- {
			if visited[j].value > i {
				down = visited[j].insertPrev(i, down)
			} else {
				down = visited[j].insertNext(i, down)
			}
			// Continue this loop with 1/2 probability.
			if j > 0 && flip() {
				break
			} else if j == 0 {
				sl.root = &Element{down: down, value: i}
			}
		}
	}
	sl.len += 1
}

func (sl *SkipList) Delete(i int) {
	sl.delete(i)
}

func (sl *SkipList) delete(i int) {
	visited := sl.search(i)
	if visited != nil {
		for j := len(visited) - 1; j >= 0; j-- {
			if visited[j].value != i {
				break
			}
			if visited[j].prev != nil {
				visited[j].prev.next = visited[j].next
			}
			if visited[j].next != nil {
				visited[j].next.prev = visited[j].prev
			}
		}
		// If the root node also has this value, replace it with the neighbor of its direct child.
		if visited[0] == sl.root {
			if down := visited[0].down; down != nil {
				if down.prev != nil {
					sl.root = down.prev
				} else {
					sl.root = down.next
				}
			}
		}
		sl.len -= 1
	}
}
