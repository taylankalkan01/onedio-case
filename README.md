# onedio-case

I stumbled upon ONEDIO's backend-developer case study PDF and decided to give it a shot.

## Run Locally

Clone the project

```bash
git clone https://github.com taylankalkan01/onedio-case
```

## Run Redis Stack on Docker

```bash
docker run -d --name onedio-redis -p 6379:6379 redis/redis-stack-server:latest
```

## Go to the project directory

### Terminal 1

```bash
  cd onedio-case
  cd api
```

```bash
  npm install
```

```bash
  npm run dev
```

### Terminal 2

```bash
  cd onedio-case
  cd cli
```

```bash
  go mod tidy
```

```
go run . parseAndSave ../data/1819-E0.csv
```

```
go run . parseAndSave ../data/1718-E0.csv
```

## Usage/Examples and API Reference

I built a command-line app using Golang. It lets you read data from a file and save it to both Redis and MongoDB databases.

### API Reference

#### Get all fixtures

```http
GET
http://localhost:3000/fixtures?limit=15&page=1
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`MONGO_URL`

`APP_PORT`

## Feedback

If you have any feedback, please reach out to me.

## Authors

- [@taylankalkan01](https://github.com/taylankalkan01)
