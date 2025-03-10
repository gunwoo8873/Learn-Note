# CentOS docker install
* ## 1. Repository set up
  ```bash
  sudo dnf -y install dnf-plugins-core
  sudo dnf config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
  ```

* ## 2. Docker engine install
  ```bash
  sudo dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
  ```
  > **Docker install version search** : `dnf list docker-ce --showduplicates | sort -r`  
  > **Custom version CMD** : `sudo dnf install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io docker-buildx-plugin docker-compose-plugin`

* ## 3. Linux system enable docker service
  ```bash
  sudo systemctl enable --now docker
  ```

* ## 4. Uninstall Docker engine
  ```bash
  sudo dnf remove docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
  ```
  ```bash
  sudo rm -rf /var/lib/docker
  sudo rm -rf /var/lib/containerd
  ```

# Ubuntu Docker install
* ## 1. Using the apt repository
  ```bash
  # Add Docker's official GPG key:
  sudo apt-get update
  sudo apt-get install ca-certificates curl
  sudo install -m 0755 -d /etc/apt/keyrings
  sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
  sudo chmod a+r /etc/apt/keyrings/docker.asc
  ```
  ```bash
  # Add the repository to Apt sources:
  echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
    $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
  sudo apt-get update
  ```

* ## 2. Install the docker package
  ```bash
  sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
  ```
  > **Docker install version search** : `apt-cache madison docker-ce | awk '{ print $3 }'`  
  > **Custom version CMD** : `sudo apt-get install docker-ce=$VERSION_STRING docker-ce-cli=$VERSION_STRING containerd.io docker-buildx-plugin docker-compose-plugin`

* ## 3. User access control
  ```bash
  # Step 1 : Docker access create group check
  cat /etc/group

  # Step 1.2 : Docker not create access group
  sudo group add docker

  # Step 2 : If not user docker access group
  sudo usermod -aG docker $USER

  # Step 3 : System reboot
  sudo reboot
  exit
  ```