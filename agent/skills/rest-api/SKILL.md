---
name: rest-api
description: REST API Standard and Best Practices. Use BEFORE modifying REST APIs.
---

<rest_api_standard>

  <endpoint_naming>
    1. Use nouns, not verbs (e.g., `/users` not `/getUsers`)
    2. Use plural nouns for collections (e.g., `/users`, `/orders`)
    3. Use lowercase with hyphens (kebab-case) for multi-word resources (e.g., `/user-profiles`)
    4. Keep URIs short and readable, avoid abbreviations
    5. DON'T include file extensions in URIs (e.g., no `.json`)
    6. Include API version at base path (e.g., `/v1/users`)
  </endpoint_naming>

  <http_methods>
    1. `POST /api/v1/my-items` - creation
    2. `GET /api/v1/my-items/{id}` - retrieval
    3. `PATCH /api/v1/my-items/{id}` - update
    4. `DELETE /api/v1/my-items/{id}` - deletion
    5. `GET /api/v1/my-items` - list with pagination
  </http_methods>

  <response_structure>
    1. Use JSON as standard data format
    2. Return appropriate HTTP status codes (200, 201, 204, 400, 401, 403, 404, 500)
    3. Provide consistent error responses with meaningful messages
    4. Use standard fields like `id`, `created_time`, `updated_time` for resources
  </response_structure>

  <query_parameters>
    1. Use for filtering, sorting, and pagination (e.g., `/users?role=admin&page=2&limit=10`)
    2. DON'T embed filtering in the path
    3. Use snake_case for parameter names (e.g., `user_id`, `created_after`)
  </query_parameters>

  <general_conventions>
    1. Minimize nesting depth (avoid more than 2-3 levels)
    2. Use hierarchical structure only when clear relationship exists (e.g., `/users/{id}/orders`)
    3. Maintain consistency across all endpoints
    4. DON'T expose internal implementation details in URLs
    5. Use American English spelling and established abbreviations
  </general_conventions>

  <examples>
  ```
  GET /v1/users
  POST /v1/users
  GET /v1/users/{user_id}
  PATCH /v1/users/{user_id}
  DELETE /v1/users/{user_id}
  GET /v1/users/{user_id}/orders
  ```

  Response format:
  ```json
  {
    "id": "user_123",
    "name": "John Doe",
    "email": "john@example.com",
    "created_time": "2023-01-01T00:00:00Z",
    "updated_time": "2023-01-02T12:30:00Z"
  }
  ```

  Error response:
  ```json
  {
    "error": {
      "code": "NOT_FOUND",
      "message": "User not found",
      "details": "User with ID 'invalid_id' does not exist"
    }
  }
  ```
  </examples>

</rest_api_standard>