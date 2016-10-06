Setup
-------------------------

### Get govendor
```
go get -u github.com/kardianos/govendor
```

### Add Govendor to your path
```
export PATH=$PATH:$GOPATH/bin
```

### Use Govendor to fetch beego
```
govendor fetch github.com/astaxie/beego
```

### Run main.go
```
go run main.go
```

### Navigate
```
http://localhost:8000/request?type=geolocation&ip=127.0.0.1
``` 

