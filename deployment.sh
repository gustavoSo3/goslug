cd ~/goslug
git pull origin main
sudo systemctl restart goslug.service
sudo systemctl restart caddy.service
