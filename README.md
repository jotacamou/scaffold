sre/scaffold
============

Requirements:
- Clone project to `$GOPATH/src/sre`
- go dep: `go get -v -u github.com/golang/dep/cmd/dep`
- rpmdevtools

To just build `scaffold` from source:

```
make
```

To create a binary RPM package of `scaffold`:
```
make rpm
```
