# A Golang client library for [siastats.info](https://siastats.info/)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/jkawamoto/go-siastats/pkg)

This client library is generated from [siastats-swagger](https://github.com/jkawamoto/siastats-swagger) with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

## Installation 
```bash
$ go get -u github.com/jkawamoto/go-siastats/pkg
```

## Usage
### Navigator API
You might need to disable verifying certificates if you get a certification error. 

```go
import	siastats "github.com/jkawamoto/go-siastats/pkg"

func main(){

	cfg := siastats.NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := siastats.NewAPIClient(cfg)

}
```

Then you can call [`client.NavigatorApi.Hash`](https://godoc.org/github.com/jkawamoto/go-siastats/pkg#NavigatorApiService.Hash)
and [`client.NavigatorApi.Status`](https://godoc.org/github.com/jkawamoto/go-siastats/pkg#NavigatorApiService.Status).
Both functions take a context, and it's fine to give `context.Background()`.

https://github.com/jkawamoto/go-siastats/tree/master/cmd/sianav has a sample command for Navigator API.

### Hosts monitor API
Since this API uses a different host, we need to set the configuration's base path to https://siastats.info:3510/hosts-api.
You might also need to disable verifying certificates.

```go
import	siastats "github.com/jkawamoto/go-siastats/pkg"

func main(){

	cfg := siastats.NewConfiguration()
	cfg.BasePath = "https://siastats.info:3510/hosts-api"
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := siastats.NewAPIClient(cfg)

}
```

Then you can call [`client.HostApi.AllHosts`](https://godoc.org/github.com/jkawamoto/go-siastats/pkg#HostApiService.AllHosts)
and [`client.HostApi.Host`](https://godoc.org/github.com/jkawamoto/go-siastats/pkg#HostApiService.Host).
Both functions take a context, and it's fine to give `context.Background()`.

https://github.com/jkawamoto/go-siastats/tree/master/cmd/siahost has a sample command for Host monitor API.
