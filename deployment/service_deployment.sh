cd ~/goslug

git pull

./deployment/build.sh

sudo systemctl stop caddy
sudo systemctl stop goslug.service

sudo systemctl disable caddy
sudo systemctl disable goslug.service

sudo systemctl daemon-reload

sudo systemctl enable caddy.service
sudo systemctl enable goslug.service

sudo systemctl start caddy.service
sudo systemctl start goslug.service