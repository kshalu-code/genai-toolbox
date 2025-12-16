---
title: "looker-create-model-set"
type: docs
weight: 1
description: >
  A "looker-create-model-set" tool creates a new model set in the Looker instance.
aliases:
- /resources/tools/looker-create-model-set
---

## About

A `looker-create-model-set` tool creates a new model set in the Looker instance. A model set defines a collection of LookML models that can be assigned to a permission set.

It's compatible with the following sources:

- [looker](../../sources/looker.md)

`looker-create-model-set` accepts the following parameters:
- `name` (required): The name of the new model set.
- `models` (required): A list of model names to include in the set.

## Example

```yaml
tools:
    create_model_set:
        kind: looker-create-model-set
        source: looker-source
        description: |
          This tool creates a new model set in the Looker instance.
          A model set defines a collection of LookML models that can be assigned to a permission set.

          Parameters:
          - name: The name of the new model set.
          - models: A list of model names to include in the set.

          Output:
          A JSON object representing the newly created model set, containing details
          such as `id`, `name`, `models` (array of strings), and `all_access` (boolean).
```

## Reference

| **field**   | **type** | **required** | **description**                                    |
|-------------|:--------:|:------------:|----------------------------------------------------|
| kind        |  string  |     true     | Must be "looker-create-model-set".                 |
| source      |  string  |     true     | Name of the source Looker instance.                |
| description |  string  |     true     | Description of the tool that is passed to the LLM. |
