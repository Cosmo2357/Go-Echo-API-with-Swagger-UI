![alt text](assetsForReadme/6.png)
# Go Echo API with Swagger UI with GitHub Actions
This is a simple REST API built with Go using the Echo framework and Swagger UI for API documentation.

## Project Structure

```
myapp/
├── cmd/
│   └── main.go            # Entry point of the application
├── config/
│   └── config.go          # Configuration settings
├── docs/
│   └── docs.go            # Swagger docs setup
├── internal/
│   ├── handler/
│   │   └── user.go        # User handler
│   └── model/
│       └── user.go        # User model
├── routes/
│   └── routes.go          # Routes setup
├── go.mod                 # Go module file
└── go.sum                 # Go dependencies file
```

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Swag](https://github.com/swaggo/swag) for generating Swagger documentation

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/myapp.git
   cd myapp
   ```

2. Initialize Go modules:
   ```sh
   go mod tidy
   ```

3. Install Swag for generating Swagger documentation:
   ```sh
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. Generate Swagger documentation:
   ```sh
   swag init -g cmd/main.go
   ```

## Running the Application

To run the application, use the following command:

```sh
go run cmd/main.go
```

The server will start at `http://localhost:8080`. You can access the Swagger UI for API documentation at:

```
http://localhost:8080/swagger/index.html
```
![alt text](assetsForReadme/swagger.png)
## API Endpoints

- `GET /users/:id` - Retrieve user details by ID

## Configuration

The application uses environment variables for configuration. You can create a `.env` file in the root directory to specify these variables.

## Project Dependencies

- [Echo](https://github.com/labstack/echo) - High performance, extensible, minimalist Go web framework
- [Swaggo](https://github.com/swaggo/echo-swagger) - Swagger integration for Echo
- [Godotenv](https://github.com/joho/godotenv) - Loads environment variables from `.env`.

## OAuth2　for the generated html
```html
<script>

const ui = SwaggerUIBundle({
  url: "./swagger.json",
  dom_id: '#swagger-ui',
  presets: [
    SwaggerUIBundle.presets.apis,
    SwaggerUIStandalonePreset
  ],
  layout: "StandaloneLayout",
  oauth2RedirectUrl: "https://yourapp.com/oauth2-redirect.html",  // Specify redirect URL
  authActionsAuthorizeOauth2: {
    useBasicAuthenticationWithAccessCodeGrant: true  // Use Basic Auth if necessary
  },
  // OAuth2 configuration
  deepLinking: true,
  authAction: {
    authorizeOauth2: {
      client_id: "YOUR_CLIENT_ID",
      client_secret: "YOUR_CLIENT_SECRET",  // Required for 'authorization_code' flow
      authorizationUrl: "https://oauthprovider.com/oauth/authorize",
      tokenUrl: "https://oauthprovider.com/oauth/token",
      scopes: {
        "read": "Read access",
        "write": "Write access"
      }
    }
  }
});

// Function to initialize OAuth authentication
ui.initOAuth({
  clientId: "YOUR_CLIENT_ID",
  clientSecret: "YOUR_CLIENT_SECRET",  // Required if necessary
  realm: "YOUR_REALM",
  appName: "Swagger UI",
  scopeSeparator: " ",
  scopes: "read write",
  additionalQueryStringParams: {}
});

</script>
```

## Best practice?
- Generate the Swagger file on each PUSH using GitHub Actions, save it in secure storage, and access it from a separate server-side rendered (SSR) application with authentication (such as Remix or Next.js) for secure display.

https://github.com/jellydn/next-swagger-doc

or Hono with Oauth or basic auth

```javascript
$ npm create hono@latest my-app
// checke here for the installation https://hono.dev/docs/getting-started/basic
$ npm run dev
// deploy to Cloudflare Workers. easy to deploy to other services. just check doc
$ npm run deploy

# index.js
import { Hono } from 'hono'
import { swaggerUI } from '@hono/swagger-ui'

const app = new Hono()

// Use the middleware to serve Swagger UI at /ui
app.get('/ui', swaggerUI({ url: '/doc' }))

export default app


```

You can specify the backend server in main.go. Then it will be written in swagger.json or swagger.And swagger UI will user the backend server that is written in swagger.yaml or json.

```go
# main.go
  // @title My API Example
  // @version 1.0
  // @description This is a sample server for demonstrating Go Echo with Swagger.
  // @host localhost:8080 ..
  // @BasePath /
```

## Actions for deploying Swagger documentation to S3
```yml
name: Deploy Swagger Documentation to S3

on:
  push:
    branches:
      - main
      - dev

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      # Step 1: Check out the code
      - name: Check out code
        uses: actions/checkout@v2

      # Step 2: Set up Go environment
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: '1.22.0'

      # Step 3: Install Swag for generating Swagger documentation
      - name: Install Swag
        run: |
          go install github.com/swaggo/swag/cmd/swag@latest
          echo "$HOME/go/bin" >> $GITHUB_PATH

      # Step 4: Generate Swagger documentation
      - name: Generate Swagger Docs
        run: |
          swag init -g cmd/main.go -o docs/swagger-ui

      # Step 5: Upload Swagger JSON to S3
      - name: Upload Swagger JSON to S3
        env:
          AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
          AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          AWS_REGION: ${{ secrets.AWS_REGION }}
        run: |
          # Set S3 bucket based on branch
          if [ "${GITHUB_REF}" == "refs/heads/main" ]; then
            S3_BUCKET="your-prod-s3-bucket-name"
          elif [ "${GITHUB_REF}" == "refs/heads/dev" ]; then
            S3_BUCKET="your-dev-s3-bucket-name"
          else
            echo "Unsupported branch"
            exit 1
          fi

          # Install AWS CLI
          sudo apt-get update
          sudo apt-get install -y awscli

          # Upload the Swagger JSON file to the specified S3 bucket
          aws s3 cp docs/swagger-ui/swagger.json s3://$S3_BUCKET/swagger.json --acl private
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to submit issues, fork the repository and send pull requests!


