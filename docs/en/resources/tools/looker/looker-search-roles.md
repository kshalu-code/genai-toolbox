# Looker - Search Roles

This tool allows you to search for roles in Looker. You can filter by role ID or name, and it supports pagination with limit and offset.

## About Looker Roles
In Looker, a **Role** defines what a user or group can do and which data they can see. It is a combination of exactly one **Permission Set** and exactly one **Model Set**.

## Example

### Prompt
"Search for roles named 'Admin'"

### Tool Call
```json
{
  "name": "looker-search-roles",
  "arguments": {
    "name": "Admin"
  }
}
```

## Reference

### Parameters

| Name | Type | Description | Required | Default |
| :--- | :--- | :--- | :--- | :--- |
| `name` | string | The name of the role to search for. | No | |
| `id` | integer | The unique ID of the role. | No | |
| `limit` | integer | The maximum number of roles to return. | No | 100 |
| `offset` | integer | The number of roles to skip before starting counts. | No | 0 |

### Output

A list of roles, each containing:
- `id`: The unique ID of the role.
- `name`: The name of the role.
- `permissions`: A list of permission strings associated with the role.
- `models`: A list of model names associated with the role.
