# Get All Sessions

Replies with all player's sessions information.

**URL** : `/Sessions/'

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data constraints** : None

**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content Type** : `application/json`

**Content example** : Response will provide all game player sessions information.

```json

[
    {
        "SessionId":"bu1sc55i7nd3mi2dbs90",
        "PlayerId":"bu1sc55i7nd3mi2dbs8g",
        "TimeSessionEnd":"2020-10-12T02:57:24Z"
    },
    {
        ...
    }
]
```

## Error Response

**Condition** : No Sessions have been recorded.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "SessionsNotFound",
    "message": "No sessions could be found."
}
```

## Notes

* Player ID is a random 20 character string generated using the xid library

  `https://github.com/rs/xid`
