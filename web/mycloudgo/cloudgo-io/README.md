# CLoudGo-io

设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

# 任务要求
1. 支持静态文件服务
2. 支持简单 js 访问
3. 提交表单，并输出一个表格
4. 对 /unknown 给出开发中的提示，返回码 5xx

# 静态文件服务
把path以"/static/"前缀的URL定位到`webRoot + "/assets/`为虚拟根目录的文件系统
在assets文件夹中不添加index.html，可以看到静态文件目录:
[](images/static.png)
