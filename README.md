<img src="./screenshots/1.png" width=200px display="block" >

# Authentication  Web Client

> Work in progress - Simple web-client for communication with gRPC authentication microservice written in Go.

- **[Web Client](https://dev.rami.im/mxweb/index/)**

- **[gRPC Service](https://github.com/rumsrami/grpc-token-service)**

## How it works

1. Client redirects the browser to LinkedIn's OAuth 2.0 authorization page where the member authenticates.
2. After authentication, LinkedIn's authorization server passes an authorization code back to the client.
3. Client sends this code to LinkedIn and LinkedIn returns an access token.
4. Client sends this token to the gRPC authentication microservice
5. Service saves this token & creates another one and sends it back to the client.
6. Client uses new token to communicate with other services

## Objectives
1. Learn how oauth2 works with different providers and Go.
2. Learn how to integrate the below technologies in one unit.

> * Docker multistage builds, scratch images with Go binaries.
> * TLS communication between client and oauth2 provider.
> * Google Cloud Build
> * Google Container Registry
> * Google Compute Engine
> * CI/CD with Github, Docker, Cloud.yaml
> * Encryption with Google Keys

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](http://opensource.org/licenses/mit-license.php)**
