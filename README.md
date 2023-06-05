# Caddy API Go Wrapper

This Go package provides a wrapper for the Caddy server's administration endpoint, which can be accessed via HTTP using a REST API. The wrapper provides a simple and idiomatic way to interact with the Caddy server's API from your Go programs.

## Installation

To use this package in your Go program, you can install it using `go get` :

```bash
go get -u github.com/multisig-labs/caddyapi
```

## Usage

First, import the `caddyapi` package in your Go file:

```go
import "github.com/multisig-labs/caddyapi"
```

Then, create a new `CaddyAPI` instance:

```go
caddy := caddyapi.NewCaddyAPI("http://localhost:2019")
```

You can now use the `caddy` instance to interact with the Caddy server's API.

### Load Configuration

To load a new configuration, use the `LoadConfig` method:

```go
config := caddyapi.Config{
	// Fill in your configuration here
}
err := caddy.LoadConfig(config)
if err != nil {
	fmt.Println("Failed to load config:", err)
}
```

### Stop Server

To stop the Caddy server, use the `StopServer` method:

```go
err := caddy.StopServer()
if err != nil {
	fmt.Println("Failed to stop server:", err)
}
```

### Get Configuration

To get the current configuration, use the `GetConfig` method:

```go
config, err := caddy.GetConfig("/path/to/config")
if err != nil {
	fmt.Println("Failed to get config:", err)
}
```

### Post Configuration

To set or replace a configuration, use the `PostConfig` method:

```go
config := caddyapi.Config{
	// Fill in your new configuration here
}
err := caddy.PostConfig("/path/to/config", config)
if err != nil {
	fmt.Println("Failed to post config:", err)
}
```

### Patch Configuration

To replace an existing configuration, use the `PatchConfig` method:

```go
config := caddyapi.Config{
	// Fill in your new configuration here
}
err := caddy.PatchConfig("/path/to/config", config)
if err != nil {
	fmt.Println("Failed to patch config:", err)
}
```

### Put Configuration

To create a new configuration, use the `PutConfig` method:

```go
config := caddyapi.Config{
	// Fill in your new configuration here
}
err := caddy.PutConfig("/path/to/config", config)
if err != nil {
	fmt.Println("Failed to put config:", err)
}
```

### Delete Configuration

To delete a configuration, use the `DeleteConfig` method:

```go
err := caddy.DeleteConfig("/path/to/config")
if err != nil {
	fmt.Println("Failed to delete config:", err)
}
```

## Documentation

For more information about the Caddy server's API, see the [official documentation](https://caddyserver.com/docs/api).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

This wrapper is based on the Caddy server's API, which is developed and maintained by the Caddy project.
