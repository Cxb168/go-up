1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

   环境：CPU：i7-10700,   内网：32G， Window 10 64位

   命令：.\redis-benchmark.exe -t set,get -n 100000 -d {字节大小}

   结论：数据大小对读写性能影响不大。对比10k以下不同数据大小的Value，redis读写性能基本没变；即使达到100k后，99线也在1ms内；

2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

   写入50w数据，key大小 10 byte，

   字节value大小：          10，     20，     50，   100，   200，  1k，      5k ，

   平均内存占用分别为：120.4,  120.4,  152.4,  200.4,  312.4,  1368.4, 8280.4

   结论：随着value增大，平均内存占用有增长，但并非线性增长。 

   





| 大小（k） | get                                                          | set                                                          |
| :-------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 10        | ====== GET ======<br/>  100000 requests completed in 0.67 seconds<br/>  50 parallel clients<br/>  10 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>150150.14 requests per second | ====== SET ======<br/>  100000 requests completed in 0.68 seconds<br/>  50 parallel clients<br/>  10 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>147058.83 requests per second |
| 20        | ====== GET ======<br/>  100000 requests completed in 0.67 seconds<br/>  50 parallel clients<br/>  20 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>149476.83 requests per second | ====== SET ======<br/>  100000 requests completed in 0.70 seconds<br/>  50 parallel clients<br/>  20 bytes payload<br/>  keep alive: 1<br/><br/>99.90% <= 1 milliseconds<br/>99.95% <= 3 milliseconds<br/>100.00% <= 3 milliseconds<br/>142857.14 requests per second |
| 50        | ====== GET ======<br/>  100000 requests completed in 0.66 seconds<br/>  50 parallel clients<br/>  50 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>152439.02 requests per second | ====== SET ======<br/>  100000 requests completed in 0.70 seconds<br/>  50 parallel clients<br/>  50 bytes payload<br/>  keep alive: 1<br/><br/>99.90% <= 1 milliseconds<br/>99.95% <= 3 milliseconds<br/>100.00% <= 3 milliseconds<br/>143472.02 requests per second |
| 100       | ====== GET ======<br/>  100000 requests completed in 0.66 seconds<br/>  50 parallel clients<br/>  100 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>150829.56 requests per second | ====== SET ======<br/>  100000 requests completed in 0.69 seconds<br/>  50 parallel clients<br/>  100 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>143884.89 requests per second |
| 200       | ====== GET ======<br/>  100000 requests completed in 0.66 seconds<br/>  50 parallel clients<br/>  200 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>150829.56 requests per second | ====== SET ======<br/>  100000 requests completed in 0.69 seconds<br/>  50 parallel clients<br/>  200 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>144508.67 requests per second |
| 1k        | ====== GET ======<br/>  100000 requests completed in 0.67 seconds<br/>  50 parallel clients<br/>  1024 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>148809.53 requests per second | ====== SET ======<br/>  100000 requests completed in 0.70 seconds<br/>  50 parallel clients<br/>  1024 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>142045.45 requests per second |
| 5k        | ====== GET ======<br/>  100000 requests completed in 0.71 seconds<br/>  50 parallel clients<br/>  5120 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>140252.45 requests per second | ====== SET ======<br/>  100000 requests completed in 0.75 seconds<br/>  50 parallel clients<br/>  5120 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>132978.73 requests per second |
| 10k       | ====== GET ======<br/>  100000 requests completed in 0.77 seconds<br/>  50 parallel clients<br/>  10240 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>130548.30 requests per second | ====== SET ======<br/>  100000 requests completed in 0.79 seconds<br/>  50 parallel clients<br/>  10240 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>126262.63 requests per second |
| 20k       | ====== GET ======<br/>  100000 requests completed in 1.49 seconds<br/>  50 parallel clients<br/>  20480 bytes payload<br/>  keep alive: 1<br/><br/>100.00% <= 0 milliseconds<br/>66934.41 requests per second | ====== SET ======<br/>  100000 requests completed in 1.09 seconds<br/>  50 parallel clients<br/>  20480 bytes payload<br/>  keep alive: 1<br/><br/>99.86% <= 1 milliseconds<br/>99.91% <= 2 milliseconds<br/>99.95% <= 3 milliseconds<br/>100.00% <= 4 milliseconds<br/>100.00% <= 4 milliseconds<br/>91491.30 requests per second |
| 50k       | ====== GET ======<br/>  100000 requests completed in 2.00 seconds<br/>  50 parallel clients<br/>  51200 bytes payload<br/>  keep alive: 1<br/><br/>91.80% <= 1 milliseconds<br/>100.00% <= 1 milliseconds<br/>50050.05 requests per second | ====== SET ======<br/>  100000 requests completed in 1.62 seconds<br/>  50 parallel clients<br/>  51200 bytes payload<br/>  keep alive: 1<br/><br/>99.54% <= 1 milliseconds<br/>99.99% <= 2 milliseconds<br/>100.00% <= 2 milliseconds<br/>61842.92 requests per second |
| 100k      | ====== GET ======<br/>  100000 requests completed in 2.80 seconds<br/>  50 parallel clients<br/>  102400 bytes payload<br/>  keep alive: 1<br/><br/>99.97% <= 1 milliseconds<br/>100.00% <= 1 milliseconds<br/>35676.06 requests per second | ====== SET ======<br/>  100000 requests completed in 2.48 seconds<br/>  50 parallel clients<br/>  102400 bytes payload<br/>  keep alive: 1<br/><br/>0.02% <= 1 milliseconds<br/>99.67% <= 2 milliseconds<br/>99.84% <= 3 milliseconds<br/>99.91% <= 4 milliseconds<br/>99.93% <= 5 milliseconds<br/>99.98% <= 6 milliseconds<br/>100.00% <= 6 milliseconds<br/>40273.86 requests per second |





