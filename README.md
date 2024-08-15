# Jooto Go WebApp Template

This is a template for a Go web application.

## Development

To run the application in development mode with hot-reload:

```
make dev
```

This will start the server on `http://localhost:8080` and automatically reload when you make changes to the code.

If this is your first time running the project, you may need to install the required tools:

```
make install-tools
```

## Building

To build the application:

```
make build
```

This will create a binary in the `bin` directory.

## Running

To build and run the application:

```
make run
```

## Testing

To run tests:

```
make test
```

## Endpoints

- `/scheduled-task`: Simulates a scheduled task
- `/webhook1`: Webhook endpoint 1
- `/webhook2`: Webhook endpoint 2
- `/health`: Health check endpoint
- `/`: Development endpoint