# Deployments with Kubernetes

本文将介绍如何使用 kubernetes 部署微服务，包括 服务发现，监控，路由，日志。用实际的例子来演示自动化流程。主要分为以下几个部分:

1. 5分钟搭建 K8S 集群
2. 部署 CNI 网络
3. 部署监控服务
4. 部署网关
5. 部署日志服务
6. 部署一个应用


## 5分钟搭建 K8S 集群

第一次完全手动搭建集群大约花了一周时间，主要的问题是在于  

1. K8S的组件多，每个程序的参数有不少，哪些是关键的参数需要花时间搞清楚。
2. 万恶的墙，代理访问外网比较慢
3. CNI网络问题，主要是 CNI 网段和云上的局域网网段冲突了，基础知识缺失导致
4. K8S 的证书和验证方式不清楚

可以参考我之前的[博文]()，即便是完全熟悉部署流程，不写脚本的情况下，如果纯手动 setup 或者 tear down 一个集群，都是比较耗时间的。

直到，发现了这个工具 kubeadm, 世界美好了。

这个工具对操作系统有限制， ubuntu 16.4 或 centos 7 以上。其实当初也看到了这个工具， 不过 因为系统限制，并且kubeadm还在alpha版本，又想手动撸一遍部署过程，所以没直接采用。 不过 kubeadm 不建议在生产环境中使用，在 官方文档中的 limitation 中有详细解释.

[文档](https://kubernetes.io/docs/getting-started-guides/kubeadm/) 中第一点就说了， kubeadm部署的是 single master，意味着不是高可用，谨慎使用。 但是作为演示实例再合适不过。

开始部署步骤:

1. 在 Digital Ocean 中开三台机器, centos 7，建议2C2G，按小时计费用不了多少钱，用完就销毁。 如果还没有注册账号，并且觉得本文对你有帮助，可以用我的 referral link 注册，可以得到 10美金, [链接](https://m.do.co/c/821db079aed2)

2. 登录三台机器，安装必要组件.
	```
	yum clean
	yum update
	cat <<EOF > /etc/yum.repos.d/kubernetes.repo
	[kubernetes]
	name=Kubernetes
	baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
	enabled=1
	gpgcheck=1
	repo_gpgcheck=1
	gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
		https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
	EOF
	setenforce 0
	yum makecache
	yum install -y docker kubelet kubeadm kubectl kubernetes-cni
	systemctl enable docker && systemctl start docker
	systemctl enable kubelet && systemctl start kubelet
	```
3. 选择一台作为master, 运行
	```
	kubeadm init
	```
	完成后会看到提示: `kubeadm join --token=311971.7260777a25d70ac8 192.168.1.100`
4. 在其他两台机器上分别运行以上提示的命令
5. 在 master 上查看状态, `kubectl get nodes`, 如果看到一共有2个node，一个master， 则表示集群创建成功。

## 部署CNI网络

kubeadm 自动部署了一个插件，就是 kube-dns, 用于服务发现，但是到这里你会发现 kube-dns 这个服务没有启动成功，因为我们还没有部署CNI网络。

```
kubectl get pods --all-namespaces | grep dns
```

这里有比较多的选择，我使用了 calico，因为性能比较好，支持一键部署。 这里有一篇对比容器网络的文章，优缺点介绍比较全面， [Battlefield: Calico, Flannel, Weave and Docker Overlay Network](http://chunqi.li/2015/11/15/Battlefield-Calico-Flannel-Weave-and-Docker-Overlay-Network/)

配置文件在cni目录下，或者可以直接在master运行： 
`kubectl apply -f http://docs.projectcalico.org/v2.0/getting-started/kubernetes/installation/hosted/kubeadm/calico.yaml`

再次查看 dns 服务是否运行成功吧。

## 监控

### Prometheus



### Grafana

