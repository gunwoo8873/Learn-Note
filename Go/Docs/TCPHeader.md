# TCP Header byte
Network Protocol는 캡슐화된 페이로드[Encapsulated Playload]의 전달 정보[Forwarding information]와 사실을 포함하는 헤더 집합으로 표현되는 결정론적 구조[Deterministic sturcture]이다.

| Name                       |                 Option                 | Byte  |
| -------------------------- | :------------------------------------: | :---: |
| Source port address        |                                        | 2byte |
| Destination port address   |                                        | 2byte |
| Sequence number            |                                        | 4byte |
| Acknowledment number       |                                        | 4byte |
| Header length and reserved |                                        | 1byte |
| Countrol flag              |                                        | 1byte |
|                            |    CWR [Congestion window reduced]     |       |
|                            | ECE [Explicit congestion notification] |       |
|                            |          URG [Urgent pointer]          |       |
|                            |  ACK [Ackonwledgment number is valid]  |       |
|                            |         PSH [Request for push]         |       |
|                            |       RST [Rest the connection]        |       |
|                            |   SYN [Synchronize sequence numbers]   |       |
|                            |     FIN [Terminate the connection]     |       |
| Window size                |                                        | 2byte |
| Checksum                   |                                        | 2byte |
| Urgent pointer             |                                        | 2byte |
