package main

import (
	"golang.org/x/tools/cover"
	"os"
	"fmt"
	"errors"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax: gocov-merge FILENAME\n")
		os.Exit(1)
	}
	profiles, err := cover.ParseProfiles(os.Args[1])
	if err != nil {
		fmt.Println("Unable to parse file")
		os.Exit(1)
	}
	for _, profile := range profiles {
		p := &Profile{profile, make([]*ProfileBlock, 0)}
		p.MergeBlocks()
		fmt.Print(p.Format())
	}
}

var ErrNoBlock = errors.New("no block")

func (p *Profile) getBlock(otherBlock cover.ProfileBlock) (*ProfileBlock, error) {
	for _, block := range p.newBlocks {
		if block.sameRange(otherBlock) {
			return block, nil
		}
	}
	return nil, ErrNoBlock
}

func (b *ProfileBlock) sameRange(other cover.ProfileBlock) bool {
	if b.StartLine != other.StartLine {
		return false
	}
	if b.StartCol != other.StartCol {
		return false
	}
	if b.EndLine != other.EndLine {
		return false
	}
	if b.EndCol != other.EndCol {
		return false
	}
	if b.NumStmt != other.NumStmt {
		return false
	}
	return true
}

func (b *ProfileBlock) ImportCount(other cover.ProfileBlock) {
	if !b.sameRange(other) {
		panic("not the same range")
	}
	b.Count += other.Count
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

type Profile struct {
	*cover.Profile
	newBlocks []*ProfileBlock
}

type ProfileBlock struct {
	*cover.ProfileBlock
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
