/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 3:49 PM
 *
 */

package com.tencent.attribution.validation

import com.tencent.attribution.proto.click.Click.ClickLog
import com.tencent.attribution.proto.conv.Conv.ConversionLog

/**
 * 判断点击对于转化是否有效
 */
class ClickLogValidation extends Serializable {
  final val ConvEventDelayInterval: Long = 86400 * 7

  def check(conv: ConversionLog, click: ClickLog): Boolean = {
    if (click.getClickTime >= conv.getEventTime) {
      return false
    }

    if (conv.getEventTime - click.getClickTime > ConvEventDelayInterval) {
      return false
    }

    true
  }
}
