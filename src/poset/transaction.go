package poset

import (
	"bytes"
	"encoding/json"
	"github.com/andrecronje/lachesis/src/peers"
)

type TransactionType uint8

const (
	PEER_ADD TransactionType = iota
	PEER_REMOVE
)

type InternalTransaction struct {
	Type TransactionType
	Peer peers.Peer
}

func NewInternalTransaction(tType TransactionType, peer peers.Peer) InternalTransaction {
	return InternalTransaction{
		Type: tType,
		Peer: peer,
	}
}

//json encoding of body only
func (t *InternalTransaction) Marshal() ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b) //will write to b
	if err := enc.Encode(t); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (t *InternalTransaction) Unmarshal(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := json.NewDecoder(b) //will read from b
	if err := dec.Decode(t); err != nil {
		return err
	}
	return nil
}
