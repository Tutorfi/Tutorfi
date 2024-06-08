If you got this problem:
failed to solve: golang:1.22.0-alpine3.19: failed to resolve source metadata for docker.io/library/golang:1.22.0-alpine3.19: error getting credentials - err: exit status 1, out: ``
Run this:
sudo rm ~/.docker/config.json
