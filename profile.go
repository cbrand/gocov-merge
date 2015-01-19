package main

import (
	"errors"
	"fmt"

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
	for i, otherBlock := range p.Blocks {
		myBlock, err := p.getBlock(otherBlock)
		fmt.Println("loop")
		if err == nil {
			myBlock.ImportCount(otherBlock)
			continue
		}
		if err == ErrNoBlock {
			fmt.Println("new block")
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
