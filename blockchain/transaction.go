package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	ID      []byte
	Inputs  []TXInput
	Outputs []TXOutput
}

type TXInput struct {
	ID  []byte
	Out int
	Sig string
}

type TXOutput struct {
	Value  int
	PubKey string
}

func (tx *Transaction) SetId() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

}

func CoinBaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{100, to}

	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetId()

	return &tx
}

func NewTransaction(from, amount, to string) *Transaction {}

func (tx *Transaction) IsCoinBase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TXInput) CanUnlock(data string) bool {
	return in.Sig == data
}
func (out *TXOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
