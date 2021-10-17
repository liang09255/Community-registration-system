# 社团报名系统说明
项目体验地址：http://39.108.81.214:8080/

前端套用了网上找的两个模板，一个前台一个后台，稍微改了改，然后加了点后端

文件结构有点乱，懒得改了

## 前端文件说明
主目录下的几个html对应的前台几个页面

backstage用于存放后台代码

js下的Personal内代码为自己写的ajax请求，用于与后端通信，大部分采用原生js，少部分利用了jQuey语法

toastr为通知插件，具体用法自行百度

图片文件由于当时怕服务器带宽不够，全都手动压缩了

woff2文档是字体文件，模板好像是国外的，忘记这个字体是要翻墙还是咋滴了，反正测试的时候发现外链经常加载失败，所以就手动下载下来了

前端我自己改动的大概就是这些了，其他大都是模板自带的，很多东西我也不懂，没敢动，反正，能跑就行哈哈哈哈哈

## 后端文件说明
后台基本使用的都是原生Golang，Gin是后面才学的，就没用上了

前台的后端就在golang那个文件夹里，后台功能在backstage里的golang

projectMain主要用于响应前端请求，监听端口为80，可手动改

emailCode里用了gomail包，用于发送邮件，需要自己配置好邮箱才能用

phoneCode用的阿里云的短信接口，按量计费，阿里云短信自带有流控，返回的json不太会处理，就直接匹配了下字符串来判断是否触发流控

pcode用于生成短信验证码并调用phoneCode里的内容发送短信

searchValues用于用户查询报名信息

submitValues用于用户提交报名信息，若数据库中存在则为更新

后台的入口写在了前端查询界面那，登录后保存一个session，检测到session存在才能进行操作后台管理系统的功能

changeStatus用于改变报名状态

getDocument用于报名信息导出为excel

seeSignup用于管理员查看报名信息

sendEmail用于批量发送通知信息

## 使用建议
建议不要使用哈哈哈哈，我只是放在这记录一下

好多东西我都不懂，当时是直接把前端资源扔服务器上，后端直接编译成了一个linux上可运行的文件直接跑的

服务器我是直接用宝塔配好来的，遇到一个bug就是宝塔装好的那个Nginx会占用80端口，要把它关了才能用80，Nginx有啥用我到现在也还没搞清楚哈哈哈

如果非想跑下我这玩意，记得手动配下发送邮件和短信的东西，要不然用不了

数据库建表代码如下

```
CREATE TABLE IF NOT EXISTS `sign_up` (
  `id` int(11) NOT NULL COMMENT '报名表id',
  `my_name` varchar(255) DEFAULT NULL COMMENT '姓名',
  `my_phone` varchar(255) DEFAULT NULL COMMENT '电话',
  `my_email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `my_part` varchar(255) DEFAULT NULL COMMENT '部门',
  `my_about` text COMMENT '自我介绍',
  `my_status` varchar(255) DEFAULT NULL COMMENT '0报名完成；1一面完成；2二面完成；3三面完成；4已被录取'
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
```
