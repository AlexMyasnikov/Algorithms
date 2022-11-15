package lists

import (
	"fmt"
)

type Person struct {
	Number   string
	Lastname string
}

type LayerNode struct {
	Number   string
	Lastname string
	Next     *LayerNode
	ByLetter *LayerNode
}

type Layerd struct {
	Head *LayerNode
}

// Ситуации:
// Head - nil x
// Такой буквы еще нет
// Добавляемая фамилия должна быть самой первой
// Добавляемая фамилия должна быть в середине
// Добавляемая фамилия должна быть в конце

// Функция проверяет, должна ли строка first стоять перед second
func isFirstAhead(first, second string) bool {
	// Проход по i-му символу каждой строки
	for i := 0; i < len(first) || i < len(second); i++ {
		// Можно рассматривать эту условную конструкцию так.
		// Если i-й символ (буква) first в алфавите стоит раньше чем, i-я из second, то мы значит first должна стоять раньше second
		if []rune(first)[i] < []rune(second)[i] {
			return true
		} else if []rune(first)[i] > []rune(second)[i] {
			return false
		}
	}
	return true
}

// Функция достает последний элемент из буквенной группы.
// Входной аргумент - первый элемент этой группы.
// Пример:  Алексеев, Андреев, Федотов, Федеров, Фучик, Фрунзе, Харитонов.
// Чтобы достать последний элемент из буквы Ф надо вызвать lastInLetter(Федотов)
func lastInLetter(v *LayerNode) *LayerNode {
	if v.Lastname == "" {
		return v
	}
	letter := []rune(v.Lastname)[0]
	// Цикл работает пока не дойдет до конца или не встретится другая буква
	for {
		if v.Next == nil || letter != []rune(v.Next.Lastname)[0] {
			return v
		}
		v = v.Next
	}
}

func (l *Layerd) Add(addedNode *LayerNode) {
	if l.Head == nil {
		l.Head = addedNode
		return
	}
	prev := &LayerNode{}
	for v := l.Head; v != nil; {
		if []rune(v.Lastname)[0] == []rune(addedNode.Lastname)[0] {
			// Добавляемая фамилия должна быть самой первой
			if isFirstAhead(addedNode.Lastname, v.Lastname) {
				last := lastInLetter(prev)
				addedNode.Next = v
				last.Next = addedNode
				addedNode.ByLetter = v.ByLetter
				v.ByLetter = nil
				prev.ByLetter = addedNode
				if l.Head == v {
					l.Head = addedNode
				}
				return
			}
			// Добавляемая фамилия должна быть в середине
			prevB := &LayerNode{}
			for b := v.Next; b != nil; b = b.Next {
				if isFirstAhead(addedNode.Lastname, b.Lastname) {
					addedNode.Next = prevB.Next
					prevB.Next = addedNode
					return
				}
				prevB, b = b, b.Next
			}
			// Добавляемая фамилия должна быть в конце
			prevB.Next = addedNode
			return
		}
		prev, v = v, v.ByLetter
	}
	// Такой буквы еще нет
	prev = &LayerNode{}
	// Циклом проверяю, не должна ли добавляемая запись (буква которой еще не встречалась, см стр 102) быть в начале или в где-то в центре
	for v := l.Head; v != nil; {
		if isFirstAhead(addedNode.Lastname, v.Lastname) {
			last := lastInLetter(prev)
			addedNode.Next = last.Next
			addedNode.ByLetter = prev.ByLetter
			last.Next = addedNode
			prev.ByLetter = addedNode
			return
		}
		v, prev = v.ByLetter, v
	}
	// Добавляемая запись должна быть в конце (из уже имеющихся в списке букв буква добавляемой записи самая "дальняя")
	last := lastInLetter(prev)
	addedNode.Next = last.Next
	addedNode.ByLetter = prev.ByLetter
	last.Next = addedNode
	prev.ByLetter = addedNode
}

func (l *Layerd) Delete(lastname, number string) {
	prev, prevL := &LayerNode{}, &LayerNode{}
	for v := l.Head; v != nil; {
		if []rune(v.Lastname)[0] == []rune(lastname)[0] {
			for l := v; l != nil; {
				if l.Number == number {
					// скраю слева
					if prevL.Lastname == "" {
						prevLast := lastInLetter(prev)
						l.Next.ByLetter = l.ByLetter
						prev.ByLetter = l.Next
						prevLast.Next = l.Next
					} else if isFirstAhead(l.Lastname, l.Next.Lastname) { // скраю справа
						prevL.Next = l.Next
					} else { // где-то в центре
						prevL.Next = l.Next
					}
				}
				l, prevL = l.Next, l
			}
		}
		v, prev = v.ByLetter, v
	}
}

func (l *Layerd) Clear() {
	l.Head = nil
}

func (l *Layerd) GetByIndex(index int) *LayerNode {
	i := 0
	for v := l.Head; v != nil; v = v.Next {
		if index == i {
			return v
		}
		i++
	}
	return nil
}

func (l *Layerd) Show() {
	i := 1
	for v := l.Head; v != nil; v = v.Next {
		fmt.Println("-----------------------------------")
		fmt.Printf("Person #%d\n", i)
		fmt.Printf("Lastname: %s Number: %s\n", v.Lastname, v.Number)
		if v.ByLetter != nil {
			fmt.Printf("ByLetter: %s\n", v.ByLetter.Lastname)
		}
		if v.Next != nil {
			fmt.Printf("Next: %s", v.Next.Lastname)
		}
		fmt.Printf("\n")
		i++
	}
}
