# mycode (GoLang)

GOLang training

## Getting Started

These instructions will get you a copy of the project up and running on your local machine
for development and testing purposes. See deployment for notes on how to deploy the project
on a live system.

### Prerequisites
synch b/w local and git repo:

	cd /home/student/.ssh
	ssh-keygen -f id_rsa_github
	Generates public/private rsa key pair.
	copy id_rsa_github.pub content to git hub project settings - Deploy keys.

Install Go on linux machine:

	wget -c https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
	echo "export PATH=$PATH:/usr/local/go/bin" > .bash_training/gopath
	source ~/.profile
	go version

## Built With

* [GoLang](https://go.dev/doc/effective_go) - The coding language used

#Compile and Run
  go run <name>.go
#compile
  go vet <name>.go

## Authors

* **Lavanya** - *GOLang* - [YourWebsite](https://github.com/chvlavanya/mycode)

