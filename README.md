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
