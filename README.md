# Technopark_3_Security

### How to use

1. Build docker container
```
docker build -t server .
```
2. Run docker container
```
docker run -it --rm -p 8080:8080 server
```
3. cURL (in other terminal)
```
curl -x http://127.0.0.1:8080 http://mail.ru
```