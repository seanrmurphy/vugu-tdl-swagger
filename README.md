# vugu-tdl-swagger

This repo contains a fullstack Go application comprising of
- a vugu based frontend
- an AWS Lambda based backend
- a swagger interface which supports FE-BE communication and hooks into AWS cognito

The main point of this repo is to demonstrate implementation of a secure backend
and how it can integrate with a vugu/Go frontend.

More details on this work is provided in this Medium post:

To build the application, follow the instructions in each of the directories;
- deploy the backend first (see the backend directory),
- create the swagger client library (in the swagger directory)
- create the frontend (see the frontend directory).
