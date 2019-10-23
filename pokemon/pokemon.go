package main

import "fmt"

type Pokemon struct {
	name        string
	pokemonType string
}

type Trainer struct {
	name      string
	pokeballs []Pokemon
}

type PokemonTrainer interface {
	Add(pokemon Pokemon) error
	//ChooseYou(pokemonName string) (Pokemon, error)
	//Free(pokemonName string) error
}

type PokemonLimitError struct{}

func (pokemonLimitError *PokemonLimitError) Error() string {
	return "Pokemon limit reached"
}

type PokemonNotFoundError struct{}

func (pokemonNotFound *PokemonNotFoundError) Error() string {
	return "Pokemon not found"
}

func (t *Trainer) Add(pokemon Pokemon) error {
	if len(t.pokeballs) >= 4 {
		return &PokemonLimitError{}
	}

	t.pokeballs = append(t.pokeballs, pokemon)

	return nil
}

func (t *Trainer) ChooseYou(pokemon string) (Pokemon, error) {

	for _, p := range t.pokeballs {
		if p.name == pokemon {
			return p, nil
		}
	}

	return Pokemon{}, &PokemonNotFoundError{}
}

func (t *Trainer) Free(pokemon string) error {
	for i, p := range t.pokeballs {
		if p.name == pokemon {
			t.pokeballs[i] = t.pokeballs[len(t.pokeballs)-1]
			t.pokeballs[len(t.pokeballs)-1] = Pokemon{}
			t.pokeballs = t.pokeballs[:len(t.pokeballs)-1]

			return nil
		}
	}

	return &PokemonNotFoundError{}
}

func main() {
	trainer := Trainer{}
	trainer.name = "Ash"

	// Initial pokemon
	fmt.Println(trainer.Add(Pokemon{"Squirtle", "water"}))

	// extra 1
	fmt.Println(trainer.Add(Pokemon{"Butterfree", "insect/psychic"}))

	// extra 2
	fmt.Println(trainer.Add(Pokemon{"Onix", "rock"}))

	// extra 3
	fmt.Println(trainer.Add(Pokemon{"Gengar", "ghost"}))

	// limit
	fmt.Println(trainer.Add(Pokemon{"Arceus", "wathever"}))

	// not found
	fmt.Println(trainer.ChooseYou("Pikachu"))

	// found
	fmt.Println(trainer.ChooseYou("Squirtle"))

	fmt.Println(trainer.pokeballs)

	fmt.Println(trainer.Free("Squirtle"))

	fmt.Println(trainer.pokeballs)

	fmt.Println(trainer.Free("Pikachu"))
}
