package models

type MatchPlayer struct {
	Id             int
	MatchId        int
	PlayerId       int
	CivilizationId int
	RedScore       int // Military
	GoldScore      int // Treasury (Money)
	WondersScore   int // Wonder
	BlueScore      int // Civilian
	OrangeScore    int // Commercial
	PurpleScore    int // Guilds
	GreenScore     int // Scientific
}

// SQL:
// CREATE TABLE match_players(id integer not null primary key, match_id integer not null, player_id integer not null, civilization_id integer not null, red_score integer not null, gold_score integer not null, wonders_score integer not null, blue_score integer not null, orange_score integer not null, purple_score integer not null, green_score integer not null);

func GetMatchPlayersForMatch(matchId int) ([]*MatchPlayer, error) {
	mplayers := []*MatchPlayer{}
	rows, err := db.Query("SELECT id, match_id, player_id, civilization_id, red_score, gold_score, wonders_score, blue_score, orange_score, purple_score, green_score FROM match_players WHERE match_id = ?", matchId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var matchId int
		var playerId int
		var civilizationId int
		var redScore int
		var goldScore int
		var wondersScore int
		var blueScore int
		var orangeScore int
		var purpleScore int
		var greenScore int
		err = rows.Scan(&id, &matchId, &playerId, &civilizationId, &redScore, &goldScore, &wondersScore, &blueScore, &orangeScore, &purpleScore, &greenScore)
		if err != nil {
			return nil, err
		}
		mplayers = append(mplayers, &MatchPlayer{
			Id:             id,
			MatchId:        matchId,
			PlayerId:       playerId,
			CivilizationId: civilizationId,
			RedScore:       redScore,
			GoldScore:      goldScore,
			WondersScore:   wondersScore,
			BlueScore:      blueScore,
			OrangeScore:    orangeScore,
			PurpleScore:    purpleScore,
			GreenScore:     greenScore,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return mplayers, nil
}

func (mp *MatchPlayer) Player() (*Player, error) {
	return GetPlayer(mp.PlayerId)
}

func (mp *MatchPlayer) Civilization() *Civilization {
	return GetCivilization(mp.CivilizationId)
}

func (mp *MatchPlayer) TotalScore() int {
	return mp.RedScore + mp.GoldScore + mp.WondersScore + mp.BlueScore + mp.OrangeScore + mp.PurpleScore + mp.GreenScore
}
