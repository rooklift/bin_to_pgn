package gochessboard

import (
	"fmt"
)

const (						// Bitmasks
	castle_white_kingside	= 0b0001
	castle_white_queenside	= 0b0010
	castle_black_kingside 	= 0b0100
	castle_black_queenside	= 0b1000
	castle_white_both		= castle_white_kingside | castle_white_queenside
	castle_black_both		= castle_black_kingside | castle_black_queenside
	castle_all				= castle_white_both | castle_black_both
)

type Board struct {
	State		[64]piecetype
	Wmove		bool
	Castling	uint8			// bitmask
	EP			boardindex		// which square, or 0 if N/A (although 0 is a real square, it's never a valid EP square)
}

func (self *Board) Copy() *Board {
	ret := new(Board)
	*ret = *self
	return ret
}

func (self *Board) ForceMove(mv string) (*Board, error) {		// No legality checks, just a few syntax checks

	if len(mv) != 4 && len(mv) != 5 {
		return nil, fmt.Errorf("Invalid move string")
	}

	start, err1 := IndexFromString(mv[0:2])
	end, err2 := IndexFromString(mv[2:4])
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("Invalid move string")
	}

	start_x := XFromIndex(start)
	start_y := YFromIndex(start)
	end_x := XFromIndex(end)
	end_y := YFromIndex(end)

	ret := self.Copy()

	// Simply move the piece...

	piece := ret.State[start]
	ret.State[end] = piece
	ret.State[start] = EMPTY

	// Move rook if this is castling...

	if piece == K_w || piece == K_b {
		if start == e8 && end == g8 {			// Black: O-O
			ret.State[h8] = EMPTY
			ret.State[f8] = R_b
		} else if start == e8 && end == c8 {	// Black: O-O-O
			ret.State[a8] = EMPTY
			ret.State[d8] = R_b
		} else if start == e1 && end == g1 {	// White: O-O
			ret.State[h1] = EMPTY
			ret.State[f1] = R_w
		} else if start == e1 && end == c1 {	// White: O-O-O
			ret.State[a1] = EMPTY
			ret.State[d1] = R_w
		}
	}

	// Clear castling permissions if the king moved...

	if piece == K_w {
		ret.Castling &= (castle_all ^ castle_white_both)
	}
	if piece == K_b {
		ret.Castling &= (castle_all ^ castle_black_both)
	}

	// Clear castling permission if rook moved (or any other piece starting at the rook's spot)...

	if start == a8 {
		ret.Castling &= (castle_all ^ castle_black_queenside)
	}
	if start == h8 {
		ret.Castling &= (castle_all ^ castle_black_kingside)
	}
	if start == a1 {
		ret.Castling &= (castle_all ^ castle_white_queenside)
	}
	if start == h1 {
		ret.Castling &= (castle_all ^ castle_white_kingside)
	}

	// Promotions...

	promotion := EMPTY
	if len(mv) == 5 {
		promotion = PromotionPieceFromChar(mv[4], self.Wmove)
	}

	if piece == P_w && end_y == 0 {		// White pawn reached top row
		if promotion != EMPTY {
			ret.State[end] = promotion
		} else {
			return nil, fmt.Errorf("Promotion move did not state promotion piece")
		}
	}

	if piece == P_b && end_y == 7 {
		if promotion != EMPTY {
			ret.State[end] = promotion
		} else {
			return nil, fmt.Errorf("Promotion move did not state promotion piece")
		}
	}

	// Delete e.p-captured pawn if applicable

	if piece == P_w && start_x != end_x && end == self.EP {
		if ret.State[IndexFromXY(end_x, start_y)] == P_b {
			ret.State[IndexFromXY(end_x, start_y)] = EMPTY
		}
	}

	if piece == P_b && start_x != end_x && end == self.EP {
		if ret.State[IndexFromXY(end_x, start_y)] == P_w {
			ret.State[IndexFromXY(end_x, start_y)] = EMPTY
		}
	}

	// Set e.p. square...

	ret.EP = 0
	if piece == P_w && YFromIndex(start) == 6 && YFromIndex(end) == 4 {
		ret.EP = boardindex(end + 8)		// i.e. one square below
	}
	if piece == P_b && YFromIndex(start) == 1 && YFromIndex(end) == 3 {
		ret.EP = boardindex(end - 8)		// i.e. one square above
	}

	// Switch colours...

	ret.Wmove = !self.Wmove
	ret.ClearBadEP()						// Requires colour to be correct, so do this now.

	return ret, nil
}

func (self *Board) ClearBadEP() {

	if self.EP == 0 {
		return
	}

	x := XFromIndex(self.EP)
	ep_y := YFromIndex(self.EP)

	if (self.Wmove && ep_y != 2) || (!self.Wmove && ep_y != 5) {
		self.EP = 0
		return
	}

	var needed_neighbour piecetype
	var opp_pawns_y int

	if ep_y == 2 {							// E.P. after a black pawn move
		needed_neighbour = P_w
		opp_pawns_y = 3
	} else {								// E.P. after a white pawn move
		needed_neighbour = P_b
		opp_pawns_y = 4
	}

	if x > 0 && self.State[IndexFromXY(x - 1, opp_pawns_y)] == needed_neighbour {
		return
	} else if x < 7 && self.State[IndexFromXY(x + 1, opp_pawns_y)] == needed_neighbour {
		return
	} else {
		self.EP = 0
		return
	}
}
