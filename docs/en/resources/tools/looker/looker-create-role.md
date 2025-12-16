---
title: looker-create-role
weight: 1
description: >
  A "looker-create-role" tool creates a new role in the Looker instance.
aliases:
- /resources/tools/looker-create-role
---

## About

A `looker-create-role` tool creates a new role in the Looker instance. A role is a combination of exactly one permission set and exactly one model set.

It's compatible with the following sources:

- [looker](../../sources/looker.md)

`looker-create-role` accepts the following parameters:
- `name` (required): The name of the new role.
- `permission_set_id` (required): The ID of the associated permission set.
- `model_set_id` (required): The ID of the associated model set.

## Example

```yaml
tools:
    create_role:
        kind: looker-create-role
        source: looker-source
        description: |
          This tool creates a new role in a Looker instance.
          A role is a combination of a permission set and a model set.

          Parameters:
          - name: The name of the new role.
          - permission_set_id: The ID of the associated permission set.
          - model_set_id: The ID of the associated model set.

          Returns the newly created role's details including id, name, permission_set, and model_set.
```

## Reference

- [Looker API: Create Role](https://docs.cloud.google.com/looker/docs/reference/looker-api/latest/methods/Role/create_role)
