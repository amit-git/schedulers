package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ResourceNeed type for working with individual client needs
type ResourceNeed struct {
	clientID string
	amount   int
}

func getAllocations(resourceNeeds []*ResourceNeed, availableResources int) (map[string]int, error) {
	var allocations = map[string]int{}
	for _, req := range resourceNeeds {
		allocations[req.clientID] = 0
	}
	remainingAllocations := availableResources
	numRemainingClients := len(resourceNeeds)
	for remainingAllocations > 0 {
		currentAllocation := remainingAllocations / numRemainingClients
		log.Printf("Current allocation size %d\n", currentAllocation)
		remainingAllocations = 0
		numRemainingClients = 0
		for _, req := range resourceNeeds {
			if allocations[req.clientID] < req.amount {
				if allocations[req.clientID]+currentAllocation > req.amount {
					// can't fully allocate current allocations
					remainingAllocations += (currentAllocation - (req.amount - allocations[req.clientID]))
					allocations[req.clientID] = req.amount
				} else {
					allocations[req.clientID] += currentAllocation
					numRemainingClients++
				}
			} else {
				log.Printf("Skipping allocations to client %s", req.clientID)
			}
		}
	}
	return allocations, nil
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
			clientID: resourceNeedParts[0],
			amount:   amountVal,
		})
	}
	return resourceNeeds
}

func main() {
	fmt.Println("TODO - Command line parameters for scheduler")
}
