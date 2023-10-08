# MiniAppContest

back
bot
miniapp(front)

## Setup guide

### build

[Go](https://go.dev/)
[Node.js](https://nodejs.org/en)

`
cd back/
// for  cross-compile use export GOOS="linux" export GOARCH="amd64"
go mod download
go build
`

`
cd bot/
// for cross-compile use export GOOS="linux" export GOARCH="amd64"
go mod download
go build
`

`
cd miniapp/
npm install
// VITE_APP_API in .env.production
npm run build
`

### deploy

#### local

*ngrok.yml*
`
tunnels:
  back:
    proto: http
    addr: 4001
  miniapp:
    proto: http
    addr: 5173
`

`
./ngrok start back miniapp
`

`
./back/back
`

`
cd miniapp/
// ngrok proxy addr for back to VITE_APP_API in .env
npm run dev
`

`
export TOKEN=""  // bot token
export TEST_ENV=""  // any value for test environment, empty for production
export URL=""  // ngrok proxy addr for miniapp
./bot/bot
`

#### hosting

Debian VPS hosting

[nginx](https://nginx.org/)

`
ssh SRV "mkdir -p /mac/back/"
scp back/back SRV:/mac/back/
ssh SRV "chmod +x /mac/back/back"

ssh SRV "mkdir -p /mac/bot/"
scp bot/bot SRV:/mac/bot/
ssh SRV "chmod +x /mac/bot/bot"

ssh SRV "mkdir -p /mac/miniapp/"
scp -r miniapp/dist/* SRV:/mac/miniapp/
scp -r misc/images/ SRV:/mac/miniapp/

scp misc/nginx/miniapp SRV:/etc/nginx/sites-available/
ssh SRV "ln -s /etc/nginx/sites-available/miniapp /etc/nginx/sites-enabled/"
ssh SRV "systemctl nginx restart"
`
`
ssh SRV "rm -f /etc/nginx/sites-enabled/default"
`

`
./back/back
`

`
export TOKEN=""
export TEST_ENV=""
export URL=""
./bot/bot
`

`
scp misc/bot_env SRV:/mac/
scp misc/systemd/miniapp-back.service SRV:/lib/systemd/system/
scp misc/systemd/miniapp-bot.service SRV:/lib/systemd/system/

ssh SRV "systemctl daemon-reload"

ssh SRV "systemctl enable miniapp-back"
ssh SRV "systemctl enable miniapp-bot"

ssh SRV "systemctl start miniapp-back"
ssh SRV "systemctl start miniapp-bot"
`
