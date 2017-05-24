/* @flow */

import type {Attr} from './CollectionTypes'

export type Stop = {
  location: number;
  color: string;
}

export type InterpolationType = "constant" | "linear"

export type Color = {
  type: InterpolationType;
  colors: Stop[];
  data_key: ?string;
  bounds: ?[number,number];
}

export type Size = {
  type: InterpolationType;
  data_key: ?string;
  bounds: [number,number];
}

export type PointDecorator = {
  points: {
    color: Color,
    size: Size,
    sprite: string
  };
  title: string;
  type: "point";
}

export type Decorator = PointDecorator

export type GroupingOperationType = "equal" | "within" | "peak"
export type GroupingOperation = {
  operation: GroupingOperationType;
  parameter: ?number;
  source_attribute: Attr;
}

export type Op = "avg" | "simple_string" | "simple_num" | "simple_location" | "max" | "min" | "median" | "first" | "last" | "sum" | "count" | "whole_group"

export type SelectionOperation = {
  id: number;
  value_name: string;
  source_attribute: Attr;
  operation: Op;
}

export type Viz = {
  source_collections: string[];
  grouping_operation: GroupingOperation;
  selection_operations: SelectionOperation[];
  decorator: Decorator
}
