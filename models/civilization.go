package models

type Civilization struct {
	ShortName string
	LongName  string
	Side      string
}

var civilizations map[int]*Civilization

func init() {
	civilizations = map[int]*Civilization{
		0:  &Civilization{ShortName: "N/A", LongName: "The Lost Availability of Civilization", Side: ""},
		1:  &Civilization{ShortName: "Alexandria", LongName: "The Lighthouse of Alexandria", Side: "A"},
		2:  &Civilization{ShortName: "Alexandria", LongName: "The Lighthouse of Alexandria", Side: "B"},
		3:  &Civilization{ShortName: "Babylon", LongName: "The Hanging Gardens of Babylon", Side: "A"},
		4:  &Civilization{ShortName: "Babylon", LongName: "The Hanging Gardens of Babylon", Side: "B"},
		5:  &Civilization{ShortName: "Ephesus", LongName: "The Temple of Artemis in Ephesus", Side: "A"},
		6:  &Civilization{ShortName: "Ephesus", LongName: "The Temple of Artemis in Ephesus", Side: "B"},
		7:  &Civilization{ShortName: "Giza", LongName: "The Pyramids of Giza", Side: "A"},
		8:  &Civilization{ShortName: "Giza", LongName: "The Pyramids of Giza", Side: "B"},
		9:  &Civilization{ShortName: "Halicarnassus", LongName: "The Mausoleum of Halicarnassus", Side: "A"},
		10: &Civilization{ShortName: "Halicarnassus", LongName: "The Mausoleum of Halicarnassus", Side: "B"},
		11: &Civilization{ShortName: "Olympia", LongName: "The Statue of Zeus in Olympia", Side: "A"},
		12: &Civilization{ShortName: "Olympia", LongName: "The Statue of Zeus in Olympia", Side: "B"},
		13: &Civilization{ShortName: "Rhodes", LongName: "The Colossus of Rhodes", Side: "A"},
		14: &Civilization{ShortName: "Rhodes", LongName: "The Colossus of Rhodes", Side: "B"},
	}
}

func GetAllCivilizations() []*Civilization {
	allCivilizations := make([]*Civilization, len(civilizations))
	for id, civilization := range civilizations {
		allCivilizations[id-1] = civilization
	}
	return allCivilizations
}

func GetCivilization(id int) *Civilization {
	return civilizations[id]
}
