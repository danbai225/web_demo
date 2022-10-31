git pull
tag=$(git describe --abbrev=0 --tags)

docker build -t web_demo:build .

if [[ $tag != "" ]]
then
    echo "push"
    docker build -t web_demo:"$tag" .
    docker build -t web_demo:latest .
    docker push web_demo:latest
    docker push web_demo:"$tag"
fi
