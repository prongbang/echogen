# gestgen

## Install

```shell script
$ go install github.com/prongbang/gestgen
```

## How to use

`-f`  feature name

```shell script
$ gestgen -f user
```
OR

```shell script
$ cd project/pkg/api && gestgen -f user
```

## Output

```
user
├── datasource.go
├── handler.go
├── provider.go
├── repository.go
├── router.go
├── usecase.go
└── user.go
```

- `datasource.go`

```go
package user

type DataSource interface {
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

- `handler.go`

```go
package user

type Handler interface {
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

- `provider.go`

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

- `repository.go`

```go
package user

type Repository interface {
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

- `router.go`

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

}

func NewRouter(handle Handler) Router {
    return &router{Handle: handle}
}
```

- `usecase.go`

```go
package user

type UseCase interface {
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

- `user.go`

```go
package user

type User struct  {

}
```