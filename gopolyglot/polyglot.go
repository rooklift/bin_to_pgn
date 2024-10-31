package gopolyglot

// TODO - remember polyglot uses e1h1 castling format.
// Remember to check such moves actually have a king there (rather than being some other piece).

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
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
		case 2: prom_string = "b"
		case 3: prom_string = "r"
		case 4: prom_string = "q"
		default: prom_string = ""
	}

	return fmt.Sprintf("%c%c%c%c%s", from_file_c, from_row_c, to_file_c, to_row_c, prom_string)
}

func Probe(book []Entry, target uint64) []Entry {

	// See the Golang docs for how sort.Search works.

	index := sort.Search(len(book), func(i int) bool { return book[i].Key >= target })

	if index < len(book) && book[index].Key == target {
		return ExtractNeighbours(book, index)
	} else {
		return nil
	}
}

func ExtractNeighbours(book []Entry, index int) []Entry {

	// Given an index into the book, return all entries that have that key.

	target := book[index].Key

	// Our Probe() function is actually guaranteed to locate the very first
	// such item, rendering the first half of this function redundant. Meh.

	for {
		if index == 0 || book[index - 1].Key != target {
			break
		}
		index -= 1
	}

	// So index is now the very first of the correct entries...

	var ret []Entry

	for {
		ret = append(ret, book[index])
		if index >= len(book) - 1 || book[index + 1].Key != target {
			break
		}
		index += 1
	}

	return ret
}
