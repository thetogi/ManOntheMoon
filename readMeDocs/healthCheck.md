# Health Check

Returns a welcome message

**URL** : `/Health-Check`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Header constraints**

- Header : `Date`
- Syntax : ```Date: <day-name>, <day> <month> <year> <hour>:<minute>:<second> GMT```

## Success Response

**Code** : `200 OK`

**Content example** : Response will provide the status of the service and the response time of the request that is sent.

```text

"Man on the Moon Game Session Review service is running normally. Response time: *Time Milliseconds*"


```text
```
