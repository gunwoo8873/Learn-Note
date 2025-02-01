# Time
```bash
# Step 1
date

> Tue Jan 28 05:36:37 PM KST 2025

# Step 1.2
timedatectl

> Local time: Tue 2025-01-28 17:37:18 KST
    Universal time: Tue 2025-01-28 08:37:18 UTC
          RTC time: Tue 2025-01-28 08:37:19
         Time zone: Asia/Seoul (KST, +0900)
  System clock synchronized: yes
       NTP service: active
   RTC in local TZ: no

# Step 2
timedatectl list-timezones | grep Seoul

# Step 3
sudo timedatectl set-timezone Asia/Seoul
```

# Smem
* ## Ubuntu
* ## CentOS 
  * EPEL repository create and access
    > Is not centos 10 version for have`t package
  ```bash
  # Step 1
  sudo dnf install epel-release -y
  sudo dnf makecache
  ```
  ```bash
  # Step 2
  yum install <pkg name>
  ```
  * smem