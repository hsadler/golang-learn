
# runs golang container with volume ./web_app
docker run -it --rm -v $PWD/web_app:/go/src/web_app -p 8090:8090 golang bash