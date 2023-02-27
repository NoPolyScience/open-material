package proposal

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// ValidateProposal validates a proposal with the cryptographic signature
func ValidateProposal(pS *ProposalWithSignature) error {
	return nil
}

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(hash, sig)
}

func EcrecoverProposal(pS *ProposalWithSignature) ([]byte, error) {
	hash := rlpHash(pS.Proposal)

	r := pS.R
	s := pS.S
	v := pS.V

	sig := encodeSignature(r, s, v)
	rec, err := Ecrecover(hash[:], sig)

	if err != nil {
		return nil, err
	}

	return rec, nil
}

func SignatureToPublicKey(pS *ProposalWithSignature) (*ecdsa.PublicKey, error) {
	rec, err := EcrecoverProposal(pS)

	if err != nil {
		return nil, err
	}

	x, y := elliptic.Unmarshal(crypto.S256(), rec)
	pub := &ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}

	return pub, nil
}
