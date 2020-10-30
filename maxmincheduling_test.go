package main

import (
	"fmt"
	"testing"
)

func TestBuildResourceNeeds(t *testing.T) {
	resourceNeeds := buildResourceNeeds("facebook=50,amazon=20,netflix=100,google=150")
	if len(resourceNeeds) == 4 {
		for _, rn := range resourceNeeds {
			t.Log(fmt.Sprintf("%s needs %d resources", rn.clientID, rn.amount))
		}
		t.Log("Looking Good")
	} else {
		t.Error("Resource needs not built")
	}
}

func TestAllocatingResources(t *testing.T) {
	_, err := getAllocations([]*ResourceNeed{})
	if err != nil {
		t.Error(fmt.Printf("Error in allocating resources - %s", err))
	}
}
