# What?

`astrolabe` is a companion to [Tiller](https://www.tillerhq.com/) sheets that intends to add some machine learning based automation.

# How?

* Read sheets directly from Google.
* Process data points through tensorflow.
* Update rows directly in Google.

# Setup

## OAuth Client

You need to [register an OAuth client](https://console.developers.google.com/apis/credentials/oauthclient) to allow this application to talk to your google drive sheets. Save the credentials JSON to `~/.config/astrolabe/client_secret.json`.

