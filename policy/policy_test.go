package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 91, Capacity: 90, Latency: 22, Risk: 5, Weight: 9}, wantScore: 238, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 77, Capacity: 82, Latency: 16, Risk: 21, Weight: 6}, wantScore: 87, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 101, Capacity: 91, Latency: 8, Risk: 18, Weight: 7}, wantScore: 186, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
