package com.attribution.datacube.common.factories.s3;

import java.io.InputStream;
import java.util.List;

public abstract class S3Client {
    public abstract List<String> getKeyListFromPrefix(String bucketName, String prefix);

    public abstract List<String> getPathFromPrefix(String bucketName, String prefix);

    public abstract InputStream getObjectInputStream(String bucketName, String key);

    public abstract void putObject(String bucketName, String key);
}
