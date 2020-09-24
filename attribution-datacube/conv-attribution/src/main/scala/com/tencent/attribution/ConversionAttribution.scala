/*
 * Copyright (c) 2020.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/7/20, 6:59 PM
 */

package com.tencent.attribution

import java.io.InputStream

import com.google.protobuf.util.JsonFormat
import com.tencent.attribution.index.RedisClickIndex
import com.tencent.attribution.matcher.ClickLogMatch
import com.tencent.attribution.proto.conv.Conv.ConversionLog
import com.tencent.attribution.util.GeneralArgParser
import com.tencent.attribution.validation.ClickLogValidation
import org.apache.hadoop.fs.Path
import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

object ConversionAttribution {
  private var clickIndexType: String = _
  private var redisConfigPath: Option[String] = None
  private var hbaseConfigPath: Option[String] = None

  private val jsonParser = JsonFormat.parser().ignoringUnknownFields()

  def convMatch(line: String, matcher: ClickLogMatch): ConversionLog = {
    val convBuilder = ConversionLog.newBuilder()
    jsonParser.merge(line, convBuilder)
    val conv = convBuilder.build()
    println("matcher: ", matcher)
    matcher
      .matchClickLog(conv)
      .getOrElse(conv)
  }

//  def initClickIndex(sc: SparkContext): ClickIndex = {
//    if (clickIndexType.equalsIgnoreCase("redis")) {
//      new RedisClickIndex(redisConfigPath.map(path => openPath(path, sc)).orNull)
//    } else if (clickIndexType.equalsIgnoreCase("hbase")) {
//      new HbaseClickIndex(hbaseConfigPath.map(path => openPath(path, sc)).orNull)
//    } else {
//      throw new Exception(s"$clickIndexType not support")
//    }
//  }

  def openPath(path: String, sc: SparkContext): InputStream = {
    val p = new Path(path)
    p.getFileSystem(sc.hadoopConfiguration).open(p)
  }

  def main(args: Array[String]) {
    val argParser = new GeneralArgParser(args)
    val inputPath = argParser.getStringValueOrThrow("input_path")
    val outputPath = argParser.getStringValueOrThrow("output_path")
    clickIndexType = argParser.argsMap.getOrElse("click_index_type", "redis")
    redisConfigPath = argParser.getStringValueOption("redis_config")
    hbaseConfigPath = argParser.getStringValueOption("hbase_config")

    val conf = new SparkConf()
      .setAppName("conversion-attribution")

    val sc = new SparkContext(conf)

    lazy val clickIndex = {
      new RedisClickIndex()
    }
    lazy val matcher = new ClickLogMatch(
      index = clickIndex,
      validation = new ClickLogValidation)

    util.Util.deletePath(sc, outputPath)

    val logData: RDD[String] = sc.textFile(inputPath).cache()
    logData
      .filter(line => !line.isEmpty)
      .map(
        line =>convMatch(line, matcher)
      )
      .map(
        joinedConv =>
          JsonFormat.printer().omittingInsignificantWhitespace().print(joinedConv)
      )
      .saveAsTextFile(outputPath)
  }
}
