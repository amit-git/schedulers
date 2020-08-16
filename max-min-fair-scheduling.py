#!/usr/local/bin/python3
import logging
import sys
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


def build_resource_needs(alloc_argv):
    resource_needs = []
    key_value_pairs = alloc_argv.split(",")
    for kv in key_value_pairs:
        k, v = kv.split("=")
        resource_needs.append(ResourceNeed(client_id=k, amount=int(v)))
    return resource_needs

# Main ############################

if len(sys.argv) != 2:
    print("Missing argument <resource_needs>")
    print("Usage: max-min-fair-scheduling.py A=24,B=45,C=15,D=10")
    exit(-1)

s = MaxMinScheduler(100)
allocations = s.get_allocations(build_resource_needs(sys.argv[1]))
logger.info(pformat(allocations))
