// ide prints output in a format easily digestible by an IDE eg Goland
package printer

import (
	"fmt"
	"io"
	"sort"

	"github.com/mibk/dupl/syntax"
)

type ide struct {
	w io.Writer
	ReadFile
}

func NewIde(w io.Writer, fread ReadFile) Printer {
	return &ide{w, fread}
}

func (ii ide) PrintHeader() error {
	return nil
}

func (ii ide) PrintClones(dups [][]*syntax.Node) error {
	clones, err := prepareClonesInfo(ii.ReadFile, dups)
	if err != nil {
		return err
	}
	sort.Sort(byNameAndLine(clones))
	for i, cl := range clones {
		nextCl := clones[(i+1)%len(clones)]
		if cl.lineEnd-cl.lineStart < 5 { // skip small differences TODO flag
			continue
		}
		if nextCl.lineStart-cl.lineStart < 3 { // skip close lines TODO flag
			continue
		}
		fmt.Fprintf(ii.w, "%s:%d: %d-%d is duplicate of %s:%d-%d\n", cl.filename, cl.lineStart, cl.lineStart, cl.lineEnd, nextCl.filename, nextCl.lineStart, nextCl.lineEnd)
	}
	return nil
}

func (ii ide) PrintFooter() error {
	return nil
}
