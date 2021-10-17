package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"

	t "github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := t.Init()
	core.SetBrand("Appeals")
	core.SetGamemode(t.ADVENTURE)

	loadConfig(core)

	if spawn != nil {
		log.Println("Using schematic world")
		if config.Spawn.Location != nil {
			spawn.Spawn.X = config.Spawn.Location.X - 8.5
			spawn.Spawn.Y = config.Spawn.Location.X - 65.5
			spawn.Spawn.Z = config.Spawn.Location.X - 8.5
			spawn.Dimension = t.OVERWORLD
		}
		core.SetMap(spawn)
	}

	core.On(func(e *t.PlayerJoinEvent) {
		m := t.ChatMessage("Start creating an appeal with /appeal")
		m.SetColor(&t.ChatColorGold)
		e.Player.SendMessage(m)
	})

	core.DeclareCommand(t.CommandNodeLiteral("appeal",
		nil,
		func(player *t.Player, args []string) {
			uuid := createAppeal(player.GetName())
			if uuid == "exists" {
				uuid = getAppeal(player.GetName())
				m := t.ChatMessage("You already have a pending appeal! Go to http://appeal.cyclize.cf/" + uuid + " to check it out.")
				m.SetColor(&t.ChatColorRed)
				m.SetClickEvent(
					&t.ChatAction{Action: "open_url", Value: "http://appeal.cyclize.cf/" + uuid},
				)
				player.SendMessage(m)
			} else {
				m := t.ChatMessage("Sucessfully created a pending appeal! Go to http://appeal.cyclize.cf/" + uuid + " to finish creating it.")
				m.SetColor(&t.ChatColorGreen)
				m.SetClickEvent(
					&t.ChatAction{Action: "open_url", Value: "http://appeal.cyclize.cf/" + uuid},
				)
				player.SendMessage(m)
			}
		},
	))

	core.Start()
}

type ChunkSave struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Bitmask int    `json:"bitmask"`
	Data    string `json:"data"`
}

func createAppeal(name string) string {
	type Payload struct {
		Name string `json:"name"`
	}

	data := Payload{
		name,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Data could not be translated to JSON! %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "SUPABASE_API_URL/appeals", body)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	errCode := gjson.Get(bodyString, "code").String()
	if errCode == "23505" {
		return "exists"
	}
	return gjson.Get(bodyString, "0.uuid").String()
}

func getAppeal(name string) string {
	req, err := http.NewRequest("GET", "SUPABASE_API_URL/appeals?name=eq."+name, nil)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	uuid := gjson.Get(bodyString, "0.uuid").String()

	return uuid
}
