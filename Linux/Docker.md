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