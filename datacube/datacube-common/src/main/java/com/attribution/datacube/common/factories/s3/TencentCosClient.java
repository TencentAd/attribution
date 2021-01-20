package com.attribution.datacube.common.factories.s3;

import com.qcloud.cos.COSClient;
import com.qcloud.cos.ClientConfig;
import com.qcloud.cos.auth.BasicCOSCredentials;
import com.qcloud.cos.auth.COSCredentials;
import com.qcloud.cos.model.PutObjectRequest;
import com.qcloud.cos.region.Region;
import com.typesafe.config.Config;

import java.io.File;
import java.io.InputStream;
import java.util.List;

public class TencentCosClient extends S3Client {
    private COSClient cosClient;

    public TencentCosClient(Config config) {
        String secretId = config.getString("secret-id");
        String secretKey = config.getString("secret-key");
        String region = config.getString("region");
        COSCredentials cred = new BasicCOSCredentials(secretId, secretKey);
        ClientConfig clientConfig = new ClientConfig(new Region(region));
        cosClient = new COSClient(cred, clientConfig);
    }

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
        File file = new File("");
        PutObjectRequest putObjectRequest = new PutObjectRequest(bucketName, key, file);
        cosClient.putObject(putObjectRequest);
    }

}
