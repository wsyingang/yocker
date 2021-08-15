# userNamespace 
process in different user namespace can show different UserId and UserGroupId
## example
for example,process A in UserNameSpace u1:userId=100,userGroupid=200
and in UserNameSpace u2:userId=200,userGroupId=400
## USE
An User but not root can create a UserNameSpace u1,in this namespace it can
perform as a root.
## demo code


## error solve
some error may error
```xml
fork/exec /usr/bin/sh: operation not permitted

```
https://blog.csdn.net/shida_csdn/article/details/84649669