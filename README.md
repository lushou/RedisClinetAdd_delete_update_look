# 使用说明
## 在redis 进群中进行添加redis 节点
!!! 注意请求建议使用工具进行获取不建议使用命令：# 这边建议使用工具进行传入参数
### 传入数据出现的问题
传入空数据出现如下报错：
![](images/2023-06-29-09-36-12.png)
### redis节点无法进行连接会出现如下报错：
```shell
curl --location --request POST 'http://119.3.229.137:31405/redisadd/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'redisMstart="<redisMstart>"' \
--form 'redisAdd="<redisAdd>"' \
--form 'passWord="<passWord>"'
```
![](images/2023-06-30-14-52-01.png)
### 这边进行传入添加的节点如图所示
```shell
# 这边建议使用工具进行传入参数
curl --location --request POST 'http://119.3.229.137:31405/redisadd/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'redisMstart="<redisMstart>"' \
--form 'redisAdd="<redisAdd>"' \
--form 'passWord="<passWord>"'
```
![](images/2023-06-30-14-47-58.png)
### 这边准备测试正确的日志
添加前
![](images/2023-06-30-16-09-52.png)
添加后
```shell
curl --location --request POST 'http://119.3.229.137:31405/redisadd/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'redisMstart="<redisMstart>"' \
--form 'redisAdd="<redisAdd>"' \
--form 'passWord="<passWord>"'
```
![](images/2023-06-30-14-56-38.png)
![](images/2023-06-30-16-11-05.png)
## 进行删除redis节点
### 注意这边只删除redis 从节点保障系统的正常运转：
### 传入参数前
```shell
curl --location --request POST 'http://119.3.229.137:31405/redisdelete/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'redisMstart="<redisMstart>"' \
--form 'redisDelete="<redisDelete>"' \
--form 'passWord="<passWord>"'
```
### 删除前:
![](images/2023-07-06-16-05-52.png)

### 当删除的节点是主的时候：
![](images/2023-07-06-16-08-36.png)
### 删除从的时候：
![](images/2023-07-06-16-09-06.png)

![](images/2023-07-06-16-09-30.png)

## redis 进行把从节点切换成主节点：
### 传入参数：
```shell
curl --location --request POST 'http://119.3.229.137:31405/redisfailover/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'redisSlave="<redisSlave>"' \
--form 'passWord="<passWord>"'
```
### 执行如下：
![](images/2023-07-06-16-01-39.png)
![](images/2023-07-06-16-00-51.png)