package main

import "fmt"

func mapde(users map[string]map[string]string, name string) {
	if users[name] != nil {

		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "小白龙"

	}

}

func main() {
	users := make(map[string]map[string]string, 3)
	users["sith"] = make(map[string]string, 3)
	users["sith"]["pwd"] = "666666"
	users["sith"]["nickname"] = "小花猫"

	mapde(users, "tom")
	mapde(users, "sith")
	fmt.Println(users)
}
