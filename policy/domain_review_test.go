package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 80, Slack: 47, Drag: 20, Confidence: 84}
	if got := DomainReviewScore(item); got != 231 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "ship" {
		t.Fatalf("lane = %s", got)
	}
}
