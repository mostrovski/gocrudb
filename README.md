# gocrudb :: an API written with Go

## A demo project showcasing usage of Go powered by the ORM and the web framework in the API context.

- Go 1.25.5

### Minimum requirements

- The application handles the following 5 operations for items in the database:
    - Getting all the inventory via `/inventory` path,
    - Getting a single item via `/inventory/{id}` path,
    - Creating an item via `/inventory` path,
    - Updating an item via `/inventory/{id}` path,
    - Deleting an item via `/inventory/{id}` path.
- Each inventory item includes:
    - ID,
    - Name,
    - Stock,
    - Price.
- Items are stored appropriately in a basic local [PostgreSQL](https://www.postgresql.org) database.
- The server returns JSON responses.
- Pagination, filtering, and sorting are enabled.
- Requests are rate limited.


### What's inside

- Maps, structs, interfaces, generics.
- Database operations with [GORM](https://gorm.io/).
- Serving and routing with [Gin](https://gin-gonic.com/).
- API documentation powered by [Stoplight Elements](https://stoplight.io/open-source/elements).

### Installation

1. Make sure you have [Go](https://go.dev/doc/install) and [PostgreSQL](https://www.postgresql.org/download) installed.
2. Clone, or download and extract the repository.
3. Change to the root of the project.
4. Run the following command to create the `.env` file:
   ```bash
   cp .env.example .env 
   ```
5. You may want to adjust the value of the `SERVICE_DB_PASSWORD` in the `.env` file, if you created one during the installation/configuration of the [PostgreSQL](https://www.postgresql.org/download).
6. Run the following command to finish the setup and start the server:  
    ```bash
    go run main.go 
    ```

### Usage

1. Head over to the http://localhost:3000 in your browser and examine the API documentation.
2. You can send API requests directly from the documentation UI.

### Kudos

The project requirements were provided as a part of the [Backend Developer with Go](https://www.udacity.com/course/backend-with-postgres-and-go--nd810) Nanodegree Program at Udacity.