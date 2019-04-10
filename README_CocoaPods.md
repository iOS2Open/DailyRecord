#### 参考文章 [CocoaPods 都做了什么？](https://www.jianshu.com/p/84936d9344ff)

这篇文章很屌的，把 CocoaPods 的原理梳理了一遍。  
我初略的概括一下我的收获：
1、CocoaPods 使用 ruby 来实现的。  
2、Podfile 文件就是一个 ruby 源代码文件。

#### 文章介绍中的案例
创建一个 Podfile 文件：
```
source 'http://source.git'
platform :ios, '8.0'

target 'HGProject' do

    pod 'AFNetworking'
    pod 'SDWebImage'
    pod 'Masonry'
    pod "Typeset"
    pod 'BlocksKit'
    pod 'Mantle'
    pod 'IQKeyboardManager'
    pod 'IQDropDownTextField'

end
```
模拟一下 CocoaPods 读取这个文件中内容的源代码（eval_pod.rb），如下：
```
$hash_value = {}

def source(url)
    $hash_value['source'] = url
end

def target(target)
    targets = $hash_value['targets']
    targets = [] if targets == nil
    targets << target
    $hash_value['targets'] = targets
    yield if block_given?
end

def platform(platform, version)
    $hash_value['platform'] = platform
    $hash_value['version'] = version
end

def pod(pod)
    pods = $hash_value['pods']
    pods = [] if pods == nil
    pods << pod
    $hash_value['pods'] = pods
end

content = File.read './Podfile'
eval content
p $hash_value
```


终端执行：
> ruby eval_pod.rb 

结果：
```
{"source"=>"http://source.git", "platform"=>:ios, "version"=>"8.0", "targets"=>["HGProject"], "pods"=>["AFNetworking", "SDWebImage", "Masonry", "Typeset", "BlocksKit", "Mantle", "IQKeyboardManager", "IQDropDownTextField"]}
```



### Podfile 常用语法

参考：[Podfile语法参考](https://www.jianshu.com/p/8af475c4f717)  
master 最新:  
```
pod 'AFNetworking', :git => 'https://github.com/gowalla/AFNetworking.git'
```

其他分支:  
```
pod 'AFNetworking', :git => 'https://github.com/gowalla/AFNetworking.git', :branch => 'dev'
```

某个tag:  
```
pod 'AFNetworking', :git => 'https://github.com/gowalla/AFNetworking.git', :tag => '0.7.0'
```


提交记录:  
```
pod 'AFNetworking', :git => 'https://github.com/gowalla/AFNetworking.git', :commit => '082f8319af'
```

项目环境配置：  
```
pod 'DebugTools', :configurations => ['Debug']
```


```
post_install do |installer|
  installer.pods_project.targets.each do |target|
    target.build_configurations.each do |config|
      config.build_settings['SWIFT_VERSION'] = '3.1'
    end
  end
end
```
