# webcounter
A webservice that returns an incremented integer every time a request is processed.

Start webservice. Default port is 6000
> $ ./webcounter -port **[port]**

Example usage:
```
$ curl http://127.0.0.1:6000
{"count":1}
```

```
$ curl http://127.0.0.1:6000
{"count":2}
```

You can reset the counter by requesting http://**[host]**:**[port]**/reset
