TODO: @josh add info about the backend dev guide
TODO: @zzq add info about seeding

## Bruno

Flowforge uses [Bruno](https://www.usebruno.com) as the API test tool. All the files can be found at `flowforge_api_bruno`.

To simplify the process of running each API, there is a pre-script created that will fetch the JWT token from Auth0 to authenticate the API requests. To get this pre-script to work correctly, simply create a `.env` file in the `flowforge_api_bruno` directory with the following content:

- `AUTH0_USERNAME`: username for the Auth0 application
- `AUTH0_PASSWORD`: password for the Auth0 application
- `AUTH0_CLIENT_SECRET`: client secret for the Auth0 application
