#! /usr/bin/env bash
#APIID=3t9ljow8x2
#RESTAPI=https://$APIID.execute-api.eu-west-1.amazonaws.com/prod
curl -X POST -d '{"title": "curl-test", "completed": false, "id": "7e09e7ae-b8ee-4084-b4ac-c8b95e6d62c5", "creationdate": "2012-04-23T18:25:43.511Z"}' $RESTAPI/todo

