# Frontend description

## How to run

- Install docker
- Build docker image
```
$ docker build -t mcki:1.0 .
```
- Run frontend locally
```
$ docker run -ti -p 127.0.0.1:3000:3000/tcp mcki:1.0
```
- Open the page http://localhost:3000 in your browser

