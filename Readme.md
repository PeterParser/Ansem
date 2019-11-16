# Requirements
YAML library
```
go get gopkg.in/yaml.v2
```

# Configuration
The configuration is stored in yaml file with the following (mandatory) attributes:
* **exploits_dir**: Path to the directory that contains the exploits
* **tick**: The tick of the gameserver in seconds
* **team_file**: Path to a plain text file with the ip of the teams
* **gameserver**: IP address and port of the gameserver in the following format "IP:port"
* **workers**: The number of workers
* **submission_type**: The type of submission required by the game server, for now only "TCP" is allowed.

Example:
```yaml
exploits_dir: "/path/to/exploits"
tick: 5
team_file : "/path/to/teams"
gameserver: "127.0.0.1:31337"
workers: 8
submission_type: "TCP"
```