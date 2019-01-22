## 参考文章
[python_gitpython_git](https://blog.csdn.net/AnYuanLzh/article/details/78803432)  
[使用 Python 操作 Git 版本库](https://www.aliyun.com/jiaocheng/434487.html)  
[GitPython](https://github.com/gitpython-developers/GitPython)

## 一、安装
> pip install GitPython

然后就可以直接使用了。

## 二、使用
### 2.1 导入
```python
#-*-coding:utf8-*-  
from git import Repo 
```


### 2.2 获取本地的仓库
```python
repoPath = './gitdir'
repo = Repo(repoPath)
```

### 2.3 简单操作
```python
print(repo.branches) 
print(repo.active_branch)   # 当前活动分支
```

### 2.4 pull 与 push
```python
remote = repo.remote()
remote.pull()
remote.push()
```

## 三、总结
相当的简单了。
