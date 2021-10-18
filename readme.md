This is a demo of Go Workspace

Please do the setup before running
1. Set $GOPATH to GoTest project directory. This is our Go Workspace
2. Set $GOBIN to $GOPATH/bin
3. Add $GOBIN to $PATH for easily accessing the application binary
4. Set $GO111MODULE to off to disable Go Modules

To run the program
1. Run 'go install pkg1 pkg2 pkg1/pkg3 pkg4' to build and install utility packages in $GOPATH/pkg/$GOOS_GOARCH
2. Run 'go install github.com/pborman/uuid' to the install the dependency package (if already downloaded in $GOPATH/src. If not, use 'go get' as discussed later)

Note that, we can simply do go install ./... in $GOPATH to install all utility packages and application binary

3. Run 'go install main' to install the binary in $GOBIN with name as 'main.exe' 
4. Run 'main' to run the application binary (if $GOBIN is in $PATH)
5. Alternatively, run 'go run main' to build and run the application binary

Note that, 'go install main' does not install utility packages in $GOPATH/pkg/$GOOS_GOARCH, it just installs the
application binary to $GOBIN

Note that, the above commands can be run from anywhere (Go finds packages / files using $GOPATH)

To demonstrate behaviour of 'go get'
1. Run 'go get github.com/pborman/uuid'from anywhere
2. Go will download source code of the above package and its dependent packages in $GOPATH/src folder
3. Go will install the above package and its dependent packages in $GOPATH/pkg/$GOOS_GOARCH