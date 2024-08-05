
# Bitcoin Runes POC

This project is a proof of concept (POC) for Bitcoin transactions using the BlockCypher API, with data storage in a PostgreSQL database using Docker Compose.

## Requirements

- [Go](https://golang.org/dl/) 1.16 or higher
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/Chxpz/bitcoin-runes-poc.git
   cd bitcoin-runes-poc
   ```

2. Set up environment variables in a `.env` file:

   ```sh
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=runes
   ```

3. Start the Docker Compose services:

   ```sh
   docker-compose up -d
   ```

4. Install Go dependencies:

   ```sh
   go get github.com/jackc/pgx/v4/pgxpool
   go get github.com/joho/godotenv
   ```

5. Run the application:

   ```sh
   go run main.go
   ```

## Project Structure

```sh
.
├── docker-compose.yml
├── .env
├── db.go
├── etch
│   ├── etch.go
├── mint
│   ├── mint.go
└── main.go
```

- `docker-compose.yml`: Docker Compose configuration for PostgreSQL.
- `.env`: Environment variable configuration file.
- `db.go`: Code to initialize the database and create the `runes_poc` table.
- `etch/etch.go`: Code to create "etch" transactions.
- `mint/mint.go`: Code to send "mint" transactions.
- `main.go`: Main function that coordinates the creation and sending of transactions and saves the results in the database.

## PostgreSQL Commands

Access the PostgreSQL database:

```sh
docker exec -it runes_postgres psql -U postgres -d runes
```

Query the `runes_poc` table:

```sql
SELECT * FROM runes_poc;
```