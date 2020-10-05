## Koho Balance - Project developed in Golang (backend) and React Js with Typescript (back-end).

This is a simple example of implementing a REST API developed in Golang and the frontend developed in React JS with
typescript, salving some data in MongoDB

As an information, this was the first time that I developed something in Golang and probably the structure used does not follow the best practices. However, I believe it can be a good source of research. Feel free to contribute in any way.

### Technologies, languages, frameworks

Backend - Golang with Echo Framework. Used swagger for API documentation (https://github.com/swaggo/swag).
Frontend - React Js with Typescript. Used Eslint and Prettier for code standardization.
Database - Used MongoDB (nosql).
Dockerfile and Docker Compose

The project can be started in an integrated way, using the docker for execution:

Backend and Frontend can be run separately, as long as there is a MongoDB server available, or even a docker container with the MongoDB image.

## Execution with Docker

- In the root folder
```
docker-compose up
```
After the server starts, the api documentation will be available at:

[http://localhost:3333/swagger/index.html](http://localhost:3333/swagger/index.html)

## Separate execution:

### Backend execution
- In the backend folder

```
go run main.go
```

To perform the tests, run:
```
go test ./...
```

### Frontend execution
- Na pasta frontend
```
yarn start
```
After the start, access the system at:

[http://localhost:3000](http://localhost:3000)


## About the challenge

The test consist of the following rules:

In finance, it's common for accounts to have so-called "velocity limits". In this task, you'll write a program that accepts or declines attempts to load funds into customers' accounts in real-time.

Each attempt to load funds will come as a single-line JSON payload, structured as follows:

```json
{
  "id": "1234",
  "customer_id": "1234",
  "load_amount": "$123.45",
  "time": "2018-01-01T00:00:00Z"
}
```

Each customer is subject to three limits:

- A maximum of $5,000 can be loaded per day
- A maximum of $20,000 can be loaded per week
- A maximum of 3 loads can be performed per day, regardless of amount

As such, a user attempting to load $3,000 twice in one day would be declined on the second attempt, as would a user attempting to load $400 four times in a day.

For each load attempt, you should return a JSON response indicating whether the fund load was accepted based on the user's activity, with the structure:

```json
{ "id": "1234", "customer_id": "1234", "accepted": true }
```

You can assume that the input arrives in ascending chronological order and that if a load ID is observed more than once for a particular user, all but the first instance can be ignored. Each day is considered to end at midnight UTC, and weeks start on Monday (i.e. one second after 23:59:59 on Sunday).

Your program should process lines from `input.txt` and return output in the format specified above, either to standard output or a file. Expected output given our input data can be found in `output.txt`.

You're welcome to write your program in a general-purpose language of your choosing, but as we use Go on the back-end and TypeScript on the front-end, we do have a preference towards solutions written in Go (back-end) and TypeScript (front-end).

We value well-structured, self-documenting code with sensible test coverage. Descriptive function and variable names are appreciated, as is isolating your business logic from the rest of your code.
