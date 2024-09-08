# EPI's GoData Tools
This project tries to augment GoData with customized tools for EPI.

The project is split into the backend(api) and the frontend.

## Backend
A Go http server that provides the endpoints for retrieving the outbreak cases:
1. `/outbreaks`: returns a list of all outbreaks
2. `/casesByOutbreak`:  returns all cases for a specific outbreak.

## Frontend
This is a React project that communicates with the `Backend` to fetch data.

## GoData Installation
These are the steps for installing GoData in a Compute Engine instance runnign Debian Linux.
It has a base disk of 10GB and 8GB of RAM. The base disk only hosts the Operating System.
We have an additional disk of 200GB in capacity.  This houses the database itself. Best practices dictate
that the operating system and storage be separate disks. This enables us to do quickly boot up another 
server and easily attach the additional disk without the risk of losing data.

- Install `lsof`:  apt-get install lsof. This is required by GoData during its startup process.


- Install certbot
- Install caddyserver:
    - sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
    - curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo tee /etc/apt/trusted.gpg.d/caddy-stable.asc
    - curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
    - sudo apt update
    - sudo apt install caddy
- Edit CaddyFile:
		godata-dev.epi.openstep.bz { 
			reverse_proxy 127.0.0.1:8000
		} 
		api.godatatools.mohw.bz { 
			reverse_proxy 127.0.0.1:9090
		}

