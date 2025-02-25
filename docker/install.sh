#!/bin/bash

sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

yum update -y
yum install -y net-tools vim curl git epel-release wget
# yum install -y iputils-ping mariadb-server npm nodejs
# yum install -y zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel gcc make libffi-devel wget mysql-devel centos-release-scl


npm install -g @angular/cli
npm install -g create-react-app
npm install -g @vue/cli

wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
rm -f go1.24.0.linux-amd64.tar.gz

# cat << EOF > /home/readme.txt
# # angular
# ng new test_angular
# cd test_angular
# ng serve --host 0.0.0.0

# # react
# create-react-app test_react
# cd test_react
# npm start

# # vue
# vue create test_vue
# cd test_vue
# npm run serve
# EOF

