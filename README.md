#  Environmental preparation

```
Get executable program
```
```
Download source code
```
```
Compile the corresponding parallel chain software executable program according to your own operating system
```
```
Executable packages contain
```

### windows
```
Gxl.exe -- parallel chain node program
```
```
Gxl-cli.exe -- parallel chain node command line tool
```
```
Gxl.para.toml -- parallel chain configuration file
```

### linux
```
GXL -- parallel chain node program
```
```
GXL cli -- parallel chain node command line tool
```
```
Gxl.para.toml -- parallel chain configuration file
```

### configuration file

[rpc]

### The jsonrpc and grpc addresses of parallel chains can be customized
```
jrpcBindAddr=":8901"
grpcBindAddr=":8902"
```

### Start node
### Window environment command
```
GXL.exe -f GXL.para.toml
```
### Linux environment commands
```
nohup ./GXL -f GXL.para.toml >/dev/null 2>&amp;1 &amp;
```

### View process
```
ps -ef | grep -v grep | grep GXL
```

### If the process has started, execute the command to query GXL network information
