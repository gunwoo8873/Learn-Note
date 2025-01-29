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
