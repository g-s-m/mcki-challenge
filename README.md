# mcki-challenge
[McKinsey full stack app challenge](https://github.com/morkro/coding-challenge)

# Assumptions
This solution based on the next assumptions not to complicate the task 
- No authorization
- Don't use SSL or TLS, just HTTP
- Backend doesn't use any DB to store reports(but can be easily modified to support DB)   

# Repository structure
```bash
├── back       - backend source code 
├── front      - frontend source code
├── LICENSE
└── README.md
```

# How to run locally

1. Build frontend following [instruction](https://github.com/g-s-m/mcki-challenge/blob/main/front/README.md)

2. Build backend following [instruction](https://github.com/g-s-m/mcki-challenge/blob/main/back/README.md)

3. Run backend container using the following command
```bash 
$ docker run --net=host -d --env BACK_HTTP_PORT=3000 --name mcki-back mcki-back:1.0
```

4. Run frontend container using the following command
```bash
$ docker run -d -ti --net host --env PORT=8080 --env BACKEND_ADDR=localhost --env BACKEND_PORT=3000 --name mcki-front mcki:1.0
```

5. Open the page http://localhost:8080 in your browser

6. Stop containers
```
$ docker container stop mcki-back
$ docker container stop mcki-front
```
