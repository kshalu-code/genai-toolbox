---
title: "looker-search-model-sets"
type: docs
weight: 1
description: >
  A "looker-search-model-sets" tool searches for model sets in the Looker instance.
aliases:
- /resources/tools/looker-search-model-sets
---

## About

A `looker-search-model-sets` tool searches for model sets in the Looker instance. A model set defines a collection of LookML models that can be assigned to a permission set.

It's compatible with the following sources:

- [looker](../../sources/looker.md)

`looker-search-model-sets` accepts the following parameters:
- `name` (optional): Filter by model set name.
- `id` (optional): Filter by specific model set ID.
- `limit` (optional): Maximum number of results to return. Default is 100.
- `offset` (optional): Starting point for pagination. Default is 0.

## Example

```yaml
tools:
    search_model_sets:
        kind: looker-search-model-sets
        source: looker-source
        description: |
          This tool searches for model sets in the Looker instance.
          A model set defines a collection of LookML models that can be assigned to a permission set.

          Parameters:
          - name (optional): Filter by model set name.
          - id (optional): Filter by specific model set ID.
          - limit (optional): Maximum number of results to return. Default is 100.
          - offset (optional): Starting point for pagination. Default is 0.

          Output:
          A JSON array of objects, where each object represents a model set and contains details
          such as `id`, `name`, `models` (array of strings), and `all_access` (boolean).
```

## Reference

| **field**   | **type** | **required** | **description**                                    |
|-------------|:--------:|:------------:|----------------------------------------------------|
| kind        |  string  |     true     | Must be "looker-search-model-sets".                |
| source      |  string  |     true     | Name of the source Looker instance.                |
| description |  string  |     true     | Description of the tool that is passed to the LLM. |
