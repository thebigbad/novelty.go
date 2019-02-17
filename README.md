Setup
===

1. Get your [dev environment](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment) setup for GAE.
2. Clone this repository, cd into the directory.
3. ```cp app.yaml.example app.yaml```

Running locally
===

1. Start the dev server: ```$ dev_appserver.py .```
2. Seed the server with starting data:
```$ curl whatever:bees@localhost:8080/yes```
3. Visit localhost:8080 in your browser.

Your novelty server is ready to go. The answer is current set to "yes", and the
password for changing the answer is "bees".

To change the answer to "no", simply visit larry:bees@localhost:8080/no

Running on appspot
===

1. Follow the [registration instuctions](https://developers.google.com/appengine/docs/go/gettingstarted/uploading) for GAE.
2. Push the app: ```$ appcfg.py .```
3. Seed the server with starting data. I'd suggest a different password than
"bees": ```$ curl http://moe:$PASSWORD@$APPID.appspot.com/yes```
4. Visit $APPID.appspot.com in your browser to behold your new novelty server.

License
===

novelty.go is released under the MIT license. See LICENSE for more details.
