#### web_demo.event_log 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 自增id | bigint | PRI | NO |  |  |
| 2 | type | 操作类型 | varchar(255) |  | NO |  |  |
| 3 | device | 操作设备 | varchar(255) |  | YES |  |  |
| 4 | username | 用户名 | varchar(255) | MUL | NO |  |  |
| 5 | user_id | 用户id | bigint | MUL | NO |  |  |
| 6 | ip | ip地址 | varchar(255) |  | YES |  |  |
| 7 | location | 位置 | text |  | YES |  |  |
| 8 | content | 具体内容 | text |  | NO |  |  |
| 9 | created_at | 创建时间 | datetime | MUL | YES |  |  |
