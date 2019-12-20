package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

type Results struct {
	Data []Clave `json:"results"`
}
type Clave struct {
	Name                   string `json:"name"`
	Model                  string `json:""model""`
	Manufacturer           string `json:"manufacturer"`
	Cost_in_credits        string `json:"cost_in_credits"`
	Length                 string `json:"length"`
	Max_atmosphering_speed string `json:"max_atmosphering_speed"`
	Crew                   string `json:"crew"`
	Passengers             string `json:"passengers"`
	Cargo_capacity         string `json:"cargo_capacity"`
	Consumables            string `json:"consumables"`
	Hyperdrive_rating      string `json:"hyperdrive_rating"`
	Mglt                   string `json:"MGLT"`
	Starship_class         string `json:"starship_class"`
	Pilots                 string `json:"pilots"`
	Films                  string `json:"films"`
	Created                string `json:"created"`
	Edited                 string `json:"edited"`
	Url                    string `json:"url"`
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	response, err := http.Get("https://swapi.co/api/starships/?page=1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Results
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject)

	//set con el nombre
	for i := 0; i < len(responseObject.Data); i++ {
		//guardar nombre
		b, _ := json.Marshal(responseObject.Data[i])
		err = client.Set("name", string(b), 0).Err()
		if err != nil {
			fmt.Println(err)
		}

		//get con el nombre
		val, err := client.Get("name").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)

	}

}
