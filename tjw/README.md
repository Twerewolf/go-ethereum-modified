使用此版本geth需要运行build/bin/geth，每次修改geth项目的某部分都需要重新编译，在项目根目录下make geth
- 会根据makefile的geth部分运行  
geth: (此处为什么代码不能显示出来呢)
	$(GORUN) build/ci.go install ./cmd/geth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."
  
设置GO环境，使用build/ci.go 来安装geth（应该即编译）  
The ci command is called from Continuous Integration scripts. 