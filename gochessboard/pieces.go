package gochessboard

type piecetype uint8

const (						// Note the IsWhitePiece and IsBlackPiece functions are sensitive to the exact order
	EMPTY piecetype = 0

	K_w piecetype = 1
	Q_w piecetype = 2
	R_w piecetype = 3
	B_w piecetype = 4
	N_w piecetype = 5
	P_w piecetype = 6

	K_b piecetype = 7
	Q_b piecetype = 8
	R_b piecetype = 9
	B_b piecetype = 10
	N_b piecetype = 11
	P_b piecetype = 12
)

func IsWhitePiece(p piecetype) bool {
	return p >= K_w && p <= P_w
}

func IsBlackPiece(p piecetype) bool {
	return p >= K_b && p <= P_b
}

func PieceRune(p piecetype) rune {
	switch p {
		case K_w: return 'K'
		case Q_w: return 'Q'
		case R_w: return 'R'
		case B_w: return 'B'
		case N_w: return 'N'
		case P_w: return 'P'
		case K_b: return 'k'
		case Q_b: return 'q'
		case R_b: return 'r'
		case B_b: return 'b'
		case N_b: return 'n'
		case P_b: return 'p'
		default: return '.'
	}
}

func PromotionPieceFromChar(c byte, wmove bool) piecetype {		// assuming the char is the last char of a valid UCI move string e.g. e7e8q
	if wmove {
		switch c {
			case 'q': return Q_w
			case 'r': return R_w
			case 'b': return B_w
			case 'n': return N_w
		}
	} else {
		switch c {
			case 'q': return Q_b
			case 'r': return R_b
			case 'b': return B_b
			case 'n': return N_b
		}
	}
	return EMPTY
}
