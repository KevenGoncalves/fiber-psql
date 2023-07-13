# Fiber Postgresql Boilerplate App

## Requirements to run the app

- Install [Taskfile](https://taskfile.dev/installation) (required)
- Install other packages with `task install`

## Env Vars

- Copy the app.env.example to app.env
- Populate the app.env

## Scripts

- To run migrations `task goose -- up`
- To run development server `task dev`

## Deployment

Railway is the recommended deployment platform, but any platform that supports Docker should work.

## TODO

- Add Swagger
