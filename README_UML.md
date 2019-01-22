# UMLLearn

### 一、类元
就是在 UML 图中的一个节点，或者说是一个元素。如下所示：  
![image.png](https://upload-images.jianshu.io/upload_images/1198135-e3d4cc153bda043c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 二、关系
上面的那些类元，在实际的 UML 中都有些什么关系呢？如下所示：    
![关系](https://upload-images.jianshu.io/upload_images/1198135-ed936a79ce3e168d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)  
这里要注意的一个问题是，箭头的指向。
![image.png](https://upload-images.jianshu.io/upload_images/1198135-e36f2fe7385416df.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

HGPerson 继承于 NSObject，所以 HGPerson 是 NSObject 的子类。箭头是 **子类指向父（超）类**。

### 三、泛化
实际就是类之间的继承关系，泛化用从子指向父的箭头表示，指向父的是一个空三角形(如图 4-7 表示)。多个泛化关系可以用箭头线组成的树来表示，每一个分支指向一个子类。样例：
![image.png](https://upload-images.jianshu.io/upload_images/1198135-17065e68f6ad041f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 四、实现
实现关系用一条带封闭空箭头的虚线来表示(如图 4-9)，且与泛化的符号很相像。
![image.png](https://upload-images.jianshu.io/upload_images/1198135-37157afbcc7d8c72.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

用一种特殊的折叠符号来表示接口(无内容)以及实现接口的类或构件。接口用一个圆圈表示，它通过实线附在表示类元的矩形上(如图4-10)。

![image.png](https://upload-images.jianshu.io/upload_images/1198135-7cf9c9a0ba32ebdb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 五、依赖
##### 描述
![image.png](https://upload-images.jianshu.io/upload_images/1198135-0699cc22f8c64719.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

##### 种类
![image.png](https://upload-images.jianshu.io/upload_images/1198135-7b1f89da1e30d6ed.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

##### 依赖关系
![image.png](https://upload-images.jianshu.io/upload_images/1198135-9a91f2787a85e26b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 六、约束
以我简朴的理解，约束就是一些通过文案来辅助描述。如下所示：
![image.png](https://upload-images.jianshu.io/upload_images/1198135-e33ead401f6a5bd8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 七、其它
![00main.png](https://upload-images.jianshu.io/upload_images/1198135-822fcd639f661056.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![01main.png](https://upload-images.jianshu.io/upload_images/1198135-dbad90c37466c210.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 八、工具
#### 8.1 OmniGraffle
![OmniGraffle.png](https://upload-images.jianshu.io/upload_images/1198135-bbeb08cef00b7010.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


#### 8.2 StarUML
![StarUML.png](https://upload-images.jianshu.io/upload_images/1198135-79f763150d4a91d8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

 