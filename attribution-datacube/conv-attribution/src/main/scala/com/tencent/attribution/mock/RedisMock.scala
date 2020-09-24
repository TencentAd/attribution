/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/8/20, 8:09 PM
 */

package com.tencent.attribution.mock

import redis.clients.jedis.Jedis

object RedisMock {
  import redis.embedded.RedisServer

  var redisServer: RedisServer = _

  def buildTestData(): Unit ={
    redisServer = RedisServer.builder()
      .port(7379)
      .setting("maxclients 10")
      .build()
    redisServer.start()

    val client = new Jedis("127.0.0.1", 7379)
    client.set("app_id1_imei1", "{\"user_data\": {\"imei\": \"imei1\"}, \"click_time\": 1597318001}")
    client.close()
  }

  def stop(): Unit = {
    redisServer.stop()
  }
}
