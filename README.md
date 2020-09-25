
# HumanID SDK for Go
humanid-golang-sdk is the official sdk for [HumanID](https://human-id.org/).

Content:
- [Getting Started](#getting-started)
- [Examples](#examples)
- [Contributing](#contributing)


## Getting Started

### Installation
```
go get github.com/bluenumberfoundation/humanid-golang-sdk
```
For the latest version of the SDK:
```
go get -u github.com/bluenumberfoundation/humanid-golang-sdk
```
### Dependencies
You can access all the package dependencies in the vendor folder

## Examples
```
import (
  "os"
  "errors"

  humanID "github.com/bluenumberfoundation/humanid-golang-sdk"
)

...

func getLoginUrl(clientID string, clientSecret string) (string, error) {
  h := humanID.New(
    os.Getenv("SERVER_ID"),
    os.Getenv("SERVER_SECRET"),
  )

  loginResp, err := h.Login(clientID, clientSecret)
  if err != nil {
    return "", err
  }

  if loginResp.Code == "401" || !loginResp.Success {
    return "", errors.New("An error has occurred!")
  }

  return loginResp.Data.WebLoginUrl, nil
}

...
```

## Contributing
### Run Tests
#### Pre-requisites (only for testing):
- docker
- docker-compose

#### Run tests:
```
docker-compose up
```
