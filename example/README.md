
## Go hooks

This repo contains an example of using the Go Dredd hook handler which provides a bridge between the Dredd API Testing Framework and Go environment to ease implementation of testing hooks provided by Dredd. Write Dredd hooks in Go to glue together API Blueprint with your Go project.

Not sure what these Dredd Hooks are? Read the Dredd documentation on them

The following are a few examples of what hooks can be used for:

- loading db fixtures
- cleanup after test step or steps
- handling authentication and sessions
- passing data between transactions (saving state from responses to stash)
- modifying request generated from blueprint
- changing generated expectations
- setting custom expectations
- debugging via logging stuff

### Installing the hooks **Caution this will not work until [this](https://github.com/snikch/goodman/pull/5) is merged**

```bash
$ go get github.com/snikch/goodman
$ cd $GOPATH/src/github.com/snikch/goodman
$ go build -o $GOPATH/bin/goodman github.com/snikch/cmd/goodman
```


### Running this example

**Caution [this](https://github.com/apiaryio/dredd/pull/505) must be merged before dredd supports the go hooks**

```bash
cd github.com/snikch/goodman/examples
npm install
```

Since Go is a compiled language, the hookfiles must be compiled before being passed to dredd via the --hookfiles flag.  **Caution**, this behavior may change in the future when hookfile globs (via bash wildcards are implemented).

The Makefile that comes with this repo does the following
- Compiles hooks/hooks.go
- Compiles the go web app (main.go)
- Runs local dredd install 

See the Makefile for more details
