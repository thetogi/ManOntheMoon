# Get All Ratings

Replies with all player's session ratings information. The information retrieved can be optionally filtered.

**URL** : `/Session/Ratings/'

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

(Optional) Provide a rating filter to only bring back ratings that match in URL query parameters.
*Description* : Integer value in range of 1-5.

```text
{
    "/Session/Ratings/?Rating=4"
}
```

(Optional) Provide a rating filter operand to filter by rating in URL query parameters.

*Description* : Valid operands (encoded) [< (%3C), <= (%3C%3D) , > (%3E), >= (%3E%3D)]. Requires the Rating parameter

```text
  ?Rating=2Filter=%3C
```

(Optional) Provide a recent flag to only return the most recent 20 ratings in URL query parameters.

*Description* : 1 (true) or 0 (false)

```text
  ?Rating=2&Filter=%3C&Recent=1
```

**Header constraints** : None

## Success Responses

**Condition** : None

**Code** : `200 OK`

**Content Type** : `application/json`

**Content example** : Response will provide bulk game player ratings based on rating and filters provided.

```json

{
    [
      {
        "SessionId":"bu1sc55i7nd3mi2dbsdg",
        "PlayerId":"bu1sc55i7nd3mi2dbs8g",
        "Rating":2,
        "Comment":"It is no doubt an optimistic enterprise. But it is good for awhile to be free from the carping note that must needs be audible when we discuss our present imperfections, to release ourselves from practical difficulties and the tangle of ways and means. It is good to stop by the track for a space, put aside the knapsack, wipe the brows, and talk a little of the upper slopes of the mountain we think we are climbing, would but the trees let us see it",
        "TimeSubmitted":"2020-10-12T02:57:24Z"
    },
    {
      ...
    }
  ]
}
```

## Error Response

**Condition** : Recent flag provided was in an unexpected format.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "InvalidFlag",
    "message": "Recent parameter can only be a 0 or 1"
}
```

**Condition** : Rating Filter is not encoded properly

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "RatingFilterError",
    "message": "<server-side-error>"
}
```

**Condition** : Rating Filter provided was not one of the defined valid operands

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "InvalidRatingFilter",
    "message": "Incorrect rating filter provided. Rating filter must be one of the following: <,<=,>,>="
}
```

**Condition** : Rating was not provided in URL parameters

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "NoRatingProvided",
    "message": "Rating was not provided with Filter."
}
```

**Condition** : No Ratings have been recorded.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "status": "NoRatings",
    "message": "No ratings were found."
}
```
## Notes

* Player ID is a random 20 character string generated using the xid library

  `https://github.com/rs/xid`
