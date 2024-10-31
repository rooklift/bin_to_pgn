package gopolyglot

// TODO - remember polyglot uses e1h1 castling format.
// Remember to check such moves actually have a king there (rather than being some other piece).

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
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

func LoadFromFile(filepath string) ([]Entry, error) {

	var previous_key uint64 = 0
	var have_warned bool = false

	var ret []Entry

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
		ret = append(ret, entry)
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
		case 2: prom_string = "n"
		case 3: prom_string = "n"
		case 4: prom_string = "n"
		default: prom_string = ""
	}

	return fmt.Sprintf("%c%c%c%c%s", from_file_c, from_row_c, to_file_c, to_row_c, prom_string)
}

func Probe(book []Entry, key uint64) []Entry {

	var ret []Entry

	// lower := 0
	// upper := len(book) - 1

	// TODO - binary search

	return ret
}