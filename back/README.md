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
## Backend API 
1. **/reports/**
   Description: returns reports in json format  
   Method: GET  
   Parameters: No  
   Example:  
```bash
http://mcki-back/reports
```  
   Returns: List of not resolved reports 

2. **/reports/reportId**
   Description: blocks or resolved reports with Id equals reportId  
   Method: PUT  
   Parameters: No  
   Example:  
```bash
http://mcki-back/reports/11c347a7-223a-4b6f-8b26-e492474873c1
```
   Body: json  
```json
{
   "ticketState": "RESOLVED"
}
```
ticketState can be either RESOLVED or BLOCKED

3. **reports/reload/**
   Description: reloads initial state of the reports  
   Method: POST  
   Parameters: No  
   Example:  
```bash
http://mcki-back/reports/reload
```
