# Create Game Session

Creates a game session using a random sessionId for a given player.

**URL** : `/Session/Create`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

Provide the game PlayerId for player finishing their game session

```text
{
    "/Session/bu1qeqti7nd4jl190s7g/Create"
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

**Content example** : Response will provide a return status with a message.

```json
{
    "Status": "OK",
    "Message": "New Session Successfully created. ID: bu1qeqti7nd4jl190s7g"
}
```

## Error Response

**Condition** : Service did not process server-side generated rating information.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "Status": "FAILED",
    "Message": "Unable to create new session. ID: bu1qeqti7nd4jl190s7g"
}
```

## Notes

* Player ID and Session IDs are random 20 character strings generated using the xid library

  `https://github.com/rs/xid`
