#!/bin/bash

set -e

echo -e "Deleting 'dist' directory...\n"
rm -rf dist

CGO_ENABLED=0 go build -o dist/its .

npm -C web-frontend run build

echo -e "\nCopying frontend app assets...\n"
cp -r pb_public dist/pb_public

echo -e "\nUploading binaries and frontend app assets...\n"
scp -r ./dist/pb_public christian@83.171.248.78:/home/christian/projects/trading-stats/production/pb_public_new
scp -r ./dist/its christian@83.171.248.78:/home/christian/projects/trading-stats/production/its_new

ssh christian@83.171.248.78 'rm -rf ~/projects/trading-stats/production/pb_public && mv -f ~/projects/trading-stats/production/pb_public_new ~/projects/trading-stats/production/pb_public'  
ssh christian@83.171.248.78 'mv -f ~/projects/trading-stats/production/its_new ~/projects/trading-stats/production/its'
ssh christian@83.171.248.78 'chmod +x ~/projects/trading-stats/production/its'

echo -e "\nSuccessfully uploaded!\n"
