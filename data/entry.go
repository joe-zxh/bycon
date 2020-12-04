package data

import (
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"sync"
)

type Command string

// BlockHash represents a SHA256 hashsum of a Block
type EntryHash [64]byte

func (d EntryHash) String() string {
	return hex.EncodeToString(d[:])
}

func (d EntryHash) ToSlice() []byte {
	return d[:]
}

type EntryID struct {
	V uint32
	N uint32
}

func (id1 *EntryID) IsOlder(id2 *EntryID) bool {
	if (id1.V == id2.V && id1.N < id2.N) || id1.V < id2.V {
		return true
	} else {
		return false
	}
}

func (id1 *EntryID) IsOlderOrEqual(id2 *EntryID) bool {
	if (id1.V == id2.V && id1.N <= id2.N) || id1.V < id2.V {
		return true
	} else {
		return false
	}
}

func (id1 *EntryID) IsNewerOrEqual(id2 *EntryID) bool {
	if (id1.V == id2.V && id1.N >= id2.N) || id1.V > id2.V {
		return true
	} else {
		return false
	}
}

type Entry struct {
	Mut       sync.Mutex
	PP        *PrePrepareArgs
	P         []*PrepareArgs
	Prepared  bool
	C         []*CommitArgs
	Committed bool

	PreEntryHash *EntryHash
	Digest       *EntryHash
}

func (e *Entry) String() string {
	return fmt.Sprintf("Entry{View: %d, Seq: %d, Committed: %v}",
		e.PP.View, e.PP.Seq, e.Committed)
}

// Hash returns a hash digest of the block.
func (e *Entry) GetDigest() EntryHash {
	// return cached hash if available
	if e.Digest != nil {
		return *e.Digest
	}

	if e.PP == nil {
		panic(`PrePrepare args of entry is empty!!!`)
	}

	s512 := sha512.New()

	s512.Write(e.PreEntryHash.ToSlice())

	byte4 := make([]byte, 4)
	binary.LittleEndian.PutUint32(byte4, uint32(e.PP.View))
	s512.Write(byte4[:])

	binary.LittleEndian.PutUint32(byte4, uint32(e.PP.Seq))
	s512.Write(byte4[:])

	for _, cmd := range *e.PP.Commands {
		s512.Write([]byte(cmd))
	}

	e.Digest = new(EntryHash)
	sum := s512.Sum(nil)
	copy(e.Digest[:], sum)

	return *e.Digest
}
