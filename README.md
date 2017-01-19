# 前提
- Google Cloud SDKインストール済み
- kubectl Componentインストール済み
- gcloud configセットアップ済み
- minikubeインストール済み

# ingress検証

## minikube でingress アドオンを有効にする

```
$ minikube addons list
- addon-manager: enabled
- dashboard: enabled
- kube-dns: enabled
- heapster: disabled
- ingress: disabled
- awsecr-creds: disabled
```

addonを有効にする

```
$ minikube addons enable ingress
```

## 検証内容

minikubeでingress(443)⇨replicaset(80)⇨pods(80)
ブラウザでhttps://ingressのアドレスが動作するか

サービスを作る

```
kubectl create -f k8s/service.yml
```

replicastを作る

```
kubectl create -f k8s/rs.yml
```

ingressを作る

```
kubectl create -f k8s/ingress.yml
```


ingressの詳細を見る

```
$ kubectl describe ing simplelb
Name:			simplelb
Namespace:		default
Address:		192.168.99.101
Default backend:	web-server:80 (172.17.0.4:80,172.17.0.5:80)
Rules:
  Host	Path	Backends
  ----	----	--------
  *	* 	web-server:80 (172.17.0.4:80,172.17.0.5:80)
Annotations:
Events:
  FirstSeen	LastSeen	Count	From				SubObjectPath	Type		Reason	Message
  ---------	--------	-----	----				-------------	--------	------	-------
  13m		13m		1	{nginx-ingress-controller }			Normal		CREATE	default/simplelb
  13m		13m		1	{nginx-ingress-controller }			Normal		CREATE	ip: 192.168.99.101
  13m		13m		1	{nginx-ingress-controller }			Normal		UPDATE	default/simplelb
```

Addressに出ているIPへブラウザからアクセスしてみる

```
default backend - 404
```
ごなる

minikube sshし、ifconfigしてみると

```
eth1      Link encap:Ethernet  HWaddr 08:00:27:D7:9F:C4
          inet addr:192.168.99.101  Bcast:192.168.99.255  Mask:255.255.255.0
          inet6 addr: fe80::a00:27ff:fed7:9fc4/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:11842 errors:0 dropped:0 overruns:0 frame:0
          TX packets:8518 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:1604998 (1.5 MiB)  TX bytes:5838792 (5.5 MiB)
```

defaultbackendコンテナーへへルーティングされてしまう
```
2272052282ef        gcr.io/google_containers/defaultbackend:1.0                  "/server"                18 hours ago        Up 18 hours                                                                              k8s_default-http-backend.54c2c7bd_default-http-backend-50f8j_kube-system_1d2a4476-dd52-11e6-b480-080027e92c9f_db67437e
```

minikubeのnode自分自身の80ポートにアクセスしたとみなされる模様。。
ingress自体のIPを変える方法が分からない、、
一旦ingressない方向で他の構成を検討する