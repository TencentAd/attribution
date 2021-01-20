package com.attribution.datacube.proto.tool;

import com.google.protobuf.Message;
import org.apache.flink.api.common.functions.RuntimeContext;
import org.apache.flink.api.common.serialization.DeserializationSchema;
import org.apache.flink.api.common.serialization.SerializationSchema;
import org.apache.flink.api.common.typeinfo.TypeInformation;
import org.apache.flink.api.java.typeutils.TypeExtractor;
import org.apache.flink.metrics.Counter;

import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.Arrays;

/**
 * 这个类是对ProtoMessage进行反序列化和序列化的封装类
 *
 * @author v_zefanliu
 */
public class ProtoMessageSchema implements DeserializationSchema<Message>, SerializationSchema<Message> {
    private static final long serialVersionUID = 1L;

    final private Class<?> pbClass;
    private transient Method parseFunction;
    private transient Counter invokeSuccessCount;
    private transient Counter invokeFailCount;

    public ProtoMessageSchema(Class pbClass) {
        this.pbClass = pbClass;
    }

    public void initInvokeFunc() throws NoSuchMethodException {
        this.parseFunction = this.pbClass.getMethod("parseFrom", byte[].class);
    }

    public void initMetrics(RuntimeContext runtimeContext) {
        invokeSuccessCount = runtimeContext.getMetricGroup().counter("ug_invoke_success_count");
        invokeFailCount = runtimeContext.getMetricGroup().counter("ug_invoke_fail_count");
    }

    @Override
    public Message deserialize(byte[] bytes) throws IOException {
        try {
            Message message = (Message) parseFunction.invoke(null, bytes);
            invokeSuccessCount.inc();
            return message;
        } catch (IllegalAccessException | InvocationTargetException e) {
            e.printStackTrace();
            invokeFailCount.inc();
        }
        return null;
    }

    @Override
    public byte[] serialize(Message message) {
        return message.toByteArray();
    }

    @Override
    public boolean isEndOfStream(Message message) {
        return false;
    }

    @Override
    public TypeInformation<Message> getProducedType() {
        return TypeExtractor.getForClass(Message.class);
    }
}
