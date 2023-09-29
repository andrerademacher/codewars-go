# codewars
Tasting the golang greatness with katas!

## Build Golang 1.20 Docker image
Use the build script in order to build the custom Golang 1.20 Docker image
`andrerademacher/codewars-go120`.

```bash
container/build.sh
```

## Run command in container
The `container.sh` script makes running commands in the Docker container easy!
To open a shell, just add the "sh" command:
```bash
container.sh sh
```

The current Go version can be shown like this:
```bash
container.sh go version
```

Go mod is already enabled and configured.
Downloading the dependencies can be done using go mod download:
```bash
container.sh go mod -x download
```

The test suite can be run by calling:
```bash
container.sh go test ./...
```
