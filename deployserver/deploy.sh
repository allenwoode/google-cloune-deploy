#! /bin/sh

kill -9 $(pgrep server)
git pull https://github.com/allenwoode/google-cloune-deploy.git
cd ~/google-cloune-deploy
./server/server &