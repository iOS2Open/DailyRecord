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


git verify-pack -v .git/objects/pack/pack-*.idx | sort -k 3 -g | tail -5
361a7a9c3b6f77259bcc3dfa324cd9a0f7c9b40e blob   226825912 225867471 5203499071
245dfde9a8b7509fdc35e1bcfced712fd7b13339 blob   234450792 106992793 2018835956
5f93ef90c96e40babe44f08b1b1ca0d5d4ea93a0 blob   1875291376 763124904 1208048085
6bd818e1734e2ec1e5f5747add441c93ea561392 blob   1875530096 763160682 2331350241
21b77172ed1ef8f0b6cf0fbd2468ef16523eb8e6 blob   1875746760 763242366 10587782
HG: zhuhong$ git rev-list --objects --all | grep 361a7a9c3b6f77259bcc3dfa324cd9a0f7c9b40e
361a7a9c3b6f77259bcc3dfa324cd9a0f7c9b40e Distribution-iphoneos.zip





