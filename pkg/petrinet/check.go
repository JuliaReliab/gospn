package petrinet

import (
	"fmt"
	"sort"
	"strings"
)

// CheckExpressions evaluates every guard, rate, weight and reward once, at the initial
// marking, and reports what fails rather than letting it fail later.
//
// An undefined variable is not a parse error: the expression language resolves variables
// lazily, and an unresolved one evaluates to its own name as a string. Nothing notices
// until a rate is asked for a number, which happens when the transition matrices are
// built -- after the whole reachability search. example/cold_vm_reju.spn (94 states) did
// exactly that: it spent the search and then died with a Go stack trace naming
// `Tvreset.rate`, which is a typo for `Tvrestart.rate`.
//
// A rate that is a string at the initial marking is a string at every marking, so one
// evaluation is enough. The errors are collected rather than returned one at a time: a
// file with two typos in it should say so once.
func (net *Net) CheckExpressions(imark []MarkInt) error {
	var problems []string
	// The message already says which transition and which part of it: the closures the
	// parser builds carry that context, so `what` is only used when they do not.
	check := func(what string, f func()) {
		if f == nil {
			return
		}
		defer func() {
			// The evaluator reports this class by panicking (log.Panic), so a recover
			// is the only way to catch it without rewriting the evaluator.
			if r := recover(); r != nil {
				msg := fmt.Sprint(r)
				if !strings.HasPrefix(msg, what) {
					msg = what + ": " + msg
				}
				problems = append(problems, msg)
			}
		}()
		f()
	}

	for _, tr := range net.translist {
		tr := tr
		if tr.guard != nil {
			check(fmt.Sprintf("transition %s: guard", tr.label), func() { tr.guard(imark) })
		}
		if tr.ratefunc != nil {
			kind := "rate"
			if tr.isImm(net) {
				kind = "weight"
			}
			check(fmt.Sprintf("transition %s: %s", tr.label, kind), func() { tr.ratefunc(imark) })
		}
	}
	for _, label := range sortedRewardLabels(net) {
		f := net.rewardfunc[label]
		check(fmt.Sprintf("reward %s", label), func() { f(imark) })
	}

	if len(problems) == 0 {
		return nil
	}
	return fmt.Errorf("the net does not evaluate:\n  %s", joinLines(problems))
}

// isImm reports whether this transition is one of the net's immediate transitions, which
// only matters for calling the value a weight rather than a rate in a message.
func (tr *Trans) isImm(net *Net) bool {
	for _, im := range net.immlist {
		if im.Trans == tr {
			return true
		}
	}
	return false
}

func sortedRewardLabels(net *Net) []string {
	labels := make([]string, 0, len(net.rewardfunc))
	for label := range net.rewardfunc {
		labels = append(labels, label)
	}
	sort.Strings(labels)
	return labels
}

func joinLines(s []string) string {
	out := ""
	for i, x := range s {
		if i > 0 {
			out += "\n  "
		}
		out += x
	}
	return out
}
