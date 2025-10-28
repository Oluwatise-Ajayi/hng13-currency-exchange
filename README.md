# hng13 Currency Exchange API

## Overview
This project is a Go-based RESTful API built with the Gin framework, designed to provide comprehensive country information and currency exchange rates. It leverages GORM for PostgreSQL database interactions, integrating with external APIs to fetch and cache global country data, including population, capital, region, and dynamic exchange rates.

## Features
- `Gin Framework`: High-performance HTTP web framework for building robust APIs in Go.
- `GORM ORM`: Elegant and developer-friendly ORM for Go, simplifying database operations with PostgreSQL.
- `PostgreSQL`: A powerful, open-source object-relational database system used for reliable data storage.
- `RESTful API Design`: Implements a clear and intuitive API structure for interacting with country and currency data.
- `External API Integration`: Seamlessly fetches and processes real-time country data from [restcountries.com](https://restcountries.com) and exchange rates from [open.er-api.com](https://open.er-api.com).
- `Dynamic Data Refresh`: Provides an endpoint to refresh and update cached country and exchange rate data in the database.
- `Data Filtering & Sorting`: Supports advanced querying for countries based on region, currency, and various sorting options (e.g., GDP, population).

## Getting Started

### Installation
To get a local copy up and running, follow these steps.

1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/Oluwatise-Ajayi/hng13-currency-exchange.git
    cd hng13-currency-exchange
    ```

2.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Set up PostgreSQL Database:**
    Ensure you have a PostgreSQL instance running. Create a new database for the project (e.g., `currency_exchange_db`).

4.  **Run Migrations:**
    The application will automatically perform database migrations (creating the `country_infos` table) upon startup.

### Environment Variables
Create a `.env` file in the root directory of the project and populate it with the following environment variables:

```env
PORT=8080
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=currency_exchange_db
DB_PORT=5432
DB_SSLMODE=disable
```

## Usage
After setting up your environment variables and installing dependencies, you can run the application.

1.  **Start the Application:**
    ```bash
    go run main.go
    ```
    Alternatively, if you have `air` installed (recommended for development for hot reloading):
    ```bash
    air
    ```

2.  **Access the API:**
    The API will be available at `http://localhost:[PORT]`, where `[PORT]` is the value specified in your `.env` file (e.g., `http://localhost:8080`).

    *   **Refresh Country Data:** Before querying, it's recommended to refresh the data:
        ```bash
        curl -X POST http://localhost:8080/countries/refresh
        ```
        This will fetch country and exchange rate data from external APIs and cache them in your database.

    *   **Retrieve All Countries:**
        ```bash
        curl http://localhost:8080/countries
        ```

    *   **Filter Countries by Region:**
        ```bash
        curl http://localhost:8080/countries?region=Africa
        ```

    *   **Sort Countries by GDP (Descending):**
        ```bash
        curl http://localhost:8080/countries?sort=gdp_desc
        ```

    *   **Get a Specific Country by Name:**
        ```bash
        curl http://localhost:8080/countries/Nigeria
        ```

    *   **Check API Status:**
        ```bash
        curl http://localhost:8080/status
        ```

## API Documentation

### Base URL
`http://localhost:[PORT]` (e.g., `http://localhost:8080`)

### Endpoints

#### GET /
**Request**:
No parameters required.

**Response**:
```json
{
  "message": "Welcome to hng13 Currency Exchange API"
}
```

**Errors**:
- `N/A`: Always returns a 200 OK.

#### POST /countries/refresh
Refreshes the country and exchange rate data in the database by fetching from external APIs.

**Request**:
No payload required.

**Response**:
```json
{
  "message": "Countries refreshed successfully"
}
```

**Errors**:
- `503 Service Unavailable`: External data source (restcountries.com or open.er-api.com) could not be reached or processed.
  ```json
  {
    "error": "External data source unavailable",
    "details": "Could not fetch data from restcountries API"
  }
  ```

#### GET /countries
Retrieves a list of all cached countries, with optional filtering and sorting.

**Request**:
Query Parameters:
- `region`: (string, optional) Filter by country region (e.g., `Africa`).
- `currency`: (string, optional) Filter by currency code (e.g., `NGN`).
- `sort`: (string, optional) Sort order.
  - `gdp_desc`: Sort by estimated GDP in descending order.
  - `gdp_asc`: Sort by estimated GDP in ascending order.
  - `population_desc`: Sort by population in descending order.
  - `population_asc`: Sort by population in ascending order.
  - `name_asc`: Sort by country name in ascending order.
  - `name_desc`: Sort by country name in descending order.

**Response**:
```json
[
  {
    "id": "1",
    "name": "Nigeria",
    "capital": "Abuja",
    "region": "Africa",
    "population": 206139587,
    "currency_code": "NGN",
    "exchange_rate": 870.0,
    "estimated_gdp": 500000000000.0,
    "flag_url": "https://restcountries.com/data/nga.svg",
    "last_refreshed_at": "2024-07-30T10:00:00Z"
  },
  {
    "id": "2",
    "name": "Ghana",
    "capital": "Accra",
    "region": "Africa",
    "population": 31072940,
    "currency_code": "GHS",
    "exchange_rate": 14.0,
    "estimated_gdp": 70000000000.0,
    "flag_url": "https://restcountries.com/data/gha.svg",
    "last_refreshed_at": "2024-07-30T10:00:00Z"
  }
]
```
Returns an empty array `[]` if no countries match the criteria.

**Errors**:
- `500 Internal Server Error`: An unexpected error occurred while querying the database.
  ```json
  {
    "error": "Internal server error"
  }
  ```

#### GET /countries/:name
Retrieves detailed information for a single country by its name.

**Request**:
Path Parameter:
- `name`: (string, required) The full name of the country.

**Response**:
```json
{
  "id": "1",
  "name": "Nigeria",
  "capital": "Abuja",
  "region": "Africa",
  "population": 206139587,
  "currency_code": "NGN",
  "exchange_rate": 870.0,
  "estimated_gdp": 500000000000.0,
  "flag_url": "https://restcountries.com/data/nga.svg",
  "last_refreshed_at": "2024-07-30T10:00:00Z"
}
```

**Errors**:
- `400 Bad Request`: No country name provided in the path.
  ```json
  {
    "error": "Validation failed",
    "details": {
      "name": "is required"
    }
  }
  ```
- `404 Not Found`: No country found with the given name.
  ```json
  {
    "error": "Country not found"
  }
  ```
- `500 Internal Server Error`: An unexpected error occurred while querying the database.
  ```json
  {
    "error": "Internal server error"
  }
  ```

#### DELETE /countries/:name
Deletes a country record from the database by its name.

**Request**:
Path Parameter:
- `name`: (string, required) The full name of the country to delete.

**Response**:
```json
{
  "message": "Country deleted successfully"
}
```

**Errors**:
- `400 Bad Request`: No country name provided in the path.
  ```json
  {
    "error": "Validation failed",
    "details": {
      "name": "is required"
    }
  }
  ```
- `404 Not Found`: No country found with the given name to delete.
  ```json
  {
    "error": "Country not found"
  }
  ```
- `500 Internal Server Error`: An unexpected error occurred while deleting from the database.
  ```json
  {
    "error": "Internal server error"
  }
  ```

#### GET /status
Provides a summary of the database, including the total number of cached countries and the timestamp of the last data refresh.

**Request**:
No parameters required.

**Response**:
```json
{
  "total_countries": 10,
  "last_refreshed_at": "2024-07-30T10:00:00Z"
}
```
If no countries are in the database, `last_refreshed_at` will be an empty string.

**Errors**:
- `500 Internal Server Error`: An unexpected error occurred while querying the database.
  ```json
  {
    "error": "Internal server error"
  }
  ```

#### GET /countries/image
Placeholder endpoint for serving a summary image. Currently returns an error.

**Request**:
No parameters required.

**Response**:
```json
{
  "error": "Summary image not found"
}
```

**Errors**:
- `404 Not Found`: This endpoint is currently unimplemented and will always return this error.

---

## Technologies Used
| Technology        | Description                                     | Link                                              |
| :---------------- | :---------------------------------------------- | :------------------------------------------------ |
| **Go (Golang)**   | Primary programming language                    | [golang.org](https://golang.org)                  |
| **Gin Framework** | High-performance HTTP web framework             | [gin-gonic.com](https://gin-gonic.com)            |
| **GORM**          | Elegant ORM for Go                              | [gorm.io](https://gorm.io)                        |
| **PostgreSQL**    | Robust relational database system               | [postgresql.org](https://www.postgresql.org)      |
| **`godotenv`**    | Loads environment variables from `.env` files   | [github.com/joho/godotenv](https://github.com/joho/godotenv) |
| **Air**           | Live reload for Go applications (development)   | [github.com/cosmtrek/air](https://github.com/cosmtrek/air) |

## Contributing
We welcome contributions to enhance this project! To contribute:

1.  **Fork the repository.** 🍴
2.  **Create a new branch** for your feature or bug fix: `git checkout -b feature/your-feature-name`.
3.  **Make your changes** and ensure the code adheres to the project's style guidelines.
4.  **Write clear, concise commit messages.**
5.  **Push your branch** to your forked repository.
6.  **Open a Pull Request** against the `main` branch of this repository, providing a detailed description of your changes.

## License
This project is licensed under the [LICENSE_NAME] License.

## Author Info
- **Oluwatise-Ajayi**
  - LinkedIn: [linkedin.com/in/Oluwatise-ajayi](https://www.linkedin.com/in/oluwatise-ajayi-26697b1ba/)
  - Twitter: #####

---

### Badges
[![Go Version](https://img.shields.io/badge/Go-1.25.0-00ADD8?logo=go)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-00ADD8?logo=gin)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-v1.31.0-E91E63?logo=gorm)](https://gorm.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791?logo=postgresql)](https://www.postgresql.org/)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)](https://example.com/build-status)
[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)