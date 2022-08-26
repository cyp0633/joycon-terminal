# `joycon-terminal` 协议文档 / Protocol Docs

波特率 4800。

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

device_num: 设备编号，连接 PC 的机器固定为 0x01，其余机器<del>未来</del>可以自选，从 0x02 到 0x08。

device_num is the device number. For the board connected to the PC, it's fixed to 0x01; otherwise it's customizable <del>in the future</del>.

key_num 是按键编号，Key1-3、上、下、左、右分别为 0x01-0x07。

key_num is the key number. Key1-3, up, down, left, right are 0x01-0x07.

key_action 是按键动作，0x01 按下，0x02 抬起。

key_action is the key action. 0x01 is pressed, 0x02 is released.

## 特殊协议 / Special protocol

设备号为 0x00 的均为特殊协议。

device_num == 0x00 is the special protocol.

### PC 握手识别 / PC Handshake Recognition

使用按键号 0x00 代表。

Represented by device number #0x00.

```
01 00 00 00 00 00 00 00
```

由 PC 主动发送，进行握手。

Sent by the PC to handshake with the board.

```
01 00 00 01 xx xx xx xx
```

回复代表成功握手，后 8 位从前到后代表从 0x01 到 0x08 的设备在线状况。

### 板间握手识别 / Board Handshake Recognition

使用按键号 0x01 代表。

Represented by device number #0x01.

```
01 00 01 00 xx 00 00 00
```

由辅助板主动发送，请求握手。xx 为请求者序号。

Sent by slave, to request handshake.

```
01 00 01 01 xx 00 00 00
```

由主板回复，代表握手成功。

Reply by the master, to indicate handshake success.
