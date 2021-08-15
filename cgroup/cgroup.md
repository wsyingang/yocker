# Cgroups
Cgroup=Linux Control Group. Use for 

## look support subSystem
```shell
lssubsys -a
```
## Three Component
### cgroup
a kind of group by process
### subsystem
subsystem will link to a cgroup use for restrict process ,like memory restrict,disk restriction
### hierarchy
hierarchy is a tree-like cgroup
use for hierarchy Cgroup can inherite .
## Example for hierarchy
```shell
mkdir cgroup-test
```
```shell
 sudo mount -t cgroup -o none,name=cgroup-test cgroup-test ./cgroup-test
```
## understand
Up to now ,I think hierarchy is a tree,which has many node. each node is a cgroup.
As for each cgroup,It's function is depend on his subsystem,
## Move process between cgroup
(1)get terminal process
```shell
echo $$
```
for my example it is 4013

then do this operation:
```shell
sudo sh -c "echo $$ >> tasks"
cat /proc/4013/cgroup 
```
you will get
```text
13:name=cgroup-test:/cgroup-1
12:perf_event:/
11:pids:/user.slice/user-1000.slice/user@1000.service
10:freezer:/
9:rdma:/
8:cpuset:/
7:devices:/user.slice
6:cpu,cpuacct:/
5:memory:/user.slice/user-1000.slice/user@1000.service
4:hugetlb:/
3:net_cls,net_prio:/
2:blkio:/
1:name=systemd:/user.slice/user-1000.slice/user@1000.service/apps.slice/apps-org.gnome.Terminal.slice/vte-spawn-a05fd39a-12e6-40ce-aa23-185e8053aa0e.scope
0::/user.slice/user-1000.slice/user@1000.service/apps.slice/apps-org.gnome.Terminal.slice/vte-spawn-a05fd39a-12e6-40ce-aa23-185e8053aa0e.scope
```
cgroup-1 already effect.
## restrict resource
when we create hierarchy,we just create cgroup without restriction.In this step
I will link the subsystem.
(1)get memory restriction
/sys/fs/cgroup/memory
bootstarp stress process:
```shell
stress --vm-bytes 200m --vm-keep -m 1
```
then you can top the procss info
then
(1) create cgroup
```shell
sudo mkdir test-limit-memory
cd test-limit-memory
```
(2) write limit
```shell
sudo sh -c "echo "100m" > memory.limit_in_bytes"
```
(3)write pid
```shell
sudo sh -c "echo $$ >tasks"
```
## How docker used
```shell
docker run -m 'memory limit'
```
