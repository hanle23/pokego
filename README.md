# PokeGO

API wrapper for [Poke API](https://pokeapi.co/), written in Go. *Supports PokeAPI v2.* Inspired by [PokeAPI-GO](https://github.com/mtslzr/pokeapi-go)

- [Documentation](#Documentation)
- [Getting Started](#Getting-Started)


## Documentation

All API endpoints is following the official [PokeAPI documentation](https://pokeapi.co/docs/v2) as of October 29, 2024, including all official type.

## Getting Started

- Importing the package
```bash
go get github.com/hanle23/pokego
```

- Example
```go
client := pokego.NewClient()

pokemon, err := client.Pokemon("pikachu")
if err != nil {
	log.Fatal(err)
}

fmt.Println(pokemon.Name)

```

- Initialize a client is required to use the API
```go
client := pokego.NewClient()
```

- Client is following singleton pattern, so you can use the same client for all API endpoints
```go
pokemon, err := client.Pokemon("pikachu")
if err != nil {
	log.Fatal(err)
}
```

- Client can have custom configuration (except for base URL)
```go
client := pokego.NewClient(
	WithUseCache(false),
	WithExpireTime(5 * time.Minute),
)
```
## Endpoints

### Berries

<details>
  <summary>Berries</summary>

  #### Berry

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Berry("cheri")
  ```

  #### Berries (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Berries("0", "20")
  ```

</details>

<details>
  <summary>Berry Firmnesses</summary>

  #### Berry Firmness

  *Must pass ID or Name (string)*

  ```go
  result, err := client.BerryFirmness("very-soft")
  ```

  #### Berry Firmness (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.BerryFirmnesses("0", "20")
  ```

</details>

<details>
  <summary>Berry Flavors</summary>

  #### Berry Flavor

  *Must pass ID or Name (string)*

  ```go
  result, err := client.BerryFlavor("spicy")
  ```

  #### Berry Flavors (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.BerryFlavors("0", "20")
  ```

</details>

### Contests

<details>
  <summary>Contest Types</summary>

  #### Contest Types

  *Must pass ID or Name (string)*

  ```go
  result, err := client.ContestType("cool")
  ```

  #### Contest Types (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.ContestTypes("0", "20")
  ```

</details>

<details>
  <summary>Contest Types</summary>

  #### Contest Effect

  *Must pass ID (string)*

  ```go
  result, err := client.ContestEffect("1")
  ```

  #### Contest Effects (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.ContestEffects("0", "20")
  ```

</details>

<details>
  <summary>Super Contest Effects</summary>

  #### Super Contest Effect

  *Must pass ID (string)*

  ```go
  result, err := client.SuperContestEffect("1")
  ```

  #### Super Contest Effects (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.SuperContestEffects("0", "20")
  ```

</details>

### Encounters

<details>
  <summary>Encounter Method</summary>

  #### Encounter Method

  *Must pass ID or Name (string)*

  ```go
  result, err := client.EncounterMethod("walk")
  ```

  #### Encounter Methods (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.EncounterMethods("0", "20")
  ```

</details>

<details>
  <summary>Encounter Conditions</summary>

  #### Encounter Condition

  *Must pass ID or Name (string)*

  ```go
  result, err := client.EncounterCondition("swarm")
  ```

  #### Encounter Conditions (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.EncounterConditions("0", "20")
  ```

</details>

<details>
  <summary>Encounter Condition Values</summary>

  #### Encounter Condition Value

  *Must pass ID or Name (string)*

  ```go
  result, err := client.EncounterConditionValue("swarm-yes")
  ```

  #### Encounter Condition Values (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.EncounterConditionValues("0", "20")
  ```

</details>


### Evolution

<details>
  <summary>Evolution Chains</summary>

  #### Evolution Chain

  *Must pass ID (string)*

  ```go
  result, err := client.EvolutionChain("1")
  ```

  #### Encounter Chains (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.EvolutionChains("0", "20")
  ```

</details>

<details>
  <summary>Evolution Triggers</summary>

  #### Evolution Trigger

  *Must pass ID or Name (string)*

  ```go
  result, err := client.EvolutionTrigger("level-up")
  ```

  #### Encounter Triggers (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.EvolutionTriggers("0", "20")
  ```

</details>

### Games

<details>
  <summary>Generation</summary>

  #### Generation

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Generation("generation-i")
  ```

  #### Generations (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Generations("0", "20")
  ```

</details>

<details>
  <summary>Pokedex</summary>

  #### Pokedex

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Pokedex("kanto")
  ```

  #### Pokedexes (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Pokedexes("0", "20")
  ```

</details>

<details>
  <summary>Version</summary>

  #### Version

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Version("red")
  ```

  #### Versions (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Versions("0", "20")
  ```

</details>

<details>
  <summary>Version Groups</summary>

  #### Version Group

  *Must pass ID or Name (string)*

  ```go
  result, err := client.VersionGroup("red-blue")
  ```

  #### Version Groups (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.VersionGroups("0", "20")
  ```

</details>

### Items

<details>
  <summary>Items</summary>

  #### Item

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Item("master-ball")
  ```

  #### Items (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Items("0", "20")
  ```

</details>

<details>
  <summary>Item Attributes</summary>

  #### Item Attribute

  *Must pass ID or Name (string)*

  ```go
  result, err := client.ItemAttribute("countable")
  ```

  #### Item Attributes (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.ItemAttributes("0", "20")
  ```

</details>

<details>
  <summary>Item Categories</summary>

  #### Item Category

  *Must pass ID or Name (string)*

  ```go
  result, err := client.ItemCategory("stat-boosts")
  ```

  #### Item Categories (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.ItemCategories("0", "20")
  ```

</details>

<details>
  <summary>Item Fling Effects</summary>

  #### Item Fling Effect

  *Must pass ID or Name (string)*

  ```go
  result, err := client.ItemFlingEffect("badly-poison")
  ```

  #### Item Fling Effects (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.ItemFlingEffects("0", "20")
  ```

</details>

<details>
  <summary>Item Pockets</summary>

  #### Item Pocket

  *Must pass ID or Name (string)*

  ```go
  result, err := client.ItemPocket("misc")
  ```

  #### Item Pockets (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.ItemPockets("0", "20")
  ```

</details>

### Locations

<details>
  <summary>Locations</summary>

  #### Location

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Location("canalave-city")
  ```

  #### Items (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Locations("0", "20")
  ```

</details>

<details>
  <summary>Location Areas</summary>

  #### Location Area

  *Must pass ID or Name (string)*

  ```go
  result, err := client.LocationArea("canalave-city-area")
  ```

  #### Location Areas (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.LocationArea("0", "20")
  ```

</details>

<details>
  <summary>Pal Park Areas</summary>

  #### Pal Park Area

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PalParkArea("forest")
  ```

  #### Pal Park Areas (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.LocationArea("0", "20")
  ```

</details>

<details>
  <summary>Regions</summary>

  #### Region

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Region("kanto")
  ```

  #### Regions (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Regions("0", "20")
  ```

</details>

### Machines

<details>
  <summary>Machines</summary>

  #### Machine

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Machine("pound")
  ```

  #### Machines (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Machines("0", "20")
  ```

</details>

### Moves

<details>
  <summary>Moves</summary>

  #### Move

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Move("pound")
  ```

  #### Moves (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Moves("0", "20")
  ```

</details>

<details>
  <summary>Move Ailment</summary>

  #### Move Ailment

  *Must pass ID or Name (string)*

  ```go
  result, err := client.MoveAilment("paralysis")
  ```

  #### Move Ailments (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.MoveAilments("0", "20")
  ```

</details>

<details>
  <summary>Move Battle Styles</summary>

  #### Move Battle Style

  *Must pass ID or Name (string)*

  ```go
  result, err := client.MoveBattleStyle("attack")
  ```

  #### Move Battle Styles (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.MoveBattleStyles("0", "20")
  ```

</details>

<details>
  <summary>Move Category</summary>

  #### Move Category

  *Must pass ID or Name (string)*

  ```go
  result, err := client.MoveCategory("ailment")
  ```

  #### Move Categories (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.MoveCategories("0", "20")
  ```

</details>

<details>
  <summary>Move Damage Class</summary>

  #### Move Damage Class

  *Must pass ID or Name (string)*

  ```go
  result, err := client.MoveDamageClass("status")
  ```

  #### Move Damage Classes (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.MoveDamageClasses("0", "20")
  ```

</details>

<details>
  <summary>Move Learn Method</summary>

  #### Move Learn Method

  *Must pass ID or Name (string)*

  ```go
  result, err := client.MoveLearnMethod("level-up")
  ```

  #### Move Learn Methods (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.MoveLearnMethods("0", "20")
  ```

</details>

<details>
  <summary>Move Targets</summary>

  #### Move Target

  *Must pass ID or Name (string)*

  ```go
  result, err := client.MoveTarget("specific-move")
  ```

  #### Move Targets (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.MoveTargets("0", "20")
  ```

</details>

### Pokemon

<details>
  <summary>Abilities</summary>

  #### Ability

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Ability("stench")
  ```

  #### Abilities (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Abilities("0", "20")
  ```

</details>

<details>
  <summary>Characteristics</summary>

  #### Characteristic

  *Must pass ID (string)*

  ```go
  result, err := client.Characteristic("1")
  ```

  #### Characteristics (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Characteristics("0", "20")
  ```

</details>

<details>
  <summary>Egg Groups</summary>

  #### Egg Group

  *Must pass ID or Name (string)*

  ```go
  result, err := client.EggGroup("monster")
  ```

  #### Egg Groups (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.EggGroups("0", "20")
  ```

</details>

<details>
  <summary>Genders</summary>

  #### Gender

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Gender("female")
  ```

  #### Genders (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Genders("0", "20")
  ```

</details>

<details>
  <summary>Growth Rates</summary>

  #### Growth Rate

  *Must pass ID or Name (string)*

  ```go
  result, err := client.GrowthRate("slow")
  ```

  #### Growth Rates (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.GrowthRates("0", "20")
  ```

</details>

<details>
  <summary>Natures</summary>

  #### Nature

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Nature("bold")
  ```

  #### Natures (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Natures("0", "20")
  ```

</details>

<details>
  <summary>Pokeathlon Stats</summary>

  #### Pokeathlon Stat

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokeathlonStat("speed")
  ```

  #### Pokeathlon Stats (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.PokeathlonStats("0", "20")
  ```

</details>

<details>
  <summary>Pokemons</summary>

  #### Pokemons

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Pokemon("clefairy")
  ```

  #### Pokemons (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Pokemons("0", "20")
  ```

</details>

<details>
  <summary>Pokemon Location Areas</summary>

  #### Pokemon Location Area

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokemonLocationAreas("clefairy")
  ```

</details>

<details>
  <summary>Pokemon Colors</summary>

  #### Pokemon Color

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokemonColor("black")
  ```

  #### Pokemon Colors (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.PokemonColors("0", "20")
  ```

</details>

<details>
  <summary>Pokemon Forms</summary>

  #### Pokemon Form

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokemonForm("arceus-bug")
  ```

  #### Pokemon Forms (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.PokemonForms("0", "20")
  ```

</details>

<details>
  <summary>Pokemon Habitats</summary>

  #### Pokemon Habitat

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokemonHabitat("cave")
  ```

  #### Pokemon Habitats (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.PokemonHabitats("0", "20")
  ```

</details>


<details>
  <summary>Pokemon Shapes</summary>

  #### Pokemon Shape

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokemonShape("ball")
  ```

  #### Pokemon Shapes (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.PokemonShapes("0", "20")
  ```

</details>

<details>
  <summary>Pokemon Species</summary>

  #### Pokemon Species

  *Must pass ID or Name (string)*

  ```go
  result, err := client.PokemonSpecies("wormadam")
  ```

  #### Pokemon Species Pagination (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.PokemonSpeciesPagination("0", "20")
  ```

</details>

<details>
  <summary>Stats</summary>

  #### Stat

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Stat("attack")
  ```

  #### Stats (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Stats("0", "20")
  ```

</details>

<details>
  <summary>Types</summary>

  #### Type

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Type("ground")
  ```

  #### Types (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Stats("0", "20")
  ```

</details>

### Utility

<details>
  <summary>Languages</summary>

  #### Language

  *Must pass ID or Name (string)*

  ```go
  result, err := client.Language("ja")
  ```

  #### Languages (pagination)

  *Must pass offset and limit (string)*

  ```go
  result, err := client.Languages("0", "20")
  ```

</details>

## Additional Information

- Config options is currently unable to change on the fly, if you want to change config setting, make sure to removeClient first and do NewClient
```go
client := pokego.NewClient(
	WithUseCache(false),
	WithExpireTime(5 * time.Minute),
)
RemoveClient()
client = pokego.NewClient(
	WithUseCache(true),
	WithExpireTime(10 * time.Minute),
)
```

- If there are some issue with the client, you can also try to Reset the client (keep all currentl config) or remove and recreate

```go
client.ResetClient()
```

```go
client.RemoveClient()
client = pokego.NewClient()
```

- If there is a specific query that you want to make and not available in the package, you can use the `Fetch` method
```go
result, err := client.Fetch("/berry/1/")
```

- Or when you want to get all available endpoints
```go
result, err := client.Fetch("")
```
