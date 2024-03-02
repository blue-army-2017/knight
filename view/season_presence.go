package view

import (
	"fmt"
	"io"
	"sort"

	"github.com/blue-army-2017/knight/model"
)

type SeasonPresence struct {
	Member    *model.Member
	Position  int8
	HomeGames int8
	AwayGames int8
}

func (p *SeasonPresence) TotalGames() int8 {
	return p.HomeGames + p.AwayGames
}

type SeasonPresencePage struct {
	Season  *model.Season
	Games   []model.SeasonGame
	Members []model.Member
	Flash   *Flash
}

func (p *SeasonPresencePage) Render(w io.Writer) {
	page := pages["season_presence"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}

func (p *SeasonPresencePage) Statistic() []SeasonPresence {
	presence := make(map[string]int8)
	for _, game := range p.Games {
		for _, member := range game.PresentMembers {
			if game.Home {
				presence[fmt.Sprintf("%s_home", member.ID)] += 1
			} else {
				presence[fmt.Sprintf("%s_away", member.ID)] += 1
			}
		}
	}

	var statistic []SeasonPresence
	for _, member := range p.Members {
		memberStat := SeasonPresence{
			Member:    &member,
			HomeGames: presence[fmt.Sprintf("%s_home", member.ID)],
			AwayGames: presence[fmt.Sprintf("%s_away", member.ID)],
		}
		statistic = append(statistic, memberStat)
	}

	sort.SliceStable(statistic, func(i, j int) bool {
		return statistic[i].TotalGames() > statistic[j].TotalGames()
	})

	for i := 0; i < len(statistic); i++ {
		statistic[i].Position = int8(i) + 1

		if i > 0 {
			prev := statistic[i-1]
			if prev.TotalGames() == statistic[i].TotalGames() {
				statistic[i].Position = prev.Position
			}
		}
	}

	return statistic
}
