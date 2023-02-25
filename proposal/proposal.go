package proposal

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type Proposal struct {
	Nonce        uint64          //1
	From         *common.Address //0x2ab0e13bff7f2c07fe4f2c4c28db4a792c83834b - UNAM Address
	Title        string          //Changing the value of copper diselenide from 1.0 to 1.1
	Pos          *big.Int        //1.1
	Height       *big.Int        //-
	FWHMLeft     *big.Int        //-
	DSpacing     *big.Int        //-
	RelIntensity *big.Int        //-
}

func (p *Proposal) Hash() common.Hash {
	h := rlpHash(p)
	return h
}

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

// rlpHash encodes x and hashes the encoded bytes.
func rlpHash(x interface{}) (h common.Hash) {
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h[:])
	return h
}

type ProposalList []*Proposal

func (p ProposalList) Len() int { return len(p) }

//RLP Encoding -> Other Nodes -> RLP Decoding -> Proposal = Signature - UNAM
//Harvard => Reject
//MIT => Accept
//Stanford => Accept
//66% of the nodes accept the proposal
//Proposal is accepted
//Data is added to the blockchain

//168132 * 10^6 = 168132000000
