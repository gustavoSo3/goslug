cd ~/goslug

git pull origin main

sudo cp ~/goslug/goslug.service /etc/systemd/system/goslug.service

sudo systemctl daemon-reload

sudo systemctl restart goslug.service
sudo systemctl restart caddy.service

sudo systemctl status goslug.service
sudo systemctl status caddy.service


