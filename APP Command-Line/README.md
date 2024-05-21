# Usage:

## Initialize the Go module
Since we will be using more than one package, initialize a new module:

```sh
go mod init app-command-line
```

## Install the required package
Install the CLI package dependency:

```sh
go get github.com/urfave/cli
```

## Build the code
Build the code inside the Go directory:

```sh
go build
```

## Running the Application
To run the application, use the following command structure:

### Search for IPs
To search for public IPs of a given DNS host:

```sh
./main.exe ip --host <hostname>
```

For example, to search for the IPs of `google.com`:

```sh
./main.exe ip --host google.com
```

### Search for Server Names
To search for the server names of a given DNS host:

```sh
./main.exe servers --host <hostname>
```

For example, to search for the server names of `google.com`:

```sh
./main.exe servers --host google.com
```

### Help
To display the help menu:

```sh
./main.exe help
```