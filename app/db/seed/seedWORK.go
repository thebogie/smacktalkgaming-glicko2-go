// seed project seed.go
package seed

import (
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
	//"log"
	"mitchgottlieb.com/smacktalkgamingtalk/objects/game"
	"mitchgottlieb.com/smacktalkgamingtalk/objects/player"
	//"bufio"
	"io/ioutil"
	//"os"
	//"reflect"
)

type Seed struct {
	Games   []game.Game
	Players []player.Player
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ProcessSeed(f string, db *neoism.Database) {

	/*
		file, err := os.Open(f) // For read access.

		if err != nil {
			panic(err)
		}

		defer file.Close()
	*/

	/* READ FILE
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines)
	*/

	//g := game.Game{}
	//dec := json.NewDecoder(file)

	var jsoncollection Seed

	//var v map[string]interface{}
	//if err := dec.Decode(&v); err != nil {
	contents, err := ioutil.ReadFile(f)
	check(err)

	json.Unmarshal(contents, &jsoncollection)
	//fmt.Println("%v", v)
	//fmt.Println(reflect.TypeOf(v))

	for _, data := range jsoncollection.Games {
		//fmt.Println(k)
		//fmt.Println(data)
		//fmt.Println(reflect.TypeOf(data))
		game.Create(data, db)
	}

	/*
		for k, data := range v {
			fmt.Println(k)
			fmt.Println(v[k][1])
			fmt.Println(data)
			fmt.Println(reflect.TypeOf(data))

			for i := 0; i < cap(data); i++ {
				fmt.Printf("(i=%d) %v\n", i, data[i])
			}

			if k == "games" {

					for x, data2 := range v[k] {
						fmt.Println(x)
						fmt.Println(data2)

					}

			}

		}
	*/

	fmt.Println("HERE")
}
