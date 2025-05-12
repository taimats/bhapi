cd ${{ secrets.SERVER_DEPLOY_DIR }}
git pull origin main
go get . && go mod tidy
go clean -cache && go build -o ${{ secrets.SERVER_EXECUTE_PATH }} -trimpath -ldflags "-w -s"
chmod +x ${{ secrets.SERVER_EXECUTE_PATH }}
nohup ${{ secrets.SERVER_EXECUTE_PATH }} &
mv ./nohup.out ${HOME}/