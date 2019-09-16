package htmlutil

import (
	"html"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// ChartColor1 is the color palette for Google Charts as collected by
// Craig Davis here: https://gist.github.com/there4/2579834
var ChartColor1 = [...]string{
	"#3366CC",
	"#DC3912",
	"#FF9900",
	"#109618",
	"#990099",
	"#3B3EAC",
	"#0099C6",
	"#DD4477",
	"#66AA00",
	"#B82E2E",
	"#316395",
	"#994499",
	"#22AA99",
	"#AAAA11",
	"#6633CC",
	"#E67300",
	"#8B0707",
	"#329262",
	"#5574A6",
	"#3B3EAC",
}

// Link is a struct to hold information for an HTML link.
type Link struct {
	Href      string
	InnerHtml string
}

const (
	Color2GreenHex       = "#00FF2A"
	Color2YellowHex      = "#DFDD13"
	Color2RedHex         = "#FF0000"
	RingCentralOrangeHex = "#FF8800"
	RingCentralBlueHex   = "#0073AE"
	RingCentralGreyHex   = "#585858"
)

var (
	bluemondayStrictPolicy                      = bluemonday.StrictPolicy()
	rxHtmlToTextNewLines         *regexp.Regexp = regexp.MustCompile(`(?i:</?p>)`)
	rxCarriageReturnLineFeed     *regexp.Regexp = regexp.MustCompile(`\r\n`)
	rxLineFeedMore2              *regexp.Regexp = regexp.MustCompile(`\n\n+`)
	rxCarriageReturnLineFeedMore *regexp.Regexp = regexp.MustCompile(`[\r\n]+`)
	rxEndingSpacesLineFeed       *regexp.Regexp = regexp.MustCompile(`\s+\n`)
)

// HTMLToTextCondensed removes HTML tags, unescapes HTML entities,
// and removes extra whitespace including non-breaking spaces.
func HTMLToTextCondensed(s string) string {
	return strings.Join(
		strings.Fields(
			strings.TrimSpace(
				html.UnescapeString(
					bluemondayStrictPolicy.Sanitize(s),
				),
			)),
		" ",
	)
}

// HTMLToText converts HTML to multi-line text.
func HTMLToText(s string) string {
	return rxLineFeedMore2.ReplaceAllString(
		strings.TrimSpace(
			html.UnescapeString(
				bluemondayStrictPolicy.Sanitize(
					rxHtmlToTextNewLines.ReplaceAllString(s, "$1\n\n"),
				),
			),
		),
		"\n\n",
	)
}
