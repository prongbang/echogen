# echogen 🚀

[![Go Reference](https://pkg.go.dev/badge/github.com/prongbang/echogen.svg)](https://pkg.go.dev/github.com/prongbang/echogen)
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/echogen)](https://goreportcard.com/report/github.com/prongbang/echogen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/prongbang/echogen.svg)](https://golang.org)

> Code generator for Echo web framework following clean architecture principles. Generate complete feature structures with a single command.

## ✨ Features

- 🏗️ **Clean Architecture** - Automatically generates layered architecture structure
- 🔌 **Wire Integration** - Built-in support for Google Wire dependency injection
- 📦 **Echo Framework** - Optimized for Echo web framework
- 🎯 **Feature-Based** - Generates complete feature modules
- ⚡ **Fast Development** - Accelerate your development workflow
- 🧩 **Modular Design** - Well-organized and maintainable code structure

## 📦 Installation

```shell
go get -u github.com/prongbang/echogen
go install github.com/prongbang/echogen
```

## 🚀 Quick Start

Generate a new feature module with a single command:

```shell
echogen -f user
```

Or generate within a specific directory:

```shell
cd project/pkg/api && echogen -f user
```

## 📁 Generated Structure

When you run `echogen -f user`, it creates the following structure:

```
user/
├── datasource.go    # Database operations
├── handler.go       # HTTP handlers
├── provider.go      # Wire dependency providers
├── repository.go    # Business logic repository
├── router.go        # Route definitions
├── usecase.go       # Use case/business logic
└── user.go          # Domain model
```

## 📝 Generated Code Examples

### 1. DataSource Layer

`datasource.go`
```go
package user

type DataSource interface {
    // Add your database operations here
}

type dataSource struct {
    DbSource database.DataSource
}

func NewDataSource(dbSource database.DataSource) DataSource {
    return &dataSource{
        DbSource: dbSource,
    }
}
```

### 2. Repository Layer

`repository.go`
```go
package user

type Repository interface {
    // Add your repository methods here
}

type repository struct {
    Ds DataSource
}

func NewRepository(ds DataSource) Repository {
    return &repository{
        Ds: ds,
    }
}
```

### 3. UseCase Layer

`usecase.go`
```go
package user

type UseCase interface {
    // Add your business logic methods here
}

type useCase struct {
    Repo Repository
}

func NewUseCase(repo Repository) UseCase {
    return &useCase{
        Repo: repo,
    }
}
```

### 4. Handler Layer

`handler.go`
```go
package user

type Handler interface {
    // Add your HTTP handlers here
}

type handler struct {
    Uc UseCase
}

func NewHandler(uc UseCase) Handler {
    return &handler{
        Uc: uc,
    }
}
```

### 5. Router Configuration

`router.go`
```go
package user

import "github.com/labstack/echo"

type Router interface {
    Initial(e *echo.Echo)
}

type router struct {
    Handle Handler
}

func (r *router) Initial(e *echo.Echo) {
    // Add your routes here
    // e.GET("/users", r.Handle.GetUsers)
    // e.POST("/users", r.Handle.CreateUser)
}

func NewRouter(handle Handler) Router {
    return &router{Handle: handle}
}
```

### 6. Wire Provider

`provider.go`
```go
package user

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
    NewDataSource,
    NewRepository,
    NewUseCase,
    NewHandler,
    NewRouter,
)
```

### 7. Domain Model

`user.go`
```go
package user

type User struct {
    // Add your user fields here
    ID        string `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    CreatedAt int64  `json:"created_at"`
}
```

## 🔧 Customization

After generating the basic structure, you can customize each layer:

1. **Add Methods** - Define interfaces and implement methods
2. **Add Fields** - Extend structs with necessary fields
3. **Add Dependencies** - Inject additional dependencies as needed
4. **Add Validations** - Implement input validation logic
5. **Add Tests** - Write unit tests for each layer

## 🎯 Best Practices

1. **Follow Clean Architecture** - Keep dependencies pointing inward
2. **Use Interfaces** - Program to interfaces, not implementations
3. **Error Handling** - Implement proper error handling at each layer
4. **Logging** - Add logging where appropriate
5. **Documentation** - Document your code and APIs

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 💖 Support the Project

If you find this package helpful, please consider supporting it:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

## 🔗 Related Projects

- [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework
- [Wire](https://github.com/google/wire) - Compile-time dependency injection for Go
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Architecture pattern by Uncle Bob

---
