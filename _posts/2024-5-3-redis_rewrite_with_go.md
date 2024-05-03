---
layout: post
title: "用Go手写Redis"
date:   2024-5-3
tags: [Redis]
comments: true
author: Shayne
---

# 源码结构

|          代码文件          | 备注                    |
| :------------------------: | ----------------------- |
|     **redis.c**     | redis服务端核心流程代码 |
|        redis-cli.c        | 客户端代码              |
|            ae.c            | 事件库                  |
|           anet.c           | 网络库                  |
|         zmalloc.c         | 内存库                  |
|   sds.c adlist.c dict.c   | 数据结构                |
| pqsort.c benchmark.c lzf.c | 辅助工具代码            |

学习redis应该遵循从整体到局部的思路，所以应该重点关注服务端的流程代码以及事件库网络库的实现方式

# Redis核心概念
