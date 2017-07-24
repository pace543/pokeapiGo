package pokeapiGo

type APIResource struct {
	Url 	string
}

type NamedAPIResource struct {
	Name 	string
	Url 	string
}

type Description struct {
	Description	string
	Language 	NamedAPIResource
}

type Effect struct {
	Effect		string
	Language	NamedAPIResource
}

type Encounter struct {
	Min_Level 		int
	Max_Level 		int
	Condition_Values	[]NamedAPIResource
	Chance 			int
	Method 			NamedAPIResource
}

type FlavorText struct {
	Flavor_Text 	string
	Language	NamedAPIResource
}

type GenerationGameIndex struct {
	Game_Index	int
	Generation	NamedAPIResource
}

type MachineVersionDetail struct {
	Machine 	APIResource
	Version_Group   NamedAPIResource
}

type Name struct {
	Name 		string
	Language 	NamedAPIResource
}

type VerboseEffect struct {
	Effect 		string
	Short_Effect 	string
	Language	NamedAPIResource
}

type VersionEncounterDetail struct {
	Version 		NamedAPIResource
	Max_Chance 		int
	Encounter_Details 	[]Encounter
}

type VersionGameIndex struct {
	Game_Index 	int
	Version 	NamedAPIResource
}

type VersionGroupFlavorText struct {
	Text 		string
	Language 	NamedAPIResource
	Version_Group   NamedAPIResource
}