package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AlexMyasnikov/algorithms/lists"
	"github.com/AlexMyasnikov/algorithms/query"
	stackdeck "github.com/AlexMyasnikov/algorithms/stack-deck"
)

//  ДЗ: Упражнение с поиском в ширину! 4 задание Упражнение 1, 4

func main() {
	// doubly()
	// singly()
	// s()
	// d()
	l()
}

func l() {
	l8 := &lists.LayerNode{
		Lastname: "Кац",
		Number:   "8",
	}
	l7 := &lists.LayerNode{
		Lastname: "Краснов",
		Number:   "7",
	}
	l6 := &lists.LayerNode{
		Lastname: "Федотов",
		Number:   "6",
	}
	l5 := &lists.LayerNode{
		Lastname: "Фёдоров",
		Number:   "5",
	}
	l3 := &lists.LayerNode{
		Lastname: "Мясников",
		Number:   "3",
	}
	l4 := &lists.LayerNode{
		Lastname: "Иванов",
		Number:   "4",
	}
	l2 := &lists.LayerNode{
		Lastname: "Андреев",
		Number:   "2",
	}
	l1 := &lists.LayerNode{
		Lastname: "Алексеев",
		Number:   "1",
	}
	l := lists.Layerd{}
	l.Add(l2)
	l.Add(l3)
	l.Add(l1)
	l.Add(l4)
	l.Add(l5)
	l.Add(l6)
	l.Add(l7)
	l.Add(l8)
	l.Show()
	l.Delete(l8.Lastname, l8.Number)
	l.Show()
}

func d() {
	deck := stackdeck.Deck{}
	deck.AddFront(1)
	deck.AddFront(2)
	deck.AddBack(-1)
	deck.Show()

	var input string
	var el float64
	for {
		fmt.Scanf("%s", &input)
		fmt.Printf("Введенная команда: %s\n", input)
		switch input {
		case "add-front":
			_, err := fmt.Scanln(&el)
			if err != nil {
				log.Fatal(err)
			}
			deck.AddFront(el)
		case "add-back":
			_, err := fmt.Scanln(&el)
			if err != nil {
				log.Fatal(err)
			}
			deck.AddBack(el)
		case "pop-front":
			v, err := deck.PopFront()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)
		case "pop-back":
			v, err := deck.PopBack()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)
		case "peek-front":
			v, err := deck.PeekFront()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)
		case "peek-back":
			v, err := deck.PeekBack()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)
		case "show":
			deck.Show()
		case "clear":
			deck.Clear()
		}
	}
}

func s() {
	stack := stackdeck.Stack{}
	f, err := os.Open("stack-deck/expressions.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		for _, el := range line {
			v, err := strconv.ParseFloat(string(el), 64)
			if err != nil {
				v1, err := stack.Pop()
				if err != nil {
					fmt.Println(err)
					return
				}
				v2, err := stack.Pop()
				if err != nil {
					fmt.Println(err)
					return
				}
				switch s := string(el); s {
				case "+":
					stack.Push(v2 + v1)
				case "-":
					stack.Push(v2 - v1)
				case "*":
					stack.Push(v2 * v1)
				case "/":
					stack.Push(v2 / v1)
				}
			} else {
				stack.Push(v)
			}
		}
		ans, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Expression #%d: %s. Answer: ans = %.3f\n", i, line, ans)
	}
}

func q() {
	first, second := &query.Player{}, &query.Player{}
	query.Shuffle(first, second)
	fmt.Println(first.Cards)
	fmt.Println(second.Cards)
	for (len(first.Cards) != 0) && (len(second.Cards) != 0) {
		f, s := first.Cards[len(first.Cards)-1], second.Cards[len(second.Cards)-1]
		fmt.Printf("First's card: %d\n", f)
		fmt.Printf("Second's card: %d\n\n", s)
		if f > s || (f == 0 && s == 9) {
			first.Cards = append([]int32{s, f}, first.Cards...)
			first.Cards = first.Cards[:len(first.Cards)-1]
			second.Cards = second.Cards[:len(second.Cards)-1]
		} else if s > f || (f == 9 || s == 0) {
			second.Cards = append([]int32{f, s}, second.Cards...)
			first.Cards = first.Cards[:len(first.Cards)-1]
			second.Cards = second.Cards[:len(second.Cards)-1]
		}
		fmt.Println("First")
		fmt.Println(first.Cards)
		fmt.Println("Second")
		fmt.Println(second.Cards)
		time.Sleep(time.Second * 1)
	}
	if len(first.Cards) == 0 {
		fmt.Println("Первый игрок проиграл")
	} else {
		fmt.Println("Второй игрок проиграл")
	}
}

