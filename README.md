# vod-CLI

This project is a CLI for managing  vod Service.
The command vod-cli can get video detail or update video detail.

## How its created?

I use Go programming language and its library for building CLI applications, cobra.
Cobra is built on a structure of commands, arguments & flags.

### Prerequisites
To create this CLI, first,  it needed to install **Go** and **cobra** .
1. Using [this](https://go.dev/dl/) documentation to install Go.
2. Installing Cobra:
   ```bash
   go get -u github.com/spf13/cobra/cobra
   ```

### Initialising the project
 First, create a directory for the project:
  ```bash
   cd go/src/
   mkdir vod-cli
   ```
  Initialize the go project and add the go.mode file:
  ```bash
   go mode init avarn-vod-cli
   ```
**Note:** The *go.mod* file defines the module's module path, the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.

With the below command, the CLI initialized for the project :

```bash
cobra init vod-cli
```

It created a few files in the project:

> vod-cli/
>>&nbsp; cmd/
>>>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;root.go
>>
>>main.go

In the main.go, the Execute() function of the cmd/root.go is called. In root.go, the root command is initialized and is  ```vod-cli```.All the other command in the CLI is the child of the root command.

There are three child commands: login, getDetail, and updateDetail.
Each of them was created with ``` cobra add <CHILD-COMMAND-NAME>```, and new files for them were created in the cmd directory.

## Login command

To use this CLI and connect to the API,  give the API key for authentication.
To set the APIKey,  use the command below and give the APIKey:
```bash
vod-cli login <API-KEY>
```

### How does it work?
To save the APIKey, the directory ``` .vod```  was created with a config file in it  . After that, save the given key in this file.
There are three functions for it:

- getdirectory(): It gets the user's home directory address and adds the config address to it, and returns it like this:
```User/mona/.vod/config```
- addapikey(): This function saves the APIKey, that the user gave to the address that the getdirectory() returns.
- readapikey(): It defines reading the API key from the file and using it in the API requests.

## GetDetail
With this command, get the detail of the videos by passing the video-id.
```bash
vod-cli getDetail <VIDEO-ID>
```
### How does it work?
 It gets the video-id from the CLI and the APIKey from the readapikey(), sets it as an "Authorization" header for the request, and sends a GET HTTP request.

**Note:** I use the ```net/http``` package to send the HTTP requests. It provides HTTP client and server implementations.


## UpdateDetail

To update the title and description of the videos, use this command.
Two flags were defined for it  :

- title: get the new title of the video with a type of string
- description: get the new description of the video with a type of string

To update these two parameters,  use the below command:

```bash
vod-cli updateDetail <VIDEO-ID> --title <NEW-TITLE> --description <NEW-DESCRIPTION>
```
**Note:** To set a new description, it is necessary to use " " to send it like :
```bash
vod-cli updateDetail 2374633874 --title change-title --description "it is to change a title"
```
 ### How does it work?
Get the video-id, title, and description from CLI args and flags and create an HTTP PATCH request.

## Using this CLI
This command generates the binary or executable file of the project in the ```$GOPATH/bin``` folder.
```bash
go install vod-cli
```
Now run `vod-cli`  in the terminal. As it is saved in the bin folder, there is no need to set the environment variable for this.

## Challenges

1. Creating modules inside the  `GOPATH` is disabled by default, so I got this error when running the above command:
```bash
   go: modules disabled inside GOPATH/src by GO111MODULE=auto; see ‘go help modules’
   ```
 To enable creating modules in GOPATH, set the `GO111MODULE` environment variable to `on`:
 ```bash
   export GO111MODULE=on
   ```

 2. Getting the APIKey is one of the challenges of this project. At first, I get it as a flag in the command, but it is not secure, so I find a way to save it as a file in the user's local system, and implementing this method is challenging.
