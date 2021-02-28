# go部署

## 参考文件

学长给的[参考文件](https://laily.feishu.cn/docs/doccnuT12i9IJWMrZbuS7ckyBjh)

## 步骤

* 确定自己的部署后的文件放在那个操作系统中,以及本地的操作系统
* 选择对应的条件进行编译
* 编译完成放在服务器中

我常用的命令

```bash
# Win
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

本地是win

服务器是centos

## 问题

1. ![image-20210223092505765](https://i.loli.net/2021/02/23/jTzaKOdMtoI9cHh.png)==win中需要用set== 

2. ![image-20210223092601439](https://i.loli.net/2021/02/23/uxZWykFsvCjon32.png)

   原因:

   ![image-20210223092652508](https://i.loli.net/2021/02/23/piWgLTDfO8seqxz.png)

   做法:

   ==在文件夹内写一个bat文件把上面的set命令写在里面,然后双击点击,在去`go build` 就会生成一个main(没有后缀)文件==

3. 问题:

   > 编译后的文件,部署上了服务器后,如果服务起不来,也就是无法调用api

   原因:

   * 本地无法访问 ------------->测试:`curl <127.0.0.1:8080>` 
   * 防火墙没关---------------->[教程](https://blog.csdn.net/u011846257/article/details/54707864)
   * 服务器端口没有开放
   * 宝塔没开端口

   总之一句话:就是端口被墙了

