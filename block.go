package main

import (
	"golang.org/x/tools/cover"
)

// ProfileBlock wraps the cover.ProfileBlock structure and adds
// additional merge and validation functionality to it.
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

// ImportCount adds the count from the given cover.ProfileBlock
// struct. This only works if the rest of the ProfileBlock fields
// are the same.
func (b *ProfileBlock) ImportCount(other cover.ProfileBlock) {
	if !b.sameRange(other) {
		panic("not the same range")
	}
	b.Count += other.Count
}
