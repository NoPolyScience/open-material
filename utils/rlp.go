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

func EncodeProposalWithSig(p *proposal.ProposalWithSignature) ([]byte, error) {
	out, err := rlp.EncodeToBytes(p)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func DecodeProposalWithSig(encoded []byte, p *proposal.ProposalWithSignature) error {
	err := rlp.DecodeBytes(encoded, p)
	if err != nil {
		return err
	}
	return nil
}
