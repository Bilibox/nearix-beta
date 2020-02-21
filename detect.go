// This file will receive a Pokemon Image and will return the appropriate pokemon.
// This will get called from main.go

package main

import (
	"net/http"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"io"
)
// BEGIN structs
type Pokemons struct{
	Pokemon map[string]string
}
// ENDOF structs
// BEGIN function definition
func receive(url string) string{
	var pokemons Pokemons
	err := Download("images/template.jpg",url)
	logErr(err)
	hash := Hash("images/template.jpg")
	hashHex := hex.EncodeToString(hash)
	readPokemonList(&pokemons)
	pokemonName := Compare(hashHex, &pokemons)
	return(pokemonName)
}
func Download(path string, url string) error{
	response, err := http.Get(url)
	logErr(err)
	var output *os.File

	if _, err := os.Stat(path); os.IsNotExist(err){
		logErr(err)
		output, err = os.Create(path)
		logErr(err)
	}else{
		err = os.Remove(path)
		logErr(err)
		output, err = os.Create(path)
		logErr(err)
	}
	// copy contents of response to output
	_, err = io.Copy(output, response.Body)
	defer response.Body.Close()
	defer output.Close()
	return err
}
func readPokemonList(pokemonStruct *Pokemons){
	reader, err := ioutil.ReadFile("config/hashes.yaml")
	logErr(err)
	yaml.Unmarshal(reader, pokemonStruct)
}
func Compare(hash string, pokemonStruct *Pokemons) string{
	return pokemonStruct.Pokemon[hash]
}
func Hash(path string) []byte{
	output, err := ioutil.ReadFile(path)
	logErr(err)
	hash := md5.New()
	hash.Write(output)
	hashSum := hash.Sum(nil)
	hexEncoded := make([]byte, hex.EncodedLen(len(hashSum)))
	hex.Encode(hexEncoded, hashSum)


	return hexEncoded
}
// ENDOF function definition
