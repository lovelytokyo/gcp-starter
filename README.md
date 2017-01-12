# 前提
- Google Cloud SDKインストール済み
- kubectl Componentインストール済み
- gcloud configセットアップ済み

# kubernetesを使ってnginx+go構成のpodとserviceを作る

## クラスタを作る
```
gcloud container clusters create lovelytokyo-server
```

## gcloudコマンドで接続
```
gcloud container clusters get-credentials lovelytokyo-server --zone asia-northeast1-a --project cyberagent-018
```

## deploy, service作成(LBも自動作成される)
```
kubectl apply -f k8s/deployment.yml
kubectl apply -f k8s/service.yml
```

## service作成結果確認
```
kubectl describe services web-server
Name:			web-server
Namespace:		default
Labels:			name=web-server
Selector:		name=web-server
Type:			LoadBalancer
IP:			10.215.251.53
LoadBalancer Ingress:	104.198.112.212
Port:			http	80/TCP
NodePort:		http	31909/TCP
Endpoints:		10.212.1.4:80,10.212.2.5:80
Session Affinity:	None
Events:
  FirstSeen	LastSeen	Count	From			SubObjectPath	Type		Reason			Message
  ---------	--------	-----	----			-------------	--------	------			-------
  1m		1m		1	{service-controller }			Normal		CreatingLoadBalancer	Creating load balancer
  39s		39s		1	{service-controller }			Normal		CreatedLoadBalancer	Created load balancer
```

## 外部からアクセス可能か確認
- 10.215.251.53にアクセスすると「Hello World by Go!」が表示される

## 参考サイト
http://qiita.com/techeten/items/ebb0833d50c882398b0f

