package models

import "time"

type Match struct {
	Id   int
	Date time.Time
}

// SQL:
// CREATE TABLE matches(id integer not null primary key, date varchar(255) not null);

func GetMatches() ([]*Match, error) {
	matches := []*Match{}
	rows, err := db.Query("SELECT id, date FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var dateStr string
		err = rows.Scan(&id, &dateStr)
		if err != nil {
			return nil, err
		}
		date, err := time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			return nil, err
		}
		matches = append(matches, &Match{Id: id, Date: date})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetMatch(id int) (*Match, error) {
	var rowId int
	var dateStr string
	err := db.QueryRow("SELECT id, date FROM matches WHERE id = ? LIMIT 1", id).Scan(&rowId, &dateStr)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return nil, err
	}
	return &Match{Id: rowId, Date: date}, nil
}

func (m *Match) Save() error {
	if m.Id == 0 {
		res, err := db.Exec("INSERT INTO matches (date) VALUES (?)", m.Date)
		if err != nil {
			return err
		}
		i64, err := res.LastInsertId()
		if err != nil {
			return err
		}
		m.Id = int(i64)
	} else {
		_, err := db.Exec("UPDATE matches SET date = ? WHERE id = ?", m.Date, m.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Match) MatchPlayers() ([]*MatchPlayer, error) {
	return GetMatchPlayersForMatch(m.Id)
}
