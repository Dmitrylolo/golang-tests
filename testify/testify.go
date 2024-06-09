package testify

type Developer struct {
	Name string
}

func FilterUnique(developers []Developer) []string {
	var uniques []string
	check := make(map[string]int)
	for _, developer := range developers {
		check[developer.Name] = 1
	}

	for name := range check {
		uniques = append(uniques, name)
	}
	return uniques
}

// func testify() {
// 	developers := []Developer{
// 		{"John", 25},
// 		{"Jane", 26},
// 		{"John", 25},
// 		{"Jane", 26},
// 		{"Jack", 27},
// 		{"Jill", 28},
// 		{"Jack", 27},
// 		{"Jill", 28},
// 	}

// 	uniques := FilterUnique(developers)

// 	for _, name := range uniques {
// 		fmt.Println(name)
// 	}
// }
