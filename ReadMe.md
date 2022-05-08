# Steps to run #
## Creating image ##
docker build -t simple_app .
## Lunch docker app ##
docker run -p 8080:8080 simple_app


## Request test ##
http://localhost:8080/random/mean?requests=2&length=5
