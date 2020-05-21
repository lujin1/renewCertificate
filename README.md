# renewCertificate

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
### 1.3 ingress 中配置
1. 配置 `cert-manager.io/cluster-issuer`
```bash
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
## 2. 安装 renewCertificate
### 2.1 配置 cronjob 中 `namespace` 和 `certificate`
> 多个 certificate 用逗号隔开  

```bash
          containers:
          - name: renew-cert
            image: harbor.wise-paas.com/library/renewcertificate:v0.0.6
            imagePullPolicy: IfNotPresent #Always
            command: ["/renewCertificate"]
            args: ["--namespace=harbor", "--certificate=harbor-cert,notary-cert"]
```
### 2.2 创建 cronjob
```bash
$ kubectl create -f cronjob-renew-cert.yaml
```
