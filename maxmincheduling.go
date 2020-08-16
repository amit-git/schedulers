package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ResourceNeed struct {
	clientId string
	amount   int
}

func buildResourceNeeds(allocArgs string) []*ResourceNeed {
	var resourceNeeds []*ResourceNeed
	var amountVal int
	var err error
	resourceParts := strings.Split(allocArgs, ",")
	for _, resourcePart := range resourceParts {
		resourceNeedParts := strings.Split(resourcePart, "=")
		amountVal, err = strconv.Atoi(resourceNeedParts[1])
		if err != nil {
			panic(fmt.Sprintf("%s :: Bad argument", resourceNeedParts[1]))
		}
		resourceNeeds = append(resourceNeeds, &ResourceNeed{
			clientId: resourceNeedParts[0],
			amount:   amountVal,
		})
	}
	return resourceNeeds
}

func main() {
	fmt.Println("vim-go")
}
