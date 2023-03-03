// Package pagination contains the pagination struct and methods
package pagination

import "math"

// Front Binding from JSON
type Front struct {
	NumPages         int  `json:"pages"`
	HasPrev, HasNext bool `json:"-"`
	Offset           int  `json:"-"`
	NumItems         int  `json:"total"`
	ItemsPerPage     int  `json:"items"`
	CurrentPage      int  `json:"current"`
	NextPage         int  `json:"next,omitempty"`
	PrevPage         int  `json:"prev,omitempty"`
}

// Calculate initialize value for pagination purpose
// pages start at 1 - not 0
func (p *Front) Calculate(numItems int) {
	p.NumItems = numItems

	// calculate number of pages
	d := float64(p.NumItems) / float64(p.ItemsPerPage)
	p.NumPages = int(math.Ceil(d))

	// Return the right offset
	p.Offset = (p.CurrentPage - 1) * p.ItemsPerPage

	// HasPrev, HasNext?
	p.HasPrev = p.CurrentPage > 1
	p.HasNext = p.CurrentPage < p.NumPages

	// calculate them
	if p.HasPrev {
		p.PrevPage = p.CurrentPage - 1
	}

	if p.HasNext {
		p.NextPage = p.CurrentPage + 1
	}
}
