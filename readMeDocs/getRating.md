# Get Rating

Replies with a game player's session rating information from the provided `sessionId` and `playerId`

**URL** : `/Session/{SessionId}/Rating?PlayerId=playerid

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

Provide the game SessionId for player rating in path

```text
{
    "/Session/bu1qeqti7nd4jl190s7g/"
}
```

Provide Player Id to retrieve player information in URL Parameters.

*Description* : 20 character alpha-numeric string identify a game player.

```text
  ?PlayerId=bu1qeqti7nd4jl190s70
```

**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content Type** : `application/json`

**Content example** : Response will provide game player session information.

```json

{
    "SessionId":"bu1qeqti7nd4jl190s7g",
    "PlayerId":"bu1qeqti7nd4jl190s70",
    "Rating":3,
    "Comment":"The secular cooling that must someday overtake our planet has already gone far indeed with our neighbour.",
    "TimeSubmitted":"2020-10-12T00:46:36Z"
}
```

## Error Response

**Condition** : Incorrect session or player IDs provided.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "SessionRatingNotFound",
    "message": "Could not find rating by player for session using PlayerId: bu1qeqti7nd4jl190s70"
}
```

## Notes

* Player ID is a random 20 character string generated using the xid library

  `https://github.com/rs/xid`
