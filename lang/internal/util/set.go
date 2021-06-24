package util

type Set map[string]bool

func (o *Set) Add(item string) {
	(*o)[item] = true
}

func (o *Set) Delete(item string) {
	delete(*o, item)
}

func (o *Set) Contains(item string) bool {
	_, ok := (*o)[item]
	return ok
}

func (o *Set) Length() int {
	return len(*o)
}

func (o *Set) Keys() []string {
	keys := make([]string, len(*o))
	i := 0
	for k := range *o {
		keys[i] = k
		i++
	}
	return keys
}

func NewSet(items ...string) *Set {
	set := make(Set)
	for _, item := range items {
		set.Add(item)
	}
	return &set
}
