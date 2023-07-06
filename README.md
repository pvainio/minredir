# minredir

Minimal http redirect server.

Use to redirect http requests to another url.  For example deploy as Google Cloud Run Service to redirect requests to your naked domain (https://example.org) to another url (https://www.example.org) or to redirect alias domains (https://example.net) to primary domains (https://example.org). 

Docker image size is about 4Mb.

# usage

Specify environment variable REDIRECT with target URL like for example:
```
docker run --rm -ti -p 8080:8080 -e REDIRECT=https://example.org pekka/minredir
```

# build
```
docker build -t minredir .
```
