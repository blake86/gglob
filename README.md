# GGLOB

gglob用于将符合特定规则的字符串，展开成列表形式，从功能上类似于glob操作，目前支持的规则有："[]"用来划定规则范围，"-"进行范围匹配，","表示“或者”。

通过gglob的expand操作，可以使得服务器列表管理更为便捷，能够显著缩短配置文件长度，提升可读性，减少出错概率

## Usage
1.连字符
```go
Expand("prefix[1-3]suffix")
```
result:
```
prefix1suffix 
prefix2suffix 
prefix3suffix
```

2.连字符 && ,
```go
Expand("prefix[1-3,5]suffix")
```
result:
```
prefix1suffix 
prefix2suffix 
prefix3suffix
prefix5suffix
```

3.多模式展开 && 零补齐
```go
Expand("prefix[1-3]mid[004-7]suffix")
``` 
result:
```
prefix1mid004suffix
prefix1mid005suffix
prefix1mid006suffix
prefix1mid007suffix
prefix2mid004suffix
prefix2mid005suffix
prefix2mid006suffix
prefix2mid007suffix
prefix3mid004suffix
prefix3mid005suffix
prefix3mid006suffix
prefix3mid007suffix
```