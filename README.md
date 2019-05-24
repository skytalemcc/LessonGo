# LessonGo
本项目是提供个人练习Golang语言使用，练习内容包含基本语法和自带库及第三方库及框架等。形成学习课程和常用的脚手架代码库。  
基于Golang 1.12.4 ,通过windows vscode 在windows golang docker容器中进行编码学习，并进行提交代码。  
### 0.0.1 建立学习环境
1，windows机器安装windows版本的docker ，并下载最新版本的golang镜像 docker pull golang 。  
2，启动windows golang docker后，对Vscode insider版本安装Visual Studio Code Remote - Containers 组件。然后通过Vscode查看-命令面板菜单--Remote-containers :Attach to the running container 来登陆选中的golang容器。  
3，进入golang容器内部，创建git仓库。  
   ```  
        git config --global user.name "skytalemcc"  
        git config --global user.email "79785200@163.com"  
        ssh-keygen -t rsa -C "79785200@163.com"  
        将/root/.ssh/id_rsa.pub内容添加到github个人账号的ssh keys，然后执行ssh  -T git@github.com来建立git账户连通性。  
        验证后进入/go/src 目录：  
        执行 git init  
            git clone https://github.com/skytalemcc/LessonGo.git  
            git add README.md  
            git commit -m "init to add the contex and check if it works"  
            git push origin master  
   ```  
### 0.0.2 Golang学习内容
1  
golang学习指南，学习golang的基本语法。   
https://tour.go-zh.org/welcome/1  
计划时间：2019-05-15(Wed) to 2019-05-19(Sun)  
完成时间：2019-05-15(Wed) to 2019-05-23(Thu)  
2  
golang高级特性，学习golang的高级语法。    
https://books.studygolang.com/gopl-zh/ go语言圣经  
https://books.studygolang.com/advanced-go-programming-book/ go语言高级编程  
https://books.studygolang.com/Mastering_Go_ZH_CN/  
https://books.studygolang.com/go42/  
https://books.studygolang.com/Go-Blog-In-Action/  
https://books.studygolang.com/The-Golang-Standard-Library-by-Example/  
https://studygolang.com/subject/2  golang系列教程  (完成)  
https://studygolang.com/subject/1  golang中文翻译组  
https://studygolang.com/subject/3  深入理解golang  
计划时间：2019-05-24(Fri) to 2019-05-27(Mon)  
完成时间：  

