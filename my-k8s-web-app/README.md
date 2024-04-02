# My K8s Web App

This is a web application that uses the Kubernetes client to interact with the cluster. It serves a simple Vue.js frontend and provides an API to create cert-manager Certificate resources, delete Certificate resources, and download data from the created TLS secrets.

## Project Structure

```
my-k8s-web-app
├── cmd
│   └── main.go
├── api
│   ├── handler.go
│   ├── router.go
│   └── server.go
├── k8s
│   ├── client.go
│   ├── certificate.go
│   └── secret.go
├── web
│   ├── package.json
│   ├── vue.config.js
│   ├── src
│   │   ├── main.js
│   │   ├── App.vue
│   │   ├── components
│   │   │   └── HelloWorld.vue
│   │   └── assets
│   │       └── logo.png
│   └── public
│       ├── index.html
│       └── favicon.ico
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Usage

1. Build the Docker image:

   ```bash
   docker build -t my-k8s-web-app .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 my-k8s-web-app
   ```

3. Access the web app in your browser at `http://localhost:8080`.

## API Endpoints

- `POST /certificates`: Create a cert-manager Certificate resource.
- `DELETE /certificates/{name}`: Delete a cert-manager Certificate resource.
- `GET /secrets/{name}`: Download data from a TLS secret.

## Development

To run the application locally for development, follow these steps:

1. Install the dependencies for the Vue.js frontend:

   ```bash
   cd web
   npm install
   ```

2. Start the Vue.js development server:

   ```bash
   npm run serve
   ```

3. In a separate terminal, run the Go server:

   ```bash
   go run cmd/main.go
   ```

4. Access the web app in your browser at `http://localhost:8080`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.