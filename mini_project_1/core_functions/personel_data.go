package CoreFunctions

const MAX = 5

type Personels map[string]string

var personel_data = [...]Personels{
	{
		"name":  "Lee Adama",
		"alias": "Apollo",
		"rank":  "Liutenant",
		"role":  "Viper Pilot",
	},
	{
		"name":  "Kara Thrace",
		"alias": "Starbuck",
		"rank":  "Liutenant",
		"role":  "Viper Pilot",
	},
	{
		"name":  "Jay Finnegan",
		"alias": "Shark",
		"rank":  "Jr. Liutenant",
		"role":  "Raptor Pilot",
	},
	{
		"name":  "Samuel Anders",
		"alias": "Longshot",
		"rank":  "Ensign",
		"role":  "Viper Pilot",
	},
}

func GetPersonels() [MAX]Personels {

	var newData [MAX]Personels

	for i := 0; i < len(personel_data); i++ {
		var temp = make(map[string]string)
		temp["role"] = personel_data[i]["role"]
		temp["rank"] = personel_data[i]["rank"]
		temp["alias"] = personel_data[i]["alias"]
		temp["name"] = personel_data[i]["name"]

		newData[i] = temp
	}
	return newData
}
