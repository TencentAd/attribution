//package com.attribution.datacube.join.process;
//
//import com.attribution.datacube.proto.pageview.PageviewService;
//import org.apache.flink.api.common.state.ListState;
//import org.apache.flink.api.common.state.ListStateDescriptor;
//import org.apache.flink.api.common.state.ValueState;
//import org.apache.flink.api.common.state.ValueStateDescriptor;
//import org.apache.flink.api.java.tuple.Tuple;
//import org.apache.flink.api.java.tuple.Tuple2;
//import org.apache.flink.configuration.Configuration;
//import org.apache.flink.streaming.api.functions.KeyedProcessFunction;
//import org.apache.flink.util.Collector;
//
//import java.util.Iterator;
//
//public class JoinProcessFunction
//        extends KeyedProcessFunction<String, Tuple2<String, PageviewService.Pageview>, Tuple2<String, PageviewService.Pageview>> {
//    private transient ValueState<PageviewService.Pageview> requestState;
//    private transient ListState<PageviewService.Pageview> strategyState;
//    private transient ListState<PageviewService.Pageview> impressionState;
//
//    @Override
//    public void open(Configuration parameters) throws Exception {
//        ValueStateDescriptor<PageviewService.Pageview> requestStateDescriptor =
//                new ValueStateDescriptor<>("request state", PageviewService.Pageview.class);
//        ListStateDescriptor<PageviewService.Pageview> strategyStateDescriptor =
//                new ListStateDescriptor<PageviewService.Pageview>("strategy state", PageviewService.Pageview.class);
//        ListStateDescriptor<PageviewService.Pageview> impressionStateDescriptor =
//                new ListStateDescriptor<PageviewService.Pageview>("impression statae", PageviewService.Pageview.class);
//        requestState = getRuntimeContext().getState(requestStateDescriptor);
//        strategyState = getRuntimeContext().getListState(strategyStateDescriptor);
//        impressionState = getRuntimeContext().getListState(impressionStateDescriptor);
//
//        // todo 这里可能后期添加prometheus监控
//    }
//
//    @Override
//    public void processElement(
//            Tuple2<String, PageviewService.Pageview> value,
//            Context ctx,
//            Collector<Tuple2<String, PageviewService.Pageview>> out)
//            throws Exception {
//        // todo 这里实现join的逻辑
//        String dataSource = value.f0;
//        PageviewService.Pageview message = value.f1;
//
//        // TODO 添加用户画像
//
//        // 这里添加策略的数据
//        if ("request".equals(dataSource)) {
//            PageviewService.Pageview request = PageviewService.Pageview.newBuilder()
//                    .setProcessTime(message.getProcessTime())
//                    .setTraceId(message.getTraceId())
//                    .setDevice(message.getDevice())
//                    .setFromPlatform(message.getFromPlatform())
//                    .setInnerCostMs(message.getInnerCostMs())
//                    .build();
//
//            requestState.update(request);
//
//            // 设置时间窗口为30分钟
//            // TODO 这里的30分钟应该添加到配置
//            ctx.timerService().registerEventTimeTimer(request.getProcessTime() * 1000L + 30 * 60 * 1000);
//
//            Iterable<PageviewService.Pageview> strategies = strategyState.get();
//            if (strategies != null) {
//                Iterator<PageviewService.Pageview> iterator = strategies.iterator();
//                while (iterator.hasNext()) {
//                    PageviewService.Pageview strategy = iterator.next();
//                    out.collect(Tuple2.of("strategy", paddingData(request, strategy, "strategy")));
//                }
//            }
//            strategyState.clear();
//        }
//
//        if ("strategy".equals(dataSource)) {
//            PageviewService.Pageview request = requestState.value();
//            if (request == null) {
//                strategyState.add(message);
//                ctx.timerService().registerEventTimeTimer(message.getProcessTime() * 1000L + 30 * 60 * 1000);
//            } else {
//                out.collect(Tuple2.of("strategy", paddingData(request, message, "strategy")));
//            }
//        }
//        // TODO 后续还有imp需要处理
//        if ("impression".equals(dataSource)) {
//            PageviewService.Pageview request = requestState.value();
//            if (request == null) {
//                impressionState.add(message);
//                ctx.timerService().registerEventTimeTimer(message.getProcessTime() * 1000L + 30 * 60 * 1000);
//            } else {
//                out.collect(Tuple2.of("impression", paddingData(request, message, "impression")));
//            }
//        }
//    }
//
//    /**
//     * 定时器，负责清理状态
//     *
//     * @param timestamp
//     * @param ctx
//     * @param out
//     * @throws Exception
//     */
//    @Override
//    public void onTimer(long timestamp, OnTimerContext ctx, Collector<Tuple2<String, PageviewService.Pageview>> out) throws Exception {
//        requestState.clear();
//        Iterable<PageviewService.Pageview> strategies = strategyState.get();
//        for (PageviewService.Pageview strategy : strategies) {
//            out.collect(Tuple2.of("strategy_empty", strategy));
//        }
//        strategyState.clear();
//    }
//
//    /**
//     * 这个方法是将数据(比如策略、曝光、点击等)追加到请求的数据后面
//     *
//     * @param request
//     * @param paddingItem
//     * @return
//     */
//    private PageviewService.Pageview paddingData(PageviewService.Pageview request, PageviewService.Pageview paddingItem, String type) throws Exception {
//        PageviewService.Pageview.Builder builder = PageviewService.Pageview.newBuilder();
//        builder.setProcessTime(request.getProcessTime())
//                .setTraceId(request.getTraceId())
//                .setDevice(request.getDevice())
//                .setFromPlatform(request.getFromPlatform())
//                .setInnerCostMs(request.getInnerCostMs());
//
//        switch (type) {
//            case "strategy": {
//                PageviewService.Pageview result = builder
//                        .addAllHitStrategy(paddingItem.getHitStrategyList())
//                        .build();
//                return result;
//            }
//            case "impression": {
//                PageviewService.Pageview result = builder
//                        .addAllImp(paddingItem.getImpList())
//                        .build();
//                return result;
//            }
//            default: {
//                throw new Exception("no such type");
//            }
//        }
//    }
//}
