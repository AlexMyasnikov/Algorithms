package lists

import "fmt"

/*
	Добавление элемента в конец списка.
	Добавление элемента в начало списка.
	Добавление элемента в определенную позицию.
	Удаление элемента по его значению.
	Удаление элемента по его номеру в односвязном списке.
	Очистка списка.
	Поиска номера элемента в списке.
	Просмотр списка.
*/

// Дано предложение, оканчивающееся точкой. Из букв предложения построить линейный однонаправленный список.

type Singly struct {
	Head *ChNode
}

func (s *Singly) PushBack(node ChNode) {
	if s.Head == nil {
		s.Head = &node
		return
	}
	var el *ChNode
	for el = s.Head; el.Next != nil; el = el.Next {
	}
	el.Next = &node
}

func (s *Singly) PushForward(node ChNode) {
	if s.Head == nil {
		s.Head = &node
		return
	}
	el := s.Head
	s.Head = &node
	node.Next = el
}

func (s *Singly) AddByIndex(idx int32, node ChNode) error {
	if idx == 0 {
		s.PushForward(node)
		return nil
	}
	var cur *ChNode
	var i int32
	prev := s.Head
	for cur = s.Head; cur != nil; cur = cur.Next {
		if i == idx {
			el := cur
			cur = &node
			cur.Next = el
			prev.Next = cur
			return nil
		}
		prev = cur
		i++
	}
	return ErrNotFound
}

func (s *Singly) DeleteByVal(node ChNode) {
	if s.Head == nil {
		return
	}
	if node.Symbol == s.Head.Symbol {
		if s.Head.Next != nil {
			s.Head = s.Head.Next
			return
		}
		s.Head = nil
		return
	}
	var prev *ChNode
	for cur := s.Head; cur != nil; cur = cur.Next {
		if cur.Symbol == node.Symbol {
			prev.Next = cur.Next
			return
		}
		prev = cur
	}
}

func (s *Singly) DeleteByIdx(idx int32) {
	if s.Head == nil {
		return
	}
	if idx == 0 {
		if s.Head.Next != nil {
			s.Head = s.Head.Next
			return
		}
		s.Head = nil
		return
	}
	var prev *ChNode
	var i int32
	for cur := s.Head; cur != nil; cur = cur.Next {
		if i == idx {
			prev.Next = cur.Next
			return
		}
		prev = cur
		i++
	}
}

func (s *Singly) Clear() {
	s.Head = nil
}

func (s *Singly) GetByVal(node ChNode) int32 {
	if s.Head == nil {
		return -1
	}
	var i int32
	for cur := s.Head; cur != nil; cur = cur.Next {
		if cur.Symbol == node.Symbol {
			return i
		}
		i++
	}
	return -1
}

func (s *Singly) Print() {
	for el := s.Head; el != nil; el = el.Next {
		if el.Next == nil {
			fmt.Printf("Current: %s. Next: %s\n", el.Symbol, "")
			continue
		}
		fmt.Printf("Current: %s. Next: %s\n", el.Symbol, el.Next.Symbol)
	}
}
