package main

var (
	digits = [][]string{
		{"I", "V", "X"},
		{"X", "L", "C"},
		{"C", "D", "M"},
		{"M", "_", "_"},
	}
)

func toRoman(num int, d int) string {
	units := digits[d]
	one := units[0]
	five := units[1]
	ten := units[2]

	switch num {
	case 0:
		return ""
	case 1:
		return one
	case 2:
		return one + one
	case 3:
		return one + one + one
	case 4:
		return one + five
	case 5:
		return five
	case 6:
		return five + one
	case 7:
		return five + one + one
	case 8:
		return five + one + one + one
	case 9:
		return one + ten
	}
	return ""
}

func intToRoman(num int) string {
	div := num
	tmp := make([]string, 4)
	for i := 0; i < 4; i++ {
		mod := div % 10
		if mod == 0 && div == 0 {
			break
		}
		tmp[i] = toRoman(mod, i)
		div = div / 10
	}

	ret := ""
	for i := len(tmp) - 1; i >= 0; i-- {
		ret += tmp[i]
	}
	return ret
}
