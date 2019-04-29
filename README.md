# LetsGo
[![CircleCI](https://circleci.com/gh/AoLab/LetsGo.svg?style=svg)](https://circleci.com/gh/AoLab/LetsGo)

## Introduction
The aim of this project is to provide a basis to write a simple HTTP service in Go.
Do not push anything on this project, you must fork it then write your code and make a pull request to it.

## Step by Step
0. Fork this repository to start
1. Write a simple HTTP webserver using [echo](https://github.com/labstack/echo). This webserver has a single endpoint 
that is named `ping` and returns no content with 200 status.
2. Create docker-compose file that runs mongodb and expose its port.
3. Add new endpoint with POST method that accept the following JSON and insert it into mongodb `geo` collection.

```json
{
  "lat": 35.8061619,
  "lng": 51.3987635
}
```

For conntecting to Mongodb use the offical driver that can be found [here](https://github.com/mongodb/mongo-go-driver).
