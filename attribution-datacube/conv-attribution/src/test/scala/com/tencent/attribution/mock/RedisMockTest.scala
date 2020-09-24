/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/8/20, 8:15 PM
 */

package com.tencent.attribution.mock

import org.scalatest.funsuite.AnyFunSuite

class RedisMockTest extends AnyFunSuite {
  test("embedded redis test") {
    RedisMock.buildTestData()
  }
}
