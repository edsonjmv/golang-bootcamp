package main

func (o *order) addItem(item string) {
	if (*o)[item] > 0 {
		(*o)[item]++
	} else {
		(*o)[item] = 1
	}
}

func (o *order) removeItem(item string) {
	delete(*o, item)
}

func (o *order) clearOrder() {
	for item := range *o {
		delete(*o, item)
	}
}
