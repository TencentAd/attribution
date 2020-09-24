package com.tencent.attribution.index

import com.tencent.attribution.proto.click.Click.ClickLog
import com.tencent.attribution.proto.user.User.UserData
import org.scalatest.funsuite.AnyFunSuite

class ClickLogJsonCodecTest extends AnyFunSuite {

  test("ClickLog json codec") {
    val userData = UserData.newBuilder().setImei("abcdefg").build()
    val click = ClickLog.newBuilder()
      .setAppId("com.tencent.qq")
      .setClickTime(150123415)
      .setUserData(userData)
      .build()

    val jsonCodec = new ClickLogJsonCodec()

    val encodeBytes = jsonCodec.encode(click)
    val decodeClick = jsonCodec.decode(encodeBytes).get

    assertResult(click.getAppId, "") {
      decodeClick.getAppId
    }
    assertResult(click.getUserData.getImei) {
      decodeClick.getUserData.getImei
    }
  }
}
