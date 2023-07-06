# minredir
Minimal http redirect server

# build
docker build -t minredir .

# usage
docker run --rm -ti -p 8080:8080 -e REDIRECT=https://example.org minredir
