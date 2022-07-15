git pull
tag=$(git describe --abbrev=0 --tags)
if [[ $tag == "" ]]
then
    tag="build"
fi
docker build -t web_demo:"$tag" .

if [[ $tag != "build" ]]
then
    docker build -t web_demo:latest .
    docker push web_demo:latest
    docker push web_demo:"$tag"
fi
