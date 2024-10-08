package CoreFunctions

const MAX = 10
const MAX_PAGE = 5

type Personel struct {
	ID      int
	Name    string
	Alias   string
	Rank    string
	Role    string
	Faction string
}

var PilotLists = [MAX]Personel{
	{
		ID:      1,
		Name:    "Lee Adama",
		Alias:   "Apollo",
		Rank:    "Liutenant",
		Role:    "Viper Pilot",
		Faction: "Human",
	},
	{
		ID:      2,
		Name:    "Kara Thrace",
		Alias:   "Starbuck",
		Rank:    "Liutenant",
		Role:    "Viper Pilot",
		Faction: "Human",
	},
	{
		ID:      3,
		Name:    "Jay Finnegan",
		Alias:   "Shark",
		Rank:    "Jr. Liutenant",
		Role:    "Raptor Pilot",
		Faction: "Human",
	},
	{
		ID:      4,
		Name:    "Samuel Anders",
		Alias:   "Longshot",
		Rank:    "Ensign",
		Role:    "Viper Pilot",
		Faction: "Human",
	},
}

var RaiderLists = [MAX]Personel{
	{
		ID:      1,
		Name:    "Cylon Raider #1",
		Alias:   "Raider One",
		Rank:    "Ensign",
		Role:    "Raider Pilot",
		Faction: "Cylon",
	},
	{
		ID:      2,
		Name:    "Cylon Raider #2",
		Alias:   "Raider Two",
		Rank:    "Sergeant",
		Role:    "Raider Pilot",
		Faction: "Cylon",
	},
	{
		ID:      3,
		Name:    "Cylon Raider #3",
		Alias:   "Raider Three",
		Rank:    "Liutenant",
		Role:    "Raider Pilot",
		Faction: "Cylon",
	},
}

func GetPilots(page_num int) *[MAX_PAGE]Personel {

	defer Recovery()

	pilots := new([MAX_PAGE]Personel)

	available_pages := MAX / MAX_PAGE

	if page_num == 0 {
		panic("Invalid page number.")
	} else if page_num <= available_pages {
		// page_num 1
		// MAX 15
		// MAX_PAGE 5
		// page_highest = 5 * 1 = 5
		// page_lowest = page_highest - MAX_PAGE = 5 - 5 = 0

		page_highest := MAX_PAGE * page_num
		page_lowest := page_highest - MAX_PAGE

		for i := page_lowest; i < page_highest; i++ {

			if PilotLists[i].ID != 0 {
				pilots[i].ID = PilotLists[i].ID
				pilots[i].Name = PilotLists[i].Name
				pilots[i].Alias = PilotLists[i].Alias
				pilots[i].Rank = PilotLists[i].Rank
				pilots[i].Role = PilotLists[i].Role
			}
		}
	}
	return pilots
}

func GetRaiders(page_num int) *[MAX_PAGE]Personel {

	defer Recovery()

	pilots := new([MAX_PAGE]Personel)

	available_pages := MAX / MAX_PAGE
	
	if page_num == 0 {
		panic("Invalid page number.")
	} else if page_num <= available_pages {
		// page_num 1
		// MAX 15
		// MAX_PAGE 5
		// page_highest = 5 * 1 = 5
		// page_lowest = page_highest - MAX_PAGE = 5 - 5 = 0

		page_highest := MAX_PAGE * page_num
		page_lowest := page_highest - MAX_PAGE

		for i := page_lowest; i < page_highest; i++ {

			if RaiderLists[i].ID != 0 {
				pilots[i].ID = RaiderLists[i].ID
				pilots[i].Name = RaiderLists[i].Name
				pilots[i].Alias = RaiderLists[i].Alias
				pilots[i].Rank = RaiderLists[i].Rank
				pilots[i].Role = RaiderLists[i].Role
			}
		}
	}
	return pilots
}

func GetPilot(id int) *Personel {

	defer InvalidRecovery()

	pilot := new(Personel)

	if id == 0 {
		panic("Invalid ID number, please input the correct number.")
	} else {
		defer NotFoundRecovery()

		for i := 0; i < len(PilotLists); i++ {
			if PilotLists[i].ID == id {
				*pilot = PilotLists[i]
			}
		}

		if pilot.ID == 0 {
			panic("ID number not found.")
		}
		return pilot
	}
}

func AddPilot(personel *Personel) *Personel {

	pilot := new(Personel)

	for i := 0; i < len(PilotLists); i++ {
		if PilotLists[i].ID == 0 {
			PilotLists[i].ID = personel.ID
			PilotLists[i].Name = personel.Name
			PilotLists[i].Alias = personel.Alias
			PilotLists[i].Rank = personel.Rank
			PilotLists[i].Role = personel.Role

			pilot = personel
		}
	}
	return pilot
}

func RemovePilot(id int) string {

	defer InvalidRecovery()

	found := false

	if id == 0 {
		panic("Invalid ID number.")
	} else {
		for i := 0; i < len(PilotLists); i++ {
			if PilotLists[i].ID == id {
				PilotLists[i].ID = 0
				PilotLists[i].Name = ""
				PilotLists[i].Alias = ""
				PilotLists[i].Rank = ""
				PilotLists[i].Role = ""

				found = true
			}
		}
	}

	defer NotFoundRecovery()

	if found {
		return "Pilot removal success."
	} else {
		panic("ID number was not found.")
	}
}
