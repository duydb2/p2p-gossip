package ttl

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const DECREMENT_VALUE = 1

type PeerTTL struct {
	ttl int //time to live
}

//SetPeerTTL sets the ttl attribute.
func (t *PeerTTL) SetTTL(ttl int) error {
	if ttl < 0 {
		return fmt.Errorf("Invalid time to live: %v", ttl)
	}
	t.ttl = ttl
	return nil
}

//GetPeerTTL returns the ttl attribute.
func (t *PeerTTL) GetTTL() int {
	return t.ttl
}

//GetBytes converts the TTL to a byte slice. Returns an error if the byte slice isn't four bytes.
func (t *PeerTTL) GetBytes() []byte {
	buf := make([]byte, 4) //ints are 4 bytes
	binary.PutVarint(buf, int64(t.ttl))
	return buf
}

//CreateTTL takes an integer, and converts it into a TTL.
func (t *PeerTTL) CreateTTL(time int) (TTL, error) {
	if time < 0 {
		return &PeerTTL{}, fmt.Errorf("Invalid time to live: %v", time)
	}
	return &PeerTTL{time}, nil
}

// CreateFromBytes takes a byte slice and turns it into a TTL.
func (t *PeerTTL) CreateFromBytes(time []byte) (TTL, error) {
	if time == nil {
		return &PeerTTL{}, fmt.Errorf("TTL byte slice is nil.")
	}
	var ret int32 //to hold the decoded value
	buf := bytes.NewBuffer(time)
	err := binary.Read(buf, binary.LittleEndian, &ret)
	//	decoded, err := binary.ReadVarint(buf)
	if err != nil {
		return &PeerTTL{}, fmt.Errorf("Error deocding ttl: %v", err)
	}
	return &PeerTTL{int(ret)}, nil
}

//DecrementTTL decrements the TTL by the constant value defined in the package.
func (t *PeerTTL) DecrementTTL() {
	t.ttl = t.ttl - DECREMENT_VALUE
}
