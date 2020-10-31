# Golang examples

Repository for basic go example programs.

# GO HELP

If you’re ever stuck without internet access, you can get the documentation running locally
via: 
godoc -http=:6060 </br>
and
pointing your browser to 
http://localhost:6060

# GO Commands
```
 $ go version               print GO version
 $ go env                   print GO environment variables set
 $ go run main.go           runs the main.go code and shows the result
 $ go run --work main.go    displays the temp binary path and results
 $ go build                 build the go binaries of modules in the current folder
 $ go build -o executable   build the go binaries and specify its name as "executable"
 $ go install               build and install binaries in bin
 $ go get                   download and install packages and dependencies
 $ go clean                 removes object files in directory
 $ go doc                   show documentations for package or symbol
```
 # go install / go build
  - Both will compile the package in current directory
  - Install put binaries in the GOPATH/bin
