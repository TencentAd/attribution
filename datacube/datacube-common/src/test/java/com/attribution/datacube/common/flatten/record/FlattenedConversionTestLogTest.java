package com.attribution.datacube.common.flatten.record;

import junit.framework.TestCase;
import org.junit.Test;

import java.lang.reflect.Field;

public class FlattenedConversionTestLogTest {
    @Test
    public void testLog() throws ClassNotFoundException {
        Class<?> aClass = Class.forName("com.attribution.datacube.common.flatten.record.TestClass");
        Field[] fields = aClass.getFields();
        System.out.println(fields.length);
    }

}