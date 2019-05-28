# Sentiment API

## Resources

### persons

Persons (person plural for the purposes of RESTful API design best practices adherence) in the system represent the body of people from whom we wish to gather anonymous sentiment.

#### Model

**DRAFT**: more fields to be added as appropriate

```json
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

* `GET /v1/periods` - get periods
* `POST /v1/periods`
*

### requests

`GET /v1/persons/{personId}/requests`
`GET /v1/periods/{periodId}/requests`
`GET /v1/requests`
`POST /v1/requests`

#### individual requests

`GET /v1/requests/{requestId}`
`DELETE /v1/requests/{requestId}`

### responses

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