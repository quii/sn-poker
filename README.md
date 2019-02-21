# SN Poker

SN poker gives you two applications

## poker-app

When started launches a web server

- [http://localhost:5000](http://localhost:5000) will manage a game of poker. Tell it how many players are playing and it will inform you of what the blind value should be over the course of the game. At the end type in the winner and it will be recorded in the database
- [http://localhost:5000/league](http://localhost:5000/league) shows the league

The app stores the data in an online JSON store.

If you run the app it will use a default bin.

If you want to provide a different one you can run it with a `BIN` env var set

e.g. `BIN=http://mybin/foo go run .`

## jsbin

If you want to make a new bin, run the jsbin app. It will print the URL of the new bin. 


## Test & Build

`./build.sh` or in docker `docker-compose up`

# Run
`./run.sh` or in docker `docker-compose run --service-ports app ./run.sh`

## Deploy to CloudFoundry

```
./build.sh
cf push -p cmd/poker-app
```
