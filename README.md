# Goggle
### Webcam alert dameon, monitors webcam activity and sends SMS when turned on 
---------
## Requirements
- Go 1.21+
- Linux (no mac or windows support, sorry)

# Clone & build
```bash 
git clone https://github.com/dylandhw/goggle.git
cd webcam-alert-daemon
go build -o goggle
```
# Copy and enable the servicefile
```bash
sudo cp goggle.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable goggle.service
sudo systemctl start goggle.service
sudo systemctl status goggle.service
```
