# This is homework of TWS : Merchant's Guide to the Galaxy
环境：
   golang 1.8 以上

   正确配置 GOPATH

   在 GOPATH 目录下解压或下载本代码

Run
    直接执行 go run main.go

注：
   测试数据在主目录下的文本 galaxy.txt，可以直接修改

代码说明：
1. 我的程序定义了两个小的模块，一个是 roman，作为最基本的罗马字符和数字之间进行转换的工具库
   另一个是 parse_roman, 主要是处理和分析读取的每一行输入，以及计算产生的结果
2. 两个模块的单测可以直接运行，通过在该模块下直接运行 go test
3. 最终程序的输出如下：
    ------OUTPUT-----
    pish tegj glob glob is 42
    glob prok Silver is 68 Credits
    glob prok Gold is 57800 Credits
    glob prok Iron is 780 Credits
    I have no idea what you are talking about
