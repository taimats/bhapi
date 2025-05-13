cd ${{ secrets.SERVER_DEPLOY_DIR }}
git pull origin main
go get . && go mod tidy
go clean -cache && go build -o ${{ secrets.SERVER_EXECUTE_PATH }} -trimpath -ldflags "-w -s"
chmod +x ${{ secrets.SERVER_EXECUTE_PATH }}
nohup ${{ secrets.SERVER_EXECUTE_PATH }} & > "${HOME}/app_$(date '+%Y_%m_%d_%H:%M:%S').log"