sudo cp ./RaspberryWeather.service /etc/systemd/system/RaspberryWeather.service

chmod +x /home/pi/RaspberryWeather/RaspberryWeather

sudo systemctl daemon-reload
sudo systemctl enable --now RaspberryWeather.service