# 原始内存情况

used_memory:712848
used_memory_human:696.14K
used_memory_rss:675920
used_memory_peak:54083472
used_memory_peak_human:51.58M
used_memory_lua:35840
mem_fragmentation_ratio:0.95
mem_allocator:jemalloc-3.6.0

| value大小（byte） | 平均内存占用 | 内存情况                                                     |
| ----------------- | ------------ | ------------------------------------------------------------ |
| 10                | 120.4        | used_memory:60907192<br/>used_memory_human:58.09M<br/>used_memory_rss:60870264<br/>used_memory_peak:61105032<br/>used_memory_peak_human:58.27M<br/>used_memory_lua:35840<br/> |
| 20                | 120.4        | used_memory:60907632<br/>used_memory_human:58.09M<br/>used_memory_rss:60870704<br/>used_memory_peak:61105472<br/>used_memory_peak_human:58.27M<br/>used_memory_lua:35840<br/> |
| 50                | 152.4        | used_memory:76908040<br/>used_memory_human:73.35M<br/>used_memory_rss:76871112<br/>used_memory_peak:77105912<br/>used_memory_peak_human:73.53M<br/>used_memory_lua:35840<br/> |
| 100               | 200.4        | used_memory:100908152<br/>used_memory_human:96.23M<br/>used_memory_rss:100871224<br/>used_memory_peak:100908152<br/>used_memory_peak_human:96.23M<br/>used_memory_lua:35840<br/> |
| 200               | 312.4        | used_memory:156908632<br/>used_memory_human:149.64M<br/>used_memory_rss:157106472<br/>used_memory_peak:157106472<br/>used_memory_peak_human:149.83M<br/>used_memory_lua:35840<br/> |
| 1k                | 1368.4       | used_memory:684909312<br/>used_memory_human:653.18M<br/>used_memory_rss:684872384<br/>used_memory_peak:4140909160<br/>used_memory_peak_human:3.86G<br/>used_memory_lua:35840<br/> |
| 5k                | 8280.4       | used_memory:4140909160<br/>used_memory_human:3.86G<br/>used_memory_rss:4140872232<br/>used_memory_peak:4140909160<br/>used_memory_peak_human:3.86G<br/>used_memory_lua:35840<br/> |

