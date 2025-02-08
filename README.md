# markdowneditor
本地部署的markdown编辑器，支持截图保存

### 部署方式

#### 1 编译main.go

> go build main.go -o mymds

注：端口可以自行修改

#### 2 服务器安装
将编译好的程序文件（mymds文件）,放到指定的目录，如\~/mymd/下，
将完整的static目录也放到~/mymd/下（和mymds同层目录），类似：

```
~/mymd/
├── static/
│   ├── ... 
├── mymds
```

#### 3 启动服务
nohup ./mymds & 

#### 4 访问服务
浏览器访问 http://xxxx:18080/
注：服务器地址和端口依据实际情况
