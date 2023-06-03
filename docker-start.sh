sudo apt-get update

sudo apt install docker.io -y

sudo dockerd

# sudo docker kill $(sudo docker ps -a -q) &   # Switch this to docker kill container ID of audit-log if there are other existing docker containers

# sudo docker rm $(sudo docker ps -a -q) &     # Switch this to docker rm container ID of audit-log if there are other existing docker containers 

docker pull codingcookierookie/audit-log &

sudo docker run -p 3000:3000 -d codingcookierookie/audit-log &