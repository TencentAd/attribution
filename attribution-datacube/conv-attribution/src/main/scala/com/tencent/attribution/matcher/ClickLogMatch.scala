/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 3:49 PM
 *
 */

package com.tencent.attribution.matcher

import com.tencent.attribution.index.{ClickIndex, UserIndexKeyBuilder}
import com.tencent.attribution.proto.click.Click.ClickLog
import com.tencent.attribution.proto.conv.Conv.ConversionLog
import com.tencent.attribution.validation.ClickLogValidation

import scala.collection.mutable.ListBuffer

/**
 * 匹配点击日志
 *
 * @param index      点击日志使用的索引
 * @param validation 点击日志校验
 */
class ClickLogMatch(index: ClickIndex,
                    validation: ClickLogValidation) extends Serializable {

  /**
   * 匹配转化日志
   *
   * @param conv 转化数据
   * @return 假如匹配上，输出转化后的日志，假如没有匹配上，
   */
  def matchClickLog(conv: ConversionLog): Option[ConversionLog] = {
    if (conv.getAppId.isEmpty) {
      return None
    }
    if (!conv.hasUserData) {
      return None
    }

    val keyBuilder = new UserIndexKeyBuilder(conv.getAppId)
    val indexKeys = keyBuilder.build(conv.getUserData)

    var candidates = new ListBuffer[ClickLog]()
    indexKeys.foreach(key => {
      index.Get(key.idType, key.id) // 获取相关点击
        .filter(click => validation.check(conv, click)) // 点击校验
        .foreach(click => {
          candidates += click
        })
    })

    candidates.toList.sortWith((c1: ClickLog, c2: ClickLog) => {
      c1.getClickTime > c2.getClickTime
    })

    if (candidates.nonEmpty) {
      return Some(conv.toBuilder.setMatchClick(candidates.head).build())
    }
    None
  }
}
