#! /bin/sh

kill -9 $(pgrep server)
cd ~/google-cloune-deploy
git pull
./server/server &