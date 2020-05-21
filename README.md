# 免费签发证书，到期自动续订

## 1. cert-manager
### 1.1 Installing cert-manager with Helm

```bash
$ kubectl create namespace cert-manager
$ helm repo add jetstack https://charts.jetstack.io
$ helm repo update
$ helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v0.15.0 \
  --set installCRDs=true
```
[参考](https://cert-manager.io/docs/installation/kubernetes/)

### 1.2 创建 letsencrypt-prod
```bash
$ kubectl create -f letsencrypt-prod.yaml
```
- letsencrypt-prod.yaml
```yaml
apiVersion: cert-manager.io/v1alpha2
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    # You must replace this email address with your own.
    # Let's Encrypt will use this to contact you about expiring
    # certificates, and issues related to your account.
    email: lu.jin@advantech.com.cn
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      # Secret resource that will be used to store the account's private key.
      name: letsencrypt-prod
    # Add a single challenge solver, HTTP01 using nginx
    solvers:
    - http01:
        ingress:
          class: nginx
```
[参考](https://cert-manager.io/docs/configuration/acme/)

### 1.3 配置 ingress
1. 配置 `cert-manager.io/cluster-issuer`
```yaml
metadata:
  annotations:
    cert-manager.io/cluster-issuer: letsencrypt-prod
    ingress.kubernetes.io/proxy-body-size: "0"
    ingress.kubernetes.io/ssl-redirect: "true"
    nginx.ingress.kubernetes.io/proxy-body-size: "0"
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
```
2.  配置 tls
```bash
  tls:
  - hosts:
    - <your domain>
    secretName: <your secret>
```

[参考](https://cert-manager.io/docs/usage/ingress/)

## 2. 安装 renewCertificate
### 2.1 创建 service Account
```bash
$ kubectl create sa renew-cert
```
### 2.2 创建 clusterrolebingding
```bash
$ kubectl create clusterrolebinding renew-cert \
  --clusterrole=cluster-admin \
  --serviceaccount=cert-manager:renew-cert
```
### 2.3 配置 cronjob 中的 `namespace` 和 `certificate`
> 多个 certificate 用逗号隔开  

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: renew-cert
spec:
  schedule: "0 0 * * *"
  failedJobsHistoryLimit: 1
  successfulJobsHistoryLimit: 1
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccount: renew-cert
          containers:
          - name: renew-cert
            image: harbor.wise-paas.com/library/renewcertificate:v0.0.6
            imagePullPolicy: IfNotPresent #Always
            command: ["/renewCertificate"]
            args: ["--namespace=harbor", "--certificate=harbor-cert,notary-cert"]
          restartPolicy: OnFailure
```
### 2.4 创建 cronjob
```bash
$ kubectl create -f cronjob-renew-cert.yaml
```

## 注意事项
- letsencrypt 限制:  
https://letsencrypt.org/docs/rate-limits/
