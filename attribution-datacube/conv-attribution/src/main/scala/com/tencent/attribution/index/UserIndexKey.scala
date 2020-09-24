package com.tencent.attribution.index

import com.tencent.attribution.proto.user.User.{IdType, UserData}

import scala.collection.mutable.ListBuffer

/**
 * 基于用户id构建的key
 *
 * @param id id字符串
 * @param idType id类型
 */
class UserIndexKey(private[attribution] val id: String, private[attribution] val idType: IdType) {}

class UserIndexKeyBuilder(val appId :String) {
  val buf = new ListBuffer[UserIndexKey]()

  def build(user: UserData): ListBuffer[UserIndexKey] ={
    addUserIndexKey(getDeviceId(user), IdType.Device)

    buf
  }

  private def addUserIndexKey(userId :Option[String], idType: IdType): Unit ={
    userId.foreach(s => addUserIndexKey(s, idType))
  }

  private def addUserIndexKey(userId :String, idType: IdType): Unit ={
    if (!checkUserIdValid(userId)) {
      return
    }
    buf += new UserIndexKey(formatKey(userId), idType)
  }

  private def formatKey(id :String) :String = appId + "_" +  id

  private def getDeviceId(user :UserData): Option[String] = {
    if (!user.getImei.isEmpty) {
      Some(user.getImei)
    } else if (!user.getIdfa.isEmpty) {
      Some(user.getIdfa)
    } else {
      None
    }
  }

  private def checkUserIdValid(userId :String):Boolean = {
    if (userId.isEmpty) {
      return false
    }

    true
  }

}
