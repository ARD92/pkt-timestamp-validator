Simple tool to validate packet timestamp with nanosecond resolution 

## Usage
```
root@server:~/TS_Validate# ./TSvalidate --help
usage: Required-args [-h|--help] -i|--intf "<value>"


============
packet-timestamp-validator
============

Arguments:

  -h  --help  Print help information
  -i  --intf  interface to bind to


root@server:~/TS_Validate# ./TSvalidate -i eth1
Timestamp: 1696598801973156000
Timestamp: 1696598801973221000
Timestamp: 1696598802965295000
Timestamp: 1696598802965349000
Timestamp: 1696598803965975000
Timestamp: 1696598803966029000
Timestamp: 1696598804971462000
Timestamp: 1696598804971511000
```
