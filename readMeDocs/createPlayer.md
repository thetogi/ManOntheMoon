# Create Player

Creates a randomly generated player.

**URL** : `/Player/Create`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints** : None

**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content example** : Response will provide a return status with a message.

```json
{
    "status": "OK",
    "message": "New Player Successfully created. ID: <player ID>"
}
```

## Error Response

**Condition** : Service did not process server-side generated player information.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "FAILED",
    "message": "New Player was not created. ID: <player ID>"
}
```

## Notes

* Player ID is a random 20 character string generated using the xid library

  `https://github.com/rs/xid`
