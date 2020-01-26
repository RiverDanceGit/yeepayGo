# 生成公私钥对

#### 自行生成
生成私钥
```
openssl genrsa -out app_private_key.pem 2048
```

生成公钥
```
openssl rsa -in app_private_key.pem -pubout -out app_public_key.pem
```

Java开发者需要将私钥转换成 PKCS8 格式
```
openssl pkcs8 -topk8 -inform PEM -in app_private_key.pem -outform PEM -nocrypt -out app_private_key_pkcs8.pem
```

#### PKCS1 秘钥解析，适用C#、PHP、Python
将下载的 p12 或 pfx 证书文件转化成 temp.key 文件
```
openssl pkcs12 -in 10000466938.p12 -nocerts -nodes -out temp.key
```

导出私钥
```
openssl rsa -in temp.key -out cfca_private_key.pem
```

导出公钥
```
openssl rsa -in temp.key -pubout -out cfca_public_key.pem
```

#### PKCS8 秘钥解析，适用Java、PHP
将下载的 p12 或 pfx 证书文件转化成 temp.key 文件
```
openssl pkcs12 -in 10000466938.p12 -nocerts -nodes -out temp.key
```

导出私钥
```
openssl rsa -in temp.key -out cfca_private_key.pem
```

转化成 PKCS8 格式，并且显示信息
```
openssl pkcs8 -topk8 -inform PEM -in cfca_private_key.pem -outform PEM -nocrypt
```

#### PKCS8 转换 PKCS1（此步骤按需选择，非必须）
```
openssl rsa -in pkcs8.pem -out pri_key.pem
```
