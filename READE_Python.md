## Python 使用小技能

创建一个目录为 **HG**。在这个目录中创建一个 **__init__.py**，内容如下：
```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

def hgTest():
    print('Hello HG')
```

在与 **HG** 目录同级的目录中创建一个 **test.py**。内容如下：
```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from infer import *

hgTest()
```

或者
```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import infer

infer.hgTest()
```

## 注意  
要注意在 from 的时候，如果目录很深的话，目录之间使用 **.** 连接，而不是 **/**。

## 错误
### 错误1
> TabError: inconsistent use of tabs and spaces in indentation  

解释：    
> Python文件运行时报TabError: inconsistent use of tabs and spaces in indentation 原因:说明Python文件中混有Tab和Space用作格式缩进。

### 错误二
> TypeError: write() argument 1 must be unicode, not str

解释：  
[python提示错误TypeError: write() argument must be str, not bytes](https://blog.csdn.net/qq_33363973/article/details/77881168)

### 错误三
> UnicodeDecodeError: 'ascii' codec can't decode byte 0xe6 in position 2: ordi

解释：  
[解决UnicodeDecodeError: 'ascii' codec can't decode byte 0xba in position 31: ordinal not in range(128)](https://blog.csdn.net/jiayk2016/article/details/79393067)

这个问题在 **python3** 中不会出现。

### 错误四
> TypeError: 'encoding' is an invalid keyword argument for this function

解释：  
[python 2.7版本解决TypeError: 'encoding' is an invalid keyword argument for this function](https://www.cnblogs.com/milian0711/p/7132377.html)

这个问题在 **python3** 中不会出现。

