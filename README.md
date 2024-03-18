# go-deck-of-cards

A Golang API server that returns a hand/deck of cards, the server uses Gin as a web framework and Mongodb as database

## Starting the Server

To start all you have to do is run the `docker-compose` up command, and the server will be served on `localhost:8080`

## Testing

To run the tests just run the `go test ./...` and all test files should run


## Endpoints
There are three endpoints on this server, `POST /deck`, `GET /deck/:uuid` and `POST /deck/:uuid/draw`, below is how you can use each one

### `POST /deck`

This endpoint creates a new deck of cards returning a 52 card unshuffled standard deck, you can request the card to be shuffled by passing the `shuffled` param to the request like so `POST /deck?shuffled=true`. You can also request a subset of the cards by sending out a comma separated string of the codes of the specific cards you want in the `cards` param like so `POST /deck?cards=AC,KH,1S`, however the codes are case sensitive

Expected Response:

```
{
    "remaining": 52,
    "shuffled": false,
    "deck_id": "314b0530-ddf6-4851-b22c-a9319dfafaef"
}
```
Expected Error:

```
{
    "error": "invalid code(s): AI, AB, 13S, 123S"
}
```


### `POST /deck/:uuid/draw`

This endpoint creates returns a card from a predefined deck, you can request the number of cards  by passing the `numberOfCards` param to the request like so `POST /deck/:uuid/draw?numberOfCards=5`. 

Expected Response:

```
{
    "cards": {
        [
            {
                "suit": "SPADES",
                "value": "ACE",
                "code": "AS"
            }
        ]
    }
}
```

Expected Error:

```
{
    "error": "The deck you are trying to open does not exist"
}
```


### `GET /deck`

This endpoint returns the deck and all the cards in it. 


Expected Response:

```
{
    "remaining": 2,
    "shuffled": false,
    "deck_id": "1bf37cb0-b91e-459b-a6aa-e10ebd907feb",
    "cards": [
        {
            "suit": "DIAMOND",
            "value": "4",
            "code": "4D"
        },
        {
            "suit": "DIAMOND",
            "value": "5",
            "code": "5D"
        }
    ]
}
```

Expected Error:

```
{
    "error": "The deck you are trying to open does not exist"
}
```
