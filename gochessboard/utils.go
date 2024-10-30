package gochessboard

import "fmt"

func IndexFromString(s string) (boardindex, error) {		// "a8" --> 0
	if len(s) != 2 {
		return 0, fmt.Errorf("IndexFromCoord: invalid argument %s", s)
	}
	x := s[0] - 97
	y := 8 - (s[1] - 48)
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return 0, fmt.Errorf("IndexFromCoord: invalid argument %s", s)
	}
	return boardindex(y * 8 + x), nil
}

func StringFromIndex(i boardindex) string {					// 0 --> "a8"
	xc := int('a')
	yc := int('8')
	x := XFromIndex(i)
	y := YFromIndex(i)
	xc += x
	yc -= y
	return fmt.Sprintf("%c%c", rune(xc), rune(yc))
}

func XFromIndex(i boardindex) int {							// Where x (col) is 0-7
	return int(i % 8)
}

func YFromIndex(i boardindex) int {							// Where y (row) is 0-7, from top
	return int(i / 8)
}

func IndexFromXY(x, y int) boardindex {						// 0, 0 --> 0 (top left)
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		panic("IndexFromXY: bad arg")
	}
	return boardindex(y * 8 + x)
}