# {{ name }} 

## install

- install requres

```
$ npm install -g concurrently
$ npm install -g lite-server
```

- run go server

```
$ go run main.go
```

- run web app in another terminal

```
$ cd web
$ yarn install
$ yarn start
```

## with docker

- create docker image and run it.

```
$ docker build -t my-golang-app .
$ docker run -it -p 8080:8080 --rm --name my-running-app my-golang-app
```

- you can access http://localhost:8080/web/
