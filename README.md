# Go API with Web Scraping

This Go API is designed to aggregate RSS feeds and store their posts in a database. It provides RESTful endpoints for various operations and includes a web scraping component to fetch posts from feeds. This README provides an overview of the application's structure and usage.

## Table of Contents

- [Go API with Web Scraping](#go-api-with-web-scraping)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Endpoints](#api-endpoints)
  - [Web Scraping](#web-scraping)
  - [Contributing](#contributing)
  - [License](#license)

## Prerequisites

Before using this API, ensure you have the following prerequisites installed on your system:

- Go (1.16 or later)
- PostgreSQL (for database)
- [chi](https://pkg.go.dev/github.com/go-chi/chi/v5) - Router library
- [github.com/joho/godotenv](https://pkg.go.dev/github.com/joho/godotenv) - Library for loading environment variables
- [github.com/lib/pq](https://pkg.go.dev/github.com/lib/pq) - PostgreSQL driver
- Other dependencies as specified in `go.mod`

## Installation

1. Clone this repository:

   ```shell
   git clone https://github.com/AlanCanalesM/rss-aggregator
    ```
2. Navigate to the project directory:

    ```shell
    cd rss-aggregator
    ```
3. Install the Go dependencies:
   
    ```shell
    go mod tidy
    ```
## Configuration
This application uses environment variables for configuration. You can set these variables by creating a .env file in the project root with the following content:
    
```shell
PORT=8080
DB_URL=postgres://username:password@localhost/database_name?sslmode=disable    
```
PORT: Port on which the API server will run.

DB_URL: PostgreSQL database URL.

## Usage
To start and build the API server and begin web scraping, run the following command:
    
```shell
go build && ./rss-aggregator
```

The API server will start on the specified port, and the web scraping process will begin in the background.

## API Endpoints

This API provides the following endpoints:

- GET /v1/health: Health check endpoint.
- GET /v1/error: Endpoint to trigger an error for testing purposes.
- GET /v1/feeds: Retrieve a list of feeds.
- GET /v1/users: Retrieve user information(requires authentication).
- GET /v1/feed_follows: Retrieve feed follow relationships (requires authentication).
- GET /v1/posts: Retrieve posts for a user (requires authentication).
- GET /v1/allPosts: Retrieve all posts.
- POST /v1/users: Create a new user.
- POST /v1/feeds: Create a new feed (requires authentication).
- POST /v1/feed_follows: Create a new feed follow relationship (requires authentication).
- DELETE /v1/feed_follows/{feedFollowID}: Delete a feed follow relationship (requires authentication).
  
For more details on each endpoint and their request/response formats, refer to the API code.

## Web Scraping
This application includes a web scraping component that periodically fetches posts from RSS feeds and stores them in the database. The web scraping process is initiated when the API server starts and runs in the background. You can configure the scraping frequency and other parameters in the startScraping function.

## Contributing
If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them.
4. Create a pull request with a clear description of your changes.

## License
This project is licensed under the MIT License. You are free to use and modify it as per the license terms.