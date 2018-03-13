# webcounter
A webservice that returns an incremented integer every time a request is processed.

Start service
$ ./webcounter -port [port]

Example usage:
$ curl http://127.0.0.1:6000
{"count":1}

$ curl http://127.0.0.1:6000
{"count":2}

To reset timer go to curl http://127.0.0.1:6000/reset
