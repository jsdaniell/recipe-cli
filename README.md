# recipe-cli
> The fastest and easy way to generate boilerplate projects with structure for production-ready.

### Why use *recipe-cli* instead common technologies' cli?

- Projects structured to different use cases.
- Diferent platform deployment and CI CD Scripts configured.
- Built-in helper scripts (build.sh in some projects).
- Self-learning helper comments (Detailed instructions to customize the files, like instructions to controllers, views, models)

### Installing

To install you can use npm to get globally on your machine:

`npm i -g recipe-cli`

After this you can get the list of project typing `recipe-cli` or init some project with `recipe-cli {language}`, like:

`recipe-cli golang`

A bunch of options will be showed to configure your project.

## Golang Projects Boilerplates

**1. API Project Structure (Mux Router)**

- Two database pre-configured configurations and models (Firebase || MongoDB) || Empty structure.
- Heroku configuration files and build scripts.
- MVP Project Structure with routes example, ready to customize.
- Utilities package for JSON responses and Token validation already created.
- Pre-configured Github Action script for deploy on heroku.
- Pre-configured CORS.

```
project
├── Procfile
├── api
│   ├── controllers
│   │   └── entity.go (controller used on entity_routes inside routes package)
│   ├── db
│   │   └── database.go (database connection - Firebase || MongoDB || Empty)
│   ├── middlewares
│   │   └── middlewares.go (middleware functions to deal with cors configuration)
│   ├── models
│   │   ├── Entity.go (example of model in accord with database choosed)
│   │   └── Token.go (token model to be used with token validation function in security package)
│   ├── repository
│   │   └── entity_repository (example of repository function used inside controllers)
│   ├── responses
│   │   └── responses.go (utility to format JSON responses from the API)
│   ├── router
│   │   ├── router.go
│   │   └── routes (route pre-configured for each controller)
│   ├── server.go
│   └── utils
│       ├── json_utility (utility to work with Structs -> JSON management)
│       └── security (utility for validate token)
├── build.sh
├── config
│   └── config.go
├── deploy_to_heroku.sh (deploy to heroku with sh deploy_to_heroku.sh file)
├── go.mod
├── go.sum
├── heroku.yml (heroku configuration file for go projects)
├── main.go
└── vendor (vendoring of the dependencies)
    ...

30 directories, 17 files
```

**2. CLI Project Structure**

- Utilities for command-line interface, like **selectors** and input user commands with validation.
- Utility to integrate shell commands inside your application.
- Pre-configured release configuration and script.
- CI CD Scripts for publish release pre-configured.
- NPM deploy script configured, production-ready.
