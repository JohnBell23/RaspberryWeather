chmod +x /home/pi/RaspberryWeather/RaspberryWeather

sudo systemctl daemon-reload
sudo systemctl enable RaspberryWeather
sudo systemctl start RaspberryWeather