package stackdeck

type Deck struct {
	front Stack
	back  Stack
}

func (d *Deck) AddFront(el float64) {
	d.front.s = append(d.front.s, el)
}

func (d *Deck) AddBack(el float64) {
	d.back.s = append(d.back.s, el)
}

func (d *Deck) PopFront() (float64, error) {
	if d.IsEmpty() {
		return 0, errEmpty
	}
	if d.front.IsEmpty() {
		firstBack := d.back.s[0]
		d.back.s = d.back.s[1:]
		return firstBack, nil
	}
	last := d.front.s[len(d.front.s)-1]
	d.front.s = d.front.s[:len(d.front.s)-1]
	return last, nil
}

func (d *Deck) PopBack() (float64, error) {
	if d.IsEmpty() {
		return 0, errEmpty
	}
	if d.back.IsEmpty() {
		firstFront := d.front.s[0]
		d.front.s = d.front.s[1:]
		return firstFront, nil
	}
	last := d.back.s[len(d.back.s)-1]
	d.back.s = d.back.s[:len(d.back.s)-1]
	return last, nil
}

func (d *Deck) PeekFront() (float64, error) {
	if d.IsEmpty() {
		return 0, errEmpty
	}
	if d.front.IsEmpty() {
		firstBack := d.back.s[0]
		return firstBack, nil
	}
	last := d.front.s[len(d.front.s)-1]
	return last, nil
}

func (d *Deck) PeekBack() (float64, error) {
	if d.IsEmpty() {
		return 0, errEmpty
	}
	if d.back.IsEmpty() {
		firstFront := d.front.s[0]
		return firstFront, nil
	}
	last := d.back.s[len(d.back.s)-1]
	return last, nil
}

func (d *Deck) IsEmpty() bool {
	return d.front.IsEmpty() && d.back.IsEmpty()
}

func (d *Deck) Show() {
	d.front.Show()
	d.back.Show()
}

func (d *Deck) Clear() {
	d.front = Stack{}
	d.back = Stack{}
}
