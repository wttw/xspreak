package result

import (
	"fmt"
	"go/token"

	"golang.org/x/tools/go/packages"

	"github.com/vorlif/spreak/pkg/po"
)

type Issue struct {
	FromExtractor string

	Domain   string
	Context  string
	MsgID    string
	PluralID string

	Comments []string
	Flags    []string

	Pkg *packages.Package

	Pos token.Position

	Message *po.Message
}

func (i *Issue) FilePath() string {
	return i.Pos.Filename
}

func (i *Issue) Line() int {
	return i.Pos.Line
}

func (i *Issue) Column() int {
	return i.Pos.Column
}

func (i *Issue) Description() string {
	return fmt.Sprintf("%s: %s", i.FromExtractor, i.MsgID)
}
