package pairing

import (
	"testing"
	"github.com/fanik05/tcg-pocket-api/internal/models"
)

func TestGenerateMatches(t *testing.T) {
	// Setup: 4 players where 1 & 2 have already played
	players := []models.Player{
		{ID: "1", Name: "Anik", Points: 3, Opponents: []string{"2"}},
		{ID: "2", Name: "Cousin", Points: 3, Opponents: []string{"1"}},
		{ID: "3", Name: "Guest A", Points: 0, Opponents: []string{"4"}},
		{ID: "4", Name: "Guest B", Points: 0, Opponents: []string{"3"}},
	}

	SortPlayersBySwiss(players)
	matches := GenerateMatches(players)

	// Logic Check: With 4 players, we must have exactly 2 matches
	if len(matches) != 2 {
		t.Errorf("Expected 2 matches, got %d", len(matches))
	}

	// Rematch Check: Player 1 should NOT be playing Player 2
	for _, m := range matches {
		if (m.Player1ID == "1" && m.Player2ID == "2") || (m.Player1ID == "2" && m.Player2ID == "1") {
			t.Errorf("Critical Failure: Engine paired Player 1 and 2 again!")
		}
	}
}