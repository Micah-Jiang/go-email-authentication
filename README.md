# go-email-authentication
## 1. introduction
The system of user login email authentication based on go lang

## 2. design ideas
user input username， password and eamil.
user click button to get code. the code is exist in the redis? if tue, try it agin after 1 minute.
for another, send code to user email, store code to redis, set the expired time for one minute.

## 3. technical framework
- mvc framework: gin
- language: go lang
- orm framework: gorm
- database: mysql
- cache: redis

## 4. feature
### (1) 针对password采用BCrypt算法进行加密
> BCrypt算法: 一种加盐的单向 Hash，不可逆的加密算法，同一种明文，每次加密后的密文都不一样，而且不可反向破解生成明文，破解难度很大。每次加密的时候首先会生成一个随机数就是盐，之后将这个盐值与明文密码进行 hash，得到 一个hash值存到数据库中。其中生成的 hash 值中包含了之前生成的盐值(22个字符)，用于后续 hash 值验证。
> MD5 加密：是不加盐的单向 Hash，不可逆的加密算法，同一个密码经过 hash 的时候生成的是同一个 hash 值，在大多数的情况下，有些经过 md5 加密的方法将会被破解。
> 密码破解的方法：
> - 暴力破解法
> - 大型数据库： 明文 -> 密文的映射
> - 彩虹表： 

