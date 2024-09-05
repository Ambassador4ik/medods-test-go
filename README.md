# medods-test-go

## Installation
### Docker 
Make sure that docker is installed. Build your container with the following command:
```shell
docker build -t medods-test-go . 
```
Run your container:
```shell
docker run -it --rm --name medods-app-container -e DB_SOURCE="<your-db-source>" medods-test-go 
```
Provide a valid [PostgreSQL database source](https://entgo.io/docs/getting-started#create-your-first-entity) inside `DB_SOURCE` env. 

### Manual
Make sure that you have go 1.23 installed.\
Generate ent schema:
```shell
go generate ./ent
```
Build the application:
```shell
go build -o app github.com/Ambassador4ik/medods-test-go/cmd/server
```
Run the binary:
```shell
./app
```

## Configuration
Make sure to properly configure client secret and Mailgun tokens.\
Database connection string can be passed to the environment as stated above.

## Usage
By default, server is configured to run on port 3000. The available handlers are:
### POST /tokens/get
**Accepts:** User GUID\
**Returns:** A pair of connected access/refresh tokens\
**Example:**
```shell
curl --location --request POST 'http://127.0.0.1:3000/tokens/get?guid=56a0d47c-4c3a-4c70-be91-7f2d5e264538'
```
Returns on success:
```shell
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjlhNWRjZmYyLWYzN2QtNDQ1OS1hYzMyLWNmMzk0OTRmNGI5YiIsImd1aWQiOiI1NmEwZDQ3Yy00YzNhLTRjNzAtYmU5MS03ZjJkNWUyNjQ1MzgiLCJpcCI6IjEyNy4wLjAuMSIsImV4cCI6MTcyNTUzMjYzMH0.qnc4ZaB_4LATeCv6ZVgiS0ejtj5UPa2W6X6oktz3S9oxojbKcYvJwiG2MWZIwusy9Sdj6XkcCP2qlLqnjOna7w",
    "refresh_token": "nUy3LMM4VNS+TWdOxxTI6tGPmMp/+o7VJQxpCe/vDMY="
}
```

### POST /tokens/refresh
**Accepts:** A pair of connected, URI-encoded access/refresh tokens\
**Returns:** A new pair of connected access/refresh tokens\
**Example:**
```shell
curl --location --request POST 'http://127.0.0.1:3000/tokens/refresh?accessToken=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjlhNWRjZmYyLWYzN2QtNDQ1OS1hYzMyLWNmMzk0OTRmNGI5YiIsImd1aWQiOiI1NmEwZDQ3Yy00YzNhLTRjNzAtYmU5MS03ZjJkNWUyNjQ1MzgiLCJpcCI6IjEyNy4wLjAuMSIsImV4cCI6MTcyNTUzMjYzMH0.qnc4ZaB_4LATeCv6ZVgiS0ejtj5UPa2W6X6oktz3S9oxojbKcYvJwiG2MWZIwusy9Sdj6XkcCP2qlLqnjOna7w&refreshToken=nUy3LMM4VNS%2BTWdOxxTI6tGPmMp%2F%2Bo7VJQxpCe%2FvDMY%3D'
```
Returns on success:
```shell
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjI2ZGZjNGQxLTFkYmYtNGFjMi05MDY3LWNjNTUwOTc2YzJjZSIsImd1aWQiOiI1NmEwZDQ3Yy00YzNhLTRjNzAtYmU5MS03ZjJkNWUyNjQ1MzgiLCJpcCI6IjEyNy4wLjAuMSIsImV4cCI6MTcyNTUzMjY1NX0.-8mU111eD_AVcJyVv-a80r_SJGggeWoPZ2HumkzazBxgFHDNn86Ep9GAnxAE69CNB83eHoKW4YuCxqDz0eeJIQ",
    "refresh_token": "8mExy7/CJiWBwu39dYUgpYursITBRL3S4Bg5VJUjDXI="
}
```
