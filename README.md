# Distributed-Crawler
Realized a distributed crawler by native code, with Micro-Service architecture.

Implemented crawler in 3 ways: 

1. Single task version
BFS Algo fetch HTML text, regex engine parsed useful info, append new valid links into task queue. Docker+ElasticSerach built the data repository, realized data persistence. Combined with Go template to create front-end demonstration webpages.

2. Concurrent version
Created multiple workers and fetchers to do the task simultaneously, wrote serveral schedulers that could realize different concurrency strategies

3. Distributed version
Decomposed the system into micro-serivce architecture, using jsonrpc communicates through different serivces
