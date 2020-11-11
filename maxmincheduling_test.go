package main

import (
	"testing"
)

func TestBuildResourceNeeds(t *testing.T) {
	resourceNeeds := buildResourceNeeds("facebook=50,amazon=20,netflix=100,google=150")
	if len(resourceNeeds) == 4 {
		for _, rn := range resourceNeeds {
			t.Logf("%s needs %d resources", rn.clientID, rn.amount)
		}
	} else {
		t.Error("Resource needs not built")
	}
}

func TestAllocatingResources(t *testing.T) {
	resourceNeeds := buildResourceNeeds("facebook=50,amazon=20,netflix=100,google=150")
	allocations, err := getAllocations(resourceNeeds, 100)
	if err != nil {
		t.Errorf("Error in allocating resources - %s", err)
	}
	for clientID, allocation := range allocations {
		t.Logf("%s => %d\n", clientID, allocation)
	}
	expectedAllocations := map[string]int{"facebook": 26, "amazon": 20, "netflix": 26, "google": 26}
	for resourceName, expectedAllocation := range expectedAllocations {
		if allocations[resourceName] != expectedAllocation {
			t.Errorf("Error in allocating resources to %s ( %d vs expected %d )", resourceName, allocations[resourceName], expectedAllocation)
		}
	}
}
