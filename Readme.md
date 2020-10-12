# Man on the Moon Review Service ![Moon Man](moon_man.jpg)

Hello, welcome to the Man on the Moon Game Review Service by Derek Askham!

This service has several RESTful API endpoints that allow:

- Players of the online game "Man on the Moon" to submit feedback for their last game session.
- Members of an operations team to view and filter the feedback.
- Users can give their session a rating on a scale of 1 to 5, and can leave a comment.
- Multiple players can rate the same session, but each player can only rate a given session once.

# Instructions to set up service

## Requirements
- Goland or similar IDE
- MySQL Server 8.0 with user credentials that has at a minimum read and write access.
- MySQL Service running on `localhost:3306`
- ManOnTheMoon schema and tables script
- Docker

## Deployment Files
  `{project-root}/DockerFile`

  `{project-root}/Docker-Compose.yaml `


## Test Files

  `{project-root}/controllers/controllers_test.go`

  *Description*: Runs tests on each API endpoint

## Testing

    docker-compose run man-on-the-moon go test -v ./...

## Install & Deploy
  1. Clone or download project to local machine

  `git clone <path-to-repo>`

  2. If you haven't already, install MySQL Server 8.0 on your host machine. Using a user with create permissions, run the provided script `ManOnTheMoon.sql` to build the database structures and populate sample data. The script is found in the root directory of the go project.

  3. Update the USER and PASSWORD environment variables in `dbConfig.env.dev` in the root directory of the go project with a user from you instance of MySQL with read and write permissions. You will not be able to connect to MySQL if this step isn't done.

  4. You can re-build the pre-built project executable or skip and continue to step 5.
    go build .

  5. Use the below Docker command to build and deploy the service on port `:8080`.

    `docker-compose --env-file ./.env up`

  6. Use a browser, such as Chrome, or your preferred RESTful endpoint testing tools for testing the defined endpoints referenced below.

# Assumptions
- Users can only submit feedback on a completed game session. If the game connection was interrupted, no feedback would be prompted.
- Ratings are used by operations to monitor and report app health on a regular basis, but perhaps not used as the sole source of app health in real time.

# REST API

* [Show info](readMeDocs/home.md) : `GET /`
* [Show info](user/get.md) : `GET /api/user/``
Read more [here](API Endpoint Documentation/Home.md) # It works!

## Get All Ratings
[See Get All Ratings documentation](readMeDocs/getAllRatings.md)

### Request

`GET /Session/Ratings/`

    curl -i "http://localhost:8080/Session/Ratings/"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:02:28 GMT
    Transfer-Encoding: chunked

    []

### Request (Optional Filters)

`GET /Session/Ratings/?{params}`

    curl -i "http://localhost:8080/Session/Ratings/?Rating=3&Recent=1"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:02:28 GMT
    Transfer-Encoding: chunked

    []

## Create a new Rating
[See create a new rating documentation](readMeDocs/createRating.md)
### Request

`POST /Session/{SessionId}/CreateRating/{params}`

    curl -i -X POST "http://localhost:8080/Session/bu1sc5di7nd3mi2dbua0/CreateRating?PlayerId=bu1sc5di7nd3mi2dbu9g&Rating=3&Comment=TestComment"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 03:04:15 GMT
    Content-Length: 115

    {"Status":"OK","Message":"Rating Successfully submitted for Session ID: bu1sc5di7nd3mi2dbua0 rating: 3 Good game moon man!"}

## Get a rating
[See get rating documentation](readMeDocs/getRating.md)
### Request

`GET /Session/{SessionId}/Rating{params}`

    curl -i "http://localhost:8080/Session/bu1sc55i7nd3mi2dbs90/Rating?PlayerdId=bu1sc55i7nd3mi2dbs8g"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:18:20 GMT
    Content-Length: 141

    {"Status":"SessionRatingNotFound","Message":"Could not find rating by player for session using PlayerId: and SessionId bu1sc55i7nd3mi2dbs90"}

## Create a Player
[See create a player documentation](readMeDocs/createPlayer.md)
### Request

`POST /Player/Create`

    curl -i -X POST "http://localhost:8080/Player/Create"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:20:59 GMT
    Content-Length: 85

    {"Status":"OK","Message":"New Player Successfully created. ID: bu1tjavr3qbtb1m73dn0"}

## Find a Player
[See get a player documentation](readMeDocs/getPlayerById.md)
### Request

`POST /Player/{PlayerId}`

    curl -i "http://localhost:8080/Player/bu1sc9ti7nd4r54i04h0"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:46:08 GMT
    Content-Length: 99

    {"PlayerId":"bu1sc9ti7nd4r54i04h0","Name":"Anthony Taylor","TimeRegistered":"2020-10-12T02:57:43Z"}

## Get All Players
[See get all players documentation](readMeDocs/getAllPlayers.md)
### Request

`GET /Players/

    curl -i "http://localhost:8080/Players/

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:23:59 GMT
    Transfer-Encoding: chunked

    ["PlayerId":"bu1sc85i7nd28g415fsg","Name":"Elizabeth Davis","TimeRegistered":"2020-10-12T02:57:36Z"}, ...]

## Find a Session
[See get a session documentation](readMeDocs/getSessionById.md)
### Request

`GET /Session/{SessionId}{params}`

    curl -i "http://localhost:8080/Session/bu1sc55i7nd3mi2dbs90?PlayerId=bu1sc55i7nd3mi2dbs8g"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:29:14 GMT
    Content-Length: 110

    {"SessionId":"bu1sc55i7nd3mi2dbs90","PlayerId":"bu1sc55i7nd3mi2dbs8g","TimeSessionEnd":"2020-10-12T02:57:24Z"}

## Create a Session
[See create a session documentation](readMeDocs/createSession.md)
### Request

`POST /Session/Create{params}`

    curl -i -X POST "http://localhost:8080/Session/Create?PlayerId=bu1sc55i7nd3mi2dbs8g"

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Mon, 12 Oct 2020 04:32:38 GMT
    Content-Length: 86

    {"Status":"OK","Message":"New Session Successfully created. ID: bu1topnr3qbtb1m73dng"}

## Get All Sessions
[See get all sessions documentation](readMeDocs/getAllSessions.md)
### Request

`GET /Sessions/`

    curl -i "http://localhost:8080/Sessions/"

### Response

HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 12 Oct 2020 04:34:58 GMT
Transfer-Encoding: chunked

[{"SessionId":"bu1sc55i7nd3mi2dbs90","PlayerId":"bu1sc55i7nd3mi2dbs8g","TimeSessionEnd":"2020-10-12T02:57:24Z"}, ...]

## Tools

[See Home documentation](readMeDocs/home.md)

[See Health-Check sessions documentation](readMeDocs/healthCheck.md)

[See Get Session Rating Random documentation](readMeDocs/gameSessionRating.md)
