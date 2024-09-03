package model

type PresenceRepository interface {
	GetSeasonPresence() ([]SeasonPresence, error)
	GetMemberPresence() ([]MemberPresence, error)
	SavePresentMembers(gameId string, presentMemberIds []string) error
}

type DefaultPresenceRepository struct {
}

func NewPresenceRepository() PresenceRepository {
	return &DefaultPresenceRepository{}
}

func (r *DefaultPresenceRepository) GetSeasonPresence() ([]SeasonPresence, error) {
	statement := `
SELECT s.name,
       SUM(g.home) AS home_games,
       COUNT(g.id) AS total_games
FROM seasons s
INNER JOIN season_games g ON g.season_id = s.id
GROUP BY s.id
ORDER BY s.created DESC`

	var presence []SeasonPresence
	result := db.Raw(statement).Scan(&presence)
	return presence, result.Error
}

func (r *DefaultPresenceRepository) GetMemberPresence() ([]MemberPresence, error) {
	statement := `
SELECT s.name AS season,
       m.last_name,
       m.first_name,
       SUM(g.home) AS home_games,
       COUNT(g.id) AS total_games
FROM present_members p
INNER JOIN members m ON m.id = p.member_id
INNER JOIN season_games g ON g.id = p.season_game_id
INNER JOIN seasons s ON s.id = g.season_id
WHERE m.active = TRUE
GROUP BY s.id,
         m.id
ORDER BY s.created DESC,
         total_games DESC,
         m.last_name,
         m.first_name`

	var presence []MemberPresence
	result := db.Raw(statement).Scan(&presence)
	return presence, result.Error
}

func (r *DefaultPresenceRepository) SavePresentMembers(gameId string, presentMemberIds []string) error {
	deleteStatement := "DELETE FROM present_members WHERE season_game_id = ?"
	result := db.Exec(deleteStatement, gameId)
	if result.Error != nil {
		return result.Error
	}

	saveStatement := "INSERT INTO present_members (season_game_id, member_id) VALUES (?, ?)"
	for _, memberId := range presentMemberIds {
		result = db.Exec(saveStatement, gameId, memberId)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
