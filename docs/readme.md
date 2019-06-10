# Sentiment API

The sentiment API is a simple REST endpoint for storage and retrieval of sentiment requests and responses for a specified group of people. This API is designed to support team feedback and to serve as a mechanism for gaining insight into current team morale.

## Resources

This API supports a number of resources for data retrieval and modification:

* [persons](#persons)
* [periods](#periods)
* [categories](#categories)
* [requests](#requests)
* [responses](#responses)

### persons

Persons (person plural for the purposes of RESTful API design best practices adherence) in the system represent the body of people from whom we wish to gather anonymous sentiment.

#### Model

**DRAFT**: more fields to be added as appropriate

```javascript
{
  "personId": "<guid>", /* primary key */
  "email": "jthomas@apptio.com", /* secondary key - used for upsert */
  "name": "Joshua Thomas"
}
```

#### Routes

* `GET /v1/persons` - retrieve all persons
* `POST /v1/persons` - create new persons
* `PUT /v1/persons` - create/update persons (upsert)
* `DELETE /v1/persons/{personId}` - delete single person
* `GET /v1/persons/{personId}` - get single person
* `PUT /v1/persons/{personId}` - update single person

### periods

A period is a timeframe for which we wish to collect sentiment.

#### periods model

```javascript
{
  "periodId": "<guid>", /* primary key */
  "requestCount": 0, /* count of requests created for sentiment */
  "responseCount": 0, /* count of responses received with sentiment */
  "startDate": "2019-05-29T07:00:00.000Z", /* composite secondary key (startDate + durationHours) */
  "durationHours": 168
}
```

#### periods routes

* `GET /v1/periods` - get periods
* `POST /v1/periods` - create new periods
* `PUT /v1/periods` - create/update periods (upsert)
* `DELETE /v1/periods/{periodId}` - delete single person
* `GET /v1/periods/{periodId}` - get single person
* `PUT /v1/periods/{periodId}` - update single person

### categories

Categories exist in order to allow more than one type of request for sentiment to be issued per period.

#### categories model

```javascript
{
  "categoryId": "<guid>", /* primary key */
}
```

#### categories routes

### requests

A request for sentiment references the period and person to whom the request is being issued. Within the request, when a response is received, the responsePeriodId is populated with a periodId, but not the exact date of the response so that correllation between the request and the person responding is not traceable when more than one request has been issued and responded to.

#### requests model

```javascript
{
  "requestId": "<guid>", /* primary key */
  "periodId": "<guid>", /* composite secondary key (periodId + personId) */
  "personId": "<guid>",
  "createdAt": "2019-05-29T07:00:00.000Z",
  "responsePeriodId":  "<guid>", /* period within which response was received */
  "data": {}
}
```

#### requests routes

* `GET /v1/persons/{personId}/requests`
* `GET /v1/periods/{periodId}/requests`
* `GET /v1/requests`
* `POST /v1/requests`
* `GET /v1/requests/{requestId}`
* `DELETE /v1/requests/{requestId}`

### responses

Responses apply to requests and refer to the period of the request.

TODO: add a grouping for requests and apply to responses so that multiple types of sentiment requests can be sent per period

#### responses model

```javascript
{
  "responseId": "<guid>", /* primary key */
  "periodId": "<guid>", /* period for which the sentiment applies */
  "createdAt": "2019-05-29T07:00:00.000Z",
  "data": {}
}
```

#### responses routes

`GET /v1/periods/{periodId}/responses`
`GET /v1/responses`

## Prerequisites

* `go` (`brew install go`)
* `dep` (`brew install dep`)
* `mongodb` (`brew install mongodb`)

### Build

```bash
git clone git@github.com:brozeph/sentimentapi.git
cd sentimentapi
dep ensure
make
```

### Test

Coming soon.
