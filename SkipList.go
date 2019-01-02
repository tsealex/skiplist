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
	root Element
	len  int

	levelLen []int
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
		node := &sl.root
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

func (sl *SkipList) insert(i int) {
	visited := sl.search(i)
	if visited == nil {
		sl.root.value = i
		sl.levelLen = append(sl.levelLen, 1)
	} else {
		var down *Element
		for j := len(visited) - 1; j >= 0; j-- {
			if visited[j].value > i {
				down = visited[j].insertPrev(i, down)
			} else {
				down = visited[j].insertNext(i, down)
			}
			sl.levelLen[j] += 1
			// If the next level has space, than continue this loop with 1/2 probability.
			if j > 0 && sl.levelLen[j-1] < sl.levelLen[j] / 2 {
				if flip() {
					break
				}
			} else if j == 0 && sl.levelLen[j] > 1 {
				sl.levelLen = append([]int{1}, sl.levelLen...)
				sl.root = Element{down: down, value: i}
				break
			}
		}
	}
	sl.len += 1
}
