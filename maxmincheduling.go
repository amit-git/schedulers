package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ResourceNeed type for working with individual client needs
type ResourceNeed struct {
	clientID string
	amount   int
}

func getAllocations(resourceNeeds []*ResourceNeed) (map[string]int, error) {
	var allocations = map[string]int{}
	allocations["random-client-id"] = 10
	return allocations, errors.New("Not Implemented")
}

/**
def get_allocations(self, resource_needs):
        allocations = {rn.client: 0 for rn in resource_needs}

        rem_allocations = self.available_resource
        num_remaining_clients = len(resource_needs)
        while rem_allocations > 0:
            current_allocation = rem_allocations / num_remaining_clients
            logger.info(f"current allocation size {current_allocation}")
            rem_allocations = 0
            num_remaining_clients = 0

            for rn in resource_needs:
                if allocations[rn.client] < rn.amount:
                    if (allocations[rn.client] +
                            current_allocation) > rn.amount:
                        # can't fully allocate current allocation
                        rem_allocations += (
                                current_allocation - (
                                    rn.amount - allocations[rn.client]))
                        allocations[rn.client] = rn.amount
                    else:
                        allocations[rn.client] += current_allocation
                        num_remaining_clients += 1
                else:
                    logger.info(f"Skipping allocation to client {rn.client}")

        return allocations

*/

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
	fmt.Println("Hello VS Coding for golang")
}
