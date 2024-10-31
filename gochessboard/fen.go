package gochessboard

import (
	"fmt"
	"strings"
)

func BoardFromFEN(fen string) (*Board, error) {

	tokens := strings.Fields(fen)
	if len(tokens) != 6 {
		return nil, fmt.Errorf("BoardFromFEN: got %d tokens", len(tokens))
	}

	rows := strings.Split(tokens[0], "/")
	if len(rows) != 8 {
		return nil, fmt.Errorf("BoardFromFEN: got %d rows", len(rows))
	}

	board := new(Board)

	for y := 0; y < 8; y++ {
		x := 0
		for _, c := range rows[y] {
			if x >= 8 {
				return nil, fmt.Errorf("BoardFromFEN: row too long")
			}
			if c >= '1' && c <= '8' {
				x += int(c - 48)
			} else {
				switch c {
					case 'K': board.State[IndexFromXY(x, y)] = K_w
					case 'Q': board.State[IndexFromXY(x, y)] = Q_w
					case 'R': board.State[IndexFromXY(x, y)] = R_w
					case 'B': board.State[IndexFromXY(x, y)] = B_w
					case 'N': board.State[IndexFromXY(x, y)] = N_w
					case 'P': board.State[IndexFromXY(x, y)] = P_w
					case 'k': board.State[IndexFromXY(x, y)] = K_b
					case 'q': board.State[IndexFromXY(x, y)] = Q_b
					case 'r': board.State[IndexFromXY(x, y)] = R_b
					case 'b': board.State[IndexFromXY(x, y)] = B_b
					case 'n': board.State[IndexFromXY(x, y)] = N_b
					case 'p': board.State[IndexFromXY(x, y)] = P_b
					default: return nil, fmt.Errorf("BoardFromFEN: invalid character '%c' in board", c)
				}
				x += 1
			}
		}
	}

	switch tokens[1] {
		case "W": fallthrough
		case "w": board.Wmove = true
		case "B": fallthrough
		case "b": board.Wmove = false
		default: return nil, fmt.Errorf("BoardFromFEN: invalid to-move character")
	}

	if strings.Contains(tokens[2], "K") { board.Castling |= CastleWhiteKingside }
	if strings.Contains(tokens[2], "Q") { board.Castling |= CastleWhiteQueenside }
	if strings.Contains(tokens[2], "k") { board.Castling |= CastleBlackKingside }
	if strings.Contains(tokens[2], "q") { board.Castling |= CastleBlackQueenside }

	i, err := IndexFromString(tokens[3])
	if err != nil {
		board.EP = i
	} else {
		board.EP = 0
	}
	board.ClearBadEP()

	return board, nil
}
