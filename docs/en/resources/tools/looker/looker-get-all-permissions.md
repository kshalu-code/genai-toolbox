---
title: "looker-get-all-permissions"
type: docs
weight: 1
description: >
  A "looker-get-all-permissions" tool retrieves all permissions in a Looker instance.
aliases:
- /resources/tools/looker-get-all-permissions
---

## About

A `looker-get-all-permissions` tool retrieves all permissions in a Looker instance. Permissions define the actions that can be performed in the system.

It's compatible with the following sources:

- [looker](../../sources/looker.md)

`looker-get-all-permissions` accepts no parameters.

## Example

```yaml
tools:
    get_all_permissions:
        kind: looker-get-all-permissions
        source: looker-source
        description: |
          This tool retrieves all permissions in a Looker instance.

          Parameters:
          This tool accepts no parameters.

          Output:
          A JSON array of objects, each representing a permission and containing details
          such as `permission` (string), `parent` (string), and `description` (string).
```

## Reference

| **field**   | **type** | **required** | **description**                                    |
|-------------|:--------:|:------------:|----------------------------------------------------|
| kind        |  string  |     true     | Must be "looker-get-all-permissions".              |
| source      |  string  |     true     | Name of the source Looker instance.                |
| description |  string  |     true     | Description of the tool that is passed to the LLM. |
