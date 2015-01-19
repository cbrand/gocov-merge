package main

import (
	"io/ioutil"
	"path/filepath"

	"golang.org/x/tools/cover"

	. "gopkg.in/check.v1"
)

type ProfileTest struct {
	dir      string
	filePath string
}

var _ = Suite(&ProfileTest{})

func (t *ProfileTest) SetUpTest(c *C) {
	t.dir = c.MkDir()
	t.filePath = filepath.Join(t.dir, "acc.out")
}

func (t *ProfileTest) TestImportFile(c *C) {
	err := ioutil.WriteFile(t.filePath, 
			[]byte("mode: count\n"+
			"github.com/hoffie/larasync/helpers/ed25519/ed25519.go:13.87,19.2 5 1\n"+
			"github.com/hoffie/larasync/helpers/ed25519/ed25519.go:13.87,19.2 5 1\n"+
 			"github.com/hoffie/larasync/helpers/ed25519/ed25519.go:14.87,19.2 5 1"),
		0600)
	
	profiles, err := cover.ParseProfiles(t.filePath)
	c.Assert(err, IsNil)
	
	c.Assert(len(profiles) == 1, Equals, true)
	profile := profiles[0]
	p := &Profile{profile, []*ProfileBlock{}}
	p.MergeBlocks()
	c.Assert(len(p.Blocks), Equals, 2)
	block := p.Blocks[0]
	c.Assert(block.Count, Equals, 2)
	c.Assert(block.StartLine, Equals, 13)
	c.Assert(block.StartCol, Equals, 87)
	c.Assert(block.EndLine, Equals, 19)
	c.Assert(block.EndCol, Equals, 2)
	c.Assert(block.NumStmt, Equals, 5)
	
	block = p.Blocks[1]
	c.Assert(block.Count, Equals, 1)
	c.Assert(block.StartLine, Equals, 14)
	c.Assert(block.StartCol, Equals, 87)
	c.Assert(block.EndLine, Equals, 19)
	c.Assert(block.EndCol, Equals, 2)
	c.Assert(block.NumStmt, Equals, 5)
}
