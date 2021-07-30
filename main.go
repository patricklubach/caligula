package main

import (
	"caligula/cmd"
)

// type Health struct {
// 	Status string `json:"status"`
// }

// func getJson(url string, target interface{}) error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 {
// 		os.Exit(1)
// 	}
// 	return json.NewDecoder(resp.Body).Decode(target)
// }

func main() {
	cmd.Execute()
	// data := new(Health)
	// getJson("http://localhost:8080/admin/health", data)
	// if data.Status != "UP" {
	// 	os.Exit(1)
	// }
}
