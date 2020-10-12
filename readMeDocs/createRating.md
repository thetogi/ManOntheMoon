# Create Rating

Creates a rating from a given game session and player. This would be for reviews submitted by a user after a game session completes.

**URL** : `/Session/{Session}/CreateRating`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

Provide the game SessionId for player creating rating in path

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

Provide rating in URL Parameters.

*Description* : Integer from 1-5.

```text
  ?PlayerId=bu1qeqti7nd4jl190s70Rating=3
```

Provide a Comment in URL Parameters.

*Description* : Character string up to 512 characters. Any beyond 512 will be truncated.

```text
  ?PlayerId=bu1qeqti7nd4jl190s70Rating=3
```


**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content example** : Response will provide a return status with a message.

```json
{
    "Status":"OK",
    "Message":"Rating Successfully submitted for sessionID: bu1sc5di7nd3mi2dbua0 Rating: 3 <player_comment>"
}
```

## Error Response

**Condition** : Rating was not provided, contained non-numeric values or non-integer values.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "Status": "FAILED_INVALID_RATING",
    "Message": "<server-side error description>"
}
```

**Condition** : Rating integer was not between the values of 1-5.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "Status": "FAILED_INVALID_RATING_QTY",
    "Message": "Rating was unable to be submitted for Session ID: bu1qeqti7nd4jl190s7g"
}
```

**Condition** : Service did not process server-side generated rating information.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "Status": "FAILED",
    "Message": "Rating was unable to be submitted for Session ID: bu1qeqti7nd4jl190s7g"
}
```

**Condition** : Player rating for game session already exists.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "Status":"FAILED_DUPLICATE",
    "Message":"Player has already submitted a rating for the session. Cannot submit more than one rating for a session. Session: bu1qeqti7nd4jl190s7g Player: bu1qeqti7nd4jl190s70 rating: 3 Comment: The secular cooling that must someday overtake our planet has already gone far indeed with our neighbour."
}

```

## Notes

* Player ID and Session IDs are random 20 character strings generated using the xid library

  `https://github.com/rs/xid`
