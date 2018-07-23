#! /bin/sh

kill -9 $(pgrep server)

cd ~/google-cloune-deploy

git pull https://github.com/allenwoode/google-cloune-deploy.git

./server/server &

echo "relanch server"
