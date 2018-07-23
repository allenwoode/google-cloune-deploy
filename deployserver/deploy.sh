#! /bin/sh

kill -9 $(pgrep server)
cd ~/google-cloud-deploy
git pull
./server/server &