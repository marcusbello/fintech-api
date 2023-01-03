# Fintech Bank API
Welcome to My Fintech Bank Api, a financial management application written in Go!, following uncle BOb's Domain Driven Design and Couchbase database.

## Features
- Open a Bank account
- Deposit money from an external payment gateway
- Send money
- Receive money
- Track transactions

## Getting Started
1. Install Go on your local machine if you don't already have it (instructions can be found [here](https://www.example.com)).
2. Also, install Docker on your local machine if you don't already have it (instructions can be found [here](https://www.example.com)).
3. Clone this repository and navigate to the root directory of the project.
4. Use docker-compose to build and launch the app.
```bash
docker compose up -d
```
5. If you have Make installed on you server, run
```bash
make build
```
to build the app and then run it with:
```bash
make run
```

## Configuration
You can configure certain aspects of the app by modifying the values in the `docker-compose.yml` file in the root directory of the project.
The file is self-documented, so you can see what each value does.

## Usage
After launching the app, you will be able to access it from your localhost on port 3030, test with:
```bash
curl -g http://localhost/3030/api/v1/ping
```
* NB: Edit the default port (:3030) to another port of your choice following the configuration.

## Documentation

Official OpenAPI documentation is available after launching the app on localhost, visit
```bash
http://localhost:3030/swagger/index.html
```
on any browser.

| Endpoints                                            | Description                      | Requirements (JSON Body) |
|------------------------------------------------------|----------------------------------|-------|
| POST http://localhost:3030/api/v1/register           | User register                    | {"user_name":"marcus", "password":"marcusPassword" "email":"marcus@email.com",} |
| POST http://localhost:3030/api/v1/signin             | User login                       | {"user_name":"marcus", "password":"marcusPassword"} |
| GET http://localhost:3030/api/v1/marcus/user         | Marcus's Profile                 | empty |
| POST http://localhost:3030/api/v1/marcus/transfer    | Transfer money                   | {"to": "another_user", "amount":20} |
| POST http://localhost:3030/api/v1/marcus/deposit     | Deposit money (work in progress) | {"amount":20} |
| GET http://localhost:3030/api/v1/marcus/transactions | Marcus's transaction history     | empty |

## Tests
To run the tests for the app, navigate to the root directory of the project and run:
```bash
go test ./...
```

## Dependencies
This project has the following dependencies:

- Gin HTTP library
- Couchbase Database package

## Contributing
If you would like to contribute to the project, please fork the repository and create a pull request with your changes on a new branch. Thank you for considering contributing!

## License
This project is licensed under the [MIT](https://en.wikipedia.org/wiki/MIT_License) License. See the [LICENSE](https://github.com/marcusbello/fintech-api/blob/master/LICENSE) file for more details.

## Author

* **Marcus Bello** - *Software Engineer* - [Marcus Bello](https://github.com/marcusbello/)

