package lists

import "fmt"

// 1 - структура Node - информация, ссылка на следующий эл.
// 2 - LinkedList - ссылка на Head, End

type Node struct {
	Number float64
	Prev   *Node
	Next   *Node
}

// Для основного задания по одно- и двунаправленным
type ChNode struct {
	Symbol string
	Prev   *ChNode
	Next   *ChNode
}

// Для задания по варианту
type Doubly struct {
	Head *Node
	Foot *Node
}

// Для основного задания
type ChDoubly struct {
	Head *ChNode
	Foot *ChNode
}

func (d *Doubly) PushBack(node Node) {
	if d.Head == nil {
		d.Head = &node
		d.Foot = &node
		return
	}
	el := d.Foot
	el.Next = &node
	node.Prev = el
	d.Foot = &node
}

func (d *Doubly) PushForward(node Node) {
	if d.Head == nil {
		d.Head = &node
		d.Foot = &node
		return
	}
	el := d.Head
	el.Prev = &node
	node.Next = el
	d.Head = &node
}

func (d *Doubly) DeleteByVal(number float64) {
	for el := *d.Head; ; {
		if el.Number == number {
			prev, next := el.Prev, el.Next
			if prev != nil {
				prev.Next = next
			} else {
				d.Head = next
			}
			if next != nil {
				next.Prev = prev
			} else {
				d.Foot = prev
			}
		}
		if el.Next == nil {
			return
		}
		el = *el.Next
	}
}

func (d *Doubly) Clear() {
	*d = Doubly{}
}

func (d *Doubly) FindIndex(number float64) int {
	idx := 0
	for el := *d.Head; ; {
		if el.Number == number {
			return idx
		}
		idx++
		if el.Next == nil {
			return -1
		}
		el = *el.Next
	}
}

func (d *Doubly) Print() {
	fmt.Println("------------------")
	for v := d.Head; v != nil; v = v.Next {
		if v.Prev != nil {
			fmt.Printf("Prev: %.2f\n", v.Prev.Number)
		}
		if v.Next != nil {
			fmt.Printf("Next: %.2f\n", v.Next.Number)
		}
		fmt.Printf("Value: %.2f\n", v.Number)
		fmt.Println("---")
	}
	fmt.Println("------------------")
}

func (d *Doubly) Mean() float64 {
	var n int32
	for v := d.Head; v != nil; v = v.Next {
		n++
	}
	var target *Node
	for v, c := d.Head, 0; v != nil; v, c = v.Next, c+1 {
		if c == int(n)/2 {
			break
		}
		target = v
	}
	mins := []float64{}
	for prev, next := target, target.Next; (prev != nil) && (next != nil); prev, next = prev.Prev, next.Next {
		if prev != nil && next != nil {
			mins = append(mins, min(prev.Number, next.Number))
		}
	}
	var sum float64
	for _, v := range mins {
		sum += v
	}
	return sum / float64(len(mins))
}

func (d *ChDoubly) PushBack(node ChNode) {
	if d.Head == nil {
		d.Head = &node
		d.Foot = &node
		return
	}
	el := d.Foot
	el.Next = &node
	node.Prev = el
	d.Foot = &node
}

func (d *ChDoubly) DeleteLastSymbol() {
	if d.Head == nil {
		d.Head = nil
		return
	}
	if d.Head.Next == nil {
		d.Head = nil
		return
	}
	el := d.Foot
	d.Foot = el.Prev
	d.Foot.Next = nil
}

func (d *ChDoubly) Print() {
	fmt.Println("------------------")
	for v := d.Head; v != nil; v = v.Next {
		if v.Prev != nil {
			fmt.Printf("Prev: %s\n", v.Prev.Symbol)
		}
		if v.Next != nil {
			fmt.Printf("Next: %s\n", v.Next.Symbol)
		}
		fmt.Printf("Value: %s\n", v.Symbol)
		fmt.Println("---")
	}
	fmt.Println("------------------")
}

func min(first, second float64) float64 {
	if first < second {
		return first
	}
	return second
}
