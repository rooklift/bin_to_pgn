package gochessboard

import (
	"fmt"
	"strings"
)

func (self *Board) String() string {

	if self == nil {
		return "<nil>"
	}

	var substrings []string

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			substrings = append(substrings, fmt.Sprintf("%c", PieceRune(self.State[y * 8 + x])))
		}
		if y == 7 {
			substrings = append(substrings, fmt.Sprintf("    Castling: %v    E.P: %v", self.CastlingString(), self.EnPassantString()))
		} else {
			substrings = append(substrings, fmt.Sprintf("\n"))
		}
	}

	return strings.Join(substrings, "")
}

func (self *Board) CastlingString() string {
	s := ""
	if self.Castling & castle_white_kingside > 0 { s += "K" }
	if self.Castling & castle_white_queenside > 0 { s += "Q" }
	if self.Castling & castle_black_kingside > 0 { s += "k" }
	if self.Castling & castle_black_queenside > 0 { s += "q" }
	return s
}

func (self *Board) EnPassantString() string {
	if self.EP == 0 {
		return "-"
	}
	return StringFromIndex(self.EP)
}
