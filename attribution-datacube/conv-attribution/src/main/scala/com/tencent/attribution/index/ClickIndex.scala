/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 3:49 PM
 *
 */

package com.tencent.attribution.index

import com.tencent.attribution.proto.click.Click.ClickLog
import com.tencent.attribution.proto.user.User.IdType

/**
 * 点击日志的索引
 */
abstract class ClickIndex(val codecClass: String =
                          "com.tencent.attribution.index.ClickLogJsonCodec") extends Serializable {
  val codec: ClickLogCodec[Array[Byte]] =
    Class.forName(codecClass).newInstance().asInstanceOf[ClickLogCodec[Array[Byte]]]

  /**
   * 接收到点击日志时，将点击日志存入索引
   *
   * @param idType 用户id类型
   * @param id     id值
   * @param click  点击日志
   */
  def Set(idType: IdType, id: String, click: ClickLog)

  /**
   * 根据用户id获取历史最近的点击信息
   *
   * @param idType 用户id类型
   * @param id     id值
   * @return 最近的点击信息
   */
  def Get(idType: IdType, id: String): Option[ClickLog]

  /**
   * 生成请求的key
   *
   * @param idType 用户id类型
   * @param id     id值
   * @return 请求索引的key
   */
  def formatKey(idType: IdType, id: String): String = {
    s"$id"
  }
}

