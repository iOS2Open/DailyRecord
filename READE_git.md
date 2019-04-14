参考简书:[为什么你的 Git 仓库变得如此臃肿](https://www.jianshu.com/p/7231b509c279)

https://www.cnblogs.com/lout/p/6111739.html


### 一、垃圾回收  
```
git gc --prune=now
```

6.35G， 变成了 5.89G ？  


### 二、filter-branch    
```
git filter-branch --force --index-filter   'git rm --cached --ignore-unmatch bigfile'   --prune-empty --tag-name-filter cat -- --all
```


窝草~，没有变化 

### 三、
```
git for-each-ref --format='delete %(refname)' refs/original | git update-ref --stdin  
git reflog expire --expire=now --all  
git gc --prune=now  
```

窝草~，没有变化 




