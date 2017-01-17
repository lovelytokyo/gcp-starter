# 前提
- Google Cloud SDKインストール済み
- kubectl Componentインストール済み
- gcloud configセットアップ済み

# 概要
- kubernetesを使ってjobを実行し、job実行結果の標準出力を受け取れるか検証

## クラスタを作る
```
gcloud container clusters create lovelytokyo-server
```

## gcloudコマンドで接続
```
gcloud container clusters get-credentials lovelytokyo-server --zone asia-northeast1-a --project lovelytokyo-018
```

## イメージを作成する
```
cd docker/go-batch
docker build -t gcr.io/lovelytokyo-018/go-batch:v1 .
gcloud docker push gcr.io/lovelytokyo-018/go-batch:v1
```

## jobを作成する
```
kubectl create -f k8s/job.yml
```

## jobが実行されたpodのログを確認する
```
kubectl get pods --show-all
--
NAME                 READY     STATUS      RESTARTS   AGE
batch-sample-ecyx9   0/1       Completed   0          1m
```

```
kubectl logs batch-sample-ecyx9                                                                                                                                                                                (feature_batch✱)
+ exec app
各処理時間合計 18.000365 sec
実時間 10.000184 sec
```

## kubectlで直接にjobを実行する
```
kubectl run -it  batch-sample --image=gcr.io/lovelytokyo-018/go-batch:v1 --rm=true go run main.go
--
Waiting for pod default/batch-sample-3619662959-yj4qh to be running, status is Pending, pod ready: false
If you don't see a command prompt, try pressing enter.
各処理時間合計 18.000395 sec
実時間 10.000203 sec
Session ended, resume using 'kubectl attach batch-sample-3619662959-yj4qh -c batch-sample -i -t' command when the pod is running
deployment "batch-sample" deleted
```


