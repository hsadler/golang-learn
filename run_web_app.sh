
# runs golang container with volume ./web_app
docker run -it --rm -v $PWD/web_app:/go/src/web_app golang bash