package com.attribution.datacube.common.factories.s3;

import java.io.InputStream;
import java.io.Serializable;
import java.util.List;

public class AmazonS3Client extends S3Client implements Serializable {
    @Override
    public List<String> getKeyListFromPrefix(String bucketName, String prefix) {
        return null;
    }

    @Override
    public List<String> getPathFromPrefix(String bucketName, String prefix) {
        return null;
    }

    @Override
    public InputStream getObjectInputStream(String bucketName, String key) {
        return null;
    }

    @Override
    public void putObject(String bucketName, String key) {

    }
}
