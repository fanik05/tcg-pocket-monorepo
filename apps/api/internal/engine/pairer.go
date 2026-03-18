package pairing

import (
	"slices"
	"github.com/fanik05/tcg-pocket-api/internal/models"
)

func SortPlayersBySwiss(players []models.Player) {
	slices.SortFunc(players, func(a, b models.Player) int {
		if a.Points != b.Points {
			return b.Points - a.Points // Sort by points descending
		}
		if a.OMWP != b.OMWP {
			if a.OMWP > b.OMWP {
				return -1 // a has higher OMWP, so it should come first
			} else {
				return 1 // b has higher OMWP, so it should come first
			}
		}
		return 0 // They are equal in points and OMWP
	})
}	

func GenerateMatches(players []models.Player) []models.Match {
    var matches []models.Match
    // used tracks IDs of players already paired this round
    used := make(map[string]bool)

    for i := range players {
        p1 := players[i]
        
        // Skip if this player was already paired as someone's opponent
        if used[p1.ID] {
            continue
        }

        foundOpponent := false
        // Look for the next best available opponent (j starts after i)
        for j := i + 1; j < len(players); j++ {
            p2 := players[j]
            
            // Skip if p2 is already paired or if they have played p1 before
            if used[p2.ID] || hasPlayed(p1, p2.ID) {
                continue
            }

            // Valid Match Found!
            matches = append(matches, models.Match{
                Player1ID: p1.ID,
                Player2ID: p2.ID,
            })
            
            used[p1.ID] = true
            used[p2.ID] = true
            foundOpponent = true
            break
        }

        // Handle the "Bye" if no opponent was found
        if !foundOpponent {
            matches = append(matches, models.Match{
                Player1ID: p1.ID,
                Player2ID: "BYE",
            })
            used[p1.ID] = true
        }
    }

    return matches
}

// Helper to check the slice of previous IDs
func hasPlayed(p models.Player, opponentID string) bool {
    return slices.Contains(p.Opponents, opponentID)
}