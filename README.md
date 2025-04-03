# 05-go-rest-api

# user-service

# REST API

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST / user/:id -- create user -- 204, 4xx, Header Location: url
PUT / users/:id -- fully update user -- 204/200, 404, 400, 500
PATCH /users/:id -- partially update user -- 2-4/200, 404, 400, 500
DELETE /users/:id -- delete uese by id -- 204, 404, 400

julienschmidt/httprouter
ilyakaznacheev/cleanenv
/mongo-driver/v2/mongo
/sirupsen/logrus
