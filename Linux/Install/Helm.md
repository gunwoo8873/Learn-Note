# Helm

* ## 1. install helm
  ```bash
  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
  chmod 700 get_helm.sh
  ./get_helm.sh
  ```

* ## 2. Helm package
  ```bash
  curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null
  sudo apt-get install apt-transport-https --yes
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
  sudo apt-get update
  sudo apt-get install helm
  ```
* ## 3. snap version of the helm pkg
  ```bash
  sudo snap install helm --classic
  ```

* ## 4. Getting source
  ```bash
  git clone https://github.com/helm/helm.git
  cd helm
  make
  ```
  > Note : Need to make installed [`sudo apt install make`]