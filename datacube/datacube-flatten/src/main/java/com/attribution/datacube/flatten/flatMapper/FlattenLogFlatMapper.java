package com.attribution.datacube.flatten.flatMapper;

import com.attribution.datacube.common.flatten.FlattenedJoinedLog;
import com.attribution.datacube.proto.pageview.PageviewService;
import com.attribution.datacube.proto.rta.RTAStrategyService;
import org.apache.flink.api.common.functions.RichFlatMapFunction;
import org.apache.flink.util.Collector;

import java.util.List;

public class FlattenLogFlatMapper extends RichFlatMapFunction<PageviewService.Pageview, FlattenedJoinedLog> {
    // TODO 这里暂时将strategy的数据扁平化到
    @Override
    public void flatMap(PageviewService.Pageview pageview, Collector<FlattenedJoinedLog> collector) throws Exception {
        String traceId = pageview.getTraceId();
        long processTime = pageview.getProcessTime();
        PageviewService.Device device = pageview.getDevice();
        String fromPlatform = pageview.getFromPlatform();
        int innerCostMs = pageview.getInnerCostMs();
        List<RTAStrategyService.RTAStrategy> hitStrategyList = pageview.getHitStrategyList();
        if (hitStrategyList.size() > 0) {
            for (RTAStrategyService.RTAStrategy rtaStrategy : hitStrategyList) {
                collector.collect(FlattenedJoinedLog.builder()
                        .traceId(traceId)
                        .processTime(processTime)
                        .fromPlatform(fromPlatform)
                        .innerCostMs(innerCostMs)
                        .imei(device.getImei())
                        .idfa(device.getIdfa())
                        .oaid(device.getOaid())
                        .rtaStrategyId(rtaStrategy.getId())
                        .rtaValidPlatform(rtaStrategy.getPlatform())
                        .strategyIds(rtaStrategy.getStrategyIdList())
                        .hitStrategyCount(1)
                        .build());
            }
        } else {
            collector.collect(FlattenedJoinedLog.builder()
                    .traceId(traceId)
                    .processTime(processTime)
                    .fromPlatform(fromPlatform)
                    .innerCostMs(innerCostMs)
                    .imei(device.getImei())
                    .idfa(device.getIdfa())
                    .oaid(device.getOaid())
                    .requestCount(1)
                    .hitStrategyCount(0)
                    .build());
        }
    }
}
