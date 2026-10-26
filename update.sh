cd ~/goslug

git fetch --all
git reset --hard origin/main

docker build -t goslug .

docker stop goslug
docker rm goslug

docker run -d -p 8080:8080 --name goslug goslug:latest
