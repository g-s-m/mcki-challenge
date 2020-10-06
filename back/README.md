# Backend description

## How to run

- Install docker
- Run following command to build backend docker image
```
$ docker build -t mcki-back:1.0 . 
```
- Run backend locally
```
$ docker run -ti --env BACK_HTTP_PORT=3000 -p 127.0.0.1:3000:3000/tcp mcki-back:1.0
```
- Try to get reports
```
$ curl http://localhost:3000/reports
```

