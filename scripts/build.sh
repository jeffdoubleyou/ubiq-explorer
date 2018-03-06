echo "Generating native ABI token bindings"
abigen --abi ./daemon/tokens/token.abi --pkg tokens --type Token --out ./daemon/tokens/tokens.go

echo "Building bin files"
go build -o ./bin/blockdaemon ./daemon/blockdaemon.go
go build -o ./bin/poolstats ./daemon/pool.go

echo "Building www files"
cd www
npm install || exit 255

cd ..
echo "Packing application"
bee pack
if [ ! -e "/opt/ubiq-explorer" ]; then
	sudo mkdir -p /opt/ubiq-explorer
fi

echo "Stopping daemon and wallet services"
if [ -e "/etc/systemd/system/blockdaemon.service" ]; then
	sudo service blockdaemon stop
fi

if [ -e "/etc/systemd/system/wallet.service" ]; then
	sudo service wallet stop
fi

if [ -e "/etc/systemd/system/ubiq-api.service" ]; then
	sudo service ubiq-api stop
fi

echo "Deploying packed application"
sudo tar -C /opt/ubiq-explorer -xf ./ubiq-explorer.tar.gz

echo "Creating services"
sudo /bin/cp -af /opt/ubiq-explorer/scripts/systemd/blockdaemon.service /etc/systemd/system/blockdaemon.service
sudo /bin/cp -af /opt/ubiq-explorer/scripts/systemd/wallet.service /etc/systemd/system/wallet.service
sudo /bin/cp -af /opt/ubiq-explorer/scripts/systemd/ubiq-api.service /etc/systemd/system/ubiq-api.service

echo "Creating nginx symlink"
sudo ln -f -s /opt/ubiq-explorer/scripts/nginx.conf /etc/nginx/sites-enabled/ubiq-explorer.nginx

echo "Creating cron"
sudo /bin/cp -af /opt/ubiq-explorer/scripts/cron /etc/cron.d/ubiq-explorer
sudo chown root: /etc/cron.d/ubiq-explorer
sudo chmod 644 /etc/cron.d/ubiq-explorer

echo "Reload systemctl"
sudo systemctl daemon-reload

echo "Enabling wallet and deamon"
sudo systemctl enable wallet
sudo systemctl enable blockdaemon
sudo systemctl enable ubiq-api

echo "Starting services"
sudo service wallet start

while ! curl -s --unix-socket ~/.ubiq/gubiq.ipc http://test >/dev/null; do echo "Waitng for wallet to be ready..."; sleep 1; done

sudo service blockdaemon start
sudo service ubiq-api start
