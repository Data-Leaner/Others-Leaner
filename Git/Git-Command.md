## Git 命令清单

### 设置参数
```bash
# 设置全局
git config --global user.name 'ray'
git config --global user.email 'xxxxx@xxx.com'

# 设置当前库
git config --local user.name 'ray'
git config --local user.email 'xxxxx@xxx.com'

# 设置系统级别
git config --system user.name 'ray'
git config --system user.email 'xxxxx@xxx.com'
```

### 初始化命令
```bash
# 进入仓库目录
cd directory_nm
# 新创建的项目
git init your_project_nm
# 如果已经存在的项目
git init
```

### 添加暂存区
```bash
# 指定添加某一个文件到暂存区
git add file_nm
# 指定添加所有变更或者新增文件到暂存区
git add .
```

### 提交到版本库
```bash
git commit -m'注释信息'
```
