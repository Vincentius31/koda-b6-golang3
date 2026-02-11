package lib

import "strings"

func SearchPerson(users[] string, keyword string) []string{
	var results []string

	for i := 0; i < len(users); i++ {
		if strings.EqualFold(users[i], keyword) {
			results = append(results, users[i])
		}
	}

	return results
}