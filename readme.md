# Data Streaming Api

Some Streaming data sources are available on the internet, ex: [twitter stream](https://developer.twitter.com/en/docs/tutorials/consuming-streaming-data).

There are not many sources and when they are available, they are limited and strict quotas applies. Hence limiting streaming dataflow and the amount of data you can receive in a time window.

In order to load test Data ingestion pipelines, we need to be able to generate this type of service ourself.

This project aims to implement this feature.

## service detail

Code is commented but here is the big picture of what this server should enable us to do: 

```mermaid

    graph LR

        subgraph server
            generator --> channel
            channel --> handler1
            channel --> handlerN
        end

        subgraph client
            handler1 --> consumer1
            handlerN --> consumerN
            consumer1 --> processing
            consumerN --> processing
            processing --> streamToDataWarehouse
        end
```

the client part should be a scalable data pipeline built with dataflow spark or kubernetes