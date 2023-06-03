wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz

rm go1.20.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go build .

# Remember to add .env file with environment properties from email env.txt

./audit-log