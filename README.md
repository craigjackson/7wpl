# 7 Wonders Pro League

## Download

    $ git clone git@github.com:craigjackson/7wpl.git $GOPATH/src/7wpl

## Setup Database

    $ cd $GOPATH/src/7wpl/cmd/server
    $ sqlite3 database.sqlite3

Go to each models/match.go, models/player.go and match_player.go and run the CREATE TABLE statements in their respsective comments.

## Run Server

    $ cd $GOPATH/src/7wpl/cmd/server
    $ go run main.go
