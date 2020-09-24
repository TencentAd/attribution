/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/8/20, 2:19 PM
 */

package com.tencent.attribution.index

import java.io.InputStream
import java.util

import com.tencent.attribution.index.RedisClickIndex.createRedis
import com.tencent.attribution.proto.click.Click
import com.tencent.attribution.proto.user.User
import redis.clients.jedis.{HostAndPort, Jedis, JedisCluster, JedisCommands}

import scala.beans.BeanProperty
import scala.collection.JavaConverters.asScalaBufferConverter


class RedisClickIndex(var config: InputStream=null) extends ClickIndex {
  if (config == null) {
    config = getClass.getResourceAsStream("/redis.yaml")
  }

  val jedisCommands: JedisCommands = createRedis(config)

  /**
   * 接收到点击日志时，将点击日志存入索引
   *
   * @param idType 用户id类型
   * @param id     id值
   * @param click  点击日志
   */
  override def Set(idType: User.IdType, id: String, click: Click.ClickLog): Unit = {
    if (click == null) {
      return
    }
    jedisCommands.set(
      formatKey(idType, id),
      new String(codec.encode(click)))
  }

  /**
   * 根据用户id获取历史最近的点击信息
   *
   * @param idType 用户id类型
   * @param id     id值
   * @return 最近的点击信息
   */
  override def Get(idType: User.IdType, id: String): Option[Click.ClickLog] = {
    val key = formatKey(idType, id)
    val value = jedisCommands.get(key)
    println(s"get result, key: $key, value: $value")

    if (value == null || value == "") {
      return None
    }
    codec.decode(value.getBytes())
  }
}

class RedisConfig {
  @BeanProperty var isCluster: Boolean = _
  @BeanProperty var node: String = _
  @BeanProperty var nodes: util.ArrayList[String] = _
  @BeanProperty var timeout: Int = _
}

object RedisClickIndex {
  /**
   * 根据配置文件创建redis实例，根据配置决定是否创建集群
   *
   * @param config redis配置文件
   * @return redis命令接口
   */
  private def createRedis(config: InputStream): JedisCommands = {
    createRedis(loadConfig(config))
  }

  /**
   * 从yaml配置文件中加载到[[RedisConfig]] Object
   *
   * @param config 配置文件
   * @return
   */
  private def loadConfig(config: InputStream): RedisConfig = {
    import org.yaml.snakeyaml.Yaml

    new Yaml()
      .loadAs(config, classOf[RedisConfig])
  }

  /**
   * 创建redis实例，根据配置决定是否创建集群
   *
   * @param redisConfig redis配置
   * @return redis命令接口
   */
  private def createRedis(redisConfig: RedisConfig): JedisCommands = {
    if (redisConfig.isCluster) {
      createRedisCluster(redisConfig: RedisConfig)
    } else {
      createSingleNodeRedis(redisConfig: RedisConfig)
    }
  }

  /**
   * 创建redis集群
   *
   * @param redisConfig redis配置
   * @return redis命令接口
   */
  private def createRedisCluster(redisConfig: RedisConfig): JedisCommands = {
    val jedisClusterNodes = new util.HashSet[HostAndPort]()
    for (node <- redisConfig.nodes.asScala) {
      jedisClusterNodes.add(HostAndPort.parseString(node))
    }
    new JedisCluster(jedisClusterNodes, redisConfig.timeout)
  }

  /**
   * 创建redis
   *
   * @param redisConfig redis配置
   * @return redis命令接口
   */
  private def createSingleNodeRedis(redisConfig: RedisConfig): JedisCommands = {
    val hp = HostAndPort.parseString(redisConfig.node)
    new Jedis(hp.getHost, hp.getPort, redisConfig.timeout)
  }
}
