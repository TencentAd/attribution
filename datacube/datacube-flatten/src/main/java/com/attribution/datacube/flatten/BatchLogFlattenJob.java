package com.attribution.datacube.flatten;

import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;
import org.apache.flink.api.java.ExecutionEnvironment;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;


public class BatchLogFlattenJob {
    private static final Logger LOG = LoggerFactory.getLogger(BatchLogFlattenJob.class);

    public static void main(String[] args) {
        String env = args[0];
        String jobName = args[1];

        Config config = ConfigFactory.load("conv-flattened-streaming.conf")
                .getConfig(env).getConfig(jobName);

        String logType = config.getString("log-type");
        String savingPath = config.getString("saving-dir");

        Config batchConfig = config.getConfig("batch-config");

        int checkpointInterval = config.getInt("checkpoint-interval");
        String checkpointDir = config.getString("checkpoint-dir");

        ExecutionEnvironment ee = ExecutionEnvironment.getExecutionEnvironment();

    }
}
