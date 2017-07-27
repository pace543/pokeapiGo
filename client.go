package pokeapiGo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Http       *http.Client
	NumWorkers int
	Jobs       chan *Job
	Results    chan *Result
}

type Job struct {
	Endpoint string
	Id       int
	Name     string
}

func (j *Job) getItem() string {
	if j.Id != 0 {
		return strconv.Itoa(j.Id)
	} else if j.Name != "" {
		return j.Name
	} else {
		return ""
	}
}

type Result struct {
	Endpoint string
	Data     interface{}
	Error    error
}

func NewClient(client *http.Client, numWorkers int) *Client {
	clientPtr := new(Client)
	if client == nil {
		clientPtr.Http = &http.Client{Timeout: 10 * time.Second}
	} else {
		clientPtr.Http = client
	}
	if numWorkers == 0 {
		clientPtr.NumWorkers = 3
	} else {
		clientPtr.NumWorkers = numWorkers
	}
	clientPtr.Jobs = make(chan *Job)
	clientPtr.Results = make(chan *Result)
	return clientPtr
}

func (c *Client) Init() {
	for i := 0; i < c.NumWorkers; i++ {
		go c.worker()
	}
}

func (c *Client) AddJobs(j ...*Job) {
	go func() {
		for job := range j {
			c.Jobs <- job
		}
	}()
}

func (c *Client) PullResult() *Result {
	return <-c.Results
}

func (c *Client) worker() {
	for job := range c.Jobs {
		fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/%v/%v", job.Endpoint, job.getItem())
		result, err := c.Http.Get(fullUrl)
		returned := new(Result)
		if err != nil {
			returned.Endpoint = job.Endpoint
			returned.Data = ""
			returned.Error = err
		} else {
			returned.Endpoint = job.Endpoint
			returned.Data = job.getStruct(result.Body)
			returned.Error = nil
		}
		result.Body.Close()
		c.Results <- returned
	}
}

func (j *Job) getStruct(result io.ReadCloser) interface{} {
	if j.getItem() == "" {
		apiResourceList := new(APIResourceList)
		err := json.NewDecoder(result).Decode(apiResourceList)
		if err != nil {
			namedApiResourceList := new(NamedAPIResourceList)
			json.NewDecoder(result).Decode(namedApiResourceList)
			return namedApiResourceList
		} else {
			return apiResourceList
		}
	}
	switch j.Endpoint {
	case "berry":
		berry := new(Berry)
		json.NewDecoder(result).Decode(berry)
		return berry
	case "berry-firmness":
		berryFirmness := new(BerryFirmness)
		json.NewDecoder(result).Decode(berryFirmness)
		return berryFirmness
	case "berry-flavor":
		berryFlavor := new(BerryFlavor)
		json.NewDecoder(result).Decode(berryFlavor)
		return berryFlavor
	case "contest-type":
		contestType := new(ContestType)
		json.NewDecoder(result).Decode(contestType)
		return contestType
	case "contest-effect":
		contestEffect := new(ContestEffect)
		json.NewDecoder(result).Decode(contestEffect)
		return contestEffect
	case "super-contest-effect":
		superContestEffect := new(SuperContestEffect)
		json.NewDecoder(result).Decode(superContestEffect)
		return superContestEffect
	case "encounter-method":
		encounterMethod := new(EncounterMethod)
		json.NewDecoder(result).Decode(encounterMethod)
		return encounterMethod
	case "encounter-condition":
		encounterCondition := new(EncounterCondition)
		json.NewDecoder(result).Decode(encounterCondition)
		return encounterCondition
	case "encounter-condition-value":
		encounterConditionValue := new(EncounterConditionValue)
		json.NewDecoder(result).Decode(encounterConditionValue)
		return encounterConditionValue
	case "evolution-chain":
		evolutionChain := new(EvolutionChain)
		json.NewDecoder(result).Decode(evolutionChain)
		return evolutionChain
	case "evolution-trigger":
		evolutionTrigger := new(EvolutionTrigger)
		json.NewDecoder(result).Decode(evolutionTrigger)
		return evolutionTrigger
	case "generation":
		generation := new(Generation)
		json.NewDecoder(result).Decode(generation)
		return generation
	case "pokedex":
		pokedex := new(Pokedex)
		json.NewDecoder(result).Decode(pokedex)
		return pokedex
	case "version":
		version := new(Version)
		json.NewDecoder(result).Decode(version)
		return version
	case "version-group":
		versionGroup := new(VersionGroup)
		json.NewDecoder(result).Decode(versionGroup)
		return versionGroup
	case "item":
		item := new(Item)
		json.NewDecoder(result).Decode(item)
		return item
	case "item-attribute":
		itemAttribute := new(ItemAttribute)
		json.NewDecoder(result).Decode(itemAttribute)
		return itemAttribute
	case "item-category":
		itemCategory := new(ItemCategory)
		json.NewDecoder(result).Decode(itemCategory)
		return itemCategory
	case "item-fling-effect":
		itemFlingEffect := new(ItemFlingEffect)
		json.NewDecoder(result).Decode(itemFlingEffect)
		return itemFlingEffect
	case "item-pocket":
		itemPocket := new(ItemPocket)
		json.NewDecoder(result).Decode(itemPocket)
		return itemPocket
	case "machine":
		machine := new(Machine)
		json.NewDecoder(result).Decode(machine)
		return machine
	case "move":
		move := new(Move)
		json.NewDecoder(result).Decode(move)
		return move
	case "move-ailment":
		moveAilment := new(MoveAilment)
		json.NewDecoder(result).Decode(moveAilment)
		return moveAilment
	case "move-battle-style":
		moveBattleStyle := new(MoveBattleStyle)
		json.NewDecoder(result).Decode(moveBattleStyle)
		return moveBattleStyle
	case "move-category":
		moveCategory := new(MoveCategory)
		json.NewDecoder(result).Decode(moveCategory)
		return moveCategory
	case "move-damage-class":
		moveDamageClass := new(MoveDamageClass)
		json.NewDecoder(result).Decode(moveDamageClass)
		return moveDamageClass
	case "move-learn-method":
		moveLearnMethod := new(MoveLearnMethod)
		json.NewDecoder(result).Decode(moveLearnMethod)
		return moveLearnMethod
	case "move-target":
		moveTarget := new(MoveTarget)
		json.NewDecoder(result).Decode(moveTarget)
		return moveTarget
	case "location":
		location := new(Location)
		json.NewDecoder(result).Decode(location)
		return location
	case "location-area":
		locationArea := new(LocationArea)
		json.NewDecoder(result).Decode(locationArea)
		return locationArea
	case "pal-park-area":
		palParkArea := new(PalParkArea)
		json.NewDecoder(result).Decode(palParkArea)
		return palParkArea
	case "region":
		region := new(Region)
		json.NewDecoder(result).Decode(region)
		return region
	}
	return nil
}
