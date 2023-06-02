sudo dockerd

go build .

sudo docker kill $(sudo docker ps -a -q)    # Switch this to docker kill container ID of audit-log if there are other existing docker containers

sudo docker rm $(sudo docker ps -a -q)      # Switch this to docker rm container ID of audit-log if there are other existing docker containers 

sudo docker build -t audit-log:latest .

sudo docker run -p 3000:3000 -d audit-log