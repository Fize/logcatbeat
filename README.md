# Logcatbeat

Welcome to Logcatbeat. It can get logs from Android devices or emulator and send them out.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/Fize/logcatbeat`

## Getting Started with Logcatbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Logcatbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Logcatbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/Fize/logcatbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Logcatbeat run the command below. This will generate a binary
in the same directory with the name logcatbeat.

```
make
```


### Run

To run Logcatbeat with debugging output enabled, run:

```
./logcatbeat -c logcatbeat.yml -e -d "*"
```

[For security purposes the libbeat framework by default drops the ability to fork/exec. So if you are developing your own Beat that needs to do those things you should programmatically register 1 your own less restrictive policy or you can disable the protections from your config by setting seccomp.enabled: false (or on the CLI with -E seccomp.enabled=false).](https://discuss.elastic.co/t/unable-to-run-commands-using-exec-command-from-beat-linux/167360)

### Test

To test Logcatbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Logcatbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Logcatbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/Fize/logcatbeat
git clone https://github.com/Fize/logcatbeat ${GOPATH}/src/github.com/Fize/logcatbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
