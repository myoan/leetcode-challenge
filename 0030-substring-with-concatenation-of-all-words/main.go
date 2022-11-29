package main

type TrieTree struct {
	Val      int
	Children []*TrieTree
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		Val:      -1,
		Children: []*TrieTree{},
	}
}

func (t *TrieTree) Insert(str string) {
	if str == "" {
		return
	}
	i := int(str[0]) - int('a')
	for _, child := range t.Children {
		if child.Val == i {
			child.Insert(str[1:])
			return
		}
	}
	tt := &TrieTree{
		Val:      i,
		Children: []*TrieTree{},
	}
	tt.Insert(str[1:])
	t.Children = append(t.Children, tt)
}

func (t *TrieTree) Find(str string) bool {
	if str == "" {
		if t.IsLeaf() {
			return true
		}
		return false
	}
	i := int(str[0]) - int('a')
	for _, child := range t.Children {
		if child.Val == i {
			return child.Find(str[1:])
		}
	}
	return false
}

func (t *TrieTree) IsLeaf() bool {
	return len(t.Children) == 0
}

type Dict struct {
	raw map[string]int
}

func NewDict(words []string) *Dict {
	ret := &Dict{
		raw: make(map[string]int, len(words)),
	}
	for _, str := range words {
		ret.raw[str] += 1
	}
	return ret
}

func (d *Dict) initVal(words []string) {
	for k, _ := range d.raw {
		d.raw[k] = 0
	}
	for _, str := range words {
		d.raw[str] += 1
	}
}

func (d *Dict) findAll() bool {
	for _, v := range d.raw {
		if v != 0 {
			return false
		}
	}
	return true
}

func (d *Dict) find(str string) {
	d.raw[str] -= 1
}

func (d *Dict) get(str string) bool {
	return d.raw[str] == 0
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	i := 0
	wlen := len(words[0])
	wsize := len(words)
	tree := NewTrieTree()
	dict := NewDict(words)
	for _, str := range words {
		tree.Insert(str)
	}

	ret := make([]int, 0)
	for {
		if i > len(s)-wlen*wsize {
			return ret
		}

		j := i
		for {
			eos := j + wlen
			if eos > len(s) {
				return ret
			}

			substr := s[j:eos]
			if tree.Find(substr) && !dict.get(substr) {
				dict.find(substr)

				if dict.findAll() {
					ret = append(ret, i)
					break
				}
			} else {
				break
			}
			j = eos
		}

		dict.initVal(words)
		i++
	}
}
