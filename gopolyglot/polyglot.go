package gopolyglot

import (
	"bufio"
	"encoding/binary"
	"os"
)

type Entry struct {
	Key		uint64
	Move	uint16
	Weight	uint16
	Learn	uint32
}

func LoadFromFile(filepath string) ([]Entry, error) {

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
		ret = append(ret, entry)
	}

	return ret, nil
}

/*

func ReadBigEndian64(b []byte) uint64 {
	var ret uint64
	ret += uint64(b[0]) << 56
	ret += uint64(b[1]) << 48
	ret += uint64(b[2]) << 40
	ret += uint64(b[3]) << 32
	ret += uint64(b[4]) << 24
	ret += uint64(b[5]) << 16
	ret += uint64(b[6]) <<  8
	ret += uint64(b[7]) <<  0
	return ret
}

func ReadBigEndian32(b []byte) uint32 {
	var ret uint32
	ret += uint32(b[0]) << 24
	ret += uint32(b[1]) << 16
	ret += uint32(b[2]) <<  8
	ret += uint32(b[3]) <<  0
	return ret
}

func ReadBigEndian16(b []byte) uint16 {
	var ret uint16
	ret += uint16(b[0]) << 8
	ret += uint16(b[1]) << 0
	return ret
}

*/
