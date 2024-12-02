
# Sample Go Project

This project is an implementation of an API for domain management using the Go programming language and the lightweight `chi` router.

## Technologies Used

- **chi Router**: A lightweight and fast router for Go, which allows for efficient handling of HTTP requests.
- **API Integration**: This project integrates with the [Simtel's Picast Laravel](https://github.com/Simtel/picast-laravel/tree/master) API for domain management.

## Features

- 🌐 **View Domain List**: Users can see all available domains.
- 🔍 **View Domain Details**: Ability to retrieve information about a specific domain.
- ➕ **Create New Domain**: Functionality for adding new domains to the system.

## Installation and Usage

Follow these steps to set up and run the project on your local machine:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Simtel/go-project
   ```

2. **Copy the Configuration File for Local Development**:
   ```bash
   cp .env .env.local
   ```
   This file will be used to store your environment settings.

3. **Start the Containers**:
   To bring up the required containers, run:
   ```bash
   make up
   ```

4. **Create the Database in Adminer**:
   Once the containers are running, use Adminer to create the database.

5. **Run Migrations**:
   Apply migrations to set up the database:
   ```bash
   make migrate
   ```

6. **Run the Project**:
   To start the project, use the following command:
   ```bash
   make run
   ```
   Alternatively, run the project in watch mode for automatic restarts upon code changes:
   ```bash
   make watch
   ```

7. **Run Tests**:
   To execute the tests, run:
   ```bash
   make test
   ```

## API Endpoints

- `GET /domains`: Retrieves a list of all domains.
- `GET /domains/{id}`: Retrieves details of a specific domain by ID.
- `POST /domains`: Creates a new domain.

## Contributing

Feel free to contribute by creating issues or pull requests. To start contributing, please follow these steps:

1. Fork the project.
2. Create a feature branch.
3. Make your changes.
4. Commit your changes.
5. Push to the branch.
6. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Feel free to modify any section according to your specific needs or details about your project!