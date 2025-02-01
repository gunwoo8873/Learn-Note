# Swap
```md
물리 메모리가 부족할 경우 대비해서 생성해 놓은 영역이다 부족하면 프로세스가 더 이상 연산을 위한 공간을
확보할 수 없기 때문에 시스템이 응답 불가 상태에 빠질 수 있다 물리 메모리가 아니라 디스크의 일부분을
메모리처럼 사용하기 위해 생성한 공간이기 때문에 메모리가 부족할 때 사용하기 보다는 메모리에 비해
접근과 처리 속도가 현저하게 떨어진다
```

* ## Process ID memory data
  ```bash
  # CMD
  cd /proc/<PID>/smaps
  ```

* ## File system satus
  ```bash
  # CMD
  cat /proc/<PID>/satus

  Name:   sshd-session
  Umask:  0022
  State:  S (sleeping)
  Tgid:   2491
  Ngid:   0
  Pid:    2491
  PPid:   2487
  TracerPid:      0
  Uid:    1000    1000    1000    1000
  Gid:    1000    1000    1000    1000
  {...}
  ```

