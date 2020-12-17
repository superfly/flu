# flu - FLy Utilities

Flu is the start of a utility command line tool for Fly developers. It is centered around managing your fly.toml files and other local assets. Fly/Flyctl is the core tool for interacting with Fly itself. 

This is an early release to gague usefulness.

Currently implemented commands: 1

flu ls

Scans the directories immediatly below the current directory and reports on the app names in any fly.toml file. e.g.

```
❯ flu ls
APPNAME           	DIRECTORY NAME
6pn-demo          	6pn-demo
appkata-6pn-nats  	appkata-6pn-nats
appkata-flyterm   	appkata-flyterm
appkata-gogs      	appkata-gogs
appkata-gogs      	appkata-gogsplus
appkata-graphql   	appkata-graphql
appkata-minio     	appkata-minio
appkata-mqtt      	appkata-mqtt
appkata-nats      	appkata-nats
appkata-node-red  	appkata-node-red
appkata-redistls  	appkata-redistls
appkata-theia-go  	appkata-theia
long-dawn-1446    	appkata-vscode-remote
```

Optional flag `-r`/`--recurse` - will recurse down the directories too.

Pull Requests and Issues welcome - Dicussions on [community.fly.io](https://community.fly.io/t/looking-for-feedback-on-new-tool/468)
