# Go Project 
This project is for learning Golang. I used microservice architecture in this project just for better understanding how to use microservices in golang.

## Authentication
For authentication this project uses OpenID Connect, can use any OpenID Connect provider like Auth0, Keycloak, Okta, etc. I tested this project with Auth0.

### Setup with Auth0
1. Create an Auth0 account and create a new application.
2. Configure the application with the following settings:
   - Allowed Callback URLs: http://localhost:8080/callback
   - Allowed Web Origins: http://localhost:8080
3. Create a new API and configure it with the following settings:
   - Identifier: https://your-auth0-domain/api/v1
   - Allowed Web Origins: http://localhost:8080
4. Create a new rule to add the user's email to the JWT payload.
5. Create a new client secret and copy it to your project's .env file.
6. Run the project and navigate to http://localhost:8080/auth/login to test the authentication flow.
7. The .env file should contain the following variables:
   - AUTH0_DOMAIN: Your Auth0 domain
   - AUTH0_CLIENT_ID: Your Auth0 client ID
   - AUTH0_CLIENT_SECRET: Your Auth0 client secret
   - AUTH0_REDIRECT_URI: Your Auth0 redirect URI


## Storage
For storage this project uses AWS S3. For local development you can simulate S3 using localstack.

### Setup with Localstack
1. Install localstack using pip: `pip install localstack`
2. Start localstack: `localstack start`
3. The .env file should contain the following variable:
   - AWS_REGION=us-east-1
   - AWS_ENDPOINT=http://localhost:4566
   - AWS_CREDENTIALS_KEY=test
   - AWS_CREDENTIALS_SECRET=test 
