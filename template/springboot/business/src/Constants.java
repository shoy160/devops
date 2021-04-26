package com.yunzhicloud.network.dialtest;

/**
 * @author shay
 * @date 2021/3/18
 */
public interface DialTestConstants {
    String SERVICE_NAME = "dial-test";
    String VERSION = "1.0.0";
    String MAPPER_PACKAGE = "com.yunzhicloud.network.dialtest.dao";
    String REST_PACKAGE = "com.yunzhicloud.network.dialtest.web.rest";


    /**
     * redis-key
     * 存储任务批次探针
     */
    String TASK_PROBE_BATCH_KEY = "task_probe_batch_";

    /**
     * 计算比例 25/40=0.625
     */
     double TARGET_DOMAIN_RATE = 0.625;


}
