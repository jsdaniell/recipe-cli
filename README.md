# recipe-cli
> The fastest and easy way to generate boilerplate projects with structure for production-ready.

### Why use *recipe-cli* instead common technologies' cli?

- Projects structured to different use cases.
- Diferent platform deployment and CI CD Scripts configured.
- Built-in helper scripts (build.sh in some projects).
- Self-learning helper comments (Detailed instructions to customize the files, like instructions to controllers, views, models)

## Golang Projects Boilerplates

**API Structure**

```
project
 ├── api
 │   ├── controllers
 │   │   └── entity.go (controller example, customize your own from here)
 │   ├── db
 │   │   └── database.go (two options pre-configured - Firebase / MongoDB | also can be boilerplated with empty if no database choosed)
 │   ├── models
 │   │   ├── Entity.go (depends of the database selected)
 │   │   └── Token.go
 │   ├── repository
 │   │   └── entity_repository
 │   ├── responses
 │   │   └── responses.go
 │   └── server.go
 ├── go.mod
 └── main.go
```