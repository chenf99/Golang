# Go语言第一次代码作业

## 运行方法

```bash
go get github.com/spf13/pflag
go install selpg.go (需要先配置GOBIN)
selpg -sstart_page -eend_page [options] (需要配置PATH)
```
## 测试案例


1. `$ selpg -s1 -e1 < input_file`
```bash
$cat input
asd
sd
ca

$selpg -s1 -e1 input 
asd
sd
ca

```

2. `$ other_command | selpg -s10 -e20`
```bash
$cat input | selpg -s1 -e20
asd
sd
ca

selpg: end_page (20) greater than total pages (1), less output than expected
```

3. `$ selpg -s10 -e20 input_file >output_file`
```bash
$touch output
$selpg -s1 -e1 input > output 
$cat output 
asd
sd
ca
```

4. `$ selpg -s10 -e20 input_file >output_file 2>error_file`
```bash
$touch error
$selpg -s0 -e2 input > output 2>error 
$cat error
selpg: you must input valid startPage and endPage(>=0).
```

5. `$ selpg -s10 -e20 -l66 input_file`
```bash
$rm input 
$ echo -e "line1\nline2\nline3\nline4" > input
$ cat input
line1
line2
line3
line4
$ selpg -s2 -e4 -l1 input 
line2
line3
line4
```

6. `$ selpg -s10 -e20 -f input_file`
```bash
$rm input 
$ echo -e "page1\fpage2\fpage3\fpage4" > input
$ cat input
page1
     page2
          page3
               page4
$ selpg -s1 -e3 -f input 
page1
     page2
          
```
