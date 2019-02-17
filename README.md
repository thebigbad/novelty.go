Setup
===

1. Get your
  [dev environment](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment)
  set up for GAE.
2. Clone this repository, cd into the directory.
3. ```cp app.yaml.example app.yaml```
4. Replace EXAMPLE_PASSWORD with your desired password.

Running locally
===

1. Start the dev server: ```$ dev_appserver.py .```
2. Visit http://localhost:8080/

Your novelty server is ready to go. The answer is currently set to "no".

To change the answer to "yes", simply visit
http://larry:EXAMPLE_PASSWORD@localhost:8080/no

Running on appspot
===

1. Follow the
   [registration instuctions](https://developers.google.com/appengine/docs/go/gettingstarted/uploading)
   for GAE.
2. Push the app: ```$ appcfg.py .```
3. Visit http://$APPID.appspot.com/ to behold your new novelty server.

Deploying automatically with Cloud Build (optional)
===

1. Follow the
   [getting started](https://cloud.google.com/source-repositories/docs/quickstart-triggering-builds-with-source-repositories#grant_gae_name_access_to_the_builder_name_service_account)
   guide for Cloud Build.
2. Create a build trigger, specifying wherever you keep your copy of this repo.
3. Choose "Cloud build configuration" and leave the path the default.
4. In "Substitution variables" add _PASSWORD with whatever value you want.

License
===

novelty.go is released under the MIT license. See LICENSE for more details.
