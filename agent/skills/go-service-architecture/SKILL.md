---
name: go-service-architecture
description: Recommended Go service architecture. Use when designing, implementing, reviewing, or refactoring Go services with similar scaffolding.
---

<golang_architecture>
  <goal>Keep business behavior independent from transport, storage, SDKs, and process startup code. Each layer has one reason to change</goal>

  <directory_structure>
    ```text
    ├── cmd/                                      # process entry points; minimal code
    ├── api/                                      # external contracts: proto, Swagger
    ├── internal/
    │   ├── app/                                  # application assembly
    │   ├── config/                               # application configuration
    │   ├── controller/                           # external input handling and validation
    │   │   ├── http/                             # HTTP handlers, if HTTP is used
    │   │   │   └── <handler>/
    │   │   │       ├── interfaces.go             # use-case interfaces needed by handler; commands and response aggregates
    │   │   │       ├── interfaces_mock.go        # mocks for interfaces.go
    │   │   │       ├── service.go                # handler constructor and helper methods
    │   │   │       └── <or code>.go
    │   │   ├── grpc/                             # gRPC handlers, if gRPC is used
    │   │   └── consumer/                         # consumers
    │   ├── usecase/                              # business scenarios
    │   │   └── <usecase>/
    │   │       ├── interfaces.go                 # external dependencies; commands and response aggregates
    │   │       ├── interfaces_mock.go            # mocks for interfaces.go
    │   │       ├── service.go                    # use-case constructor and helper methods
    │   │       └── <or code>.go
    │   ├── domain/                               # domain model without external dependencies
    │   │   └── <area>/                           # business area: models, events, errors
    │   ├── repository/                           # storage access
    │   │   └── <area>/
    │   │       ├── service.go                    # repository constructor and helper methods
    │   │       └── <or code>.go
    │   ├── infra/                                # external APIs, caches, producers
    │   │   └── <adapter>/
    │   │       ├── service.go                    # adapter constructor and helper methods
    │   │       └── <or code>.go
    │   └── core/                                 # shared utilities and helpers
    ├── pkg/                                      # code for external reuse
    └── migrations/                               # database migrations
    ```

    <file_rules>
      1. Interfaces live in `interfaces.go` at consumer site.
      2. Interface method types (commands, response aggregates, etc.) live near interface.
      3. Mocks for interfaces from `interfaces.go` live in `interfaces_mock.go`.
      4. Package constructors live in `service.go` and are named `New(...)`, NOT `New<PackageName>(...)`. Keep `service.go` small and focused on construction and helper methods.
      5. Tests are white-box by default: test package matches package under test.
    </file_rules>
  </directory_structure>

  <guidelines>
    1. DTO tags such as `json`, `db`, `bson`, `yaml`, or `xml` MUST NEVER appear inside domain and usecase layers.
    2. Only commands, response aggregates, and domain models can "move" between layers.
    3. Internal models and methods of layers are NEVER exported outside layer.
    4. Repository layer does not contain business logic, but it may have filtering, sorting, aggregation, etc. logic either in form of database queries or programmatically.
    5. Usecase that calls other usecases MUST be modeled as orchestrator and depend on them through interfaces.
  </guidelines>

  <interfaces>
    1. MUST describe consumer need, not implementation capability.
    2. Method types MUST NOT use technical DTOs, e.g., HTTP, gRPC, etc.
    3. MUST contain only methods needed by specific consumer.
    4. Every implementation MUST include a compile-time assertion in implementation package: `var _ consumer.Interface = (*Service)(nil)`.
    5. Interface methods MUST contain only minimum necessary for consumer to work.
    6. Consumer owns not only interface, but also all method types: commands, response aggregates, filters, payloads.
    7. Before making a type implement an interface from another package, MUST verify that interface package does not already depend transitively on implementation package. If it does, adding required implementation-to-consumer import would create a cycle and work MUST STOP before code changes.
    8. Cycle exposed by a required compile-time interface assertion is proof of an invalid dependency graph. MUST NOT move or remove assertion, place it in a wiring package, suppress linter, relocate interface away from its consumer, or introduce a shared contract package. MUST identify and remove dependency edge that violates intended layer direction.
  </interfaces>

  <layers>
    1. `internal/domain`
        1) Stores business model types: entities, value objects, aggregates, application events, and application errors.
        2) MUST NOT depend on any other layer.
        3) MUST NOT contain command or response aggregates. Keep them near interfaces that use them.
        4) It is important not to confuse domain models and response aggregates. main criteria:
            - Domain models describe business entities and their behavior
            - Response aggregates describe data needed by a specific consumer
        5) If response aggregates are needed in several usecases, they can be moved to a separate package at top level of usecase.
    2. `internal/usecase`
        1) Implements business scenarios and orchestration.
        2) MUST NOT depend directly on repository and infra layers. Only depend on their interfaces.
    3. `internal/controller`
        1) Handles transport input (HTTP/gRPC/consumers) and validation.
        2) MUST NOT contain business logic or orchestration.
    4. `internal/repository`
        1) Works with data storage.
        2) MUST NOT contain business logic or orchestration.
    5. `internal/infra`
        1) Works with HTTP/gRPC clients, caches, producers, etc.
        2) MUST NOT contain business logic or orchestration.
    6. `internal/app`
        1) Assembles application: creates and wires all dependencies.
        2) MUST NOT contain business logic or orchestration.
        3) MUST NOT use interfaces, but only concrete implementations of dependencies.
  </layers>

  <request_flow>
    Describes execution order, not import direction:
      1. External input arrives through an HTTP handler, gRPC handler, message consumer, or event consumer.
      2. `controller` checks input format, maps external data to command type of its local interface, and calls that interface.
      3. `usecase` implementation receives controller-owned types, applies application rules, and calls external dependencies through use-case-owned interfaces.
      4. `repository` and `infra` receive use-case-owned contract types and map them to SQL rows, storage DTOs, external SDK DTOs, Kafka messages, or external HTTP/gRPC requests.
      5. `controller` receives response from its local interface and maps it to a transport response or processing acknowledgement.
  </request_flow>

  <examples>
    <example name="defining usecase contract for controller">
      ```go
      // internal/controller/http/order/interfaces.go
      package order
      import "context"
      type CreateOrderUsecase interface {
          CreateOrder(ctx context.Context, cmd CreateOrderCommand) (CreateOrderResponse, error)
      }
      // internal/controller/http/order/command.go
      type CreateOrderCommand struct {
          ClientID string
          ItemIDs []string
      }
      // internal/controller/http/order/response.go
      type CreateOrderResponse struct {
          OrderID string
      }
      ```
    </example>
     <example name="implementing usecase contract in usecase layer">
      ```go
      // internal/usecase/createorder/service.go
      package createorder
      import (
          "context"
          orderhttp "example/internal/controller/http/order"
      )
      type Service struct{}
      var _ orderhttp.CreateOrderUsecase = (*Service)(nil)
      func (s *Service) CreateOrder(
          ctx context.Context,
          cmd orderhttp.CreateOrderCommand,
      ) (orderhttp.CreateOrderResponse, error) {
          orderID := cmd.ClientID + "-order"
          return orderhttp.CreateOrderResponse{OrderID: orderID}, nil
      }
      ```
    </example>
    <example name="defining usecase contract for repository">
      ```go
      // internal/usecase/createorder/interfaces.go
      package createorder
      import "context"
      type OrderRepository interface {
          SaveOrder(ctx context.Context, cmd SaveOrderCommand) (SaveOrderResponse, error)
      }
      // internal/usecase/createorder/command.go
      type SaveOrderCommand struct {
          ClientID string
          ItemIDs []string
      }
      // internal/usecase/createorder/response.go
      type SaveOrderResponse struct {
          OrderID string
      }
      ```
    </example>
    <example name="implementing usecase contract in repository layer">
      ```go
      // internal/repository/order/service.go
      package order
      import (
          "context"
          createorder "example/internal/usecase/createorder"
      )
      type Service struct{}
      var _ createorder.OrderRepository = (*Service)(nil)
      func (s *Service) SaveOrder(
          ctx context.Context,
          cmd createorder.SaveOrderCommand,
      ) (createorder.SaveOrderResponse, error) {
          orderID := cmd.ClientID + "-saved"
          return createorder.SaveOrderResponse{OrderID: orderID}, nil
      }
      ```
    </example>
  </examples>
</golang_architecture>