# Referral API

Referral is the backend for `Referral`. A system for rewarding new signups.
___

## Requirements

If you want to run within a container, you need `docker/docker-compose`. Else, you'll need to have `mysql 5.6+` installed.

## Run
The app environment is configured via a `.env` file located in the root directly.  
The default port is `8080`.


With docker:
```shell
docker-compose up --build --abort-on-container-exit
```
Without docker,
```shell
go run main.go
```

Unit test:
```shell
 docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit 
```

## Swagger documentation
Swagger docs are available at http://localhost:8080/swagger/index.html  
![doc](https://i.postimg.cc/SQYfk0Wk/Screenshot-2022-05-08-at-11-55-16.png)

## Design
### Authentication
- A user registers with a `username` and `password`.
- Upon login, they receive a `jwt` to be used for subsequent authenticated requests.

### Referral
- An existing user can create a referal code. 
- They'll be returned a link to be shared. e.g
```shell
{
    "link": "http://localhost:8080/auth/register?code=630304864"
}
```
- A user can create any number of referrals.

### Signup with referal
- Upon signup with a valid referral code, you get 10 dollars.
- A single referral code can only be used a max of 5 times.
- The above parameters are configurable.

### packages
The diagram below shows the interaction between the various entities.

![design](https://i.postimg.cc/X7kbHz1k/Untitled-Diagram-drawio.png)

## .env
.env is committed simply to facilitate testing. In production, you wouldn't do this.