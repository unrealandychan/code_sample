package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var BaseUrl = "https://pokeapi.co/api/v2/pokemon/"

type PokemonResponse struct{
	Name string `json:"name"`
}


func get_pokemon (number int ,c chan string , wg *sync.WaitGroup){
	url := BaseUrl + strconv.Itoa(number)
	response ,err := http.Get(url)
	if err !=nil{
		log.Fatalf("Something on http reqest went wrong: %s",err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Something on ReaderClose went wrong: %s",err)
		}
	}(response.Body)
	if response.StatusCode == http.StatusOK{
		bodyBytes , err := io.ReadAll(response.Body)
		if err !=nil{
			log.Fatalf("Something on reading response went wrong: %s",err)
		}
		pokemonResponse := PokemonResponse{}
		if err := json.Unmarshal(bodyBytes,&pokemonResponse);err!=nil{
			log.Fatalf("Something on json unmarshall went wrong: %s",err)
		}
		fmt.Println(number)
		c <- pokemonResponse.Name
		wg.Done()
	}
}



func main(){
	fmt.Println("Function Start")
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(151)
	ch := make(chan string,151)

	for i:=1 ; i <152;i++{
		go get_pokemon(i ,ch,&wg)
	}
	wg.Wait()
	close(ch)
	for name := range ch {
		fmt.Println(name)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Time taken: %s",elapsed)

}