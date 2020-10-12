# Retrieve Random Rating

Generates and returns a random player rating, comment, SessionId, and PlayerId. This is used for diagnostic testing and does not interact with any database. Information returned does not persist.

**URL** : `/GameSession/Rating`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data constraints**: NONE

**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content example** : Response will provide a return status with a message.

```json
{
  "SessionId":"bu1skenr3qbtb1m73d30",
  "PlayerId":"bu1skenr3qbtb1m73d3g",
  "Rating":1,"Comment":"Near it in the field, I remember, were three faint points of light, three telescopic stars infinitely remote, and all around it was the unfathomable darkness of empty space.",
  "TimeSubmitted":"2020-10-12T03:15:06.9222322Z"
}
```

## Notes

* Player ID and Session IDs are random 20 character strings generated using the xid library

  `https://github.com/rs/xid`
