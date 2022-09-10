package main

import (
	"fmt"
)

func main() {
	// teknik untuk mengambil data konkret dari data yang terbungkus dalam interface

	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleep"},
	}

	fmt.Println(data["nama"].(string))      // john wick
	fmt.Println(data["grade"].(int))        // 2
	fmt.Println(data["height"].(float64))   // 156.5
	fmt.Println(data["isMale"].(bool))      // true
	fmt.Println(data["hobbies"].([]string)) // [eating sleep]

	// jika tidak tahu tipe datanya, bisa meng-casting ke tipe type, namun hanya bisa dilakukan pada switch
	fmt.Println("========================")
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))

		}
	}

}
