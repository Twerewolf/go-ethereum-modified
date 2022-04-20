使用此版本geth需要运行build/bin/geth，每次修改geth项目的某部分都需要重新编译，在项目根目录下make geth
- 会根据makefile的geth部分运行  
geth: (此处为什么代码不能显示出来呢)
	$(GORUN) build/ci.go install ./cmd/geth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."
  
设置GO环境，使用build/ci.go 来安装geth（应该即编译）  
The ci command is called from Continuous Integration scripts. 

### 22/4/19
1. 首先在服务器上安装时出现关于ed25519.go算法的问题，最终通过更新golang版本得以解决；  
原本1.12或1.16都存在问题，更新到1.18.1问题消失。  
更新方法参考golang.google.cn/doc/install，删除原文件再解压安装包到原位置。  
2. 在根目录执行make geth 遇到问题：  
accounts/accounts.go:24:2: import "github.com/ethereum/go-ethereum" is a program, not an importable package
util.go:46: exit status 1
exit status 1
make: *** [Makefile:12: geth] Error 1

在自己上传的go-ethereum版本里面才遇到此问题，clone到服务器的官方版本可以正常make
找到问题：之前在自己创建的go-ethereum-modified项目中最开始随手上传了一个hashzero.go，其中package是main。
最终影响了程序中的其他关系，虽不知整体原因。
删除后得以正常编译make geth。


### **22/4/20**

installation instruction:

1. please download golang(go) language first, detailed at: https://golang.google.cn/doc/install;
2. get go-ethereum-modified project from github.com/Twerewolf;
3. run ****make geth**** at the root of project;
    1. if downloading package meets some trouble, might try change proxy use: 
        
        `$export GOPROXY="[https://goproxy.cn](https://goproxy.cn/)"`
        
4. use scripts from tjw/privatechain ,detailed info from the readme in that dir.