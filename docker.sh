# Docker commands
NAME_IMAGE=proto-compiler2
DIR_LOCAL=`pwd`
docker build --platform linux/amd64 --tag $NAME_IMAGE .

docker run --platform linux/amd64 --rm -v $DIR_LOCAL:/temp/buf-gen $NAME_IMAGE
