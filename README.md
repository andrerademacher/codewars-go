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

The Composer binary is already present in the latest version supporting PHP 8.0 .
To run any composer command, like `composer install`, just type:
```bash
container.sh composer install
```

After installing the dev dependencies, the PHPUnit test suite can be run
inside the container by calling the phpunit binary in the vendor directory:
```bash
container.sh vendor/bin/phpunit
```

For convenience, this can be done also by using the composer script:
```bash
container.sh composer phpunit
```