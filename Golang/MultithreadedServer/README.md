# Demo Go Server
#### A demo server is create under ```server/ ``` folder and a http request generator to hit that server is create under "RequestGenerator/" folder

# Usage
## Server
- Open new terminal
- Git fetch code from [server/](https://github.com/thejayendrasingh/GolangPractice)
- Go to ```Golang/MultithreadedServer/Server``` folder
- Run ```make build``` command from root of server folder
- Run command ```./go_server``` to start server

## Request Generator
- Open new terminal
- Go to ```Golang/MultithreadedServer/RequestGenerator``` folder
- Run ```go run request_generator.go```
- Prompt will ask a question ```Enter number of random request to generate : ```, enter any number
- Utility will auto trigger al request with random inputs
