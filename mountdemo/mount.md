# MountNamespace 
mount namespace refer the process view of file system.
## view all proc
```shell
ls /proc
```
in this operation you will look all proc,but we should look container
```shell
mount -t proc proc /proc
```
then exec
```shell
ls /proc
```
you will find less 
and 
```shell
ps -ef 
```
only very less process is running