package main

import (
	"errors"

	"golang.org/x/tools/cover"
)

var ErrNoBlock = errors.New("no block")

type Profile struct {
	*cover.Profile
	newBlocks []*ProfileBlock
}

func (p *Profile) getBlock(otherBlock cover.ProfileBlock) (*ProfileBlock, error) {
	for _, block := range p.newBlocks {
		if block.sameRange(otherBlock) {
			return block, nil
		}
	}
	return nil, ErrNoBlock
}

func (p *Profile) MergeBlocks() {
	for _, otherBlock := range p.Blocks {
		myBlock, err := p.getBlock(otherBlock)
		if err == nil {
			myBlock.ImportCount(otherBlock)
			continue
		}
		if err == ErrNoBlock {
			p.newBlocks = append(p.newBlocks, &ProfileBlock{&otherBlock})
			continue
		}
		if err != nil {
			panic(err)
		}
	}
	p.newBlocksToBlocks()
}

func (p *Profile) newBlocksToBlocks() {
	p.Blocks = make([]cover.ProfileBlock, len(p.newBlocks))
	for i, b := range p.newBlocks {
		p.Blocks[i] = *b.ProfileBlock
	}
}