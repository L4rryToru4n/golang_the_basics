package CoreFunctions

const MAX = 5

type Personels map[string]string

var personel_data = [...]Personels{
	{
		"name": "Lee 'Apollo' Adama",
		"rank": "Liutenant",
		"role": "Viper Pilot",
	},
	{
		"name": "Kara 'Starbuck' Thrace",
		"rank": "Liutenant",
		"role": "Viper Pilot",
	},
	{
		"name": "Jay 'Shark' Finnegan",
		"rank": "Jr. Liutenant",
		"role": "Raptor Pilot",
	},
	{
		"name": "Samuel 'Longshot' Anders",
		"rank": "Ensign",
		"role": "Viper Pilot",
	},
}

func GetPersonels() [MAX]string {

	var newData [MAX]string

	for i := 0; i < len(personel_data); i++ {
		var temp = personel_data[i]["role"] + " " + personel_data[i]["rank"] + " " + personel_data[i]["name"]
		newData[i] = temp
	}
	return newData
}
