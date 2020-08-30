
# runs golang container with volume ./app
docker run -it --rm -v $PWD/app:/go/src/app golang bash
