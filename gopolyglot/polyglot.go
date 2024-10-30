package gopolyglot

import (
	"bufio"
	"io"
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

	reader := bufio.NewReader(f)

	buf := make([]byte, 16, 16)

	for {
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			break
		}

		var entry Entry

		entry.Key = ReadBigEndian64(buf[0:8])
		entry.Move = ReadBigEndian16(buf[8:10])
		entry.Weight = ReadBigEndian16(buf[10:12])
		entry.Learn = ReadBigEndian32(buf[12:16])

		ret = append(ret, entry)
	}

	return ret, nil
}

// Is there a better way than writing these functions?
// Yeah, probably using binary.Read() on the bufio buffer

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
