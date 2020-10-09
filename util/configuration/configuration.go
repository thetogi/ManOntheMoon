package configuration

import "log"

type Config struct {
	name               string
	idName             string
	idLength           int32
	configEntries      []configEntry
	configKeyStructure []configKeyDefinition
	currentEntry       int
}

func (c *Config) ConfigName() string {
	return c.name
}

func (c *Config) Configuration() []configEntry {
	return c.configEntries
}

//Load should return an instance of Config
func (c *Config) Load(cfgName string) {

	//var config string = "";
	//return nil
}

func (c *Config) new(name string, idName string, idLength int32, keys []configKeyDefinition) {
	//TODO validate if file already exists

	//TODO add configKeyStructure to internal variable

	c.name = name
	c.idName = idName
	c.idLength = idLength

	// Use map to record duplicates as we find them.
	var encountered = map[string]bool{}
	var result []configKeyDefinition

	//remove duplicates
	for i := range keys {
		if encountered[keys[i].keyDefinitionName] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[keys[i].keyDefinitionName] = false
			// Append to result slice.
			result = append(result, keys[i])
		}
	}
	c.configKeyStructure = result
}

//TODO SAVE

//TODO REMOVE ENTRY

//TODO ADD KEY

//TODO ADD KEYS

//TODO REMOVE KEY

//UPDATE KEYVALUE

type configKeyDefinition struct {
	keyDefinitionName   string
	keyDefinitionLength int32
}

type configKeyValue map[string]keyValue

type keyValue struct {
	id     string
	length int32
	value  string
}

type configEntry struct {
	entryName      string
	entryKeyValues []configKeyValue
}

//TODO ADD New ENTRY
func (c *Config) AddEntry(idName string) {
	//TODO check length on incoming name and value
	if c.FindEntry(idName) {
		log.Fatalf("%v already exists and cannot be added.", idName)
	}
	//TODO Add configKeyStructure and their default values
	if c.configKeyStructure == nil {
		log.Fatalf("No configKeyStructure have been initialized. New() must be ran before configKeyStructure can be added.")
	}

	var tempKeys []configKeyValue
	for i := range c.configKeyStructure {
		tempKey := configKeyValue{}
		tempKey[c.configKeyStructure[i].keyDefinitionName] = keyValue{c.configKeyStructure[i].keyDefinitionName, c.configKeyStructure[i].keyDefinitionLength, ""}
		tempKeys = append(tempKeys, tempKey)
	}
	c.configEntries = append(c.configEntries, configEntry{idName, tempKeys})
}

func (c *Config) FindEntry(idName string) bool {
	for i := range c.configEntries {
		if idName == c.configEntries[i].entryName {
			c.currentEntry = i
			return true
		}
	}
	return false
}

//String value getter
func (c *Config) SetValue(value string) {

	c.configEntries[c.currentEntry].

}

func (c *Config) columnIndex(columnName string) bool {
	if c.configEntries != nil {
		for i, name := range c.configEntries[0].entryKeyValues[0] {
			if name.id == columnName

			if  name["test"] == "test"	{
			return true
			}
		}
		return false
	} else {
		log.Fatalf("No configKeyStructure have been initialized. New() must be ran before configKeyStructure can be added.")

	}
}
	//Int value getter

	//Decimal value getter
