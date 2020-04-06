# postmanctl

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/kevinswiber/postmanctl)

A command-line interface to the [Postman API](https://docs.api.getpostman.com/).

## Usage

```
Controls the Postman API

Usage:
  postmanctl [command]

Available Commands:
  config      Configure access to the Postman API.
  create      Create new Postman resources.
  delete      Delete existing Postman resources.
  describe    Describe an entity in the Postman API
  fork        Create a fork of a Postman resource.
  get         Retrieve Postman resources.
  help        Help about any command
  merge       Merge a fork of a Postman resource.
  replace     Replace existing Postman resources.
  run         Execute runnable Postman resources.

Flags:
      --config string   config file (default is $HOME/.postmanctl.yaml)
  -h, --help            help for postmanctl

Use "postmanctl [command] --help" for more information about a command.
```

## Install

Currently, `postmanctl` can be installed via `go get`:

```
$ go get -u github.com/kevinswiber/postmanctl/cmd/postmanctl
```

## Getting started

### Configuring access

To start using `postmanctl`, configure access to your Postman account.

```
$ postmanctl config set-context <context-name>
```

You'll need a Postman API Key to add to your configuration, which can be generated here: https://the.postman.co/settings/me/api-keys.

Example:

```
$ postmanctl config set-context personal
Postman API Key: 
config file written to $HOME/.postmanctl.yaml
```

You're now ready to start using `postmanctl`!

### Fetching a list of Postman collections

Now that access to the Postman API has been configured, you can start playing around with different commands.

Fetch a list of Postman collections.

```
$ postmanctl get collections
UID                                             NAME
10354132-0a428e3b-4112-46ee-b57a-d2f3e1b7c860   httpbin
10354132-22f0b9af-83e6-4f4a-b14a-879342f3e582   Using data files
10354132-e02524dc-54d5-49d7-9ef8-121209316083   Demo API
```

### Get more information about a collection

```
$ postmanctl describe collection 10354132-e02524dc-54d5-49d7-9ef8-121209316083
Info:
  ID:      e02524dc-54d5-49d7-9ef8-121209316083
  Name:    Demo API
  Schema:  https://schema.getpostman.com/json/collection/v2.1.0/collection.json
Scripts:
  PreRequest:  true
  Test:        true
Variables:     baseUrl
Items:
  .
  └── Weather Forecast
      ├── Get Weather Forecast (scripts: prerequest,test)
      └── Create Weather Forecast (scripts: prerequest,test)
```

### Create a mock server

You can create resources by piping in a JSON object describing that resource or by passing in a file with the `--filename` flag.

```
$ cat << EOF | postmanctl create mock
{
  "name": "demo-api-mock",
  "collection": "10354132-e02524dc-54d5-49d7-9ef8-121209316083",
  "environment": "10354132-84e7635f-9b30-427f-bead-27901790659e"
}
EOF

10354132-588394be-63e5-4194-828c-4439fee85ca8
```

The ID of the newly created resource is returned.

You can check the new resource by running:

```
$ postmanctl describe mock 10354132-588394be-63e5-4194-828c-4439fee85ca8
```

### Using a custom JSONPath output to filter results

Now that we have a mock server, let's run a command to get its public URL.

```
$ postmanctl get mock 10354132-588394be-63e5-4194-828c-4439fee85ca8 -o jsonpath="{[].mockUrl}"
https://588394be-63e5-4194-828c-4439fee85ca8.mock.pstmn.io
```

Now that we have our mock server URL, we can start making requests.

```
$ curl https://588394be-63e5-4194-828c-4439fee85ca8.mock.pstmn.io/WeatherForecast  
[
 {
  "date": "<dateTime>",
  "temperatureC": "<integer>",
  "temperatureF": "<integer>",
  "summary": "<string>"
 },
 {
  "date": "<dateTime>",
  "temperatureC": "<integer>",
  "temperatureF": "<integer>",
  "summary": "<string>"
 }
]%  
```

## Learning more

Feel free to peruse the auto-generated [CLI docs](doc/postmanctl.md) to learn more about the commands or just explore with `postmanctl <command> <subcommand> --help`.

JSONPath syntax is borrowed from the implementation used by the Kubernetes CLI.  You can find more documentation on that in the `kubectl` documentation: https://kubernetes.io/docs/reference/kubectl/jsonpath/.

## License

Apache License 2.0 (Apache-2.0) 

Copyright © 2020 Kevin Swiber <kswiber@gmail.com>