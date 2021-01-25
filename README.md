# devops

devops config

### config deploy 占位符说明

```
# jenkinsfile

${hub-group}：镜像仓库项目名称
${image-name}：镜像名称
${git-group}：gitlab分组
${git-name}：gitlab项目名称

# deploy
${ns-name}: KS项目名称
${app-name}：KS应用名称
${deploy-name}：工作负载名称
${deploy-desc}：工作负载描述

```

### devops 工程创建步骤

```
1. 创建凭证
harbor-id：harbor账号密码
gitlab-id：gitlab账号密码
kube-id：kubeconfig
2. 创建DevOps工程
3. 将本项目对应的文件，拷贝到项目根目录，替换相应的占位符，提交代码
4. 扫描远程分支
```

### 项目相关准备

```
1. 创建私有仓库harbor

在配置中心 -> 秘钥中创建，选择镜像仓库秘钥，输入仓库地址和账号密码

2. 针对Java项目，添加jvm配置

在配置中心 -> 配置中创建，名称为统一jvm，配置如下
  JAVA_OPTS: >-
    -Xmn120m -Xms256m -Xmx768m -XX:MaxMetaspaceSize=128m
    -XX:CompressedClassSpaceSize=24m
  JAVA_OPTS_LG: >-
    -Xmn120m -Xms256m -Xmx1500m -XX:MaxMetaspaceSize=128m
    -XX:CompressedClassSpaceSize=24m
  JAVA_OPTS_SM: '-Xms64m -Xmx384m -XX:MaxMetaspaceSize=128m -XX:CompressedClassSpaceSize=24m'
```

### 工具使用

```
# Windows
> devops.exe
# 根据提示输入

```
