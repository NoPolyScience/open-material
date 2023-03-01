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
	proposer := pS.Proposal.Proposer

	if err != nil {
		return false, err
	}

	x, y := elliptic.Unmarshal(crypto.S256(), rec)
	pub := &ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}
	recoveredProposer := crypto.PubkeyToAddress(*pub)

	//validate proposer
	if recoveredProposer != *proposer {
		return false, nil
	}

	compressed := crypto.CompressPubkey(pub)

	encodedSig := encodeSignature(pS.R, pS.S, pS.V)
	encodedSigNoRecover := encodedSig[:len(encodedSig)-1] // remove recovery byte

	result := secp256k1.VerifySignature(compressed, hash[:], encodedSigNoRecover)
	return result, nil
}
