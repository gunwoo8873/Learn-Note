# System start and end
* **H/W Boot** : Power on → Firmware → Boot loader → Linux kernel and initrd
* **Firmware** :
  * **Bios** : 
  * **EFI** [`Extensible Firmware interface`]
    * Selecter address : 32 Bit -> 64 Bit
    * Partition table reading CRC used
  * **NVRAM**
  * **POST** [Power On Self Test]
  * **MBR** [Master Boot Record] and *BootLoader loading
    * BootLoader : LILO, GRUB / GRUB2, EFILinux, U-Boot  
      [GRUB2 BootLoader](https://www.notion.so/GRUB2-BootLoader-467480f78eea4fbd9fd2d708accf4b7b?pvs=21)
        
    * MBR Byte code : 0, 440, 446, 510, 511
        | 16   | 10  | Description      | Size      |
        | ---- | --- | ---------------- | --------- |
        | 0000 | 0   | Code             | 440 ~ 446 |
        | 01B8 | 440 | Disk = DDS       | 4         |
        | 01BC | 444 | Null = 0x0000    | 2         |
        | 01BE | 446 | Partition Tables | 64        |
        | 01FE | 510 | 55h =  Masic No  | 2         |
        | 01FF | 511 | AAh =  Masic No  | 2         |
        |      |     |                  | 512       |

        > 446 Bytes : Stage 1  
        > 64 Bytes : Primary Partition Table ⇒ 4ea is 16 Byte entry  
        > 2 Bytes : Masic No  