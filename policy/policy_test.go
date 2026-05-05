package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 91, Capacity: 90, Latency: 22, Risk: 5, Weight: 9}
	if got := Score(signal); got != 238 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 77, Capacity: 82, Latency: 16, Risk: 21, Weight: 6}
	if got := Score(signal); got != 87 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 101, Capacity: 91, Latency: 8, Risk: 18, Weight: 7}
	if got := Score(signal); got != 186 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
