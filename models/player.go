package models

type Player struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// SQL:
// CREATE TABLE players(id integer not null primary key, name varchar(255) not null);

func GetPlayers() ([]*Player, error) {
	players := []*Player{}
	rows, err := db.Query("SELECT id, name FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		players = append(players, &Player{Id: id, Name: name})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return players, nil
}

func GetPlayer(id int) (*Player, error) {
	var rowId int
	var name string
	err := db.QueryRow("SELECT id, name FROM players WHERE id = ? LIMIT 1", id).Scan(&rowId, &name)
	if err != nil {
		return nil, err
	}
	return &Player{Id: rowId, Name: name}, nil
}

func (p *Player) Save() error {
	if p.Id == 0 {
		res, err := db.Exec("INSERT INTO players (name) VALUES (?)", p.Name)
		if err != nil {
			return err
		}
		i64, err := res.LastInsertId()
		if err != nil {
			return err
		}
		p.Id = int(i64)
	} else {
		_, err := db.Exec("UPDATE players SET name = ? WHERE id = ?", p.Name, p.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