func singly() {
	// TODO switch case
	s := lists.Singly{}
	el := lists.ChNode{}
	var input string
	var idx int32
	for {
		fmt.Scanf("%s", &input)
		fmt.Printf("Введенная команда: %s\n", input)
		if input == "forward" {
			_, err := fmt.Scanln(&el.Symbol)
			if err != nil {
				log.Fatal(err)
			}
			s.PushForward(el)
		} else if input == "back" {
			_, err := fmt.Scanln(&el.Symbol)
			if err != nil {
				log.Fatal(err)
			}
			s.PushBack(el)
		} else if input == "addidx" {
			_, err := fmt.Scanf("%d %s", &idx, &el.Symbol)
			if err != nil {
				log.Fatal(err)
			}
			err = s.AddByIndex(idx, el)
			if errors.Is(err, lists.ErrNotFound) {
				fmt.Printf("index %d out of range\n", idx)
			}
		} else if input == "delval" {
			_, err := fmt.Scanln(&el.Symbol)
			if err != nil {
				log.Fatal(err)
			}
			s.DeleteByVal(el)
		} else if input == "delidx" {
			_, err := fmt.Scanln(&idx)
			if err != nil {
				log.Fatal(err)
			}
			s.DeleteByIdx(idx)
		} else if input == "delall" {
			s.Clear()
		} else if input == "get" {
			_, err := fmt.Scanln(&el.Symbol)
			if err != nil {
				log.Fatal(err)
			}
			idx := s.GetByVal(el)
			if idx == -1 {
				fmt.Println("not found")
				continue
			}
			fmt.Printf("Index of element %s: %d\n", el.Symbol, idx)
		} else if input == "print" {
			s.Print()
		}
	}
}

func doubly() {
	var input string
	d := lists.Doubly{}
	dch := lists.ChDoubly{}
	for {
		fmt.Scanf("%s", &input)
		if input == "forward" {
			fmt.Println("Введите n: ")
			var n int32
			_, err := fmt.Scanln(&n)
			if err != nil {
				log.Fatal(err)
			}
			var number float64
			for i := 0; i < 2*int(n); i++ {
				fmt.Printf("Введите %d-e число: ", i+1)
				_, err = fmt.Scanln(&number)
				if err != nil {
					log.Fatal(err)
				}
				d.PushForward(lists.Node{Number: number})
			}
		} else if input == "back" {
			fmt.Println("Введите n: ")
			var n int32
			_, err := fmt.Scanln(&n)
			if err != nil {
				log.Fatal(err)
			}
			var number float64
			for i := 0; i < 2*int(n); i++ {
				fmt.Printf("Введите %d-e число: ", i+1)
				_, err = fmt.Scanln(&number)
				if err != nil {
					log.Fatal(err)
				}
				d.PushBack(lists.Node{Number: number})
			}
		} else if input == "print" {
			d.Print()
		} else if input == "delete" {
			fmt.Println("Введите число: ")
			var number float64
			_, err := fmt.Scanln(&number)
			if err != nil {
				log.Fatal(err)
			}
			d.DeleteByVal(number)
		} else if input == "clear" {
			d.Clear()
		} else if input == "index" {
			fmt.Println("Введите число: ")
			var number float64
			_, err := fmt.Scanln(&number)
			if err != nil {
				log.Fatal(err)
			}
			idx := d.FindIndex(number)
			if idx == -1 {
				fmt.Println("not found")
			} else {
				fmt.Printf("Index: %d\n", idx)
			}
		} else if input == "mean" {
			mean := d.Mean()
			fmt.Println("Mean: ", mean)
		} else if input == "symbol" {
			var s string
			var bs = []string{",", "!", ":", ";", "?"}
			for {
				fmt.Println("Введите символ: ")
				_, err := fmt.Scanln(&s)
				if err != nil {
					log.Fatal(err)
				}
				if s == "." {
					break
				} else if contains(bs, s) {
					dch.DeleteLastSymbol()
				} else {
					dch.PushBack(lists.ChNode{Symbol: s})
				}
				dch.Print()
			}
		}
	}
}

func contains(arr []string, el string) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}
	return false
}
