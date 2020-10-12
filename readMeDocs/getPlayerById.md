# Get Player By ID

Replies with game player information from the provided `PlayerId`

**URL** : `/Player/{PlayerId}`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data constraints** : None

**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content Type** : `application/json`

**Content example** : Response will provide game player information.

```json

{
    "PlayerId":"bu1qf15i7nd4ms6pra3g",
    "Name":"William Miller",
    "TimeRegistered":"2020-10-12T00:47:01Z"
}
```

## Error Response

**Condition** : Service did not process server-side generated player information.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "PlayerNotFound",
    "message": "Could not find player using PlayerId: bu1qf15i7nd4ms6pra3g"
}
```

## Notes

* Player ID is a random 20 character string generated using the xid library

  `https://github.com/rs/xid`
