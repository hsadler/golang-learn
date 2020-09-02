
# build docker container
docker build --no-cache -f=Dockerfile -t=go_web_app:local .
# run docker container with exposed port
docker run -it --rm -p=8090:8090 --name=go_web_app go_web_app:local
