package main

var (
	buttons = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
)

func createWord(prefix string, list [][]string) []string {
	if len(list) == 0 {
		return []string{prefix}
	}
	ret := make([]string, 0)
	for _, s := range list[0] {
		newprefix := prefix + string(s)
		tmp := createWord(newprefix, list[1:])
		ret = append(ret, tmp...)
	}
	return ret
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	tmp := make([][]string, len(digits))
	for i, d := range digits {
		tmp[i] = buttons[byte(d)]
	}
	return createWord("", tmp)
}
