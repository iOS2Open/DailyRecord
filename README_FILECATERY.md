## 文件分类
主要用于 dump 其他 APP 之后， 将头文件分类：  

```
#!/usr/bin/env python3
# -*- coding: utf-8 -*-


import os
import shutil

lists = ["Item", "Service", "Info", "ViewController", "Model", "Cell", "Data", "Manager", "Tool", "Button", "View"]

print(lists)

for folderName, subfolders, filenames in os.walk('/Users/zhuhong/Desktop/Security/frida-ios-dump/Payload2/QQMains_0'):
    for filename in filenames:
        for name in lists:
            if name in filename:
                try:
                    shutil.move(folderName+'/'+filename, '/Users/zhuhong/Desktop/Security/frida-ios-dump/Payload2/QQMains_0/' + name + '/'+filename)
                except OSError:
                    print('OSError')

    
```
