package main

import (
	"errors"
	"fmt"

	"golang.org/x/tools/cover"
)

// ErrNoBlock gets returned if no given block with the configuration
// yet exists in the newly created blobs.
var ErrNoBlock = errors.New("no block")

// Profile wraps the cover.Profile package and adds an additional
// list which is used during the block merge process.
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

// MergeBlocks combines the Blocks in this ProfileBlock.
func (p *Profile) MergeBlocks() {
	for i, otherBlock := range p.Blocks {
		myBlock, err := p.getBlock(otherBlock)
		if err == nil {
			myBlock.ImportCount(otherBlock)
			continue
		}
		if err == ErrNoBlock {
			p.newBlocks = append(p.newBlocks, &ProfileBlock{&p.Blocks[i]})
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

// Format prints the Blocks into the common cover format which can
// be further processed.
func (p *Profile) Format() string {
	res := ""
	for _, block := range p.Blocks {
		res += fmt.Sprintf("%s:%d.%d,%d.%d %d %d\n",
			p.FileName, block.StartLine, block.StartCol,
			block.EndLine, block.EndCol,
			block.NumStmt, block.Count)
	}
	return res
}
