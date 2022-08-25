# `joycon-terminal` 协议文档 / Protocol Docs

## 按键协议 / Key Protocol

数据包共 8 字节，结构为：

The data packet is 8 bytes long, and the structure is as follows:

```c
struct {
    uint8_t head;
    uint8_t device_num;
    uint8_t key_num;
    uint8_t key_action;
    uint8_t unused1;
    uint8_t unused2;
    uint8_t unused3;
    uint8_t unused4;
}
```

head：数据包头，固定为 0x01。

head is the header of the data packet, and it's fixed to 0x01.

device_num: 设备编号，连接 PC 的机器固定为 0x01，其余机器<del>未来</del>可以自选。

device_num is the device number. For the board connected to the PC, it's fixed to 0x01; otherwise it's customizable <del>in the future</del>.

key_num 是按键编号，Key1-3、上、下、左、右分别为 0x01-0x07。

key_num is the key number. Key1-3, up, down, left, right are 0x01-0x07.

key_action 是按键动作，0x01 按下，0x02 抬起。

key_action is the key action. 0x01 is pressed, 0x02 is released.

