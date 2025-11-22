# Linux system process amount used check

* ## free memory command
  ```md
  # CMD
  free [-Option]

  # Result
                  total       used        free      shared  buff/cache   available
  Mem:         7829468      623876     7181988        9304      274988     7205592
  Swap:        6291452           0     6291452
  ```

* ## meminfo
  ```bash
  # CMD
  cat /proc/meninfo

  # Resutl
  MemTotal:        7829468 kB
  MemFree:         7181296 kB
  MemAvailable:    7204924 kB
  Buffers:            4224 kB
  Cached:           228240 kB
  SwapCached:            0 kB
  Active:           233652 kB
  Inactive:         117452 kB
  Active(anon):     127856 kB
  Inactive(anon):        0 kB
  Active(file):     105796 kB
  Inactive(file):   117452 kB
  Unevictable:           0 kB
  Mlocked:               0 kB
  SwapTotal:       6291452 kB
  SwapFree:        6291452 kB
  Dirty:                 0 kB
  {...}
  Slab:             124704 kB
  SReclaimable:      42552 kB
  SUnreclaim:        82152 kB
  {...}
  
  > Slab : 메모리 영역 중 커널이 직접 사용하는 영역을 지칭하며, Dentry cache, inode cache 등 커널이 사용하는 메모리가 포함된다
  > SReclaimable : Slab 영역 중 재사용될 수 있는 영역이며 캐시 용도로 사용하는 메모리들이 여기에 포함된다
  > SUnreclaim : Slab 영역 중 재사용될 수 없는 영역이며 커널이 현재 사용 중인 영역이기도 한다 [해제해서 다른 용도로는 사용이 불가]
  ```

* ## Slab
  ```bash
  # CMD
  slabtop -o

  # Result
  Active / Total Objects (% used)    : 619978 / 622531 (99.6%)
  Active / Total Slabs (% used)      : 15344 / 15344 (100.0%)
  Active / Total Caches (% used)     : 179 / 298 (60.1%)
  Active / Total Size (% used)       : 109786.78K / 110907.80K (99.0%)
  Minimum / Average / Maximum Object : 0.01K / 0.18K / 8.00K

  OBJS  ACTIVE  USE OBJ  SIZE  SLABS OBJ/SLAB CACHE SIZE NAME
  94690  94690 100%     0.02K    557      170      2228K avtab_node
  52980  52980 100%     0.13K   1766       30      7064K kernfs_node_cache
  48320  48320 100%     0.06K    755       64      3020K ebitmap_node
  47616  47571  99%     0.02K    186      256       744K kmalloc-16
  41727  41727 100%     0.19K   1987       21      7948K dentry
  ```