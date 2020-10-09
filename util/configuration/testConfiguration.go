package configuration

func main() {

	planets := Config{}
	var planetsConfigKeyStructure []configKeyDefinition
	planetsConfigKeyStructure = append(planetsConfigKeyStructure, configKeyDefinition{"planetOrder", 10})
	planetsConfigKeyStructure = append(planetsConfigKeyStructure, configKeyDefinition{"containsHumans", 1})
	planets.new("planetsConfig", "planetId", 6, planetsConfigKeyStructure)
	planets.AddEntry("EARTH")
}
