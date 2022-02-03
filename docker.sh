# Docker commands
NAME_IMAGE=proto-compiler2
DIR_LOCAL=`pwd`
docker build --tag $NAME_IMAGE .

docker run --rm -v $DIR_LOCAL:/temp/buf-gen $NAME_IMAGE
