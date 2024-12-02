# Example Project in Go

## Utilized Technologies
- **chi Router** - a lightweight and fast router for Go.
- **Project API** from [Simtel's Picast Laravel](https://github.com/Simtel/picast-laravel/tree/master) - API integration for managing domains.

## Features
- View a list of domains 🌐
- View details of a specific domain 🔍
- Create a new domain ➕

## Installation and Usage

1. **Clone the project**:
```bash
   git clone https://github.com/Simtel/go-project
```

2. Copy the configuration file for local development:

```bash
cp .env .env.local
```

3. Up containers
```bash
make up
```

4. Create database db in Adminer

5. Run migrations
```bash
make migrate
```

6. Run the project:You can run the project using the following command:
```bash
make run
```
Or use the watch mode for automatic restarts on code changes:
```bash
make watch
```

7. Run tests:To execute tests in the project, use the following command:
```bash
make test
```

Contributing
If you'd like to contribute to the project, please create a pull request or open an issue in the repository. We welcome any suggestions and improvements!

License
This project is licensed under the MIT License. See the LICENSE file for details.