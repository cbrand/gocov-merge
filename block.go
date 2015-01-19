package main

import (
	"fmt"

	"golang.org/x/tools/cover"
)

type ProfileBlock struct {
	*cover.ProfileBlock
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
