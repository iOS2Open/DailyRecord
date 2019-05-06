```shell
HG:usbmuxd zhuhong$ sh login.sh 
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the RSA key sent by the remote host is
SHA256:6SyfwOIv+Miy6M50jCvdl/Wn19nFFBcYuBizeI1yDsY.
Please contact your system administrator.
Add correct host key in /Users/zhuhong/.ssh/known_hosts to get rid of this message.
Offending RSA key in /Users/zhuhong/.ssh/known_hosts:3
RSA host key for [localhost]:10000 has changed and you have requested strict checking.
Host key verification failed.
```

解决方案是：  
先这样的盘她 ：   
```shell
open ~/.ssh/known_hosts
```

删除对应公钥  

然后再盘她：   

```shell
ssh-keygen -F slave
```
