cd ~/url-shortener

git fetch --all
git reset --hard origin/main

docker build -t urlshortener .

docker stop url-shortener
docker rm url-shortener

docker run -d --name url-shortener urlshortener:latest
