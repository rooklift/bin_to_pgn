package gopolyglot

// TODO - remember polyglot uses e1h1 castling format.
// Remember to check such moves actually have a king there (rather than being some other piece).

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"sort"

	gcb "github.com/rooklift/bin_to_pgn/gochessboard"
)

type Entry struct {
	Key		uint64
	Move	uint16
	Weight	uint16
	Learn	uint32
}

func (self *Entry) MoveString() string {
	return ParseMove(self.Move)
}

type Book struct {
	Entries	[]Entry
}

func LoadFromFile(filepath string) (*Book, error) {

	var previous_key uint64 = 0
	var have_warned bool = false

	ret := new(Book)

	f, err := os.Open(filepath)
	if err != nil {
		return ret, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		var entry Entry
		err := binary.Read(reader, binary.BigEndian, &entry)
		if err != nil {
			break
		}
		if !have_warned && entry.Key < previous_key {
			fmt.Fprintf(os.Stderr, "Warning: book was not in order! This will invalidate results.\n")
			have_warned = true			// TODO - maybe we should just sort the thing.
		}
		previous_key = entry.Key
		ret.Entries = append(ret.Entries, entry)
	}

	return ret, nil
}

func ParseMove(b uint16) string {

	to_file   := b >>  0 & 0b111
	to_row    := b >>  3 & 0b111
	from_file := b >>  6 & 0b111
	from_row  := b >>  9 & 0b111
	promotion := b >> 12 & 0b111

	to_file_c := 'a' + to_file
	to_row_c  := '1' + to_row

	from_file_c := 'a' + from_file
	from_row_c  := '1' + from_row

	prom_string := ""

	switch promotion {
		case 1: prom_string = "n"
		case 2: prom_string = "b"
		case 3: prom_string = "r"
		case 4: prom_string = "q"
		default: prom_string = ""
	}

	return fmt.Sprintf("%c%c%c%c%s", from_file_c, from_row_c, to_file_c, to_row_c, prom_string)
}

func (self *Book) Probe(target uint64) []Entry {

	// See the Golang docs for how sort.Search works.

	index := sort.Search(len(self.Entries), func(i int) bool { return self.Entries[i].Key >= target })

	if index < len(self.Entries) && self.Entries[index].Key == target {
		return self.ExtractNeighbours(index)
	} else {
		return nil
	}
}

func (self *Book) ExtractNeighbours(index int) []Entry {

	// Given an index into the book, return all entries that have that key.

	target := self.Entries[index].Key

	// Our Probe() function is actually guaranteed to locate the very first
	// such item, rendering the first half of this function redundant. Meh.

	for {
		if index == 0 || self.Entries[index - 1].Key != target {
			break
		}
		index -= 1
	}

	// So index is now the very first of the correct entries...

	var ret []Entry

	for {
		ret = append(ret, self.Entries[index])
		if index >= len(self.Entries) - 1 || self.Entries[index + 1].Key != target {
			break
		}
		index += 1
	}

	return ret
}

func KeyFromBoard(board *gcb.Board) uint64 {

	var key uint64

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := board.State[y * 8 + x]
			if piece == gcb.EMPTY {
				continue
			}
			var polyglot_kind int
			switch piece {
				case gcb.P_b: polyglot_kind = 0
				case gcb.P_w: polyglot_kind = 1
				case gcb.N_b: polyglot_kind = 2
				case gcb.N_w: polyglot_kind = 3
				case gcb.B_b: polyglot_kind = 4
				case gcb.B_w: polyglot_kind = 5
				case gcb.R_b: polyglot_kind = 6
				case gcb.R_w: polyglot_kind = 7
				case gcb.Q_b: polyglot_kind = 8
				case gcb.Q_w: polyglot_kind = 9
				case gcb.K_b: polyglot_kind = 10
				case gcb.K_w: polyglot_kind = 11
			}
			index := (64 * polyglot_kind) + (8 * (7 - y)) + x
			key ^= PolyglotPieceXorVals[index]
		}
	}

	if board.Castling & gcb.CastleWhiteKingside  > 0 { key ^= PolyglotCastleXorVals[0] }
	if board.Castling & gcb.CastleWhiteQueenside > 0 { key ^= PolyglotCastleXorVals[1] }
	if board.Castling & gcb.CastleBlackKingside  > 0 { key ^= PolyglotCastleXorVals[2] }
	if board.Castling & gcb.CastleBlackQueenside > 0 { key ^= PolyglotCastleXorVals[3] }

	if board.EP > 0 {
		x := gcb.XFromIndex(board.EP)
		key ^= PolyglotEnPassantXorVals[x]
	}

	if board.Wmove { key ^= PolyglotActiveXorVal }

	return key

}
