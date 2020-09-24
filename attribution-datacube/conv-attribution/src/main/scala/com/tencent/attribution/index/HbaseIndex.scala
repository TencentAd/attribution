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
import java.util.concurrent.Executors

import com.tencent.attribution.index.HbaseClickIndex.createHbaseTable
import com.tencent.attribution.proto.click.Click.ClickLog
import com.tencent.attribution.proto.user.User.IdType
import org.apache.hadoop.conf.Configuration
import org.apache.hadoop.hbase.client.{ConnectionFactory, Get, Put, Table}
import org.apache.hadoop.hbase.util.Bytes
import org.apache.hadoop.hbase.{HBaseConfiguration, TableName}
import org.yaml.snakeyaml.Yaml

import scala.beans.BeanProperty

/**
 * 基于hbase实现点击索引
 */
case class HbaseClickIndex(var configInputStream: InputStream) extends ClickIndex {
  if (configInputStream == null) {
    configInputStream = getClass.getResourceAsStream("/hbase.yaml")
  }

  val table: Table = createHbaseTable(configInputStream)

  final val cfName  = "cf".getBytes()
  final val qualifierName = "c".getBytes()

  /**
   * 将点击数据存入hbase
   *
   * @param idType 用户id类型
   * @param id     id值
   * @param click  点击日志
   */
  override def Set(idType: IdType, id: String, click: ClickLog): Unit = {
    val put = new Put(Bytes.toBytes(formatKey(idType, id)))
    put.addColumn(cfName, qualifierName, codec.encode(click))
    table.put(put)
  }

  /**
   * 从hbase取出点击数据
   *
   * @param idType 用户id类型
   * @param id     id值
   * @return 最近的点击信息
   */
  override def Get(idType: IdType, id: String): Option[ClickLog] = {
    val get = new Get(Bytes.toBytes(formatKey(idType, id)))
    val result = table.get(get)
    if (result.isEmpty) {
      return None
    }

    val data = result.getValueAsByteBuffer("cf".getBytes(), "c".getBytes()).array()
    codec.decode(data)
  }
}

object HbaseClickIndex {
  /**
   * 从配置文件中加载到 [[com.tencent.attribution.index.HBaseConfig]] 类
   *
   * @param configInputStream 配置文件路径，yaml格式
   * @return [[com.tencent.attribution.index.HBaseConfig]]
   */
  private def loadConfig(configInputStream: InputStream): HBaseConfig = {
    new Yaml()
      .loadAs(configInputStream, classOf[HBaseConfig])
  }

  /**
   * 创建访问ClickIndex的表
   *
   * @param configInputStream 配置文件路径，yaml格式
   * @return hbase table
   */
  private def createHbaseTable(configInputStream: InputStream): Table = {
    val config = loadConfig(configInputStream)
    val hbaseConf: Configuration = HBaseConfiguration.create()
    HBaseConfiguration.setWithPrefix(hbaseConf, "", config.properties.entrySet())
    val conn = ConnectionFactory.createConnection(hbaseConf)
    conn.getTable(
      TableName.valueOf(config.tableName),
      Executors.newFixedThreadPool(config.poolSize))
  }
}

class HBaseConfig {
  @BeanProperty var tableName: String = _
  @BeanProperty var poolSize: Int = _
  @BeanProperty var properties: util.Map[String,String] = _
}
