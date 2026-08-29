package result

// Provenance is what a result file records about the run that produced it.
//
// It used to be assembled inline in `gospn sim` and did not exist at all for
// `gospn mark`, so a marking graph on disk said nothing about which binary or which
// net it came from. The version matters in particular because 0.16.0 changed which
// random stream a given seed produces: from that release a seed alone no longer
// identifies a run.
type Provenance struct {
	Version  string // gospn version, from -ldflags at build time
	Revision string // vcs revision, empty when not built from a checkout
	Net      string // the definition file, empty when read from stdin
	Command  string // the subcommand that wrote the file
}

// AddTo appends the non-empty fields as text elements. Run parameters that belong to
// one subcommand only -- the seed, the horizon -- are added by that subcommand.
func (p Provenance) AddTo(r *Result) {
	if p.Version != "" {
		r.AddText("gospn_version", p.Version)
	}
	if p.Revision != "" {
		r.AddText("gospn_revision", p.Revision)
	}
	if p.Command != "" {
		r.AddText("gospn_command", p.Command)
	}
	if p.Net != "" {
		r.AddText("net", p.Net)
	}
}
