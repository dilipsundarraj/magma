#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .equipment_fragment import EquipmentFragment, QUERY as EquipmentFragmentQuery

QUERY: List[str] = EquipmentFragmentQuery + ["""
query EquipmentPositionsQuery($id: ID!) {
  equipment: node(id: $id) {
    ... on Equipment {
      equipmentType {
        positionDefinitions {
          id
          name
        }
      }
      positions {
        definition {
          id
          name
        }
        attachedEquipment {
          ...EquipmentFragment
        }
      }
    }
  }
}

"""]

@dataclass
class EquipmentPositionsQuery(DataClassJsonMixin):
    @dataclass
    class EquipmentPositionsQueryData(DataClassJsonMixin):
        @dataclass
        class Node(DataClassJsonMixin):
            @dataclass
            class EquipmentType(DataClassJsonMixin):
                @dataclass
                class EquipmentPositionDefinition(DataClassJsonMixin):
                    id: str
                    name: str

                positionDefinitions: List[EquipmentPositionDefinition]

            @dataclass
            class EquipmentPosition(DataClassJsonMixin):
                @dataclass
                class EquipmentPositionDefinition(DataClassJsonMixin):
                    id: str
                    name: str

                @dataclass
                class Equipment(EquipmentFragment):
                    pass

                definition: EquipmentPositionDefinition
                attachedEquipment: Optional[Equipment]

            equipmentType: EquipmentType
            positions: List[EquipmentPosition]

        equipment: Optional[Node]

    data: EquipmentPositionsQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> EquipmentPositionsQueryData:
        # fmt: off
        variables = {"id": id}
        response_text = client.call(''.join(set(QUERY)), variables=variables)
        return cls.from_json(response_text).data
