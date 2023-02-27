package proposal

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// ValidateProposal validates a proposal with the cryptographic signature

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(hash, sig)
}

func EcrecoverProposal(hash *common.Hash, pS *ProposalWithSignature) ([]byte, error) {
	//hash := rlpHash(pS.Proposal)

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

func ValidateProposal(pS *ProposalWithSignature) (bool, error) {
	hash := rlpHash(pS.Proposal)
	rec, err := EcrecoverProposal(&hash, pS)

	if err != nil {
		return false, err
	}

	x, y := elliptic.Unmarshal(crypto.S256(), rec)
	pub := &ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}
	compressed := secp256k1.CompressPubkey(pub.X, pub.Y)
	encodedSig := encodeSignature(pS.R, pS.S, pS.V)

	result := secp256k1.VerifySignature(compressed, hash[:], encodedSig)
	return result, nil
}
