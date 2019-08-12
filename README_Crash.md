找到本机 `symbolicatecrash` 的位置  
```
find /Applications -name symbolicatecrash -type f  
```

开始解析 crash 堆栈  

```
./symbolicatecrash x.crash x.dSYM > log.crash
```


Error: "DEVELOPER_DIR" is not defined at ./symbolicatecrash line 69.


```
export DEVELOPER_DIR="/Applications/Xcode.app/Contents/Developer"
```


/Applications/Xcode.app/Contents/SharedFrameworks/DVTFoundation.framework/Versions/A/Resources
