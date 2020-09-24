/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/8/20, 2:19 PM
 */

package com.tencent.attribution.index

import java.nio.charset.StandardCharsets

import com.google.protobuf.util.JsonFormat
import com.tencent.attribution.proto.click.Click.ClickLog

/**
 * 点击日志编码和解码
 *
 * @tparam T 编码后的类型
 */
trait ClickLogCodec[T] extends Serializable {
  /**
   * 编码
   *
   * @param clickLog 点击日志
   * @return
   */
  def encode(clickLog: ClickLog): T

  /**
   * 解码
   *
   * @param data 解码需要的数据
   * @return
   */
  def decode(data: T): Option[ClickLog]
}

/**
 * 按照protobuf进行二级制序列化和反序列化
 */
class ClickLogProtobufCodec extends ClickLogCodec[Array[Byte]] {
  /**
   * protobuf序列化成二进制数据
   *
   * @param clickLog 点击日志
   * @return 序列化后的二进制数据
   */
  override def encode(clickLog: ClickLog): Array[Byte] = {
    clickLog.toByteArray
  }

  override def decode(bytes: Array[Byte]): Option[ClickLog] = {
    try {
      val clickLog = ClickLog.parseFrom(bytes)
      Some(clickLog)
    } catch {
      case _: Throwable => {}
        None
    }
  }
}

/**
 * 按照json格式序列化和反序列化
 */
class ClickLogJsonCodec extends ClickLogCodec[Array[Byte]] {
  /**
   * 序列化成json格式的字节数组
   *
   * @param clickLog 点击日志
   * @return
   */
  override def encode(clickLog: ClickLog): Array[Byte] = {
    JsonFormat.printer()
      .omittingInsignificantWhitespace()
      .print(clickLog.toBuilder)
      .getBytes()
  }

  /**
   * 从json格式的字节数组解析点击日志
   *
   * @param bytes json格式的字节数组
   * @return
   */
  override def decode(bytes: Array[Byte]): Option[ClickLog] = {
    val builder = ClickLog.newBuilder()
    JsonFormat.parser()
      .ignoringUnknownFields()
      .merge(new String(bytes, StandardCharsets.UTF_8), builder)
    Some(builder.build())
  }
}
