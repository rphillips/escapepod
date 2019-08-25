package sanitize

import (
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

func String(s string) string {
	p := bluemonday.NewPolicy()
	p.AllowStandardAttributes()
	p.AllowStandardURLs()
	p.AllowAttrs("href").OnElements("a")
	p.AllowAttrs("usemap").Matching(
		regexp.MustCompile(`(?i)^#[\p{L}\p{N}_-]+$`),
	).OnElements("img")
	p.AllowElements("b", "i", "pre", "small", "strike", "tt", "u")
	p.AllowTables()
	p.AllowImages()
	return p.Sanitize(s)
}
