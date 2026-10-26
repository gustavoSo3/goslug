cd ~/goslug

sudo cp ~/goslug/goslug.service /etc/systemd/system/goslug.service

sudo systemctl daemon-reload

sudo systemctl restart goslug.service
sudo systemctl restart caddy.service



