package com.attribution.datacube.common.factories.s3;

import com.typesafe.config.Config;
import software.amazon.awssdk.services.securityhub.model.AwsCertificateManagerCertificateDetails;

public class S3ClientFactory {
    public static S3Client getS3Client(Config config) throws Exception {
        String clientType = config.getString("s3-type");
        switch(clientType) {
            case "s3": {
                return new AmazonS3Client(config);
            }
            case "tencent": {
                return new TencentCosClient(config);
            }
            default: {
                throw new Exception("no such kind of s3 client");
            }
        }
    }
}
