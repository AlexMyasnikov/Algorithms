package lists

import (
	"container/list"
	"errors"
	"fmt"
)

type MARSH struct {
	begin  string
	end    string
	number int32
}

var ErrNotFound = errors.New("not found")

func Begin() {
	input := ""
	marsh := MARSH{}
	l := list.New()
	for {
		fmt.Scanf("%s", &input)
		if input == "add" {
			fmt.Scanf("%s %s %d", &marsh.begin, &marsh.end, &marsh.number)
			l.PushBack(marsh)
		} else if input == "get" {
			fmt.Scanf("%d", &marsh.number)
			elem, err := find(l, marsh.number)
			if errors.Is(err, ErrNotFound) {
				fmt.Println(err)
			} else {
				fmt.Println(elem)
			}
		} else if input == "all" {
			sort(l)
			fmt.Println(l.Len())
			for e := l.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value.(MARSH))
			}
		}
	}
}

func sort(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		for j := e.Next(); j != nil; j = j.Next() {
			if e.Value.(MARSH).number > j.Value.(MARSH).number {
				l.MoveAfter(e, j)
				e = l.Front()
			}
		}
	}
}

func find(l *list.List, number int32) (MARSH, error) {
	for j := l.Front(); j != nil; j = j.Next() {
		if j.Value.(MARSH).number == number {
			return j.Value.(MARSH), nil
		}
	}
	return MARSH{}, ErrNotFound
}
