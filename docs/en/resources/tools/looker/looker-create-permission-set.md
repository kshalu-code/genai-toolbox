---
title: looker-create-permission-set
weight: 1
description: >
  A "looker-create-permission-set" tool creates a new permission set in the Looker instance.
aliases:
- /resources/tools/looker-create-permission-set
---

## About

A `looker-create-permission-set` tool creates a new permission set in the Looker instance. Permission sets define what actions a user or group can perform (e.g., `access_data`, `see_looks`).

It's compatible with the following sources:

- [looker](../../sources/looker.md)

`looker-create-permission-set` accepts the following parameters:
- `name` (required): The name of the new permission set.
- `permissions` (required): A list of permission strings to include in the set.

## Example

```yaml
tools:
    create_permission_set:
        kind: looker-create-permission-set
        source: looker-source
        description: |
          This tool creates a new permission set in a Looker instance.
          Permission sets define what actions a user or group can perform.

          Parameters:
          - name: The name of the new permission set.
          - permissions: A list of permission strings to include in the set.

          Returns the newly created permission set's details.
```

## Reference

- [Looker API: Create Permission Set](https://docs.cloud.google.com/looker/docs/reference/looker-api/latest/methods/Role/create_permission_set)
