Setup
===

1. Get your [dev environment](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment) setup for GAE.
2. Clone this repository, cd into the directory.
3. ```cp app.yaml.example app.yaml```
4. Edit app.yaml with the app name you plan to use.

Running locally
===

1. Start the dev server: ```$ dev_appserver.py .```
2. Seed the server with starting data:
```$ curl whatever:bees@localhost:8080/yes```
3. Visit localhost:8080 in your browser.

Your novelty server is ready to go. The answer is current set to "yes", and the
password for changing the answer is "bees".

To change the answer to "no", simply visit larry:bees@localhost:8080/no

Running tests
===

1. Install go.
2. ```$ cd basicauth/ && go test```

Running on appspot
===

1. Follow the [registration instuctions](https://developers.google.com/appengine/docs/go/gettingstarted/uploading) for GAE.
2. Make sure that the app id in app.yaml matches you new app id.
3. Push the app: ```$ appcfg.py .```
4. Seed the server with starting data. I'd suggest a different password than
"bees": ```$ curl http://moe:$PASSWORD@$APPID.appspot.com/yes```
5. Visit $APPID.appspot.com in your browser to behold your new novelty server.

Sample Benchmarks
===

    $ ab -n 1000 -c 20 http://$APPID.appspot.com/
    This is ApacheBench, Version 2.3 <$Revision: 655654 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking $APPID.appspot.com (be patient)
    Completed 100 requests
    Completed 200 requests
    Completed 300 requests
    Completed 400 requests
    Completed 500 requests
    Completed 600 requests
    Completed 700 requests
    Completed 800 requests
    Completed 900 requests
    Completed 1000 requests
    Finished 1000 requests


    Server Software:        Google
    Server Hostname:        $APPID.appspot.com
    Server Port:            80

    Document Path:          /
    Document Length:        730 bytes

    Concurrency Level:      20
    Time taken for tests:   13.121 seconds
    Complete requests:      1000
    Failed requests:        0
    Write errors:           0
    Total transferred:      875000 bytes
    HTML transferred:       730000 bytes
    Requests per second:    76.21 [#/sec] (mean)
    Time per request:       262.424 [ms] (mean)
    Time per request:       13.121 [ms] (mean, across all concurrent requests)
    Transfer rate:          65.12 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:       54   66   9.2     63     131
    Processing:   125  195  75.5    153     695
    Waiting:      124  195  75.4    153     695
    Total:        181  261  76.7    223     756

    Percentage of the requests served within a certain time (ms)
      50%    223
      66%    269
      75%    316
      80%    332
      90%    388
      95%    412
      98%    437
      99%    446
     100%    756 (longest request)

License
===

novelty.go is released under the MIT license. See LICENSE for more details.
