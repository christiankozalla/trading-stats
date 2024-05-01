#!/bin/bash

set -e  # This command causes the script to exit immediately if any command exits with a non-zero status.

echo -e "Deleting 'dist' directory...\n"
rm -rf dist

CGO_ENABLED=0 go build -o dist/its .

cd web-client

npm run build

cd ..

echo -e "\nCopying frontend app assets...\n"
cp -r pb_public dist/pb_public

echo -e "\nUploading binaries and frontend app assets...\n"
scp -r ./dist/pb_public christian@83.171.248.78:/home/christian/projects/trading-stats/production/pb_public_new
scp -r ./dist/its christian@83.171.248.78:/home/christian/projects/trading-stats/production/its_new

ssh christian@83.171.248.78 'rm -rf ~/projects/trading-stats/production/pb_public && mv -f ~/projects/trading-stats/production/pb_public_new ~/projects/trading-stats/production/pb_public'  
ssh christian@83.171.248.78 'mv -f ~/projects/trading-stats/production/its_new ~/projects/trading-stats/production/its'  

echo -e "\nSuccessfully uploaded!\n"
