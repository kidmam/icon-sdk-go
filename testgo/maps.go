package main

//https://www.digitalocean.com/community/tutorials/understanding-maps-in-go?utm_source=twitter&utm_medium=social&utm_campaign=do-maps
import (
	"fmt"
	"sort"
)

func main() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)
	fmt.Println(sammy["name"])

	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}

	keys := []string{}

	for key := range sammy {
		keys = append(keys, key)
	}
	fmt.Printf("%q", keys)
	fmt.Println()

	sort.Strings(keys)
	fmt.Printf("%q", keys)
	fmt.Println()

	items := make([]string, len(sammy))
	var i int

	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q", items)
	fmt.Println()

	fmt.Println(len(sammy))

	counts := map[string]int{}
	fmt.Println(counts["sammy"])

	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}

	if count, ok := counts["sammy"]; ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}

	usernames := map[string]string{"Sammy": "sammy-shark", "Jamie": "mantisshrimp54"}
	usernames["Drew"] = "squidly"
	fmt.Println(usernames)

	followers := map[string]int{"drew": 305, "mary": 428, "cindy": 918}
	followers["drew"] = 342
	fmt.Println(followers)

	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}
	delete(permissions, 16)
	fmt.Println(permissions)
}
