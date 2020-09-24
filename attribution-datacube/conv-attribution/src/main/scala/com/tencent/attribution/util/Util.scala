/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/8/20, 2:24 PM
 */

package com.tencent.attribution.util

import org.apache.hadoop.fs.Path
import org.apache.spark.SparkContext

object Util {
  // 清空输出路径
  def deletePath(sc: SparkContext, deletePath: String): Unit = {
    val path = new Path(deletePath)
    val fs = path.getFileSystem(sc.hadoopConfiguration)
    fs.delete(path, true)
  }
}
