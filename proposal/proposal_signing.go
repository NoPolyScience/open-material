package proposal

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ProposalWithSignature is a proposal with a signature
type ProposalWithSignature struct {
	Proposal Proposal
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

func (pS *ProposalWithSignature) Hash(p *Proposal) common.Hash {
	h := rlpHash(p)
	return h
}

func SignProposal(p *Proposal, privKey *ecdsa.PrivateKey) (*ProposalWithSignature, error) {
	hash := rlpHash(p)
	sig, err := crypto.Sign(hash[:], privKey)
	if err != nil {
		return nil, err
	}
	r, s, v := decodeSignature(sig)
	return &ProposalWithSignature{
		Proposal: *p,
		V:        v,
		R:        r,
		S:        s,
	}, nil
}

func decodeSignature(sig []byte) (r, s, v *big.Int) {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v
}

func encodeSignature(r, s, v *big.Int) []byte {
	return append(append(r.Bytes(), s.Bytes()...), byte(v.Uint64()-27))
}
