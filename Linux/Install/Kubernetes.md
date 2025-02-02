# 1. Install kubectl on Linux
```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```
> **Custom Install version :**
> ```bash
> $(curl -L -s https://dl.k8s.io/release/stable.txt)
> curl -LO https://dl.k8s.io/release/v1.32.0/bin/linux/amd64/kubectl
> ```

* ## Validate the binary [x86-64]
  ```bash
  # 1. Download the kubectl checksum file:
  curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
  ```

  ```bash
  # 2. Validate the kubectl binary against the checksum file:
  echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
  ```

  ```bash
  # 2.2 If valid, the output is:
  kubectl: OK

  # 2.3 If the check fails, sha256 exits with nonzero status and prints output similar to:
  kubectl: FAILED
  sha256sum: WARNING: 1 computed checksum did NOT match
  ```

# 2. Install kubectl
  ```bash
  sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
  ```
  > **User for root access on the target system**
  > ```bash
  > chmod +x kubectl
  > mkdir -p ~/.local/bin
  > mv ./kubectl ~/.local/bin/kubectl
  > ```

# 3. Install kubeadm
```bash
sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
```

```bash
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v1.32/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v1.32/rpm/repodata/repomd.xml.key
exclude=kubelet kubeadm kubectl cri-tools kubernetes-cni
EOF
```

```bash
sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
```


