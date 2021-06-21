# XChain Blockchain


## 1. 编译
代码库从vnode修改，但改为使用go module作为依赖管理, 直接执行make命令可以编译，go会自动获取依赖包

bls_lib包如果go无法自动获得，可以在go.mod文件中用replace语句切换到本地路径上的包
```bash
make xchain
```
```bash
replace github.com/MOACChain/MoacLib => /media/yifan/ssd/go/src/github.com/MOACChain/MoacLib                                                                     replace github.com/innowells/bls_lib/v2 => /media/yifan/ssd/go/src/github.com/innowells/bls_lib
```
---

## 2. 运行概述
```bash
本次demo一共涉及4条链，4个合约，资产从X链跨至Y链。包括：

1. X 链，部署vault X 合约
2. Y 链，部署vault Y 合约
3. vnode链，部署vssbase合约
4. xchain链，系统genesis自带xevents系统合约，无需部署
```
其中xchain链，采用3节点模式，xchain节点暂时为pow出块，后期会切换到bls出块。
其中xchain需要与3条链都发生交互：

1. vnode链：xchain需要与vnode链交互来完成bls签名需要的分布式密钥vss过程。
2. X链：xchain会定时监听X链上的vault合约的deposit事件，采用bls合并签名后，记录在本身链上的系统合约xevents中
3. Y链：xchain会定时读取xevents合约中记录的deposit事件，并调用Y链上的相应vault合约的mint函数

---

## 3. 配置
```bash
以上运行过程，需要事先做多步配置，主要包括5个：
1. 初始化xchain节点的本地密钥
2. 撰写配置文件
3. 调用vssbase合约
4. 在vnode, X, Y等三条链上给xchain链的节点添加资金。
5. 在X,Y中授予xchain节点minter权限
```

---

### 3.1 初始化xchain节点的本地密钥
调用以下命令生成本地密钥
```bash
./build/bin/xchain --datadir ~/.xchain1 account newx
```
调用以下命令查看本地密钥，用于后续步骤
```bash
./build/bin/xchain --datadir ~/.xchain1 account listx
```
结果为
```bash
Account #0: {0xc996264b44d35f9ae8291101760fb1ecffb445f5}, pubkey: 0x9045469e9cf0d49c4629df0221939cfd07e1719d969c26672262e5e596139ff0
```
其中0xc996264b44d35f9ae8291101760fb1ecffb445f5为节点地址，0x9045469e9cf0d49c4629df0221939cfd07e1719d969c26672262e5e596139ff0为节点公钥
记录以上两者后，用于后续 3.3, 3.4 步骤。

---

### 3.2 撰写配置文件
配置文件为两种
1. vnodeconfig.json 用于描述vnode信息与vssbase合约位置

```json
{
    "VnodeIP": "172.21.0.11",
    "VnodePort": "8545",
    "ChainId": 95125,
    "VssBaseAddr": "0x2E32C6F7630ca3f06EfAbEaDa1da0Bd28aA18FEA"
}
```
2. vaults.json 用于描述X链，Y链的RPC位置，Vault合约地址，需要监听的token的mapping信息等
```json
{
  "vaults": [
    {
      "vaultx": {
        "id": 95125,
        "rpc": "http://172.21.0.11:8545",
        "prefix": "mc",
        "address": "0xABE1A1A941C9666ac221B041aC1cFE6167e1F1D0"
      },
      "vaulty": {
        "id": 95125,
        "rpc": "http://172.21.0.11:8545",
        "prefix": "mc",
        "address": "0xcCa8BAA2d1E83A38bdbcF52a9e5BbB530f50493A"
      },
      "tokenmappings": [
        {
          "sourcechainid": 95125,
          "sourcetoken": "0x350e47237eb2515b3b30c2f232268b998e392409",
          "mappedchainid": 95125,
          "mappedtoken": "0x8553ce822a9072b5ff0992da9a61d5ce54a1f5df"
        }
      ]
    }
  ]
}

```

---

### 3.3 调用vssbase合约
vssbase合约调用与bls链设置类似，需要对于每个xchain节点，调用registerVss函数，
其中的地址与公钥为3.1步骤中获得:
```bash
registerVSS(Address, pubkey);
```
---

### 3.4 添加资金
对于每个xchain节点，需要在以下的链上打入资金，节点地址为3.1获得
1. vnode上的资金，用于xchain节点调用vssbase合约
2. X链上资金，用于xchain节点调用vault x的withdraw函数
3. Y链上资金，用于xchain节点调用vault y的mint函数

---

### 3.5 授予minter权限
调用X,Y合约中的grantMinter函数，地址为3.1中获得：
```bash
grantMinter(addr)
```

---

## 4  xchain节点运行
注意：

1. 节点的mine必须打开

2. 节点的rpc必须打开

3. 挑选其中一个作为bootnode节点，其他节点采用bootnode
```bash
./build/bin/xchain --datadir ~/.xchain1 --mine --minerthreads 1 --rpc --rpcport 18545 --rpcaddr 0.0.0.0 --rpcapi txpool,chain3,mc,net,vnode,personal,admin,miner
```


```bash
./build/bin/xchain --datadir ~/.xchain3 --mine --minerthreads 1 --rpc --rpcport 38545 --rpcaddr 0.0.0.0 --rpcapi txpool,chain3,mc,net,vnode,personal,admin,miner --port 50333 --bootnodesv4 enode://b02fff0c541506fdb9b1bc3296f8132a41b3fc5f6a5ff331f33203826b9f8275d6231ace83311c8ea34b716b9efd09c58bcca8f9a6499a3d79031fbbdb0994b3@192.168.0.156:30333
```