package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieTree_Insert(t *testing.T) {
	root := NewTrieTree()
	root.Insert("hoge")

	wordidx := int('a')
	node := root

	assert.Equal(t, node.Children[0].Val, int('h')-wordidx, "")
	node = node.Children[0]
	assert.Equal(t, node.Children[0].Val, int('o')-wordidx, "")
	node = node.Children[0]
	assert.Equal(t, node.Children[0].Val, int('g')-wordidx, "")
	node = node.Children[0]
	assert.Equal(t, node.Children[0].Val, int('e')-wordidx, "")

	root.Insert("home")
	node = root

	assert.Equal(t, node.Children[0].Val, int('h')-wordidx, "")
	node = node.Children[0]
	assert.Equal(t, node.Children[0].Val, int('o')-wordidx, "")
	node = node.Children[0]
	assert.Equal(t, node.Children[1].Val, int('m')-wordidx, "")
	node = node.Children[1]
	assert.Equal(t, node.Children[0].Val, int('e')-wordidx, "")

	assert.Equal(t, root.Find("hoge"), true, "")
	assert.Equal(t, root.Find("hog"), false, "")
	assert.Equal(t, root.Find("home"), true, "")
}

func Test_fundSubstring(t *testing.T) {
	testcases := []struct {
		msg    string
		s      string
		words  []string
		expect []int
	}{
		{
			msg:    "empty case",
			s:      "",
			words:  []string{},
			expect: []int{},
		},
		{
			msg:    "when words is one",
			s:      "barfoothefoobarman",
			words:  []string{"foo"},
			expect: []int{3, 9},
		},
		{
			msg:    "when words match at last",
			s:      "barfoothefoo",
			words:  []string{"foo"},
			expect: []int{3, 9},
		},
		{
			msg:    "when words match at first and last",
			s:      "foothefoo",
			words:  []string{"foo"},
			expect: []int{0, 6},
		},
		{
			msg:    "example 1",
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			expect: []int{0, 9},
		},
		{
			msg:    "example 2",
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			expect: []int{},
		},
		{
			msg:    "example 3",
			s:      "barfoofoobarthefoobarman",
			words:  []string{"foo", "bar", "the"},
			expect: []int{6, 9, 12},
		},
		{
			msg:    "failed answer 1",
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "good"},
			expect: []int{8},
		},
	}

	for _, tt := range testcases {
		actual := findSubstring(tt.s, tt.words)
		if !matchIntArray(actual, tt.expect) {
			t.Errorf("[%s] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
	}
}
