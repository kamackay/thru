package model

type Opts struct {
	File string `arg:"d" help:"File to print to" default:"."`

	Highlight  string `short:"h" help:"Highlight text that matches this"`
	Timestamps bool   `short:"t" help:"Print Timestamps"`
	Version    bool   `help:"Show version"`
}
