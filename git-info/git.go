package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	name, repos, err := gitinfo("https://api.github.com/users/YashJadhav1242")
	fmt.Println(name, repos, err)

}

func gitinfo(url string) (string, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return "", 0, err
	}

	var i Info

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&i); err != nil {
		return "", 0, err

	}

	return i.Name, i.Public_Repos, nil

}

type Info struct {
	Name         string
	Public_Repos int
}
