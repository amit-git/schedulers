#!/usr/local/bin/python3
import logging
from pprint import pformat

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class ResourceNeed:
    def __init__(self, client_id, amount):
        self.client = client_id
        self.amount = amount

class MaxMinScheduler:
    def __init__(self, available_resource):
        self.available_resource = available_resource
        pass

    def get_allocations(self, resource_needs):
        allocations = {rn.client:0 for rn in resource_needs}

        rem_allocations = self.available_resource
        num_remaining_clients = len(resource_needs)
        while rem_allocations > 0:
            current_allocation = rem_allocations / num_remaining_clients
            logger.info(f"current allocation size {current_allocation}")
            rem_allocations = 0
            num_remaining_clients = 0

            for rn in resource_needs:
                if allocations[rn.client] < rn.amount:
                    if (allocations[rn.client] + current_allocation) > rn.amount:
                        # can't fully allocate current allocation
                        rem_allocations += (current_allocation - (rn.amount - allocations[rn.client]))
                        allocations[rn.client] = rn.amount
                    else:
                        allocations[rn.client] += current_allocation
                        num_remaining_clients += 1
                else: 
                    logger.info(f"Skipping allocation to client {rn.client}")

        return allocations

############ Main #########################
r1 = ResourceNeed(client_id="uber", amount=34)
r2 = ResourceNeed(client_id="linkedIn", amount=10)
r3 = ResourceNeed(client_id="apple", amount=55)
r4 = ResourceNeed(client_id="google", amount=45)

s = MaxMinScheduler(100)
allocations = s.get_allocations([r1, r2, r3, r4])
logger.info(pformat(allocations))

