go build .

sudo docker kill $(sudo docker ps -a -q)

sudo docker rm $(sudo docker ps -a -q)

sudo docker build -t audit-log:latest .

sudo docker run -p 3000:3000 -d audit-log