Distributed Services with Go



```shell
$ http -v GET localhost:8080 offset:=0
GET / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 13
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/1.0.3

{
    "offset": 0
}

HTTP/1.1 404 Not Found
Content-Length: 17
Content-Type: text/plain; charset=utf-8
Date: Sun, 29 Mar 2020 00:38:37 GMT
X-Content-Type-Options: nosniff

offset not found
```

```shell
[macmini]~/Projects/distributed-services-with-go/proglog:0 (master)
$ jo record=$(jo value="TGV0J3MgR28gIzEK") | http -v POST localhost:8080
POST / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 40
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/1.0.3

{
    "record": {
        "value": "TGV0J3MgR28gIzEK"
    }
}

HTTP/1.1 200 OK
Content-Length: 13
Content-Type: text/plain; charset=utf-8
Date: Sun, 29 Mar 2020 00:46:30 GMT

{
    "offset": 3
}
```

```shell
[macmini]~/Projects/distributed-services-with-go/proglog:0 (master)
$ http -v GET localhost:8080 offset:=0
GET / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 13
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/1.0.3

{
    "offset": 0
}

HTTP/1.1 200 OK
Content-Length: 51
Content-Type: text/plain; charset=utf-8
Date: Sun, 29 Mar 2020 00:47:17 GMT

{
    "record": {
        "offset": 0,
        "value": "TGV0J3MgR28gIzEK"
    }
}

```

