# Logcatbeat

Welcome to Logcatbeat. It can get logs from Cloud Android and send them out.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/fize/logcatbeat`

## Getting Started with Logcatbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Run

To run Logcatbeat with debugging output enabled, run:

```
./logcatbeat -c logcatbeat.yml -e -d "*" -E seccomp.enabled=false
```

[For security purposes the libbeat framework by default drops the ability to fork/exec. So if you are developing your own Beat that needs to do those things you should programmatically register 1 your own less restrictive policy or you can disable the protections from your config by setting seccomp.enabled: false (or on the CLI with -E seccomp.enabled=false).](https://discuss.elastic.co/t/unable-to-run-commands-using-exec-command-from-beat-linux/167360)

### Clone

To clone Logcatbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/fize/logcatbeat
git clone https://github.com/fize/logcatbeat ${GOPATH}/src/github.com/fize/logcatbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Config

- option: this is logcat command option
- os: linux or android
- tags: a tag list
