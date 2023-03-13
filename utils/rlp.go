package utils

import (
	"github.com/Open-Material/open-material/proposal"
	"github.com/ethereum/go-ethereum/rlp"
)

func EncodeProposal(p *proposal.Proposal) ([]byte, error) {
	out, err := rlp.EncodeToBytes(p)
	if err != nil {
		return nil, err
	}
	return out, nil
}
